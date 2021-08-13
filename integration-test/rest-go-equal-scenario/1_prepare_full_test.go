package rgequalscenario

import (
	"net/http"
	"testing"

	gs "github.com/cloud-barista/poc-cicd-tumblebug/integration-test/go-scenario"
	rs "github.com/cloud-barista/poc-cicd-tumblebug/integration-test/rest-scenario"
)

func TestPrepareFull(t *testing.T) {
	t.Run("prepare full test for mock driver", func(t *testing.T) {
		gs.SetUpForGrpc()

		tc := rs.TestCases{
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
		rs.EchoTest(t, tc)

		tc = rs.TestCases{
			Name:                 "list namespace for rest",
			EchoFunc:             "RestGetAllNs",
			HttpMethod:           http.MethodGet,
			WhenURL:              "/tumblebug/ns",
			GivenQueryParams:     "",
			GivenParaNames:       nil,
			GivenParaVals:        nil,
			GivenPostData:        "",
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"ns":[`,
		}
		res, _ := rs.EchoTest(t, tc)
		RestPrepareResult["list namespace"] = res

		gtc := gs.TestCases{
			Name:                "list namespace for grpc",
			Instance:            gs.NsApi,
			Method:              "ListNS",
			Args:                nil,
			ExpectResStartsWith: `{"ns":[`,
		}
		res, _ = gs.MethodTest(t, gtc)
		GrpcPrepareResult["list namespace"] = res

		tc = rs.TestCases{
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
		rs.EchoTest(t, tc)

		tc = rs.TestCases{
			Name:                 "list vnet for rest",
			EchoFunc:             "RestGetAllResources",
			HttpMethod:           http.MethodGet,
			WhenURL:              "/tumblebug/ns/:nsId/resources/vNet",
			GivenQueryParams:     "",
			GivenParaNames:       []string{"nsId"},
			GivenParaVals:        []string{"ns-unit-01"},
			GivenPostData:        "",
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"vNet":[{"id":"mock-unit-config01-dev"`,
		}
		res, _ = rs.EchoTest(t, tc)
		RestPrepareResult["list vnet"] = res

		gtc = gs.TestCases{
			Name:     "list vnet for grpc",
			Instance: gs.McirApi,
			Method:   "ListVNet",
			Args: []interface{}{
				`{ "nsId": "ns-unit-01", "resourceType":"vNet"}`,
			},
			ExpectResStartsWith: `{"vNet":[{"id":"mock-unit-config01-dev"`,
		}
		res, _ = gs.MethodTest(t, gtc)
		GrpcPrepareResult["list vnet"] = res

		tc = rs.TestCases{
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
		rs.EchoTest(t, tc)

		tc = rs.TestCases{
			Name:                 "list security for rest",
			EchoFunc:             "RestGetAllResources",
			HttpMethod:           http.MethodGet,
			WhenURL:              "/tumblebug/ns/:nsId/resources/securityGroup",
			GivenQueryParams:     "",
			GivenParaNames:       []string{"nsId"},
			GivenParaVals:        []string{"ns-unit-01"},
			GivenPostData:        "",
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"securityGroup":[{"id":"mock-unit-config01-dev"`,
		}
		res, _ = rs.EchoTest(t, tc)
		RestPrepareResult["list security"] = res

		gtc = gs.TestCases{
			Name:     "list security for grpc",
			Instance: gs.McirApi,
			Method:   "ListSecurityGroup",
			Args: []interface{}{
				`{ "nsId": "ns-unit-01", "resourceType":"securityGroup"}`,
			},
			ExpectResStartsWith: `{"securityGroup":[{"id":"mock-unit-config01-dev"`,
		}
		res, _ = gs.MethodTest(t, gtc)
		GrpcPrepareResult["list security"] = res

		tc = rs.TestCases{
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
		rs.EchoTest(t, tc)

		tc = rs.TestCases{
			Name:                 "list sshkey for rest",
			EchoFunc:             "RestGetAllResources",
			HttpMethod:           http.MethodGet,
			WhenURL:              "/tumblebug/ns/:nsId/resources/sshKey",
			GivenQueryParams:     "",
			GivenParaNames:       []string{"nsId"},
			GivenParaVals:        []string{"ns-unit-01"},
			GivenPostData:        "",
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"sshKey":[{"id":"mock-unit-config01-dev"`,
		}
		res, _ = rs.EchoTest(t, tc)
		RestPrepareResult["list sshkey"] = res

		gtc = gs.TestCases{
			Name:     "list sshkey for grpc",
			Instance: gs.McirApi,
			Method:   "ListSshKey",
			Args: []interface{}{
				`{ "nsId": "ns-unit-01", "resourceType":"sshKey"}`,
			},
			ExpectResStartsWith: `{"sshKey":[{"id":"mock-unit-config01-dev"`,
		}
		res, _ = gs.MethodTest(t, gtc)
		GrpcPrepareResult["list sshkey"] = res

		tc = rs.TestCases{
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
		rs.EchoTest(t, tc)

		tc = rs.TestCases{
			Name:                 "list image for rest",
			EchoFunc:             "RestGetAllResources",
			HttpMethod:           http.MethodGet,
			WhenURL:              "/tumblebug/ns/:nsId/resources/image",
			GivenQueryParams:     "",
			GivenParaNames:       []string{"nsId"},
			GivenParaVals:        []string{"ns-unit-01"},
			GivenPostData:        "",
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"image":[{"namespace":"ns-unit-01","id":"mock-unit-config01-dev"`,
		}
		res, _ = rs.EchoTest(t, tc)
		RestPrepareResult["list image"] = res

		gtc = gs.TestCases{
			Name:     "list image for grpc",
			Instance: gs.McirApi,
			Method:   "ListImage",
			Args: []interface{}{
				`{ "nsId": "ns-unit-01", "resourceType":"image"}`,
			},
			ExpectResStartsWith: `{"image":[{"namespace":"ns-unit-01","id":"mock-unit-config01-dev"`,
		}
		res, _ = gs.MethodTest(t, gtc)
		GrpcPrepareResult["list image"] = res

		tc = rs.TestCases{
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
		rs.EchoTest(t, tc)

		tc = rs.TestCases{
			Name:                 "list spec for rest",
			EchoFunc:             "RestGetAllResources",
			HttpMethod:           http.MethodGet,
			WhenURL:              "/tumblebug/ns/:nsId/resources/spec",
			GivenQueryParams:     "",
			GivenParaNames:       []string{"nsId"},
			GivenParaVals:        []string{"ns-unit-01"},
			GivenPostData:        "",
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"spec":[{"namespace":"ns-unit-01","id":"mock-unit-config01-dev"`,
		}
		res, _ = rs.EchoTest(t, tc)
		RestPrepareResult["list spec"] = res

		gtc = gs.TestCases{
			Name:     "list spec for grpc",
			Instance: gs.McirApi,
			Method:   "ListSpec",
			Args: []interface{}{
				`{ "nsId": "ns-unit-01", "resourceType":"spec"}`,
			},
			ExpectResStartsWith: `{"spec":[{"namespace":"ns-unit-01","id":"mock-unit-config01-dev"`,
		}
		res, _ = gs.MethodTest(t, gtc)
		GrpcPrepareResult["list spec"] = res

		tc = rs.TestCases{
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
		rs.EchoTest(t, tc)

		tc = rs.TestCases{
			Name:                 "list mcis for rest",
			EchoFunc:             "RestGetAllMcis",
			HttpMethod:           http.MethodGet,
			WhenURL:              "/tumblebug/ns/:nsId/mcis",
			GivenQueryParams:     "",
			GivenParaNames:       []string{"nsId"},
			GivenParaVals:        []string{"ns-unit-01"},
			GivenPostData:        "",
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"mcis":[{"id":"mock-unit-config01-dev"`,
		}
		res, _ = rs.EchoTest(t, tc)
		RestPrepareResult["list mcis"] = res

		gtc = gs.TestCases{
			Name:     "list mcis for grpc",
			Instance: gs.McisApi,
			Method:   "ListMcis",
			Args: []interface{}{
				`{ "nsId": "ns-unit-01" }`,
			},
			ExpectResStartsWith: `{"mcis":[{"id":"mock-unit-config01-dev"`,
		}
		res, _ = gs.MethodTest(t, gtc)
		GrpcPrepareResult["list mcis"] = res

		tc = rs.TestCases{
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
		rs.EchoTest(t, tc)

		tc = rs.TestCases{
			Name:                 "list mcis policy for rest",
			EchoFunc:             "RestGetAllMcisPolicy",
			HttpMethod:           http.MethodGet,
			WhenURL:              "/tumblebug/ns/:nsId/policy/mcis",
			GivenQueryParams:     "",
			GivenParaNames:       []string{"nsId"},
			GivenParaVals:        []string{"ns-unit-01"},
			GivenPostData:        "",
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"mcisPolicy":[{"Name":"mock-unit-config01-dev"`,
		}
		res, _ = rs.EchoTest(t, tc)
		RestPrepareResult["list mcis policy"] = res

		gtc = gs.TestCases{
			Name:     "list mcis policy for grpc",
			Instance: gs.McisApi,
			Method:   "ListMcisPolicy",
			Args: []interface{}{
				`{ "nsId": "ns-unit-01" }`,
			},
			ExpectResStartsWith: `{"mcisPolicy":[{"Name":"mock-unit-config01-dev"`,
		}
		res, _ = gs.MethodTest(t, gtc)
		GrpcPrepareResult["list mcis policy"] = res

		tc = rs.TestCases{
			Name:             "update config",
			EchoFunc:         "RestPostConfig",
			HttpMethod:       http.MethodPost,
			WhenURL:          "/tumblebug/config",
			GivenQueryParams: "",
			GivenParaNames:   nil,
			GivenParaVals:    nil,
			GivenPostData: `{
					"name": "key01",
					"value": "value01"
				}`,
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"id":"key01","name":"key01","value":"value01"}`,
		}
		rs.EchoTest(t, tc)

		tc = rs.TestCases{
			Name:                 "list config for rest",
			EchoFunc:             "RestGetAllConfig",
			HttpMethod:           http.MethodGet,
			WhenURL:              "/tumblebug/config",
			GivenQueryParams:     "",
			GivenParaNames:       nil,
			GivenParaVals:        nil,
			GivenPostData:        "",
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"config":[{"id":"key01","name":"key01","value":"value01"}]}`,
		}
		res, _ = rs.EchoTest(t, tc)
		RestPrepareResult["list config"] = res

		gtc = gs.TestCases{
			Name:                "list config for grpc",
			Instance:            gs.TbutilApi,
			Method:              "ListConfig",
			Args:                nil,
			ExpectResStartsWith: `{"config":[{"id":"key01","name":"key01","value":"value01"}]}`,
		}
		res, _ = gs.MethodTest(t, gtc)
		GrpcPrepareResult["list config"] = res

		tc = rs.TestCases{
			Name:                 "list connection config for rest",
			EchoFunc:             "RestGetConnConfigList",
			HttpMethod:           http.MethodGet,
			WhenURL:              "/tumblebug/connConfig",
			GivenQueryParams:     "",
			GivenParaNames:       nil,
			GivenParaVals:        nil,
			GivenPostData:        "",
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"connectionconfig":[{"ConfigName":"mock-unit-config01","ProviderName":"MOCK"`,
		}
		res, _ = rs.EchoTest(t, tc)
		RestPrepareResult["list connection config"] = res

		gtc = gs.TestCases{
			Name:                "list connection config for grpc",
			Instance:            gs.TbutilApi,
			Method:              "ListConnConfig",
			Args:                nil,
			ExpectResStartsWith: `{"connectionconfig":[{"ConfigName":"mock-unit-config01","ProviderName":"MOCK"`,
		}
		res, _ = gs.MethodTest(t, gtc)
		GrpcPrepareResult["list connection config"] = res

		tc = rs.TestCases{
			Name:                 "list region for rest",
			EchoFunc:             "RestGetRegionList",
			HttpMethod:           http.MethodGet,
			WhenURL:              "/tumblebug/region",
			GivenQueryParams:     "",
			GivenParaNames:       nil,
			GivenParaVals:        nil,
			GivenPostData:        "",
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"region":[{"RegionName":"mock-unit-region01","ProviderName":"MOCK"`,
		}
		res, _ = rs.EchoTest(t, tc)
		RestPrepareResult["list region"] = res

		gtc = gs.TestCases{
			Name:                "list region for grpc",
			Instance:            gs.TbutilApi,
			Method:              "ListRegion",
			Args:                nil,
			ExpectResStartsWith: `{"region":[{"RegionName":"mock-unit-region01","ProviderName":"MOCK"`,
		}
		res, _ = gs.MethodTest(t, gtc)
		GrpcPrepareResult["list region"] = res

		tc = rs.TestCases{
			Name:                 "list object for rest",
			EchoFunc:             "RestGetObjects",
			HttpMethod:           http.MethodGet,
			WhenURL:              "/tumblebug/objects",
			GivenQueryParams:     "",
			GivenParaNames:       nil,
			GivenParaVals:        nil,
			GivenPostData:        "",
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"object":["/config/key01"`,
		}
		res, _ = rs.EchoTest(t, tc)
		RestPrepareResult["list object"] = res

		gtc = gs.TestCases{
			Name:     "list object for grpc",
			Instance: gs.TbutilApi,
			Method:   "ListObject",
			Args: []interface{}{
				`{ "key": "" }`,
			},
			ExpectResStartsWith: `{"object":["/config/key01"`,
		}
		res, _ = gs.MethodTest(t, gtc)
		GrpcPrepareResult["list object"] = res

		gs.TearDownForGrpc()
	})

}
