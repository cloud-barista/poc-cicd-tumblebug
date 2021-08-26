package goscenario

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/exec"
	"strconv"
	"sync"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"
	"xorm.io/xorm"
	"xorm.io/xorm/names"

	"github.com/cloud-barista/poc-cicd-tumblebug/src/api/grpc/config"
	"github.com/cloud-barista/poc-cicd-tumblebug/src/api/grpc/logger"
	"github.com/cloud-barista/poc-cicd-tumblebug/src/core/common"
	"github.com/cloud-barista/poc-cicd-tumblebug/src/core/mcir"
	"github.com/cloud-barista/poc-cicd-tumblebug/src/core/mcis"
	"github.com/go-resty/resty/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"

	cbstore "github.com/cloud-barista/cb-store"
	gc "github.com/cloud-barista/poc-cicd-tumblebug/src/api/grpc/common"
	pb "github.com/cloud-barista/poc-cicd-tumblebug/src/api/grpc/protobuf/cbtumblebug"
	grpc_common "github.com/cloud-barista/poc-cicd-tumblebug/src/api/grpc/server/common"
	grpc_mcir "github.com/cloud-barista/poc-cicd-tumblebug/src/api/grpc/server/mcir"
	grpc_mcis "github.com/cloud-barista/poc-cicd-tumblebug/src/api/grpc/server/mcis"

	api "github.com/cloud-barista/poc-cicd-tumblebug/src/api/grpc/request"

	"bou.ke/monkey"
)

type TestCases struct {
	Name                string
	Instance            interface{}
	Method              string
	Args                []interface{}
	ExpectResStartsWith string
	ExpectResContains   string
}

var (
	holdStdout *os.File        = nil
	nullOut    *os.File        = nil
	NsApi      *api.NSApi      = nil
	McirApi    *api.MCIRApi    = nil
	McisApi    *api.MCISApi    = nil
	TbutilApi  *api.UtilityApi = nil
	gs         *grpc.Server    = nil
)

func init() {
	logrus.SetLevel(logrus.ErrorLevel)
}

