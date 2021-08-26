package restscenario

import (
	"net/http"
	"testing"
)

func TestRestMisc(t *testing.T) {
	t.Run("rest api misc test for mock driver", func(t *testing.T) {
		SetUpForRest()

		tc := TestCases{
			Name:                 "health",
			EchoFunc:             "RestGetHealth",
			HttpMethod:           http.MethodGet,
			WhenURL:              "/tumblebug/health",
			GivenQueryParams:     "",
			GivenParaNames:       nil,
			GivenParaVals:        nil,
			GivenPostData:        "",
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"message":"API server of CB-Tumblebug is alive"}`,
		}
		EchoTest(t, tc)

		tc = TestCases{
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
			Name:                 "list namespace id",
			EchoFunc:             "RestGetAllNs",
			HttpMethod:           http.MethodGet,
			WhenURL:              "/tumblebug/ns",
			GivenQueryParams:     "?option=id",
			GivenParaNames:       nil,
			GivenParaVals:        nil,
			GivenPostData:        "",
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"idList":["ns-unit-01"]}`,
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
			Name:                 "check resource",
			EchoFunc:             "RestCheckResource",
			HttpMethod:           http.MethodGet,
			WhenURL:              "/tumblebug/ns/:nsId/checkResource/:resourceType/:resourceId",
			GivenQueryParams:     "",
			GivenParaNames:       []string{"nsId", "resourceType", "resourceId"},
			GivenParaVals:        []string{"ns-unit-01", "vNet", "mock-unit-config01-dev"},
			GivenPostData:        "",
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"exists":true}`,
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
			Name:             "connection fetch images",
			EchoFunc:         "RestFetchImages",
			HttpMethod:       http.MethodPost,
			WhenURL:          "/tumblebug/ns/:nsId/resources/fetchImages",
			GivenQueryParams: "",
			GivenParaNames:   []string{"nsId"},
			GivenParaVals:    []string{"ns-unit-01"},
			GivenPostData: `{
				"connectionName": "mock-unit-config01"
			}`,
			ExpectStatus:         http.StatusCreated,
			ExpectBodyStartsWith: `{"message":"Fetched 5 images (from 1 connConfigs)"}`,
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
			Name:             "connection fetch specs",
			EchoFunc:         "RestFetchSpecs",
			HttpMethod:       http.MethodPost,
			WhenURL:          "/tumblebug/ns/:nsId/resources/fetchSpecs",
			GivenQueryParams: "",
			GivenParaNames:   []string{"nsId"},
			GivenParaVals:    []string{"ns-unit-01"},
			GivenPostData: `{
				"connectionName": "mock-unit-config01"
			}`,
			ExpectStatus:         http.StatusCreated,
			ExpectBodyStartsWith: `{"message":"Fetched 4 specs (from 1 connConfigs)"}`,
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
			Name:                 "get mcis vm",
			EchoFunc:             "RestGetMcisVm",
			HttpMethod:           http.MethodGet,
			WhenURL:              "/tumblebug/ns/:nsId/mcis/:mcisId/vm/:vmId",
			GivenQueryParams:     "",
			GivenParaNames:       []string{"nsId", "mcisId", "vmId"},
			GivenParaVals:        []string{"ns-unit-01", "mock-unit-config01-dev", "mock-unit-config01-dev-01"},
			GivenPostData:        "",
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"id":"mock-unit-config01-dev-01"`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:                 "get mcis vm status",
			EchoFunc:             "RestGetMcisVm",
			HttpMethod:           http.MethodGet,
			WhenURL:              "/tumblebug/ns/:nsId/mcis/:mcisId/vm/:vmId",
			GivenQueryParams:     "?action=status",
			GivenParaNames:       []string{"nsId", "mcisId", "vmId"},
			GivenParaVals:        []string{"ns-unit-01", "mock-unit-config01-dev", "mock-unit-config01-dev-01"},
			GivenPostData:        "",
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"id":"mock-unit-config01-dev-01"`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:                 "suspend mcis vm",
			EchoFunc:             "RestGetMcisVm",
			HttpMethod:           http.MethodGet,
			WhenURL:              "/tumblebug/ns/:nsId/mcis/:mcisId/vm/:vmId",
			GivenQueryParams:     "?action=suspend",
			GivenParaNames:       []string{"nsId", "mcisId", "vmId"},
			GivenParaVals:        []string{"ns-unit-01", "mock-unit-config01-dev", "mock-unit-config01-dev-01"},
			GivenPostData:        "",
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"message":"Suspending the VM"}`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:                 "resume mcis vm",
			EchoFunc:             "RestGetMcisVm",
			HttpMethod:           http.MethodGet,
			WhenURL:              "/tumblebug/ns/:nsId/mcis/:mcisId/vm/:vmId",
			GivenQueryParams:     "?action=resume",
			GivenParaNames:       []string{"nsId", "mcisId", "vmId"},
			GivenParaVals:        []string{"ns-unit-01", "mock-unit-config01-dev", "mock-unit-config01-dev-01"},
			GivenPostData:        "",
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"message":"Resuming the VM"}`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:                 "reboot mcis vm",
			EchoFunc:             "RestGetMcisVm",
			HttpMethod:           http.MethodGet,
			WhenURL:              "/tumblebug/ns/:nsId/mcis/:mcisId/vm/:vmId",
			GivenQueryParams:     "?action=reboot",
			GivenParaNames:       []string{"nsId", "mcisId", "vmId"},
			GivenParaVals:        []string{"ns-unit-01", "mock-unit-config01-dev", "mock-unit-config01-dev-01"},
			GivenPostData:        "",
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"message":"Rebooting the VM"}`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:             "all benchmark",
			EchoFunc:         "RestGetAllBenchmark",
			HttpMethod:       http.MethodPost,
			WhenURL:          "/tumblebug/ns/:nsId/benchmarkall/mcis/:mcisId",
			GivenQueryParams: "",
			GivenParaNames:   []string{"nsId", "mcisId"},
			GivenParaVals:    []string{"ns-unit-01", "mock-unit-config01-dev"},
			GivenPostData: `{
				"host": "localhost"
			}`,
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"resultarray":[{"result":"1.0","unit":"unit","desc":"mrtt complete"`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:             "benchmark",
			EchoFunc:         "RestGetBenchmark",
			HttpMethod:       http.MethodPost,
			WhenURL:          "/tumblebug/ns/:nsId/benchmark/mcis/:mcisId",
			GivenQueryParams: "?action=cpus",
			GivenParaNames:   []string{"nsId", "mcisId"},
			GivenParaVals:    []string{"ns-unit-01", "mock-unit-config01-dev"},
			GivenPostData: `{
				"host": "localhost"
			}`,
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"resultarray":[{"result":"1.0","unit":"unit","desc":"cpus complete"`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:                 "check mcis",
			EchoFunc:             "RestCheckMcis",
			HttpMethod:           http.MethodGet,
			WhenURL:              "/tumblebug/ns/:nsId/checkMcis/:mcisId",
			GivenQueryParams:     "",
			GivenParaNames:       []string{"nsId", "mcisId"},
			GivenParaVals:        []string{"ns-unit-01", "mock-unit-config01-dev"},
			GivenPostData:        "",
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"exists":true}`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:                 "check mcis vm",
			EchoFunc:             "RestCheckVm",
			HttpMethod:           http.MethodGet,
			WhenURL:              "/tumblebug/ns/:nsId/mcis/:mcisId/checkVm/:vmId",
			GivenQueryParams:     "",
			GivenParaNames:       []string{"nsId", "mcisId", "vmId"},
			GivenParaVals:        []string{"ns-unit-01", "mock-unit-config01-dev", "mock-unit-config01-dev-01"},
			GivenPostData:        "",
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"exists":true}`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:                 "test add object association",
			EchoFunc:             "RestTestAddObjectAssociation",
			HttpMethod:           http.MethodPut,
			WhenURL:              "/tumblebug/ns/:nsId/testAddObjectAssociation/:resourceType/:resourceId",
			GivenQueryParams:     "",
			GivenParaNames:       []string{"nsId", "resourceType", "resourceId"},
			GivenParaVals:        []string{"ns-unit-01", "vNet", "mock-unit-config01-dev"},
			GivenPostData:        "",
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `["/ns/ns-unit-01/mcis/mock-unit-config01-dev/vm/mock-unit-config01-dev-01","/test/vm/key"]`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:                 "test delete object association",
			EchoFunc:             "RestTestDeleteObjectAssociation",
			HttpMethod:           http.MethodPut,
			WhenURL:              "/tumblebug/ns/:nsId/testDeleteObjectAssociation/:resourceType/:resourceId",
			GivenQueryParams:     "",
			GivenParaNames:       []string{"nsId", "resourceType", "resourceId"},
			GivenParaVals:        []string{"ns-unit-01", "vNet", "mock-unit-config01-dev"},
			GivenPostData:        "",
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `["/ns/ns-unit-01/mcis/mock-unit-config01-dev/vm/mock-unit-config01-dev-01"]`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:                 "test count object association",
			EchoFunc:             "RestTestGetAssociatedObjectCount",
			HttpMethod:           http.MethodGet,
			WhenURL:              "/tumblebug/ns/:nsId/testGetAssociatedObjectCount/:resourceType/:resourceId",
			GivenQueryParams:     "",
			GivenParaNames:       []string{"nsId", "resourceType", "resourceId"},
			GivenParaVals:        []string{"ns-unit-01", "vNet", "mock-unit-config01-dev"},
			GivenPostData:        "",
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"associatedObjectCount":1}`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:                 "terminate mcis vm",
			EchoFunc:             "RestGetMcisVm",
			HttpMethod:           http.MethodGet,
			WhenURL:              "/tumblebug/ns/:nsId/mcis/:mcisId/vm/:vmId",
			GivenQueryParams:     "?action=terminate",
			GivenParaNames:       []string{"nsId", "mcisId", "vmId"},
			GivenParaVals:        []string{"ns-unit-01", "mock-unit-config01-dev", "mock-unit-config01-dev-01"},
			GivenPostData:        "",
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"message":"Terminating the VM"}`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:                 "delete mcis vm",
			EchoFunc:             "RestDelMcisVm",
			HttpMethod:           http.MethodDelete,
			WhenURL:              "/tumblebug/ns/:nsId/mcis/:mcisId/vm/:vmId",
			GivenQueryParams:     "",
			GivenParaNames:       []string{"nsId", "mcisId", "vmId"},
			GivenParaVals:        []string{"ns-unit-01", "mock-unit-config01-dev", "mock-unit-config01-dev-01"},
			GivenPostData:        "",
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"message":"Deleting the VM info"}`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:             "create mcis(installMonAgent:yes,vmGroupSize>0)",
			EchoFunc:         "RestPostMcis",
			HttpMethod:       http.MethodPost,
			WhenURL:          "/tumblebug/ns/:nsId/mcis",
			GivenQueryParams: "",
			GivenParaNames:   []string{"nsId"},
			GivenParaVals:    []string{"ns-unit-01"},
			GivenPostData: `{
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
			}`,
			ExpectStatus:         http.StatusCreated,
			ExpectBodyStartsWith: `{"id":"mock-unit-config01-dev2"`,
		}
		EchoTest(t, tc)

		TearDownForRest()
	})

}
