package goscenario

import (
	"os"
	"testing"
)

func TestGoFullDocArg(t *testing.T) {

	t.Run("go api full test for mock driver by doccument args style (SPIDER_CALL_METHOD=REST)", func(t *testing.T) {
		SetUpForGrpc()

		holdEnv := os.Getenv("SPIDER_CALL_METHOD")
		os.Setenv("SPIDER_CALL_METHOD", "REST")

		GoFullDocArg(t)

		os.Setenv("SPIDER_CALL_METHOD", holdEnv)

		TearDownForGrpc()
	})

	t.Run("go api full test for mock driver by doccument args style (SPIDER_CALL_METHOD=GRPC)", func(t *testing.T) {
		SetUpForGrpc()

		holdEnv := os.Getenv("SPIDER_CALL_METHOD")
		os.Setenv("SPIDER_CALL_METHOD", "GRPC")

		GoFullDocArg(t)

		os.Setenv("SPIDER_CALL_METHOD", holdEnv)

		TearDownForGrpc()
	})

}

func GoFullDocArg(t *testing.T) {
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
		Name:                "list namespace",
		Instance:            NsApi,
		Method:              "ListNS",
		Args:                nil,
		ExpectResStartsWith: `{"ns":[`,
	}
	MethodTest(t, tc)

	tc = TestCases{
		Name:     "get namespace",
		Instance: NsApi,
		Method:   "GetNS",
		Args: []interface{}{
			`{ "nsId": "ns-unit-01" }`,
		},
		ExpectResStartsWith: `{"id":"ns-unit-01"`,
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
		Name:     "list vnet",
		Instance: McirApi,
		Method:   "ListVNet",
		Args: []interface{}{
			`{ "nsId": "ns-unit-01", "resourceType":"vNet"}`,
		},
		ExpectResStartsWith: `{"vNet":[{"id":"mock-unit-config01-dev"`,
	}
	MethodTest(t, tc)

	tc = TestCases{
		Name:     "list vnet id",
		Instance: McirApi,
		Method:   "ListVNetId",
		Args: []interface{}{
			`{ "nsId": "ns-unit-01", "resourceType":"vNet"}`,
		},
		ExpectResStartsWith: `{"idList":["mock-unit-config01-dev"]}`,
	}
	MethodTest(t, tc)

	tc = TestCases{
		Name:     "get vnet",
		Instance: McirApi,
		Method:   "GetVNet",
		Args: []interface{}{
			`{ "nsId": "ns-unit-01", "resourceType":"vNet", "resourceId":"mock-unit-config01-dev"}`,
		},
		ExpectResStartsWith: `{"id":"mock-unit-config01-dev"`,
	}
	MethodTest(t, tc)

	tc = TestCases{
		Name:     "inspect vnet",
		Instance: TbutilApi,
		Method:   "InspectMcirResources",
		Args: []interface{}{
			`{ "connectionName": "mock-unit-config01", "type":"vNet" }`,
		},
		ExpectResStartsWith: `{"resourcesOnCsp":[{"id":"mock-unit-config01-dev"`,
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
		Name:     "list security",
		Instance: McirApi,
		Method:   "ListSecurityGroup",
		Args: []interface{}{
			`{ "nsId": "ns-unit-01", "resourceType":"securityGroup"}`,
		},
		ExpectResStartsWith: `{"securityGroup":[{"id":"mock-unit-config01-dev"`,
	}
	MethodTest(t, tc)

	tc = TestCases{
		Name:     "list security id",
		Instance: McirApi,
		Method:   "ListSecurityGroupId",
		Args: []interface{}{
			`{ "nsId": "ns-unit-01", "resourceType":"securityGroup"}`,
		},
		ExpectResStartsWith: `{"idList":["mock-unit-config01-dev"]}`,
	}
	MethodTest(t, tc)

	tc = TestCases{
		Name:     "get security",
		Instance: McirApi,
		Method:   "GetSecurityGroup",
		Args: []interface{}{
			`{ "nsId": "ns-unit-01", "resourceType":"securityGroup", "resourceId":"mock-unit-config01-dev"}`,
		},
		ExpectResStartsWith: `{"id":"mock-unit-config01-dev"`,
	}
	MethodTest(t, tc)

	tc = TestCases{
		Name:     "inspect security",
		Instance: TbutilApi,
		Method:   "InspectMcirResources",
		Args: []interface{}{
			`{ "connectionName": "mock-unit-config01", "type":"securityGroup" }`,
		},
		ExpectResStartsWith: `{"resourcesOnCsp":[{"id":"mock-unit-config01-dev"`,
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
		Name:     "list sshkey",
		Instance: McirApi,
		Method:   "ListSshKey",
		Args: []interface{}{
			`{ "nsId": "ns-unit-01", "resourceType":"sshKey"}`,
		},
		ExpectResStartsWith: `{"sshKey":[{"id":"mock-unit-config01-dev"`,
	}
	MethodTest(t, tc)

	tc = TestCases{
		Name:     "list sshkey id",
		Instance: McirApi,
		Method:   "ListSshKeyId",
		Args: []interface{}{
			`{ "nsId": "ns-unit-01", "resourceType":"sshKey"}`,
		},
		ExpectResStartsWith: `{"idList":["mock-unit-config01-dev"]}`,
	}
	MethodTest(t, tc)

	tc = TestCases{
		Name:     "get sshkey",
		Instance: McirApi,
		Method:   "GetSshKey",
		Args: []interface{}{
			`{ "nsId": "ns-unit-01", "resourceType":"sshKey", "resourceId":"mock-unit-config01-dev"}`,
		},
		ExpectResStartsWith: `{"id":"mock-unit-config01-dev"`,
	}
	MethodTest(t, tc)

	tc = TestCases{
		Name:     "inspect sshkey",
		Instance: TbutilApi,
		Method:   "InspectMcirResources",
		Args: []interface{}{
			`{ "connectionName": "mock-unit-config01", "type":"sshKey" }`,
		},
		ExpectResStartsWith: `{"resourcesOnCsp":[{"id":"mock-unit-config01-dev"`,
	}
	MethodTest(t, tc)

	tc = TestCases{
		Name:     "list lookup image",
		Instance: McirApi,
		Method:   "ListLookupImage",
		Args: []interface{}{
			`{ "connectionName": "mock-unit-config01" }`,
		},
		ExpectResStartsWith: `{"image":[{"IId":{"NameId":"mock-vmimage-01"`,
	}
	MethodTest(t, tc)

	tc = TestCases{
		Name:     "lookup image",
		Instance: McirApi,
		Method:   "GetLookupImage",
		Args: []interface{}{
			`{ "connectionName": "mock-unit-config01", "cspImageId": "mock-vmimage-01" }`,
		},
		ExpectResStartsWith: `{"IId":{"NameId":"mock-vmimage-01"`,
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
		Name:     "list image",
		Instance: McirApi,
		Method:   "ListImage",
		Args: []interface{}{
			`{ "nsId": "ns-unit-01", "resourceType":"image"}`,
		},
		ExpectResStartsWith: `{"image":[{"namespace":"ns-unit-01","id":"mock-unit-config01-dev"`,
	}
	MethodTest(t, tc)

	tc = TestCases{
		Name:     "list image id",
		Instance: McirApi,
		Method:   "ListImageId",
		Args: []interface{}{
			`{ "nsId": "ns-unit-01", "resourceType":"image"}`,
		},
		ExpectResStartsWith: `{"idList":["mock-unit-config01-dev"]}`,
	}
	MethodTest(t, tc)

	tc = TestCases{
		Name:     "get image",
		Instance: McirApi,
		Method:   "GetImage",
		Args: []interface{}{
			`{ "nsId": "ns-unit-01", "resourceType":"image", "resourceId":"mock-unit-config01-dev"}`,
		},
		ExpectResStartsWith: `{"namespace":"ns-unit-01","id":"mock-unit-config01-dev"`,
	}
	MethodTest(t, tc)

	tc = TestCases{
		Name:     "search image",
		Instance: McirApi,
		Method:   "SearchImage",
		Args: []interface{}{
			`{
				"nsId":  "ns-unit-01",
				"keywords": [
					"mock"
				]					
			}`,
		},
		ExpectResStartsWith: `{"image":[{"namespace":"ns-unit-01","id":"mock-unit-config01-dev"`,
	}
	MethodTest(t, tc)

	tc = TestCases{
		Name:     "fetch images",
		Instance: McirApi,
		Method:   "FetchImage",
		Args: []interface{}{
			`{ "ConnectionName": "!all", "nsId": "ns-unit-01" }`,
		},
		ExpectResStartsWith: `{"message":"Fetched 5 images (from 1 connConfigs)"}`,
	}
	MethodTest(t, tc)

	tc = TestCases{
		Name:     "list lookup spec",
		Instance: McirApi,
		Method:   "ListLookupSpec",
		Args: []interface{}{
			`{ "connectionName": "mock-unit-config01" }`,
		},
		ExpectResStartsWith: `{"vmspec":[{"Region":"default","Name":"mock-vmspec-01"`,
	}
	MethodTest(t, tc)

	tc = TestCases{
		Name:     "lookup spec",
		Instance: McirApi,
		Method:   "GetLookupSpec",
		Args: []interface{}{
			`{ "connectionName": "mock-unit-config01", "cspSpecName": "mock-vmspec-01" }`,
		},
		ExpectResStartsWith: `{"Region":"default","Name":"mock-vmspec-01"`,
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
		Name:     "list spec",
		Instance: McirApi,
		Method:   "ListSpec",
		Args: []interface{}{
			`{ "nsId": "ns-unit-01", "resourceType":"spec"}`,
		},
		ExpectResStartsWith: `{"spec":[{"namespace":"ns-unit-01","id":"mock-unit-config01-dev"`,
	}
	MethodTest(t, tc)

	tc = TestCases{
		Name:     "list spec id",
		Instance: McirApi,
		Method:   "ListSpecId",
		Args: []interface{}{
			`{ "nsId": "ns-unit-01", "resourceType":"spec"}`,
		},
		ExpectResStartsWith: `{"idList":["mock-unit-config01-dev"]}`,
	}
	MethodTest(t, tc)

	tc = TestCases{
		Name:     "get spec",
		Instance: McirApi,
		Method:   "GetSpec",
		Args: []interface{}{
			`{ "nsId": "ns-unit-01", "resourceType":"spec", "resourceId":"mock-unit-config01-dev"}`,
		},
		ExpectResStartsWith: `{"namespace":"ns-unit-01","id":"mock-unit-config01-dev"`,
	}
	MethodTest(t, tc)

	tc = TestCases{
		Name:     "update spec",
		Instance: McirApi,
		Method:   "UpdateSpec",
		Args: []interface{}{
			`{
				"nsId":  "ns-unit-01",
				"spec": {
					"id": "mock-unit-config01-dev", 
					"description": "UpdateSpec() test"
				}					
			}`,
		},
		ExpectResStartsWith: `{"namespace":"ns-unit-01","id":"mock-unit-config01-dev"`,
	}
	MethodTest(t, tc)

	tc = TestCases{
		Name:     "filter spec",
		Instance: McirApi,
		Method:   "FilterSpec",
		Args: []interface{}{
			`{
				"nsId":  "ns-unit-01",
				"spec": {
						"num_vCPU": 4, 
						"mem_GiB": 32
				}					
			}`,
		},
		ExpectResStartsWith: `{"spec":[{"namespace":"ns-unit-01","id":"mock-unit-config01-dev"`,
	}
	MethodTest(t, tc)

	tc = TestCases{
		Name:     "range filter spec",
		Instance: McirApi,
		Method:   "FilterSpecsByRange",
		Args: []interface{}{
			`{
				"nsId":  "ns-unit-01",
				"filter": {
						"mem_GiB": {
							"min": 4
						}
				}					
			}`,
		},
		ExpectResStartsWith: `{"spec":[{"namespace":"ns-unit-01","id":"mock-unit-config01-dev"`,
	}
	MethodTest(t, tc)

	tc = TestCases{
		Name:     "fetch specs",
		Instance: McirApi,
		Method:   "FetchSpec",
		Args: []interface{}{
			`{ "ConnectionName": "!all", "nsId": "ns-unit-01" }`,
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
		Name:     "add vm to mcis",
		Instance: McisApi,
		Method:   "CreateMcisVM",
		Args: []interface{}{
			`{
				"nsId":  "ns-unit-01",
				"mcisId": "mock-unit-config01-dev",
				"mcisvm": {
					"name": "mock-unit-config01-dev",
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
			}`,
		},
		ExpectResStartsWith: `{"id":"mock-unit-config01-dev"`,
	}
	MethodTest(t, tc)

	tc = TestCases{
		Name:     "add vmgroup to mcis",
		Instance: McisApi,
		Method:   "CreateMcisVMGroup",
		Args: []interface{}{
			`{
				"nsId":  "ns-unit-01",
				"mcisId": "mock-unit-config01-dev",
				"groupvm": {
					"vmGroupSize": "3",
					"name": "mock-unit-config01-dev",
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
			}`,
		},
		ExpectResStartsWith: `{"id":"mock-unit-config01-dev"`,
	}
	MethodTest(t, tc)

	tc = TestCases{
		Name:     "list mcis",
		Instance: McisApi,
		Method:   "ListMcis",
		Args: []interface{}{
			`{ "nsId": "ns-unit-01" }`,
		},
		ExpectResStartsWith: `{"mcis":[{"id":"mock-unit-config01-dev"`,
	}
	MethodTest(t, tc)

	tc = TestCases{
		Name:     "list mcis id",
		Instance: McisApi,
		Method:   "ListMcisId",
		Args: []interface{}{
			`{ "nsId": "ns-unit-01" }`,
		},
		ExpectResStartsWith: `{"idList":["mock-unit-config01-dev"]}`,
	}
	MethodTest(t, tc)

	tc = TestCases{
		Name:     "list mcis status",
		Instance: McisApi,
		Method:   "ListMcisStatus",
		Args: []interface{}{
			`{ "nsId": "ns-unit-01" }`,
		},
		ExpectResStartsWith: `{"mcis":[{"id":"mock-unit-config01-dev"`,
	}
	MethodTest(t, tc)

	tc = TestCases{
		Name:     "get mcis",
		Instance: McisApi,
		Method:   "GetMcisInfo",
		Args: []interface{}{
			`{ "nsId": "ns-unit-01", "mcisId":"mock-unit-config01-dev"}`,
		},
		ExpectResStartsWith: `{"id":"mock-unit-config01-dev"`,
	}
	MethodTest(t, tc)

	tc = TestCases{
		Name:     "get mcis status",
		Instance: McisApi,
		Method:   "GetMcisStatus",
		Args: []interface{}{
			`{ "nsId": "ns-unit-01", "mcisId":"mock-unit-config01-dev"}`,
		},
		ExpectResStartsWith: `{"status":{"id":"mock-unit-config01-dev"`,
	}
	MethodTest(t, tc)

	tc = TestCases{
		Name:     "suspend mcis",
		Instance: McisApi,
		Method:   "ControlMcis",
		Args: []interface{}{
			`{ "nsId": "ns-unit-01", "mcisId":"mock-unit-config01-dev", "action":"suspend"}`,
		},
		ExpectResStartsWith: `{"message":"Suspending the MCIS"}`,
	}
	MethodTest(t, tc)

	tc = TestCases{
		Name:     "resume mcis",
		Instance: McisApi,
		Method:   "ControlMcis",
		Args: []interface{}{
			`{ "nsId": "ns-unit-01", "mcisId":"mock-unit-config01-dev", "action":"resume"}`,
		},
		ExpectResStartsWith: `{"message":"Resuming the MCIS"}`,
	}
	MethodTest(t, tc)

	tc = TestCases{
		Name:     "reboot mcis",
		Instance: McisApi,
		Method:   "ControlMcis",
		Args: []interface{}{
			`{ "nsId": "ns-unit-01", "mcisId":"mock-unit-config01-dev", "action":"reboot"}`,
		},
		ExpectResStartsWith: `{"message":"Rebooting the MCIS"}`,
	}
	MethodTest(t, tc)

	tc = TestCases{
		Name:     "refine mcis",
		Instance: McisApi,
		Method:   "ControlMcis",
		Args: []interface{}{
			`{ "nsId": "ns-unit-01", "mcisId":"mock-unit-config01-dev", "action":"refine"}`,
		},
		ExpectResStartsWith: `{"message":"Refined the MCIS"}`,
	}
	MethodTest(t, tc)

	tc = TestCases{
		Name:     "list vm id",
		Instance: McisApi,
		Method:   "ListMcisVmId",
		Args: []interface{}{
			`{ "nsId": "ns-unit-01", "mcisId":"mock-unit-config01-dev"}`,
		},
		ExpectResStartsWith: `{"idList":["mock-unit-config01-dev","mock-unit-config01-dev-0"`,
	}
	MethodTest(t, tc)

	tc = TestCases{
		Name:     "inspect vm",
		Instance: TbutilApi,
		Method:   "InspectVmResources",
		Args: []interface{}{
			`{ "connectionName": "mock-unit-config01", "type":"vm" }`,
		},
		ExpectResStartsWith: `{"resourcesOnCsp":[{"id":"ns-unit-01-mock-unit-config01-dev-mock-unit-config01-dev"`,
	}
	MethodTest(t, tc)

	tc = TestCases{
		Name:     "create mcis policy",
		Instance: McisApi,
		Method:   "CreateMcisPolicy",
		Args: []interface{}{
			`{
				"nsId":  "ns-unit-01",
				"mcisId": "mock-unit-config01-dev",
				"ReqInfo": {
					"description": "Tumblebug Auto Control Demo",
					"policy": [
						{
							"autoCondition": {
								"metric": "cpu",
								"operator": ">=",
								"operand": "80",
								"evaluationPeriod": "10"
							},
							"autoAction": {
								"actionType": "ScaleOut",
								"placementAlgo": "random",
								"vm": {
									"name": "AutoGen"
								},
								"postCommand": {
									"command": "wget https://raw.githubusercontent.com/cloud-barista/cb-tumblebug/master/assets/scripts/setweb.sh -O ~/setweb.sh; chmod +x ~/setweb.sh; sudo ~/setweb.sh; wget https://raw.githubusercontent.com/cloud-barista/cb-tumblebug/master/assets/scripts/runLoadMaker.sh -O ~/runLoadMaker.sh; chmod +x ~/runLoadMaker.sh; sudo ~/runLoadMaker.sh"
								}
							}
						},				
						{
							"autoCondition": {
								"metric": "cpu",
								"operator": "<=",
								"operand": "60",
								"evaluationPeriod": "10"
							},
							"autoAction": {
								"actionType": "ScaleIn"
							}
						}
					]
				}					
			}`,
		},
		ExpectResStartsWith: `{"Name":"mock-unit-config01-dev"`,
	}
	MethodTest(t, tc)

	tc = TestCases{
		Name:     "list mcis policy",
		Instance: McisApi,
		Method:   "ListMcisPolicy",
		Args: []interface{}{
			`{ "nsId": "ns-unit-01" }`,
		},
		ExpectResStartsWith: `{"mcisPolicy":[{"Name":"mock-unit-config01-dev"`,
	}
	MethodTest(t, tc)

	tc = TestCases{
		Name:     "get mcis policy",
		Instance: McisApi,
		Method:   "GetMcisPolicy",
		Args: []interface{}{
			`{ "nsId": "ns-unit-01", "mcisId":"mock-unit-config01-dev"}`,
		},
		ExpectResStartsWith: `{"Name":"mock-unit-config01-dev"`,
	}
	MethodTest(t, tc)

	tc = TestCases{
		Name:     "update config",
		Instance: TbutilApi,
		Method:   "CreateConfig",
		Args: []interface{}{
			`{ "name": "key01", "value":"value01"}`,
		},
		ExpectResStartsWith: `{"id":"key01","name":"key01","value":"value01"}`,
	}
	MethodTest(t, tc)

	tc = TestCases{
		Name:                "list config",
		Instance:            TbutilApi,
		Method:              "ListConfig",
		Args:                nil,
		ExpectResStartsWith: `{"config":[{"id":"key01","name":"key01","value":"value01"}]}`,
	}
	MethodTest(t, tc)

	tc = TestCases{
		Name:     "get config",
		Instance: TbutilApi,
		Method:   "GetConfig",
		Args: []interface{}{
			`{ "configId": "key01" }`,
		},
		ExpectResStartsWith: `{"id":"key01","name":"key01","value":"value01"}`,
	}
	MethodTest(t, tc)

	tc = TestCases{
		Name:     "install agent",
		Instance: McisApi,
		Method:   "InstallMonitorAgentToMcis",
		Args: []interface{}{
			`{
				"nsId":  "ns-unit-01",
				"mcisId": "mock-unit-config01-dev",
				"cmd": {
					"command": "echo -n [CMD] Works! [Public IP: ; curl https://api.ipify.org ; echo -n ], [HostName: ; hostname ; echo -n ]"
				}					
			}`,
		},
		ExpectResStartsWith: `{"result_array":[{"mcisId":"mock-unit-config01-dev"`,
	}
	MethodTest(t, tc)

	tc = TestCases{
		Name:     "get monitoring data",
		Instance: McisApi,
		Method:   "GetMonitorData",
		Args: []interface{}{
			`{ "nsId": "ns-unit-01", "mcisId":"mock-unit-config01-dev", "metric": "cpu"}`,
		},
		ExpectResStartsWith: `{"nsId":"ns-unit-01","mcisId":"mock-unit-config01-dev","mcisMonitoring":[{"metric"`,
	}
	MethodTest(t, tc)

	tc = TestCases{
		Name:                "list connection config",
		Instance:            TbutilApi,
		Method:              "ListConnConfig",
		Args:                nil,
		ExpectResStartsWith: `{"connectionconfig":[{"ConfigName":"mock-unit-config01","ProviderName":"MOCK"`,
	}
	MethodTest(t, tc)

	tc = TestCases{
		Name:     "get connection config",
		Instance: TbutilApi,
		Method:   "GetConnConfig",
		Args: []interface{}{
			`{ "connConfigName": "mock-unit-config01" }`,
		},
		ExpectResStartsWith: `{"ConfigName":"mock-unit-config01","ProviderName":"MOCK"`,
	}
	MethodTest(t, tc)

	tc = TestCases{
		Name:                "list region",
		Instance:            TbutilApi,
		Method:              "ListRegion",
		Args:                nil,
		ExpectResStartsWith: `{"region":[{"RegionName":"mock-unit-region01","ProviderName":"MOCK"`,
	}
	MethodTest(t, tc)

	tc = TestCases{
		Name:     "get region",
		Instance: TbutilApi,
		Method:   "GetRegion",
		Args: []interface{}{
			`{ "regionName": "mock-unit-region01" }`,
		},
		ExpectResStartsWith: `{"RegionName":"mock-unit-region01","ProviderName":"MOCK"`,
	}
	MethodTest(t, tc)

	tc = TestCases{
		Name:     "command mcis",
		Instance: McisApi,
		Method:   "CmdMcis",
		Args: []interface{}{
			`{
				"nsId":  "ns-unit-01",
				"mcisId": "mock-unit-config01-dev",
				"cmd": {
					"command": "echo -n [CMD] Works! [Public IP: ; curl https://api.ipify.org ; echo -n ], [HostName: ; hostname ; echo -n ]"
				}					
			}`,
		},
		ExpectResStartsWith: `{"result_array":[{"mcisId":"mock-unit-config01-dev"`,
	}
	MethodTest(t, tc)

	tc = TestCases{
		Name:     "command mcis vm",
		Instance: McisApi,
		Method:   "CmdMcisVm",
		Args: []interface{}{
			`{
				"nsId":  "ns-unit-01",
				"mcisId": "mock-unit-config01-dev",
				"vmId": "mock-unit-config01-dev",
				"cmd": {
					"command": "echo -n [CMD] Works! [Public IP: ; curl https://api.ipify.org ; echo -n ], [HostName: ; hostname ; echo -n ]"
				}					
			}`,
		},
		ExpectResStartsWith: `{"Result":"echo -n [CMD] Works`,
	}
	MethodTest(t, tc)

	tc = TestCases{
		Name:     "list object",
		Instance: TbutilApi,
		Method:   "ListObject",
		Args: []interface{}{
			`{ "key": "" }`,
		},
		ExpectResStartsWith: `{"object":["/config/key01"`,
	}
	MethodTest(t, tc)

	tc = TestCases{
		Name:     "list sub object",
		Instance: TbutilApi,
		Method:   "ListObject",
		Args: []interface{}{
			`{ "key": "/config" }`,
		},
		ExpectResStartsWith: `{"object":["/config/key01"]}`,
	}
	MethodTest(t, tc)

	tc = TestCases{
		Name:     "get object",
		Instance: TbutilApi,
		Method:   "GetObject",
		Args: []interface{}{
			`{ "key": "/config/key01" }`,
		},
		ExpectResStartsWith: `{"id":"key01","name":"key01","value":"value01"}`,
	}
	MethodTest(t, tc)

	//
	// Delete Resources
	//

	tc = TestCases{
		Name:     "delete object",
		Instance: TbutilApi,
		Method:   "DeleteObject",
		Args: []interface{}{
			`{ "key": "/config/key01" }`,
		},
		ExpectResStartsWith: `{"message":"The object has been deleted"}`,
	}
	MethodTest(t, tc)

	tc = TestCases{
		Name:     "delete all sub object",
		Instance: TbutilApi,
		Method:   "DeleteAllObject",
		Args: []interface{}{
			`{ "key": "/config" }`,
		},
		ExpectResStartsWith: `{"message":"Objects have been deleted"}`,
	}
	MethodTest(t, tc)

	tc = TestCases{
		Name:     "init config",
		Instance: TbutilApi,
		Method:   "InitConfig",
		Args: []interface{}{
			`{ "configId": "key01" }`,
		},
		ExpectResStartsWith: `{"message":"The config key01 has been initialized."}`,
	}
	MethodTest(t, tc)

	tc = TestCases{
		Name:                "init all config",
		Instance:            TbutilApi,
		Method:              "InitAllConfig",
		Args:                nil,
		ExpectResStartsWith: `{"message":"All configs have been initialized."}`,
	}
	MethodTest(t, tc)

	tc = TestCases{
		Name:     "delete mcis policy",
		Instance: McisApi,
		Method:   "DeleteMcisPolicy",
		Args: []interface{}{
			`{ "nsId": "ns-unit-01", "mcisId":"mock-unit-config01-dev"}`,
		},
		ExpectResStartsWith: `{"message":"Deleting the MCIS Policy info"}`,
	}
	MethodTest(t, tc)

	tc = TestCases{
		Name:     "delete all mcis policy",
		Instance: McisApi,
		Method:   "DeleteAllMcisPolicy",
		Args: []interface{}{
			`{ "nsId": "ns-unit-01" }`,
		},
		ExpectResStartsWith: `{"message":"No MCIS Policy to delete"}`,
	}
	MethodTest(t, tc)

	tc = TestCases{
		Name:     "terminate mcis",
		Instance: McisApi,
		Method:   "ControlMcis",
		Args: []interface{}{
			`{ "nsId": "ns-unit-01", "mcisId":"mock-unit-config01-dev", "action":"terminate"}`,
		},
		ExpectResStartsWith: `{"message":"Terminating the MCIS"}`,
	}
	MethodTest(t, tc)

	tc = TestCases{
		Name:     "delete mcis",
		Instance: McisApi,
		Method:   "DeleteMcis",
		Args: []interface{}{
			`{ "nsId": "ns-unit-01", "mcisId":"mock-unit-config01-dev"}`,
		},
		ExpectResStartsWith: `{"message":"Deleting the MCIS mock-unit-config01-dev"}`,
	}
	MethodTest(t, tc)

	tc = TestCases{
		Name:     "delete all mcis",
		Instance: McisApi,
		Method:   "DeleteAllMcis",
		Args: []interface{}{
			`{ "nsId": "ns-unit-01" }`,
		},
		ExpectResStartsWith: `{"message":"No MCIS to delete"}`,
	}
	MethodTest(t, tc)

	tc = TestCases{
		Name:     "delete spec",
		Instance: McirApi,
		Method:   "DeleteSpec",
		Args: []interface{}{
			`{ "nsId": "ns-unit-01", "resourceType":"spec", "resourceId":"mock-unit-config01-dev", "force":"false"}`,
		},
		ExpectResStartsWith: `{"message":"The spec mock-unit-config01-dev has been deleted"}`,
	}
	MethodTest(t, tc)

	tc = TestCases{
		Name:     "delete all spec",
		Instance: McirApi,
		Method:   "DeleteAllSpec",
		Args: []interface{}{
			`{ "nsId": "ns-unit-01", "resourceType":"spec", "force":"false"}`,
		},
		ExpectResStartsWith: `{"message":"All specs has been deleted"}`,
	}
	MethodTest(t, tc)

	tc = TestCases{
		Name:     "delete image",
		Instance: McirApi,
		Method:   "DeleteImage",
		Args: []interface{}{
			`{ "nsId": "ns-unit-01", "resourceType":"image", "resourceId":"mock-unit-config01-dev", "force":"false"}`,
		},
		ExpectResStartsWith: `{"message":"The image mock-unit-config01-dev has been deleted"}`,
	}
	MethodTest(t, tc)

	tc = TestCases{
		Name:     "delete all image",
		Instance: McirApi,
		Method:   "DeleteAllImage",
		Args: []interface{}{
			`{ "nsId": "ns-unit-01", "resourceType":"image", "force":"false"}`,
		},
		ExpectResStartsWith: `{"message":"All images has been deleted"}`,
	}
	MethodTest(t, tc)

	tc = TestCases{
		Name:     "delete sshkey",
		Instance: McirApi,
		Method:   "DeleteSshKey",
		Args: []interface{}{
			`{ "nsId": "ns-unit-01", "resourceType":"sshKey", "resourceId":"mock-unit-config01-dev", "force":"false"}`,
		},
		ExpectResStartsWith: `{"message":"The sshKey mock-unit-config01-dev has been deleted"}`,
	}
	MethodTest(t, tc)

	tc = TestCases{
		Name:     "delete all sshkey",
		Instance: McirApi,
		Method:   "DeleteAllSshKey",
		Args: []interface{}{
			`{ "nsId": "ns-unit-01", "resourceType":"sshKey", "force":"false"}`,
		},
		ExpectResStartsWith: `{"message":"All sshKeys has been deleted"}`,
	}
	MethodTest(t, tc)

	tc = TestCases{
		Name:     "delete security",
		Instance: McirApi,
		Method:   "DeleteSecurityGroup",
		Args: []interface{}{
			`{ "nsId": "ns-unit-01", "resourceType":"securityGroup", "resourceId":"mock-unit-config01-dev", "force":"false"}`,
		},
		ExpectResStartsWith: `{"message":"The securityGroup mock-unit-config01-dev has been deleted"}`,
	}
	MethodTest(t, tc)

	tc = TestCases{
		Name:     "delete all security",
		Instance: McirApi,
		Method:   "DeleteAllSecurityGroup",
		Args: []interface{}{
			`{ "nsId": "ns-unit-01", "resourceType":"securityGroup", "force":"false"}`,
		},
		ExpectResStartsWith: `{"message":"All securityGroups has been deleted"}`,
	}
	MethodTest(t, tc)

	tc = TestCases{
		Name:     "delete vnet",
		Instance: McirApi,
		Method:   "DeleteVNet",
		Args: []interface{}{
			`{ "nsId": "ns-unit-01", "resourceType":"vNet", "resourceId":"mock-unit-config01-dev", "force":"false"}`,
		},
		ExpectResStartsWith: `{"message":"The vNet mock-unit-config01-dev has been deleted"}`,
	}
	MethodTest(t, tc)

	tc = TestCases{
		Name:     "delete all vnet",
		Instance: McirApi,
		Method:   "DeleteAllVNet",
		Args: []interface{}{
			`{ "nsId": "ns-unit-01", "resourceType":"vNet", "force":"false"}`,
		},
		ExpectResStartsWith: `{"message":"All vNets has been deleted"}`,
	}
	MethodTest(t, tc)

	tc = TestCases{
		Name:     "delete namespace",
		Instance: NsApi,
		Method:   "DeleteNS",
		Args: []interface{}{
			`{ "nsId": "ns-unit-01" }`,
		},
		ExpectResStartsWith: `{"message":"The ns has been deleted"}`,
	}
	MethodTest(t, tc)

	tc = TestCases{
		Name:                "delete all namespace",
		Instance:            NsApi,
		Method:              "DeleteAllNS",
		Args:                nil,
		ExpectResStartsWith: `{"message":"All namespaces has been deleted"}`,
	}
	MethodTest(t, tc)
}
