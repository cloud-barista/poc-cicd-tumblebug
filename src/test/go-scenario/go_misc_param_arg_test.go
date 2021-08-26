package goscenario

import (
	"testing"

	api "github.com/cloud-barista/poc-cicd-tumblebug/src/api/grpc/request"
	core_common "github.com/cloud-barista/poc-cicd-tumblebug/src/core/common"
	core_mcir "github.com/cloud-barista/poc-cicd-tumblebug/src/core/mcir"
	core_mcis "github.com/cloud-barista/poc-cicd-tumblebug/src/core/mcis"
)

func TestGoMiscParamArg(t *testing.T) {
	t.Run("go api misc test for mock driver by parameter args style", func(t *testing.T) {
		SetUpForGrpc()

		tc := TestCases{
			Name:     "create namespace",
			Instance: NsApi,
			Method:   "CreateNSByParam",
			Args: []interface{}{
				&core_common.NsReq{
					Name:        "ns-unit-01",
					Description: "NameSpace for General Testing",
				},
			},
			ExpectResStartsWith: `{"id":"ns-unit-01"`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:                "list namespace id",
			Instance:            NsApi,
			Method:              "ListNSId",
			Args:                nil,
			ExpectResStartsWith: `{"idList":["ns-unit-01"]}`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "create vnet",
			Instance: McirApi,
			Method:   "CreateVNetByParam",
			Args: []interface{}{
				&api.TbVNetCreateRequest{
					NsId: "ns-unit-01",
					Item: core_mcir.TbVNetReq{
						Name:           "mock-unit-config01-dev",
						ConnectionName: "mock-unit-config01",
						CidrBlock:      "192.168.0.0/16",
						SubnetInfoList: []core_mcir.SpiderSubnetReqInfo{
							core_mcir.SpiderSubnetReqInfo{
								Name:         "mock-unit-config01-dev",
								IPv4_CIDR:    "192.168.1.0/24",
								KeyValueList: []core_common.KeyValue{},
							},
						},
						Description: "",
					},
				},
			},
			ExpectResStartsWith: `{"id":"mock-unit-config01-dev"`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "check resource",
			Instance: McirApi,
			Method:   "CheckResourceByParam",
			Args: []interface{}{
				"ns-unit-01",
				"mock-unit-config01-dev",
				"vNet",
			},
			ExpectResStartsWith: `{"exists":true}`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "create security",
			Instance: McirApi,
			Method:   "CreateSecurityGroupByParam",
			Args: []interface{}{
				&api.TbSecurityGroupCreateRequest{
					NsId: "ns-unit-01",
					Item: core_mcir.TbSecurityGroupReq{
						Name:           "mock-unit-config01-dev",
						ConnectionName: "mock-unit-config01",
						VNetId:         "mock-unit-config01-dev",
						Description:    "test description",
						FirewallRules: &[]core_mcir.SpiderSecurityRuleInfo{
							core_mcir.SpiderSecurityRuleInfo{
								FromPort:   "1",
								ToPort:     "65535",
								IPProtocol: "tcp",
								Direction:  "inbound",
							},
							core_mcir.SpiderSecurityRuleInfo{
								FromPort:   "1",
								ToPort:     "65535",
								IPProtocol: "udp",
								Direction:  "inbound",
							},
							core_mcir.SpiderSecurityRuleInfo{
								FromPort:   "-1",
								ToPort:     "-1",
								IPProtocol: "icmp",
								Direction:  "inbound",
							},
						},
					},
				},
			},
			ExpectResStartsWith: `{"id":"mock-unit-config01-dev"`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "create sshkey",
			Instance: McirApi,
			Method:   "CreateSshKeyByParam",
			Args: []interface{}{
				&api.TbSshKeyCreateRequest{
					NsId: "ns-unit-01",
					Item: core_mcir.TbSshKeyReq{
						Name:           "mock-unit-config01-dev",
						ConnectionName: "mock-unit-config01",
						Description:    "",
					},
				},
			},
			ExpectResStartsWith: `{"id":"mock-unit-config01-dev"`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "register image with id",
			Instance: McirApi,
			Method:   "CreateImageWithIDByParam",
			Args: []interface{}{
				&api.TbImageCreateRequest{
					NsId: "ns-unit-01",
					Item: core_mcir.TbImageReq{
						Name:           "mock-unit-config01-dev",
						ConnectionName: "mock-unit-config01",
						CspImageId:     "mock-vmimage-01",
						Description:    "Canonical, Ubuntu, 18.04 LTS, amd64 bionic",
					},
				},
			},
			ExpectResStartsWith: `{"namespace":"ns-unit-01","id":"mock-unit-config01-dev"`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "connection fetch images",
			Instance: McirApi,
			Method:   "FetchImageByParam",
			Args: []interface{}{
				"mock-unit-config01",
				"ns-unit-01",
			},
			ExpectResStartsWith: `{"message":"Fetched 5 images (from 1 connConfigs)"}`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "register spec",
			Instance: McirApi,
			Method:   "CreateSpecWithSpecNameByParam",
			Args: []interface{}{
				&api.TbSpecCreateRequest{
					NsId: "ns-unit-01",
					Item: core_mcir.TbSpecReq{
						Name:           "mock-unit-config01-dev",
						ConnectionName: "mock-unit-config01",
						CspSpecName:    "mock-vmspec-01",
						Description:    "",
					},
				},
			},
			ExpectResStartsWith: `{"namespace":"ns-unit-01","id":"mock-unit-config01-dev"`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "connection fetch specs",
			Instance: McirApi,
			Method:   "FetchSpecByParam",
			Args: []interface{}{
				"mock-unit-config01",
				"ns-unit-01",
			},
			ExpectResStartsWith: `{"message":"Fetched 4 specs (from 1 connConfigs)"}`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "create mcis",
			Instance: McisApi,
			Method:   "CreateMcisByParam",
			Args: []interface{}{
				&api.TbMcisCreateRequest{
					NsId: "ns-unit-01",
					Item: core_mcis.TbMcisReq{
						Name:            "mock-unit-config01-dev",
						InstallMonAgent: "no",
						Description:     "Tumblebug Demo",
						Vm: []core_mcis.TbVmReq{
							core_mcis.TbVmReq{
								Name:           "mock-unit-config01-dev-01",
								ConnectionName: "mock-unit-config01",
								SpecId:         "mock-unit-config01-dev",
								ImageId:        "mock-unit-config01-dev",
								VNetId:         "mock-unit-config01-dev",
								SubnetId:       "mock-unit-config01-dev",
								SecurityGroupIds: []string{
									"mock-unit-config01-dev",
								},
								SshKeyId:       "mock-unit-config01-dev",
								VmUserAccount:  "cb-user",
								VmUserPassword: "",
								Description:    "description",
							},
						},
					},
				},
			},
			ExpectResStartsWith: `{"id":"mock-unit-config01-dev"`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "get mcis vm",
			Instance: McisApi,
			Method:   "GetMcisVMInfoByParam",
			Args: []interface{}{
				"ns-unit-01",
				"mock-unit-config01-dev",
				"mock-unit-config01-dev-01",
			},
			ExpectResStartsWith: `{"id":"mock-unit-config01-dev-01"`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "get mcis vm status",
			Instance: McisApi,
			Method:   "GetMcisVMStatusByParam",
			Args: []interface{}{
				"ns-unit-01",
				"mock-unit-config01-dev",
				"mock-unit-config01-dev-01",
			},
			ExpectResStartsWith: `{"id":"mock-unit-config01-dev-01"`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "suspend mcis vm",
			Instance: McisApi,
			Method:   "ControlMcisVMByParam",
			Args: []interface{}{
				"ns-unit-01",
				"mock-unit-config01-dev",
				"mock-unit-config01-dev-01",
				"suspend",
			},
			ExpectResStartsWith: `{"message":"Suspending the VM"}`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "resume mcis vm",
			Instance: McisApi,
			Method:   "ControlMcisVMByParam",
			Args: []interface{}{
				"ns-unit-01",
				"mock-unit-config01-dev",
				"mock-unit-config01-dev-01",
				"resume",
			},
			ExpectResStartsWith: `{"message":"Resuming the VM"}`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "reboot mcis vm",
			Instance: McisApi,
			Method:   "ControlMcisVMByParam",
			Args: []interface{}{
				"ns-unit-01",
				"mock-unit-config01-dev",
				"mock-unit-config01-dev-01",
				"reboot",
			},
			ExpectResStartsWith: `{"message":"Rebooting the VM"}`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "all benchmark",
			Instance: McisApi,
			Method:   "GetAllBenchmarkByParam",
			Args: []interface{}{
				"ns-unit-01",
				"mock-unit-config01-dev",
				"localhost",
			},
			ExpectResStartsWith: `{"resultarray":[{"result":"1.0","unit":"unit","desc":"mrtt complete"`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "benchmark",
			Instance: McisApi,
			Method:   "GetBenchmarkByParam",
			Args: []interface{}{
				"ns-unit-01",
				"mock-unit-config01-dev",
				"cpus",
				"localhost",
			},
			ExpectResStartsWith: `{"resultarray":[{"result":"1.0","unit":"unit","desc":"cpus complete"`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "check mcis",
			Instance: McisApi,
			Method:   "CheckMcisByParam",
			Args: []interface{}{
				"ns-unit-01",
				"mock-unit-config01-dev",
			},
			ExpectResStartsWith: `{"exists":true}`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "check mcis vm",
			Instance: McisApi,
			Method:   "CheckVmByParam",
			Args: []interface{}{
				"ns-unit-01",
				"mock-unit-config01-dev",
				"mock-unit-config01-dev-01",
			},
			ExpectResStartsWith: `{"exists":true}`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "terminate mcis vm",
			Instance: McisApi,
			Method:   "ControlMcisVMByParam",
			Args: []interface{}{
				"ns-unit-01",
				"mock-unit-config01-dev",
				"mock-unit-config01-dev-01",
				"terminate",
			},
			ExpectResStartsWith: `{"message":"Terminating the VM"}`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "create mcis(installMonAgent:yes,vmGroupSize>0)",
			Instance: McisApi,
			Method:   "CreateMcisByParam",
			Args: []interface{}{
				&api.TbMcisCreateRequest{
					NsId: "ns-unit-01",
					Item: core_mcis.TbMcisReq{
						Name:            "mock-unit-config01-dev2",
						InstallMonAgent: "yes",
						Description:     "Tumblebug Demo",
						Vm: []core_mcis.TbVmReq{
							core_mcis.TbVmReq{
								VmGroupSize:    "3",
								Name:           "mock-unit-config01-dev2-01",
								ConnectionName: "mock-unit-config01",
								SpecId:         "mock-unit-config01-dev",
								ImageId:        "mock-unit-config01-dev",
								VNetId:         "mock-unit-config01-dev",
								SubnetId:       "mock-unit-config01-dev",
								SecurityGroupIds: []string{
									"mock-unit-config01-dev",
								},
								SshKeyId:       "mock-unit-config01-dev",
								VmUserAccount:  "cb-user",
								VmUserPassword: "",
								Description:    "description",
							},
						},
					},
				},
			},
			ExpectResStartsWith: `{"id":"mock-unit-config01-dev2"`,
		}
		MethodTest(t, tc)

		TearDownForGrpc()
	})
}
