package goscenario

import (
	"os"
	"testing"
)

func TestGoMiscDocArg(t *testing.T) {

	t.Run("go api misc test for mock driver by doccument args style (SPIDER_CALL_METHOD=REST)", func(t *testing.T) {
		SetUpForGrpc()

		holdEnv := os.Getenv("SPIDER_CALL_METHOD")
		os.Setenv("SPIDER_CALL_METHOD", "REST")

		GoMiscDocArg(t)

		os.Setenv("SPIDER_CALL_METHOD", holdEnv)

		TearDownForGrpc()
	})

	t.Run("go api misc test for mock driver by doccument args style (SPIDER_CALL_METHOD=GRPC)", func(t *testing.T) {
		SetUpForGrpc()

		holdEnv := os.Getenv("SPIDER_CALL_METHOD")
		os.Setenv("SPIDER_CALL_METHOD", "GRPC")

		GoMiscDocArg(t)

		os.Setenv("SPIDER_CALL_METHOD", holdEnv)

		TearDownForGrpc()
	})

}

func GoMiscDocArg(t *testing.T) {
	tc := TestCases{
		Name:     "create namespace",
		Instance: NsApi,
		Method:   "CreateNS",
		Args: []interface{}{
			`{"name":"ns-unit-01","description":"NameSpace for General Testing"}`,
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
		Method:   "CreateVNet",
		Args: []interface{}{
			`{
				"nsId":  "ns-unit-01",
				"vNet": {
					"name": "mock-unit-config01-dev",
					"connectionName": "mock-unit-config01",
					"cidrBlock": "192.168.0.0/16",
					"subnetInfoList": [ {
						"Name": "mock-unit-config01-dev",
						"IPv4_CIDR": "192.168.1.0/24"
					} ]
				}					
			}`,
		},
		ExpectResStartsWith: `{"id":"mock-unit-config01-dev"`,
	}
	MethodTest(t, tc)

	tc = TestCases{
		Name:     "check resource",
		Instance: McirApi,
		Method:   "CheckResource",
		Args: []interface{}{
			`{ "nsId": "ns-unit-01", "resourceId":"mock-unit-config01-dev", "resourceType": "vNet"}`,
		},
		ExpectResStartsWith: `{"exists":true}`,
	}
	MethodTest(t, tc)

	tc = TestCases{
		Name:     "create security",
		Instance: McirApi,
		Method:   "CreateSecurityGroup",
		Args: []interface{}{
			`{
				"nsId":  "ns-unit-01",
				"securityGroup": {
					"name": "mock-unit-config01-dev",
					"connectionName": "mock-unit-config01",
					"vNetId": "mock-unit-config01-dev",
					"description": "test description",
						"firewallRules": [
							{
								"FromPort": "1",
								"ToPort": "65535",
								"IPProtocol": "tcp",
								"Direction": "inbound"
							},
							{
								"FromPort": "1",
								"ToPort": "65535",
								"IPProtocol": "udp",
								"Direction": "inbound"
							},
							{
								"FromPort": "-1",
								"ToPort": "-1",
								"IPProtocol": "icmp",
								"Direction": "inbound"
							}
						]	
				}					
			}`,
		},
		ExpectResStartsWith: `{"id":"mock-unit-config01-dev"`,
	}
	MethodTest(t, tc)

	tc = TestCases{
		Name:     "create sshkey",
		Instance: McirApi,
		Method:   "CreateSshKey",
		Args: []interface{}{
			`{
				"nsId":  "ns-unit-01",
				"sshKey": {
					"name": "mock-unit-config01-dev",
					"connectionName": "mock-unit-config01",
					"description": ""	
				}					
			}`,
		},
		ExpectResStartsWith: `{"id":"mock-unit-config01-dev"`,
	}
	MethodTest(t, tc)

	tc = TestCases{
		Name:     "register image with id",
		Instance: McirApi,
		Method:   "CreateImageWithID",
		Args: []interface{}{
			`{
				"nsId":  "ns-unit-01",
				"image": {
					"connectionName": "mock-unit-config01",
					"name": "mock-unit-config01-dev",
					"cspImageId": "mock-vmimage-01",
					"description": "Canonical, Ubuntu, 18.04 LTS, amd64 bionic"
				}					
			}`,
		},
		ExpectResStartsWith: `{"namespace":"ns-unit-01","id":"mock-unit-config01-dev"`,
	}
	MethodTest(t, tc)

	tc = TestCases{
		Name:     "connection fetch images",
		Instance: McirApi,
		Method:   "FetchImage",
		Args: []interface{}{
			`{ "ConnectionName": "mock-unit-config01", "nsId": "ns-unit-01" }`,
		},
		ExpectResStartsWith: `{"message":"Fetched 5 images (from 1 connConfigs)"}`,
	}
	MethodTest(t, tc)

	tc = TestCases{
		Name:     "register spec",
		Instance: McirApi,
		Method:   "CreateSpecWithSpecName",
		Args: []interface{}{
			`{
				"nsId":  "ns-unit-01",
				"spec": {
					"connectionName": "mock-unit-config01",
					"name": "mock-unit-config01-dev",
					"cspSpecName": "mock-vmspec-01"		
				}					
			}`,
		},
		ExpectResStartsWith: `{"namespace":"ns-unit-01","id":"mock-unit-config01-dev"`,
	}
	MethodTest(t, tc)

	tc = TestCases{
		Name:     "connection fetch specs",
		Instance: McirApi,
		Method:   "FetchSpec",
		Args: []interface{}{
			`{ "ConnectionName": "mock-unit-config01", "nsId": "ns-unit-01" }`,
		},
		ExpectResStartsWith: `{"message":"Fetched 4 specs (from 1 connConfigs)"}`,
	}
	MethodTest(t, tc)

	tc = TestCases{
		Name:     "create mcis",
		Instance: McisApi,
		Method:   "CreateMcis",
		Args: []interface{}{
			`{
				"nsId":  "ns-unit-01",
				"mcis": {
					"name": "mock-unit-config01-dev",
					"description": "Tumblebug Demo",
					"installMonAgent": "no",
					"vm": [ {
						"name": "mock-unit-config01-dev-01",
						"imageId": "mock-unit-config01-dev",
						"vmUserAccount": "cb-user",
						"connectionName": "mock-unit-config01",
						"sshKeyId": "mock-unit-config01-dev",
						"specId": "mock-unit-config01-dev",
						"securityGroupIds": [
							"mock-unit-config01-dev"
						],
						"vNetId": "mock-unit-config01-dev",
						"subnetId": "mock-unit-config01-dev",
						"description": "description",
						"vmUserPassword": ""
					}
					]
				}					
			}`,
		},
		ExpectResStartsWith: `{"id":"mock-unit-config01-dev"`,
	}
	MethodTest(t, tc)

	tc = TestCases{
		Name:     "get mcis vm",
		Instance: McisApi,
		Method:   "GetMcisVMInfo",
		Args: []interface{}{
			`{ "nsId": "ns-unit-01", "mcisId":"mock-unit-config01-dev", "vmId":"mock-unit-config01-dev-01"}`,
		},
		ExpectResStartsWith: `{"id":"mock-unit-config01-dev-01"`,
	}
	MethodTest(t, tc)

	tc = TestCases{
		Name:     "get mcis vm status",
		Instance: McisApi,
		Method:   "GetMcisVMStatus",
		Args: []interface{}{
			`{ "nsId": "ns-unit-01", "mcisId":"mock-unit-config01-dev", "vmId":"mock-unit-config01-dev-01"}`,
		},
		ExpectResStartsWith: `{"id":"mock-unit-config01-dev-01"`,
	}
	MethodTest(t, tc)

	tc = TestCases{
		Name:     "suspend mcis vm",
		Instance: McisApi,
		Method:   "ControlMcisVM",
		Args: []interface{}{
			`{ "nsId": "ns-unit-01", "mcisId":"mock-unit-config01-dev", "vmId":"mock-unit-config01-dev-01", "action":"suspend"}`,
		},
		ExpectResStartsWith: `{"message":"Suspending the VM"}`,
	}
	MethodTest(t, tc)

	tc = TestCases{
		Name:     "resume mcis vm",
		Instance: McisApi,
		Method:   "ControlMcisVM",
		Args: []interface{}{
			`{ "nsId": "ns-unit-01", "mcisId":"mock-unit-config01-dev", "vmId":"mock-unit-config01-dev-01", "action":"resume"}`,
		},
		ExpectResStartsWith: `{"message":"Resuming the VM"}`,
	}
	MethodTest(t, tc)

	tc = TestCases{
		Name:     "reboot mcis vm",
		Instance: McisApi,
		Method:   "ControlMcisVM",
		Args: []interface{}{
			`{ "nsId": "ns-unit-01", "mcisId":"mock-unit-config01-dev", "vmId":"mock-unit-config01-dev-01", "action":"reboot"}`,
		},
		ExpectResStartsWith: `{"message":"Rebooting the VM"}`,
	}
	MethodTest(t, tc)

	tc = TestCases{
		Name:     "all benchmark",
		Instance: McisApi,
		Method:   "GetAllBenchmark",
		Args: []interface{}{
			`{ "nsId": "ns-unit-01", "mcisId":"mock-unit-config01-dev", "bm": { "host": "localhost" } }`,
		},
		ExpectResStartsWith: `{"resultarray":[{"result":"1.0","unit":"unit","desc":"mrtt complete"`,
	}
	MethodTest(t, tc)

	tc = TestCases{
		Name:     "benchmark",
		Instance: McisApi,
		Method:   "GetBenchmark",
		Args: []interface{}{
			`{ "nsId": "ns-unit-01", "mcisId":"mock-unit-config01-dev", "action":"cpus", "bm": { "host": "localhost" } }`,
		},
		ExpectResStartsWith: `{"resultarray":[{"result":"1.0","unit":"unit","desc":"cpus complete"`,
	}
	MethodTest(t, tc)

	tc = TestCases{
		Name:     "check mcis",
		Instance: McisApi,
		Method:   "CheckMcis",
		Args: []interface{}{
			`{ "nsId": "ns-unit-01", "mcisId":"mock-unit-config01-dev" }`,
		},
		ExpectResStartsWith: `{"exists":true}`,
	}
	MethodTest(t, tc)

	tc = TestCases{
		Name:     "check mcis vm",
		Instance: McisApi,
		Method:   "CheckVm",
		Args: []interface{}{
			`{ "nsId": "ns-unit-01", "mcisId":"mock-unit-config01-dev", "vmId":"mock-unit-config01-dev-01" }`,
		},
		ExpectResStartsWith: `{"exists":true}`,
	}
	MethodTest(t, tc)

	tc = TestCases{
		Name:     "terminate mcis vm",
		Instance: McisApi,
		Method:   "ControlMcisVM",
		Args: []interface{}{
			`{ "nsId": "ns-unit-01", "mcisId":"mock-unit-config01-dev", "vmId":"mock-unit-config01-dev-01", "action":"terminate"}`,
		},
		ExpectResStartsWith: `{"message":"Terminating the VM"}`,
	}
	MethodTest(t, tc)

	tc = TestCases{
		Name:     "create mcis(installMonAgent:yes,vmGroupSize>0)",
		Instance: McisApi,
		Method:   "CreateMcis",
		Args: []interface{}{
			`{
				"nsId":  "ns-unit-01",
				"mcis": {
					"name": "mock-unit-config01-dev2",
					"description": "Tumblebug Demo",
					"installMonAgent": "yes",
					"vm": [ {
						"vmGroupSize": "3",
						"name": "mock-unit-config01-dev2-01",
						"imageId": "mock-unit-config01-dev",
						"vmUserAccount": "cb-user",
						"connectionName": "mock-unit-config01",
						"sshKeyId": "mock-unit-config01-dev",
						"specId": "mock-unit-config01-dev",
						"securityGroupIds": [
							"mock-unit-config01-dev"
						],
						"vNetId": "mock-unit-config01-dev",
						"subnetId": "mock-unit-config01-dev",
						"description": "description",
						"vmUserPassword": ""
					}
					]
				}					
			}`,
		},
		ExpectResStartsWith: `{"id":"mock-unit-config01-dev2"`,
	}
	MethodTest(t, tc)
}
