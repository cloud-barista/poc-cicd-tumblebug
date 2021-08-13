package cliscenario

import (
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestCliFullYaml(t *testing.T) {
	t.Run("command full yaml in/out test for mock driver", func(t *testing.T) {
		SetUpForCli()

		tc := TestCases{
			Name: "create namespace",
			CmdArgs: []string{"namespace", "create", "--config", "../conf/grpc_conf.yaml", "-i", "yaml", "-o", "yaml", "-d", `
"name": "ns-unit-01"
"description": "NameSpace for General Testing"
`,
			},
			ExpectResStartsWith: `id: ns-unit-01
name: ns-unit-01`,
		}
		TumblebugCmdTest(t, tc)

		tc = TestCases{
			Name:    "list namespace",
			CmdArgs: []string{"namespace", "list", "--config", "../conf/grpc_conf.yaml", "-i", "yaml", "-o", "yaml"},
			ExpectResStartsWith: `ns:
- id: ns-unit-01`,
		}
		TumblebugCmdTest(t, tc)

		tc = TestCases{
			Name:    "get namespace",
			CmdArgs: []string{"namespace", "get", "--config", "../conf/grpc_conf.yaml", "-i", "yaml", "-o", "yaml", "--ns", "ns-unit-01"},
			ExpectResStartsWith: `id: ns-unit-01
name: ns-unit-01`,
		}
		TumblebugCmdTest(t, tc)

		tc = TestCases{
			Name: "create vnet",
			CmdArgs: []string{"network", "create", "--config", "../conf/grpc_conf.yaml", "-i", "yaml", "-o", "yaml", "-d", `
"nsId": "ns-unit-01"
"vNet":
  "name": "mock-unit-config01-dev"
  "connectionName": "mock-unit-config01"
  "cidrBlock": "192.168.0.0/16"
  "subnetInfoList":
    - "Name": "mock-unit-config01-dev"
      "IPv4_CIDR": "192.168.1.0/24"
`,
			},
			ExpectResStartsWith: `id: mock-unit-config01-dev
name: mock-unit-config01-dev`,
		}
		TumblebugCmdTest(t, tc)

		tc = TestCases{
			Name:    "list vnet",
			CmdArgs: []string{"network", "list", "--config", "../conf/grpc_conf.yaml", "-i", "yaml", "-o", "yaml", "--ns", "ns-unit-01"},
			ExpectResStartsWith: `vNet:
- id: mock-unit-config01-dev`,
		}
		TumblebugCmdTest(t, tc)

		tc = TestCases{
			Name:    "list vnet id",
			CmdArgs: []string{"network", "list-id", "--config", "../conf/grpc_conf.yaml", "-i", "yaml", "-o", "yaml", "--ns", "ns-unit-01"},
			ExpectResStartsWith: `idList:
- mock-unit-config01-dev`,
		}
		TumblebugCmdTest(t, tc)

		tc = TestCases{
			Name:    "get vnet",
			CmdArgs: []string{"network", "get", "--config", "../conf/grpc_conf.yaml", "-i", "yaml", "-o", "yaml", "--ns", "ns-unit-01", "--id", "mock-unit-config01-dev"},
			ExpectResStartsWith: `id: mock-unit-config01-dev
name: mock-unit-config01-dev`,
		}
		TumblebugCmdTest(t, tc)

		tc = TestCases{
			Name:    "inspect vnet",
			CmdArgs: []string{"util", "inspect-mcir", "--config", "../conf/grpc_conf.yaml", "-i", "yaml", "-o", "yaml", "--cc", "mock-unit-config01", "--type", "vNet"},
			ExpectResStartsWith: `resourcesOnCsp:
- id: mock-unit-config01-dev`,
		}
		TumblebugCmdTest(t, tc)

		tc = TestCases{
			Name: "create security",
			CmdArgs: []string{"securitygroup", "create", "--config", "../conf/grpc_conf.yaml", "-i", "yaml", "-o", "yaml", "-d", `
"nsId": "ns-unit-01"
"securityGroup":
  "name": "mock-unit-config01-dev"
  "connectionName": "mock-unit-config01"
  "vNetId": "mock-unit-config01-dev"
  "description": "test description"
  "firewallRules":
    - "FromPort": "1"
      "ToPort": "65535"
      "IPProtocol": "tcp"
      "Direction": "inbound"
    - "FromPort": "1"
      "ToPort": "65535"
      "IPProtocol": "udp"
      "Direction": "inbound"
    - "FromPort": "-1"
      "ToPort": "-1"
      "IPProtocol": "icmp"
      "Direction": "inbound"
`,
			},
			ExpectResStartsWith: `id: mock-unit-config01-dev
name: mock-unit-config01-dev`,
		}
		TumblebugCmdTest(t, tc)

		tc = TestCases{
			Name:    "list security",
			CmdArgs: []string{"securitygroup", "list", "--config", "../conf/grpc_conf.yaml", "-i", "yaml", "-o", "yaml", "--ns", "ns-unit-01"},
			ExpectResStartsWith: `securityGroup:
- id: mock-unit-config01-dev`,
		}
		TumblebugCmdTest(t, tc)

		tc = TestCases{
			Name:    "list security id",
			CmdArgs: []string{"securitygroup", "list-id", "--config", "../conf/grpc_conf.yaml", "-i", "yaml", "-o", "yaml", "--ns", "ns-unit-01"},
			ExpectResStartsWith: `idList:
- mock-unit-config01-dev`,
		}
		TumblebugCmdTest(t, tc)

		tc = TestCases{
			Name:    "get security",
			CmdArgs: []string{"securitygroup", "get", "--config", "../conf/grpc_conf.yaml", "-i", "yaml", "-o", "yaml", "--ns", "ns-unit-01", "--id", "mock-unit-config01-dev"},
			ExpectResStartsWith: `id: mock-unit-config01-dev
name: mock-unit-config01-dev`,
		}
		TumblebugCmdTest(t, tc)

		tc = TestCases{
			Name:    "inspect security",
			CmdArgs: []string{"util", "inspect-mcir", "--config", "../conf/grpc_conf.yaml", "-i", "yaml", "-o", "yaml", "--cc", "mock-unit-config01", "--type", "securityGroup"},
			ExpectResStartsWith: `resourcesOnCsp:
- id: mock-unit-config01-dev`,
		}
		TumblebugCmdTest(t, tc)

		tc = TestCases{
			Name: "create sshkey",
			CmdArgs: []string{"keypair", "create", "--config", "../conf/grpc_conf.yaml", "-i", "yaml", "-o", "yaml", "-d", `
"nsId": "ns-unit-01"
"sshKey":
  "name": "mock-unit-config01-dev"
  "connectionName": "mock-unit-config01"
  "description": ""
`,
			},
			ExpectResStartsWith: `id: mock-unit-config01-dev
name: mock-unit-config01-dev`,
		}
		TumblebugCmdTest(t, tc)

		tc = TestCases{
			Name:    "list sshkey",
			CmdArgs: []string{"keypair", "list", "--config", "../conf/grpc_conf.yaml", "-i", "yaml", "-o", "yaml", "--ns", "ns-unit-01"},
			ExpectResStartsWith: `sshKey:
- id: mock-unit-config01-dev`,
		}
		TumblebugCmdTest(t, tc)

		tc = TestCases{
			Name:    "list sshkey id",
			CmdArgs: []string{"keypair", "list-id", "--config", "../conf/grpc_conf.yaml", "-i", "yaml", "-o", "yaml", "--ns", "ns-unit-01"},
			ExpectResStartsWith: `idList:
- mock-unit-config01-dev`,
		}
		TumblebugCmdTest(t, tc)

		tc = TestCases{
			Name:    "get sshkey",
			CmdArgs: []string{"keypair", "get", "--config", "../conf/grpc_conf.yaml", "-i", "yaml", "-o", "yaml", "--ns", "ns-unit-01", "--id", "mock-unit-config01-dev"},
			ExpectResStartsWith: `id: mock-unit-config01-dev
name: mock-unit-config01-dev`,
		}
		TumblebugCmdTest(t, tc)

		tc = TestCases{
			Name:    "inspect sshkey",
			CmdArgs: []string{"util", "inspect-mcir", "--config", "../conf/grpc_conf.yaml", "-i", "yaml", "-o", "yaml", "--cc", "mock-unit-config01", "--type", "sshKey"},
			ExpectResStartsWith: `resourcesOnCsp:
- id: mock-unit-config01-dev`,
		}
		TumblebugCmdTest(t, tc)

		tc = TestCases{
			Name:    "list lookup image",
			CmdArgs: []string{"image", "list-csp", "--config", "../conf/grpc_conf.yaml", "-i", "yaml", "-o", "yaml", "--cc", "mock-unit-config01"},
			ExpectResStartsWith: `image:
- IId:`,
		}
		TumblebugCmdTest(t, tc)

		tc = TestCases{
			Name:    "lookup image",
			CmdArgs: []string{"image", "get-csp", "--config", "../conf/grpc_conf.yaml", "-i", "yaml", "-o", "yaml", "--cc", "mock-unit-config01", "--image", "mock-vmimage-01"},
			ExpectResStartsWith: `IId:
  NameId: mock-vmimage-01`,
		}
		TumblebugCmdTest(t, tc)

		tc = TestCases{
			Name: "register image with id",
			CmdArgs: []string{"image", "create-id", "--config", "../conf/grpc_conf.yaml", "-i", "yaml", "-o", "yaml", "-d", `
"nsId": "ns-unit-01"
"image":
  "connectionName": "mock-unit-config01"
  "name": "mock-unit-config01-dev"
  "cspImageId": "mock-vmimage-01"
  "description": "Canonical, Ubuntu, 18.04 LTS, amd64 bionic"
`,
			},
			ExpectResStartsWith: `namespace: ns-unit-01
id: mock-unit-config01-dev`,
		}
		TumblebugCmdTest(t, tc)

		tc = TestCases{
			Name:    "list image",
			CmdArgs: []string{"image", "list", "--config", "../conf/grpc_conf.yaml", "-i", "yaml", "-o", "yaml", "--ns", "ns-unit-01"},
			ExpectResStartsWith: `image:
- namespace: ns-unit-01`,
		}
		TumblebugCmdTest(t, tc)

		tc = TestCases{
			Name:    "list image id",
			CmdArgs: []string{"image", "list-id", "--config", "../conf/grpc_conf.yaml", "-i", "yaml", "-o", "yaml", "--ns", "ns-unit-01"},
			ExpectResStartsWith: `idList:
- mock-unit-config01-dev`,
		}
		TumblebugCmdTest(t, tc)

		tc = TestCases{
			Name:    "get image",
			CmdArgs: []string{"image", "get", "--config", "../conf/grpc_conf.yaml", "-i", "yaml", "-o", "yaml", "--ns", "ns-unit-01", "--id", "mock-unit-config01-dev"},
			ExpectResStartsWith: `namespace: ns-unit-01
id: mock-unit-config01-dev`,
		}
		TumblebugCmdTest(t, tc)

		tc = TestCases{
			Name: "search image",
			CmdArgs: []string{"image", "search", "--config", "../conf/grpc_conf.yaml", "-i", "yaml", "-o", "yaml", "-d", `{
					"nsId":  "ns-unit-01",
					"keywords": [
						"mock"
					]
				}`},
			ExpectResStartsWith: `image:
- namespace: ns-unit-01`,
		}
		TumblebugCmdTest(t, tc)

		tc = TestCases{
			Name:                "fetch images",
			CmdArgs:             []string{"image", "fetch", "--config", "../conf/grpc_conf.yaml", "-i", "yaml", "-o", "yaml", "--cc", "!all", "--ns", "ns-unit-01"},
			ExpectResStartsWith: `message: Fetched 5 images (from 1 connConfigs)`,
		}
		TumblebugCmdTest(t, tc)

		tc = TestCases{
			Name:    "list lookup spec",
			CmdArgs: []string{"spec", "list-csp", "--config", "../conf/grpc_conf.yaml", "-i", "yaml", "-o", "yaml", "--cc", "mock-unit-config01"},
			ExpectResStartsWith: `vmspec:
- Region: default`,
		}
		TumblebugCmdTest(t, tc)

		tc = TestCases{
			Name:    "lookup spec",
			CmdArgs: []string{"spec", "get-csp", "--config", "../conf/grpc_conf.yaml", "-i", "yaml", "-o", "yaml", "--cc", "mock-unit-config01", "--spec", "mock-vmspec-01"},
			ExpectResStartsWith: `Region: default
Name: mock-vmspec-01`,
		}
		TumblebugCmdTest(t, tc)

		tc = TestCases{
			Name: "register spec",
			CmdArgs: []string{"spec", "create-id", "--config", "../conf/grpc_conf.yaml", "-i", "yaml", "-o", "yaml", "-d", `
"nsId": "ns-unit-01"
"spec":
  "connectionName": "mock-unit-config01"
  "name": "mock-unit-config01-dev"
  "cspSpecName": "mock-vmspec-01"
`,
			},
			ExpectResStartsWith: `namespace: ns-unit-01
id: mock-unit-config01-dev`,
		}
		TumblebugCmdTest(t, tc)

		tc = TestCases{
			Name:    "list spec",
			CmdArgs: []string{"spec", "list", "--config", "../conf/grpc_conf.yaml", "-i", "yaml", "-o", "yaml", "--ns", "ns-unit-01"},
			ExpectResStartsWith: `spec:
- namespace: ns-unit-01`,
		}
		TumblebugCmdTest(t, tc)

		tc = TestCases{
			Name:    "list spec id",
			CmdArgs: []string{"spec", "list-id", "--config", "../conf/grpc_conf.yaml", "-i", "yaml", "-o", "yaml", "--ns", "ns-unit-01"},
			ExpectResStartsWith: `idList:
- mock-unit-config01-dev`,
		}
		TumblebugCmdTest(t, tc)

		tc = TestCases{
			Name:    "get spec",
			CmdArgs: []string{"spec", "get", "--config", "../conf/grpc_conf.yaml", "-i", "yaml", "-o", "yaml", "--ns", "ns-unit-01", "--id", "mock-unit-config01-dev"},
			ExpectResStartsWith: `namespace: ns-unit-01
id: mock-unit-config01-dev`,
		}
		TumblebugCmdTest(t, tc)

		tc = TestCases{
			Name: "update spec",
			CmdArgs: []string{"spec", "update", "--config", "../conf/grpc_conf.yaml", "-i", "yaml", "-o", "yaml", "-d", `
"nsId": "ns-unit-01"
"spec":
  "id": "mock-unit-config01-dev"
  "description": "UpdateSpec() test"
`,
			},
			ExpectResStartsWith: `namespace: ns-unit-01
id: mock-unit-config01-dev`,
		}
		TumblebugCmdTest(t, tc)

		tc = TestCases{
			Name: "filter spec",
			CmdArgs: []string{"spec", "filter", "--config", "../conf/grpc_conf.yaml", "-i", "yaml", "-o", "yaml", "-d", `
"nsId": "ns-unit-01"
"spec":
  "num_vCPU": 4
  "mem_GiB": 32
`,
			},
			ExpectResStartsWith: `spec:
- namespace: ns-unit-01`,
		}
		TumblebugCmdTest(t, tc)

		tc = TestCases{
			Name: "range filter spec",
			CmdArgs: []string{"spec", "filter-by-range", "--config", "../conf/grpc_conf.yaml", "-i", "yaml", "-o", "yaml", "-d", `
"nsId": "ns-unit-01"
"filter":
  "mem_GiB":
    "min": 4
`,
			},
			ExpectResStartsWith: `spec:
- namespace: ns-unit-01`,
		}
		TumblebugCmdTest(t, tc)

		tc = TestCases{
			Name:                "fetch specs",
			CmdArgs:             []string{"spec", "fetch", "--config", "../conf/grpc_conf.yaml", "-i", "yaml", "-o", "yaml", "--cc", "!all", "--ns", "ns-unit-01"},
			ExpectResStartsWith: `message: Fetched 4 specs (from 1 connConfigs)`,
		}
		TumblebugCmdTest(t, tc)

		tc = TestCases{
			Name: "create mcis",
			CmdArgs: []string{"mcis", "create", "--config", "../conf/grpc_conf.yaml", "-i", "yaml", "-o", "yaml", "-d", `
"nsId": "ns-unit-01"
"mcis":
  "name": "mock-unit-config01-dev"
  "description": "Tumblebug Demo"
  "installMonAgent": "no"
  "vm":
    - "name": "mock-unit-config01-dev-01"
      "imageId": "mock-unit-config01-dev"
      "vmUserAccount": "cb-user"
      "connectionName": "mock-unit-config01"
      "sshKeyId": "mock-unit-config01-dev"
      "specId": "mock-unit-config01-dev"
      "securityGroupIds":
        - "mock-unit-config01-dev"
      "vNetId": "mock-unit-config01-dev"
      "subnetId": "mock-unit-config01-dev"
      "description": "description"
      "vmUserPassword": ""
`,
			},
			ExpectResStartsWith: `id: mock-unit-config01-dev
name: mock-unit-config01-dev`,
		}
		TumblebugCmdTest(t, tc)

		tc = TestCases{
			Name: "add vm to mcis",
			CmdArgs: []string{"mcis", "add-vm", "--config", "../conf/grpc_conf.yaml", "-i", "yaml", "-o", "yaml", "-d", `
"nsId": "ns-unit-01"
"mcisId": "mock-unit-config01-dev"
"mcisvm":
  "name": "mock-unit-config01-dev"
  "imageId": "mock-unit-config01-dev"
  "vmUserAccount": "cb-user"
  "connectionName": "mock-unit-config01"
  "sshKeyId": "mock-unit-config01-dev"
  "specId": "mock-unit-config01-dev"
  "securityGroupIds":
    - "mock-unit-config01-dev"
  "vNetId": "mock-unit-config01-dev"
  "subnetId": "mock-unit-config01-dev"
  "description": "description"
  "vmUserPassword": ""
`,
			},
			ExpectResStartsWith: `id: mock-unit-config01-dev
name: mock-unit-config01-dev`,
		}
		TumblebugCmdTest(t, tc)

		tc = TestCases{
			Name: "add vmgroup to mcis",
			CmdArgs: []string{"mcis", "group-vm", "--config", "../conf/grpc_conf.yaml", "-i", "yaml", "-o", "yaml", "-d", `
"nsId": "ns-unit-01"
"mcisId": "mock-unit-config01-dev"
"groupvm":
  "vmGroupSize": "3"
  "name": "mock-unit-config01-dev"
  "imageId": "mock-unit-config01-dev"
  "vmUserAccount": "cb-user"
  "connectionName": "mock-unit-config01"
  "sshKeyId": "mock-unit-config01-dev"
  "specId": "mock-unit-config01-dev"
  "securityGroupIds":
    - "mock-unit-config01-dev"
  "vNetId": "mock-unit-config01-dev"
  "subnetId": "mock-unit-config01-dev"
  "description": "description"
  "vmUserPassword": ""
`,
			},
			ExpectResStartsWith: `id: mock-unit-config01-dev
name: mock-unit-config01-dev`,
		}
		TumblebugCmdTest(t, tc)

		tc = TestCases{
			Name:    "list mcis",
			CmdArgs: []string{"mcis", "list", "--config", "../conf/grpc_conf.yaml", "-i", "yaml", "-o", "yaml", "--ns", "ns-unit-01"},
			ExpectResStartsWith: `mcis:
- id: mock-unit-config01-dev`,
		}
		TumblebugCmdTest(t, tc)

		tc = TestCases{
			Name:    "list mcis id",
			CmdArgs: []string{"mcis", "list-id", "--config", "../conf/grpc_conf.yaml", "-i", "yaml", "-o", "yaml", "--ns", "ns-unit-01"},
			ExpectResStartsWith: `idList:
- mock-unit-config01-dev`,
		}
		TumblebugCmdTest(t, tc)

		tc = TestCases{
			Name:    "list mcis status",
			CmdArgs: []string{"mcis", "status-list", "--config", "../conf/grpc_conf.yaml", "-i", "yaml", "-o", "yaml", "--ns", "ns-unit-01"},
			ExpectResStartsWith: `mcis:
- id: mock-unit-config01-dev`,
		}
		TumblebugCmdTest(t, tc)

		tc = TestCases{
			Name:    "get mcis",
			CmdArgs: []string{"mcis", "get", "--config", "../conf/grpc_conf.yaml", "-i", "yaml", "-o", "yaml", "--ns", "ns-unit-01", "--mcis", "mock-unit-config01-dev"},
			ExpectResStartsWith: `id: mock-unit-config01-dev
name: mock-unit-config01-dev`,
		}
		TumblebugCmdTest(t, tc)

		tc = TestCases{
			Name:    "get mcis status",
			CmdArgs: []string{"mcis", "status", "--config", "../conf/grpc_conf.yaml", "-i", "yaml", "-o", "yaml", "--ns", "ns-unit-01", "--mcis", "mock-unit-config01-dev"},
			ExpectResStartsWith: `status:
  id: mock-unit-config01-dev`,
		}
		TumblebugCmdTest(t, tc)

		tc = TestCases{
			Name:                "suspend mcis",
			CmdArgs:             []string{"mcis", "suspend", "--config", "../conf/grpc_conf.yaml", "-i", "yaml", "-o", "yaml", "--ns", "ns-unit-01", "--mcis", "mock-unit-config01-dev"},
			ExpectResStartsWith: `message: Suspending the MCIS`,
		}
		TumblebugCmdTest(t, tc)

		tc = TestCases{
			Name:                "resume mcis",
			CmdArgs:             []string{"mcis", "resume", "--config", "../conf/grpc_conf.yaml", "-i", "yaml", "-o", "yaml", "--ns", "ns-unit-01", "--mcis", "mock-unit-config01-dev"},
			ExpectResStartsWith: `message: Resuming the MCIS`,
		}
		TumblebugCmdTest(t, tc)

		tc = TestCases{
			Name:                "reboot mcis",
			CmdArgs:             []string{"mcis", "reboot", "--config", "../conf/grpc_conf.yaml", "-i", "yaml", "-o", "yaml", "--ns", "ns-unit-01", "--mcis", "mock-unit-config01-dev"},
			ExpectResStartsWith: `message: Rebooting the MCIS`,
		}
		TumblebugCmdTest(t, tc)

		tc = TestCases{
			Name:    "list vm id",
			CmdArgs: []string{"mcis", "list-vm-id", "--config", "../conf/grpc_conf.yaml", "-i", "yaml", "-o", "yaml", "--ns", "ns-unit-01", "--mcis", "mock-unit-config01-dev"},
			ExpectResStartsWith: `idList:
- mock-unit-config01-dev`,
		}
		TumblebugCmdTest(t, tc)

		tc = TestCases{
			Name:    "inspect vm",
			CmdArgs: []string{"util", "inspect-vm", "--config", "../conf/grpc_conf.yaml", "-i", "yaml", "-o", "yaml", "--cc", "mock-unit-config01"},
			ExpectResStartsWith: `resourcesOnCsp:
- id: ns-unit-01-mock-unit-config01-dev-mock-unit-config01-dev`,
		}
		TumblebugCmdTest(t, tc)

		tc = TestCases{
			Name: "create mcis policy",
			CmdArgs: []string{"mcis", "create-policy", "--config", "../conf/grpc_conf.yaml", "-i", "yaml", "-o", "yaml", "-d", `
"nsId": "ns-unit-01"
"mcisId": "mock-unit-config01-dev"
"ReqInfo":
  "description": "Tumblebug Auto Control Demo"
  "policy":
    - "autoCondition":
        "metric": "cpu"
        "operator": ">="
        "operand": "80"
        "evaluationPeriod": "10"
      "autoAction":
        "actionType": "ScaleOut"
        "placementAlgo": "random"
        "vm":
          "name": "AutoGen"
        "postCommand":
          "command": "wget https://raw.githubusercontent.com/cloud-barista/cb-tumblebug/master/assets/scripts/setweb.sh -O ~/setweb.sh; chmod +x ~/setweb.sh; sudo ~/setweb.sh; wget https://raw.githubusercontent.com/cloud-barista/cb-tumblebug/master/assets/scripts/runLoadMaker.sh -O ~/runLoadMaker.sh; chmod +x ~/runLoadMaker.sh; sudo ~/runLoadMaker.sh"
    - "autoCondition":
        "metric": "cpu"
        "operator": "<="
        "operand": "60"
        "evaluationPeriod": "10"
      "autoAction":
        "actionType": "ScaleIn"
`,
			},
			ExpectResStartsWith: `Name: mock-unit-config01-dev
Id: mock-unit-config01-dev`,
		}
		TumblebugCmdTest(t, tc)

		tc = TestCases{
			Name:    "list mcis policy",
			CmdArgs: []string{"mcis", "list-policy", "--config", "../conf/grpc_conf.yaml", "-i", "yaml", "-o", "yaml", "--ns", "ns-unit-01"},
			ExpectResStartsWith: `mcisPolicy:
- Name: mock-unit-config01-dev`,
		}
		TumblebugCmdTest(t, tc)

		tc = TestCases{
			Name:    "get mcis policy",
			CmdArgs: []string{"mcis", "get-policy", "--config", "../conf/grpc_conf.yaml", "-i", "yaml", "-o", "yaml", "--ns", "ns-unit-01", "--mcis", "mock-unit-config01-dev"},
			ExpectResStartsWith: `Name: mock-unit-config01-dev
Id: mock-unit-config01-dev`,
		}
		TumblebugCmdTest(t, tc)

		tc = TestCases{
			Name: "update config",
			CmdArgs: []string{"config", "create", "--config", "../conf/grpc_conf.yaml", "-i", "yaml", "-o", "yaml", "-d", `
"name": "key01"
"value": "value01"
`,
			},
			ExpectResStartsWith: `id: key01
name: key01`,
		}
		TumblebugCmdTest(t, tc)

		tc = TestCases{
			Name:    "list config",
			CmdArgs: []string{"config", "list", "--config", "../conf/grpc_conf.yaml", "-i", "yaml", "-o", "yaml"},
			ExpectResStartsWith: `config:
- id: key01`,
		}
		TumblebugCmdTest(t, tc)

		tc = TestCases{
			Name:    "get config",
			CmdArgs: []string{"config", "get", "--config", "../conf/grpc_conf.yaml", "-i", "yaml", "-o", "yaml", "--id", "key01"},
			ExpectResStartsWith: `id: key01
name: key01`,
		}
		TumblebugCmdTest(t, tc)

		tc = TestCases{
			Name: "install agent",
			CmdArgs: []string{"mcis", "install-mon", "--config", "../conf/grpc_conf.yaml", "-i", "yaml", "-o", "yaml", "-d", `
"nsId": "ns-unit-01"
"mcisId": "mock-unit-config01-dev"
"cmd":
  "command": "echo -n [CMD] Works! [Public IP: ; curl https://api.ipify.org ; echo -n ], [HostName: ; hostname ; echo -n ]"
`,
			},
			ExpectResStartsWith: `result_array:
- mcisId: mock-unit-config01-dev`,
		}
		TumblebugCmdTest(t, tc)

		tc = TestCases{
			Name:    "get monitoring data",
			CmdArgs: []string{"mcis", "get-mon", "--config", "../conf/grpc_conf.yaml", "-i", "yaml", "-o", "yaml", "--ns", "ns-unit-01", "--mcis", "mock-unit-config01-dev", "--metric", "cpu"},
			ExpectResStartsWith: `nsId: ns-unit-01
mcisId: mock-unit-config01-dev`,
		}
		TumblebugCmdTest(t, tc)

		tc = TestCases{
			Name:    "list connection config",
			CmdArgs: []string{"util", "list-cc", "--config", "../conf/grpc_conf.yaml", "-i", "yaml", "-o", "yaml"},
			ExpectResStartsWith: `connectionconfig:
- ConfigName: mock-unit-config01`,
		}
		TumblebugCmdTest(t, tc)

		tc = TestCases{
			Name:    "get connection config",
			CmdArgs: []string{"util", "get-cc", "--config", "../conf/grpc_conf.yaml", "-i", "yaml", "-o", "yaml", "--cc", "mock-unit-config01"},
			ExpectResStartsWith: `ConfigName: mock-unit-config01
ProviderName: MOCK`,
		}
		TumblebugCmdTest(t, tc)

		tc = TestCases{
			Name:    "list region",
			CmdArgs: []string{"util", "list-region", "--config", "../conf/grpc_conf.yaml", "-i", "yaml", "-o", "yaml"},
			ExpectResStartsWith: `region:
- RegionName: mock-unit-region01`,
		}
		TumblebugCmdTest(t, tc)

		tc = TestCases{
			Name:    "get region",
			CmdArgs: []string{"util", "get-region", "--config", "../conf/grpc_conf.yaml", "-i", "yaml", "-o", "yaml", "--region", "mock-unit-region01"},
			ExpectResStartsWith: `RegionName: mock-unit-region01
ProviderName: MOCK`,
		}
		TumblebugCmdTest(t, tc)

		tc = TestCases{
			Name: "command mcis",
			CmdArgs: []string{"mcis", "command", "--config", "../conf/grpc_conf.yaml", "-i", "yaml", "-o", "yaml", "-d", `
"nsId": "ns-unit-01"
"mcisId": "mock-unit-config01-dev"
"cmd":
  "command": "echo -n [CMD] Works! [Public IP: ; curl https://api.ipify.org ; echo -n ], [HostName: ; hostname ; echo -n ]"
`,
			},
			ExpectResStartsWith: `result_array:
- mcisId: mock-unit-config01-dev`,
		}
		TumblebugCmdTest(t, tc)

		tc = TestCases{
			Name: "command mcis vm",
			CmdArgs: []string{"mcis", "command-vm", "--config", "../conf/grpc_conf.yaml", "-i", "yaml", "-o", "yaml", "-d", `
"nsId": "ns-unit-01"
"mcisId": "mock-unit-config01-dev"
"vmId": "mock-unit-config01-dev"
"cmd":
  "command": "echo -n [CMD] Works! [Public IP: ; curl https://api.ipify.org ; echo -n ], [HostName: ; hostname ; echo -n ]"
`,
			},
			ExpectResStartsWith: `Result: 'echo -n [CMD] Works`,
		}
		TumblebugCmdTest(t, tc)

		tc = TestCases{
			Name:    "list object",
			CmdArgs: []string{"util", "list-obj", "--config", "../conf/grpc_conf.yaml", "-i", "yaml", "-o", "yaml"},
			ExpectResStartsWith: `object:
- /config/key01`,
		}
		TumblebugCmdTest(t, tc)

		tc = TestCases{
			Name:    "list sub object",
			CmdArgs: []string{"util", "list-obj", "--config", "../conf/grpc_conf.yaml", "-i", "yaml", "-o", "yaml", "--key", "/config"},
			ExpectResStartsWith: `object:
- /config/key01`,
		}
		TumblebugCmdTest(t, tc)

		tc = TestCases{
			Name:    "get object",
			CmdArgs: []string{"util", "get-obj", "--config", "../conf/grpc_conf.yaml", "-i", "yaml", "-o", "yaml", "--key", "/config/key01"},
			ExpectResStartsWith: `id: key01
name: key01`,
		}
		TumblebugCmdTest(t, tc)

		//
		// Delete Resources
		//

		tc = TestCases{
			Name:                "delete object",
			CmdArgs:             []string{"util", "delete-obj", "--config", "../conf/grpc_conf.yaml", "-i", "yaml", "-o", "yaml", "--key", "/config/key01"},
			ExpectResStartsWith: `message: The object has been deleted`,
		}
		TumblebugCmdTest(t, tc)

		tc = TestCases{
			Name:                "delete all sub object",
			CmdArgs:             []string{"util", "delete-all-obj", "--config", "../conf/grpc_conf.yaml", "-i", "yaml", "-o", "yaml", "--key", "/config"},
			ExpectResStartsWith: `message: Objects have been deleted`,
		}
		TumblebugCmdTest(t, tc)

		tc = TestCases{
			Name:                "init config",
			CmdArgs:             []string{"config", "init", "--config", "../conf/grpc_conf.yaml", "-i", "yaml", "-o", "yaml", "--id", "key01"},
			ExpectResStartsWith: `message: The config key01 has been initialized.`,
		}
		TumblebugCmdTest(t, tc)

		tc = TestCases{
			Name:                "init all config",
			CmdArgs:             []string{"config", "init-all", "--config", "../conf/grpc_conf.yaml", "-i", "yaml", "-o", "yaml"},
			ExpectResStartsWith: `message: All configs have been initialized.`,
		}
		TumblebugCmdTest(t, tc)

		tc = TestCases{
			Name:                "delete mcis policy",
			CmdArgs:             []string{"mcis", "delete-policy", "--config", "../conf/grpc_conf.yaml", "-i", "yaml", "-o", "yaml", "--ns", "ns-unit-01", "--mcis", "mock-unit-config01-dev"},
			ExpectResStartsWith: `message: Deleting the MCIS Policy info`,
		}
		TumblebugCmdTest(t, tc)

		tc = TestCases{
			Name:                "delete all mcis policy",
			CmdArgs:             []string{"mcis", "delete-all-policy", "--config", "../conf/grpc_conf.yaml", "-i", "yaml", "-o", "yaml", "--ns", "ns-unit-01"},
			ExpectResStartsWith: `message: No MCIS Policy to delete`,
		}
		TumblebugCmdTest(t, tc)

		tc = TestCases{
			Name:                "terminate mcis",
			CmdArgs:             []string{"mcis", "terminate", "--config", "../conf/grpc_conf.yaml", "-i", "yaml", "-o", "yaml", "--ns", "ns-unit-01", "--mcis", "mock-unit-config01-dev"},
			ExpectResStartsWith: `message: Terminating the MCIS`,
		}
		TumblebugCmdTest(t, tc)

		tc = TestCases{
			Name:                "delete mcis",
			CmdArgs:             []string{"mcis", "delete", "--config", "../conf/grpc_conf.yaml", "-i", "yaml", "-o", "yaml", "--ns", "ns-unit-01", "--mcis", "mock-unit-config01-dev"},
			ExpectResStartsWith: `message: Deleting the MCIS mock-unit-config01-dev`,
		}
		TumblebugCmdTest(t, tc)

		tc = TestCases{
			Name:                "delete all mcis",
			CmdArgs:             []string{"mcis", "delete-all", "--config", "../conf/grpc_conf.yaml", "-i", "yaml", "-o", "yaml", "--ns", "ns-unit-01"},
			ExpectResStartsWith: `message: No MCIS to delete`,
		}
		TumblebugCmdTest(t, tc)

		tc = TestCases{
			Name:                "delete spec",
			CmdArgs:             []string{"spec", "delete", "--config", "../conf/grpc_conf.yaml", "-i", "yaml", "-o", "yaml", "--ns", "ns-unit-01", "--id", "mock-unit-config01-dev", "--force", "false"},
			ExpectResStartsWith: `message: The spec mock-unit-config01-dev has been deleted`,
		}
		TumblebugCmdTest(t, tc)

		tc = TestCases{
			Name:                "delete all spec",
			CmdArgs:             []string{"spec", "delete-all", "--config", "../conf/grpc_conf.yaml", "-i", "yaml", "-o", "yaml", "--ns", "ns-unit-01", "--force", "false"},
			ExpectResStartsWith: `message: All specs has been deleted`,
		}
		TumblebugCmdTest(t, tc)

		tc = TestCases{
			Name:                "delete image",
			CmdArgs:             []string{"image", "delete", "--config", "../conf/grpc_conf.yaml", "-i", "yaml", "-o", "yaml", "--ns", "ns-unit-01", "--id", "mock-unit-config01-dev", "--force", "false"},
			ExpectResStartsWith: `message: The image mock-unit-config01-dev has been deleted`,
		}
		TumblebugCmdTest(t, tc)

		tc = TestCases{
			Name:                "delete all image",
			CmdArgs:             []string{"image", "delete-all", "--config", "../conf/grpc_conf.yaml", "-i", "yaml", "-o", "yaml", "--ns", "ns-unit-01", "--force", "false"},
			ExpectResStartsWith: `message: All images has been deleted`,
		}
		TumblebugCmdTest(t, tc)

		tc = TestCases{
			Name:                "delete sshkey",
			CmdArgs:             []string{"keypair", "delete", "--config", "../conf/grpc_conf.yaml", "-i", "yaml", "-o", "yaml", "--ns", "ns-unit-01", "--id", "mock-unit-config01-dev", "--force", "false"},
			ExpectResStartsWith: `message: The sshKey mock-unit-config01-dev has been deleted`,
		}
		TumblebugCmdTest(t, tc)

		tc = TestCases{
			Name:                "delete all sshkey",
			CmdArgs:             []string{"keypair", "delete-all", "--config", "../conf/grpc_conf.yaml", "-i", "yaml", "-o", "yaml", "--ns", "ns-unit-01", "--force", "false"},
			ExpectResStartsWith: `message: All sshKeys has been deleted`,
		}
		TumblebugCmdTest(t, tc)

		tc = TestCases{
			Name:                "delete security",
			CmdArgs:             []string{"securitygroup", "delete", "--config", "../conf/grpc_conf.yaml", "-i", "yaml", "-o", "yaml", "--ns", "ns-unit-01", "--id", "mock-unit-config01-dev", "--force", "false"},
			ExpectResStartsWith: `message: The securityGroup mock-unit-config01-dev has been deleted`,
		}
		TumblebugCmdTest(t, tc)

		tc = TestCases{
			Name:                "delete all security",
			CmdArgs:             []string{"securitygroup", "delete-all", "--config", "../conf/grpc_conf.yaml", "-i", "yaml", "-o", "yaml", "--ns", "ns-unit-01", "--force", "false"},
			ExpectResStartsWith: `message: All securityGroups has been deleted`,
		}
		TumblebugCmdTest(t, tc)

		tc = TestCases{
			Name:                "delete vnet",
			CmdArgs:             []string{"network", "delete", "--config", "../conf/grpc_conf.yaml", "-i", "yaml", "-o", "yaml", "--ns", "ns-unit-01", "--id", "mock-unit-config01-dev", "--force", "false"},
			ExpectResStartsWith: `message: The vNet mock-unit-config01-dev has been deleted`,
		}
		TumblebugCmdTest(t, tc)

		tc = TestCases{
			Name:                "delete all vnet",
			CmdArgs:             []string{"network", "delete-all", "--config", "../conf/grpc_conf.yaml", "-i", "yaml", "-o", "yaml", "--ns", "ns-unit-01", "--force", "false"},
			ExpectResStartsWith: `message: All vNets has been deleted`,
		}
		TumblebugCmdTest(t, tc)

		tc = TestCases{
			Name:                "delete namespace",
			CmdArgs:             []string{"namespace", "delete", "--config", "../conf/grpc_conf.yaml", "-i", "yaml", "-o", "yaml", "--ns", "ns-unit-01"},
			ExpectResStartsWith: `message: The ns has been deleted`,
		}
		TumblebugCmdTest(t, tc)

		tc = TestCases{
			Name:                "delete all namespace",
			CmdArgs:             []string{"namespace", "delete-all", "--config", "../conf/grpc_conf.yaml", "-i", "yaml", "-o", "yaml"},
			ExpectResStartsWith: `message: All namespaces has been deleted`,
		}
		TumblebugCmdTest(t, tc)

		TearDownForCli()
	})
}