func SetUpForGrpc() {

	holdStdout = os.Stdout
	nullOut, _ := os.Open(os.DevNull)
	os.Stdout = nullOut

	cbstore.InitData()

	if _, err := os.Stat("../meta_db/dat/cbtumblebug.s3db"); err == nil {
		os.Remove("../meta_db/dat/cbtumblebug.s3db")
	}
	if _, err := os.Stat("./benchmarking.csv"); err == nil {
		os.Remove("./benchmarking.csv")
	}
	if _, err := os.Stat("./rttmap.csv"); err == nil {
		os.Remove("./rttmap.csv")
	}

	/**
	** Backend Server Setup
	**/
	client := resty.New().SetCloseConnection(true)

	cmd := exec.Command("./stop.sh")
	cmd.Dir = "../backend"
	cmd.Run()

	cmd = exec.Command("./start.sh")
	cmd.Dir = "../backend"
	cmd.Start()

	backendFlg := false
	for i := 0; i < 60; i++ {
		//fmt.Printf("backend server waiting... \n")
		time.Sleep(time.Second * 5)

		_, err := client.R().
			Get("http://localhost:31024/spider/")

		if err == nil {
			backendFlg = true
			break
		}
	}

	if !backendFlg {
		log.Fatalf("backend server failed...\n")
	}

	client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(`{"DriverName":"mock-unit-driver01","ProviderName":"MOCK", "DriverLibFileName":"mock-driver-v1.0.so"}`).
		Post("http://localhost:31024/spider/driver")

	client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(`{"CredentialName":"mock-unit-credential01","ProviderName":"MOCK", "KeyValueInfoList": [{"Key":"MockName", "Value":"mock_unit_name00"}]}`).
		Post("http://localhost:31024/spider/credential")

	client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(`{"RegionName":"mock-unit-region01","ProviderName":"MOCK", "KeyValueInfoList": [{"Key":"Region", "Value":"default"}]}`).
		Post("http://localhost:31024/spider/region")

	client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(`{"ConfigName":"mock-unit-config01","ProviderName":"MOCK", "DriverName":"mock-unit-driver01", "CredentialName":"mock-unit-credential01", "RegionName":"mock-unit-region01"}`).
		Post("http://localhost:31024/spider/connectionconfig")

	/**
	** Tumblebug Env Setup
	**/
	common.SPIDER_REST_URL = common.NVL(os.Getenv("SPIDER_REST_URL"), "http://localhost:1024/spider")
	common.DRAGONFLY_REST_URL = common.NVL(os.Getenv("DRAGONFLY_REST_URL"), "http://localhost:9090/dragonfly")
	common.DB_URL = common.NVL(os.Getenv("DB_URL"), "localhost:3306")
	common.DB_DATABASE = common.NVL(os.Getenv("DB_DATABASE"), "cb_tumblebug")
	common.DB_USER = common.NVL(os.Getenv("DB_USER"), "cb_tumblebug")
	common.DB_PASSWORD = common.NVL(os.Getenv("DB_PASSWORD"), "cb_tumblebug")
	common.AUTOCONTROL_DURATION_MS = common.NVL(os.Getenv("AUTOCONTROL_DURATION_MS"), "10000")

	// load the latest configuration from DB (if exist)
	fmt.Println("")
	fmt.Println("[Update system environment]")
	common.UpdateGlobalVariable(common.StrDRAGONFLY_REST_URL)
	common.UpdateGlobalVariable(common.StrSPIDER_REST_URL)
	common.UpdateGlobalVariable(common.StrAUTOCONTROL_DURATION_MS)

	// load config
	//masterConfigInfos = confighandler.GetMasterConfigInfos()

	//Setup database (meta_db/dat/cbtumblebug.s3db)
	fmt.Println("")
	fmt.Println("[Setup SQL Database]")

	err := os.MkdirAll("../meta_db/dat/", os.ModePerm)
	if err != nil {
		fmt.Println(err.Error())
	}

	//err = common.OpenSQL("../meta_db/dat/cbtumblebug.s3db") // commented out to move to use XORM
	common.ORM, err = xorm.NewEngine("sqlite3", "../meta_db/dat/cbtumblebug.s3db")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Database access info set successfully")
	}
	//common.ORM.SetMapper(names.SameMapper{})
	common.ORM.SetTableMapper(names.SameMapper{})
	common.ORM.SetColumnMapper(names.SameMapper{})

	/* // Required if using MySQL // Not required if using SQLite
	err = common.SelectDatabase(common.DB_DATABASE)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("DB selected successfully..")
	}
	*/

	// "CREATE Table IF NOT EXISTS spec(...)"
	//err = common.CreateSpecTable() // commented out to move to use XORM
	err = common.ORM.Sync2(new(mcir.TbSpecInfo))
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Table spec set successfully..")
	}

	// "CREATE Table IF NOT EXISTS image(...)"
	//err = common.CreateImageTable() // commented out to move to use XORM
	err = common.ORM.Sync2(new(mcir.TbImageInfo))
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Table image set successfully..")
	}

	//defer db.Close()

	//Ticker for MCIS Orchestration Policy
	fmt.Println("")
	fmt.Println("[Initiate Multi-Cloud Orchestration]")

	autoControlDuration, _ := strconv.Atoi(common.AUTOCONTROL_DURATION_MS) //ms
	ticker := time.NewTicker(time.Millisecond * time.Duration(autoControlDuration))
	go func() {
		for t := range ticker.C {
			//display ticker if you need (remove '_ = t')
			_ = t
			//fmt.Println("- Orchestration Controller ", t.Format("2006-01-02 15:04:05"))
			mcis.OrchestrationController()
		}
	}()
	defer ticker.Stop()

	/**
	** Tumblebug Grpc Server Setup
	**/
	listener := bufconn.Listen(1024 * 1024)

	monkey.Patch(gc.NewCBConnection, func(gConf *config.GrpcClientConfig) (*gc.CBConnection, io.Closer, error) {
		conn, _ := grpc.DialContext(context.Background(), "", grpc.WithInsecure(), grpc.WithContextDialer(
			func(context.Context, string) (net.Conn, error) {
				return listener.Dial()
			}))
		return &gc.CBConnection{Conn: conn}, nil, nil
	})

	logger := logger.NewLogger()

	tumblebugsrv := &config.GrpcServerConfig{
		Addr: "127.0.0.1:30252",
	}

	cbserver, closer, err := gc.NewCBServer(tumblebugsrv)
	if err != nil {
		logger.Fatal("failed to create grpc server: ", err)
	}

	gs = cbserver.Server
	pb.RegisterUtilityServer(gs, &grpc_common.UtilityService{})
	pb.RegisterNSServer(gs, &grpc_common.NSService{})
	pb.RegisterMCIRServer(gs, &grpc_mcir.MCIRService{})
	pb.RegisterMCISServer(gs, &grpc_mcis.MCISService{})

	go func() {

		if closer != nil {
			defer closer.Close()
		}

		if err := gs.Serve(listener); err != nil {
			logger.Fatal("failed to serve: ", err)
		}
	}()

	time.Sleep(time.Second * 2)

	/**
	** Tumblebug Grpc API Setup
	**/
	NsApi = api.NewNSManager()

	err = NsApi.SetConfigPath("../conf/grpc_conf.yaml")
	if err != nil {
		logger.Fatal(err)
	}

	err = NsApi.Open()
	if err != nil {
		logger.Fatal(err)
	}

	McirApi = api.NewMCIRManager()

	err = McirApi.SetConfigPath("../conf/grpc_conf.yaml")
	if err != nil {
		logger.Fatal(err)
	}

	err = McirApi.Open()
	if err != nil {
		logger.Fatal(err)
	}

	McisApi = api.NewMCISManager()

	err = McisApi.SetConfigPath("../conf/grpc_conf.yaml")
	if err != nil {
		logger.Fatal(err)
	}

	err = McisApi.Open()
	if err != nil {
		logger.Fatal(err)
	}

	TbutilApi = api.NewUtilityManager()

	err = TbutilApi.SetConfigPath("../conf/grpc_conf.yaml")
	if err != nil {
		logger.Fatal(err)
	}

	err = TbutilApi.Open()
	if err != nil {
		logger.Fatal(err)
	}

	/**
	** Function Patch for Testing
	**/
	monkey.Patch(mcis.CheckConnectivity, func(host string, port string) error {
		return nil
	})

	monkey.Patch(mcis.SSHRun, func(sshInfo mcis.SSHInfo, cmd string) (string, error) {
		//fmt.Printf("ssh cmd : %s\n", cmd)

		return cmd + " success", nil
	})

	monkey.Patch(mcis.SSHCopy, func(sshInfo mcis.SSHInfo, sourcePath string, remotePath string) error {
		return nil
	})

	monkey.Patch(mcis.CheckDragonflyEndpoint, func() error {
		return nil
	})

	monkey.Patch(mcis.GetCloudLocation, func(cloudType string, nativeRegion string) mcis.GeoLocation {
		location := mcis.GeoLocation{}

		location.CloudType = cloudType
		location.NativeRegion = nativeRegion
		location.BriefAddr = "South Korea (Seoul)"
		location.Latitude = "37.4767"
		location.Longitude = "126.8841"

		return location
	})

	monkey.Patch(mcis.CallMonitoringAsync, func(wg *sync.WaitGroup, nsID string, mcisID string, vmID string, givenUserName string, method string, cmd string, returnResult *[]mcis.SshCmdResult) {
		defer wg.Done() //goroutin sync done

		vmIP, _ := mcis.GetVmIp(nsID, mcisID, vmID)
		vmInfoTmp, _ := mcis.GetVmObject(nsID, mcisID, vmID)
		vmInfoTmp.MonAgentStatus = "installing"
		mcis.UpdateVmInfo(nsID, mcisID, vmInfoTmp)

		sshResultTmp := mcis.SshCmdResult{}
		sshResultTmp.McisId = mcisID
		sshResultTmp.VmId = vmID
		sshResultTmp.VmIp = vmIP

		sshResultTmp.Result = "CallMonitoringAsync result"
		sshResultTmp.Err = nil
		*returnResult = append(*returnResult, sshResultTmp)
		vmInfoTmp.MonAgentStatus = "installed"

		mcis.UpdateVmInfo(nsID, mcisID, vmInfoTmp)
	})

	monkey.Patch(mcis.CallGetMonitoringAsync, func(wg *sync.WaitGroup, nsID string, mcisID string, vmID string, vmIP string, method string, metric string, cmd string, returnResult *[]mcis.MonResultSimple) {
		defer wg.Done() //goroutin sync done

		ResultTmp := mcis.MonResultSimple{}
		ResultTmp.VmId = vmID
		ResultTmp.Metric = metric

		ResultTmp.Value = "0"
		*returnResult = append(*returnResult, ResultTmp)
	})

	monkey.Patch(mcis.CallMilkyway, func(wg *sync.WaitGroup, vmList []string, nsId string, mcisId string, vmId string, vmIp string, action string, option string, results *mcis.BenchmarkInfoArray) {
		defer wg.Done()

		resultTmp := mcis.BenchmarkInfo{}
		if action == "init" || action == "clean" || action == "install" {
			resultTmp.Result = action + " complete"
			resultTmp.Unit = ""
		} else {
			resultTmp.Result = "1.0"
			resultTmp.Unit = "unit"
		}
		resultTmp.Desc = action + " complete"
		resultTmp.Elapsed = "1.0"

		resultTmp.SpecId = mcis.GetVmSpecId(nsId, mcisId, vmId)
		results.ResultArray = append(results.ResultArray, resultTmp)
	})
}

func TearDownForGrpc() {
	NsApi.Close()
	McirApi.Close()
	McisApi.Close()
	TbutilApi.Close()
	gs.Stop()

	cmd := exec.Command("./stop.sh")
	cmd.Dir = "../backend"
	cmd.Run()

	cbstore.InitData()

	if _, err := os.Stat("../meta_db/dat/cbtumblebug.s3db"); err == nil {
		os.Remove("../meta_db/dat/cbtumblebug.s3db")
	}
	if _, err := os.Stat("./benchmarking.csv"); err == nil {
		os.Remove("./benchmarking.csv")
	}
	if _, err := os.Stat("./rttmap.csv"); err == nil {
		os.Remove("./rttmap.csv")
	}

	nullOut.Close()
	os.Stdout = holdStdout
}
