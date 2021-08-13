package restscenario

import (
	"net/http"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestRestDuplicate(t *testing.T) {
	t.Run("rest api duplicate test for mock driver", func(t *testing.T) {
		SetUpForRest()

		tc := TestCases{
			Name:                 "create namespace",
			EchoFunc:             "RestPostNs",
			HttpMethod:           http.MethodPost,
			WhenURL:              "/tumblebug/ns",
			GivenQueryParams:     "",
			GivenParaNames:       nil,
			GivenParaVals:        nil,
			GivenPostData:        `{"name":"ns-unit-01","description":"NameSpace for General Testing"}`,
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"id":"ns-unit-01"`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:                 "create namespace(duplicate)",
			EchoFunc:             "RestPostNs",
			HttpMethod:           http.MethodPost,
			WhenURL:              "/tumblebug/ns",
			GivenQueryParams:     "",
			GivenParaNames:       nil,
			GivenParaVals:        nil,
			GivenPostData:        `{"name":"ns-unit-01","description":"NameSpace for General Testing"}`,
			ExpectStatus:         http.StatusBadRequest,
			ExpectBodyStartsWith: `{"message":"CreateNs(); The namespace ns-unit-01 already exists."}`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:             "create vnet",
			EchoFunc:         "RestPostVNet",
			HttpMethod:       http.MethodPost,
			WhenURL:          "/tumblebug/ns/:nsId/resources/vNet",
			GivenQueryParams: "",
			GivenParaNames:   []string{"nsId"},
			GivenParaVals:    []string{"ns-unit-01"},
			GivenPostData: `{
				"name": "mock-unit-config01-dev",
				"connectionName": "mock-unit-config01",
				"cidrBlock": "192.168.0.0/16",
				"subnetInfoList": [ {
					"Name": "mock-unit-config01-dev",
					"IPv4_CIDR": "192.168.1.0/24"
				} ]
			}`,
			ExpectStatus:         http.StatusCreated,
			ExpectBodyStartsWith: `{"id":"mock-unit-config01-dev"`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:             "create vnet(duplicate)",
			EchoFunc:         "RestPostVNet",
			HttpMethod:       http.MethodPost,
			WhenURL:          "/tumblebug/ns/:nsId/resources/vNet",
			GivenQueryParams: "",
			GivenParaNames:   []string{"nsId"},
			GivenParaVals:    []string{"ns-unit-01"},
			GivenPostData: `{
				"name": "mock-unit-config01-dev",
				"connectionName": "mock-unit-config01",
				"cidrBlock": "192.168.0.0/16",
				"subnetInfoList": [ {
					"Name": "mock-unit-config01-dev",
					"IPv4_CIDR": "192.168.1.0/24"
				} ]
			}`,
			ExpectStatus:         http.StatusInternalServerError,
			ExpectBodyStartsWith: `{"message":"The vNet mock-unit-config01-dev already exists."}`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:             "create security",
			EchoFunc:         "RestPostSecurityGroup",
			HttpMethod:       http.MethodPost,
			WhenURL:          "/tumblebug/ns/:nsId/resources/securityGroup",
			GivenQueryParams: "",
			GivenParaNames:   []string{"nsId"},
			GivenParaVals:    []string{"ns-unit-01"},
			GivenPostData: `{
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
			}`,
			ExpectStatus:         http.StatusCreated,
			ExpectBodyStartsWith: `{"id":"mock-unit-config01-dev"`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:             "create security(duplicate)",
			EchoFunc:         "RestPostSecurityGroup",
			HttpMethod:       http.MethodPost,
			WhenURL:          "/tumblebug/ns/:nsId/resources/securityGroup",
			GivenQueryParams: "",
			GivenParaNames:   []string{"nsId"},
			GivenParaVals:    []string{"ns-unit-01"},
			GivenPostData: `{
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
			}`,
			ExpectStatus:         http.StatusInternalServerError,
			ExpectBodyStartsWith: `{"message":"The securityGroup mock-unit-config01-dev already exists."}`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:             "create sshkey",
			EchoFunc:         "RestPostSshKey",
			HttpMethod:       http.MethodPost,
			WhenURL:          "/tumblebug/ns/:nsId/resources/sshKey",
			GivenQueryParams: "",
			GivenParaNames:   []string{"nsId"},
			GivenParaVals:    []string{"ns-unit-01"},
			GivenPostData: `{
				"name": "mock-unit-config01-dev",
				"connectionName": "mock-unit-config01",
				"description": ""					
			}`,
			ExpectStatus:         http.StatusCreated,
			ExpectBodyStartsWith: `{"id":"mock-unit-config01-dev"`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:             "create sshkey(duplicate)",
			EchoFunc:         "RestPostSshKey",
			HttpMethod:       http.MethodPost,
			WhenURL:          "/tumblebug/ns/:nsId/resources/sshKey",
			GivenQueryParams: "",
			GivenParaNames:   []string{"nsId"},
			GivenParaVals:    []string{"ns-unit-01"},
			GivenPostData: `{
				"name": "mock-unit-config01-dev",
				"connectionName": "mock-unit-config01",
				"description": ""					
			}`,
			ExpectStatus:         http.StatusInternalServerError,
			ExpectBodyStartsWith: `{"message":"The sshKey mock-unit-config01-dev already exists."}`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:             "register image with id",
			EchoFunc:         "RestPostImage",
			HttpMethod:       http.MethodPost,
			WhenURL:          "/tumblebug/ns/:nsId/resources/image",
			GivenQueryParams: "?action=registerWithId",
			GivenParaNames:   []string{"nsId"},
			GivenParaVals:    []string{"ns-unit-01"},
			GivenPostData: `{
				"connectionName": "mock-unit-config01",
				"name": "mock-unit-config01-dev",
				"cspImageId": "mock-vmimage-01",
				"description": "Canonical, Ubuntu, 18.04 LTS, amd64 bionic"
			}`,
			ExpectStatus:         http.StatusCreated,
			ExpectBodyStartsWith: `{"namespace":"ns-unit-01","id":"mock-unit-config01-dev"`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:             "register image with id(duplicate)",
			EchoFunc:         "RestPostImage",
			HttpMethod:       http.MethodPost,
			WhenURL:          "/tumblebug/ns/:nsId/resources/image",
			GivenQueryParams: "?action=registerWithId",
			GivenParaNames:   []string{"nsId"},
			GivenParaVals:    []string{"ns-unit-01"},
			GivenPostData: `{
				"connectionName": "mock-unit-config01",
				"name": "mock-unit-config01-dev",
				"cspImageId": "mock-vmimage-01",
				"description": "Canonical, Ubuntu, 18.04 LTS, amd64 bionic"
			}`,
			ExpectStatus:         http.StatusInternalServerError,
			ExpectBodyStartsWith: `{"message":"The image mock-unit-config01-dev already exists."}`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:             "register spec",
			EchoFunc:         "RestPostSpec",
			HttpMethod:       http.MethodPost,
			WhenURL:          "/tumblebug/ns/:nsId/resources/spec",
			GivenQueryParams: "",
			GivenParaNames:   []string{"nsId"},
			GivenParaVals:    []string{"ns-unit-01"},
			GivenPostData: `{
				"connectionName": "mock-unit-config01",
				"name": "mock-unit-config01-dev",
				"cspSpecName": "mock-vmspec-01"				
			}`,
			ExpectStatus:         http.StatusCreated,
			ExpectBodyStartsWith: `{"namespace":"ns-unit-01","id":"mock-unit-config01-dev"`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:             "register spec(duplicate)",
			EchoFunc:         "RestPostSpec",
			HttpMethod:       http.MethodPost,
			WhenURL:          "/tumblebug/ns/:nsId/resources/spec",
			GivenQueryParams: "",
			GivenParaNames:   []string{"nsId"},
			GivenParaVals:    []string{"ns-unit-01"},
			GivenPostData: `{
				"connectionName": "mock-unit-config01",
				"name": "mock-unit-config01-dev",
				"cspSpecName": "mock-vmspec-01"				
			}`,
			ExpectStatus:         http.StatusInternalServerError,
			ExpectBodyStartsWith: `{"message":"The spec mock-unit-config01-dev already exists."}`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:             "create mcis",
			EchoFunc:         "RestPostMcis",
			HttpMethod:       http.MethodPost,
			WhenURL:          "/tumblebug/ns/:nsId/mcis",
			GivenQueryParams: "",
			GivenParaNames:   []string{"nsId"},
			GivenParaVals:    []string{"ns-unit-01"},
			GivenPostData: `{
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
			}`,
			ExpectStatus:         http.StatusCreated,
			ExpectBodyStartsWith: `{"id":"mock-unit-config01-dev"`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:             "create mcis(duplicate)",
			EchoFunc:         "RestPostMcis",
			HttpMethod:       http.MethodPost,
			WhenURL:          "/tumblebug/ns/:nsId/mcis",
			GivenQueryParams: "",
			GivenParaNames:   []string{"nsId"},
			GivenParaVals:    []string{"ns-unit-01"},
			GivenPostData: `{
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
			}`,
			ExpectStatus:         http.StatusInternalServerError,
			ExpectBodyStartsWith: `{"message":"The mcis mock-unit-config01-dev already exists."}`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:             "add vm to mcis",
			EchoFunc:         "RestPostMcisVm",
			HttpMethod:       http.MethodPost,
			WhenURL:          "/tumblebug/ns/:nsId/mcis/:mcisId/vm",
			GivenQueryParams: "",
			GivenParaNames:   []string{"nsId", "mcisId"},
			GivenParaVals:    []string{"ns-unit-01", "mock-unit-config01-dev"},
			GivenPostData: `{
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
			}`,
			ExpectStatus:         http.StatusCreated,
			ExpectBodyStartsWith: `{"id":"mock-unit-config01-dev"`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:             "add vm to mcis(duplicate)",
			EchoFunc:         "RestPostMcisVm",
			HttpMethod:       http.MethodPost,
			WhenURL:          "/tumblebug/ns/:nsId/mcis/:mcisId/vm",
			GivenQueryParams: "",
			GivenParaNames:   []string{"nsId", "mcisId"},
			GivenParaVals:    []string{"ns-unit-01", "mock-unit-config01-dev"},
			GivenPostData: `{
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
			}`,
			ExpectStatus:         http.StatusInternalServerError,
			ExpectBodyStartsWith: `{"message":"The vm mock-unit-config01-dev already exists."}`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:             "add vmgroup to mcis",
			EchoFunc:         "RestPostMcisVmGroup",
			HttpMethod:       http.MethodPost,
			WhenURL:          "/tumblebug/ns/:nsId/mcis/:mcisId/vmgroup",
			GivenQueryParams: "",
			GivenParaNames:   []string{"nsId", "mcisId"},
			GivenParaVals:    []string{"ns-unit-01", "mock-unit-config01-dev"},
			GivenPostData: `{
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
			}`,
			ExpectStatus:         http.StatusCreated,
			ExpectBodyStartsWith: `{"id":"mock-unit-config01-dev"`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:             "add vmgroup to mcis(duplicate)",
			EchoFunc:         "RestPostMcisVmGroup",
			HttpMethod:       http.MethodPost,
			WhenURL:          "/tumblebug/ns/:nsId/mcis/:mcisId/vmgroup",
			GivenQueryParams: "",
			GivenParaNames:   []string{"nsId", "mcisId"},
			GivenParaVals:    []string{"ns-unit-01", "mock-unit-config01-dev"},
			GivenPostData: `{
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
			}`,
			ExpectStatus:         http.StatusCreated,
			ExpectBodyStartsWith: `{"id":"mock-unit-config01-dev"`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:             "create mcis policy",
			EchoFunc:         "RestPostMcisPolicy",
			HttpMethod:       http.MethodPost,
			WhenURL:          "/tumblebug/ns/:nsId/policy/mcis/:mcisId",
			GivenQueryParams: "",
			GivenParaNames:   []string{"nsId", "mcisId"},
			GivenParaVals:    []string{"ns-unit-01", "mock-unit-config01-dev"},
			GivenPostData: `{
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
			}`,
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"Name":"mock-unit-config01-dev"`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:             "create mcis policy(duplicate)",
			EchoFunc:         "RestPostMcisPolicy",
			HttpMethod:       http.MethodPost,
			WhenURL:          "/tumblebug/ns/:nsId/policy/mcis/:mcisId",
			GivenQueryParams: "",
			GivenParaNames:   []string{"nsId", "mcisId"},
			GivenParaVals:    []string{"ns-unit-01", "mock-unit-config01-dev"},
			GivenPostData: `{
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
			}`,
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `The MCIS Policy Obj mock-unit-config01-dev already exists.`,
		}
		EchoTest(t, tc)

		TearDownForRest()
	})

}
