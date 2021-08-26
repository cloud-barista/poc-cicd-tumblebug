package restscenario

import (
	"net/http"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestRestFull(t *testing.T) {
	t.Run("rest api full test for mock driver", func(t *testing.T) {
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
			Name:                 "list namespace",
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
		EchoTest(t, tc)

		tc = TestCases{
			Name:                 "get namespace",
			EchoFunc:             "RestGetNs",
			HttpMethod:           http.MethodGet,
			WhenURL:              "/tumblebug/ns/:nsId",
			GivenQueryParams:     "",
			GivenParaNames:       []string{"nsId"},
			GivenParaVals:        []string{"ns-unit-01"},
			GivenPostData:        "",
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"id":"ns-unit-01"`,
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
			Name:                 "list vnet",
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
		EchoTest(t, tc)

		tc = TestCases{
			Name:                 "list vnet id",
			EchoFunc:             "RestGetAllResources",
			HttpMethod:           http.MethodGet,
			WhenURL:              "/tumblebug/ns/:nsId/resources/vNet",
			GivenQueryParams:     "?option=id",
			GivenParaNames:       []string{"nsId"},
			GivenParaVals:        []string{"ns-unit-01"},
			GivenPostData:        "",
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"idList":["mock-unit-config01-dev"]}`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:                 "get vnet",
			EchoFunc:             "RestGetResource",
			HttpMethod:           http.MethodGet,
			WhenURL:              "/tumblebug/ns/:nsId/resources/vNet/:resourceId",
			GivenQueryParams:     "",
			GivenParaNames:       []string{"nsId", "resourceId"},
			GivenParaVals:        []string{"ns-unit-01", "mock-unit-config01-dev"},
			GivenPostData:        "",
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"id":"mock-unit-config01-dev"`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:             "inspect vnet",
			EchoFunc:         "RestInspectResources",
			HttpMethod:       http.MethodPost,
			WhenURL:          "/tumblebug/inspectResources",
			GivenQueryParams: "",
			GivenParaNames:   nil,
			GivenParaVals:    nil,
			GivenPostData: `{
				"connectionName": "mock-unit-config01",
				"type": "vNet"				
			}`,
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"resourcesOnCsp":[{"id":"mock-unit-config01-dev"`,
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
			Name:                 "list security",
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
		EchoTest(t, tc)

		tc = TestCases{
			Name:                 "list security id",
			EchoFunc:             "RestGetAllResources",
			HttpMethod:           http.MethodGet,
			WhenURL:              "/tumblebug/ns/:nsId/resources/securityGroup",
			GivenQueryParams:     "?option=id",
			GivenParaNames:       []string{"nsId"},
			GivenParaVals:        []string{"ns-unit-01"},
			GivenPostData:        "",
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"idList":["mock-unit-config01-dev"]}`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:                 "get security",
			EchoFunc:             "RestGetResource",
			HttpMethod:           http.MethodGet,
			WhenURL:              "/tumblebug/ns/:nsId/resources/securityGroup/:resourceId",
			GivenQueryParams:     "",
			GivenParaNames:       []string{"nsId", "resourceId"},
			GivenParaVals:        []string{"ns-unit-01", "mock-unit-config01-dev"},
			GivenPostData:        "",
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"id":"mock-unit-config01-dev"`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:             "inspect security",
			EchoFunc:         "RestInspectResources",
			HttpMethod:       http.MethodPost,
			WhenURL:          "/tumblebug/inspectResources",
			GivenQueryParams: "",
			GivenParaNames:   nil,
			GivenParaVals:    nil,
			GivenPostData: `{
				"connectionName": "mock-unit-config01",
				"type": "securityGroup"				
			}`,
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"resourcesOnCsp":[{"id":"mock-unit-config01-dev"`,
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
			Name:                 "list sshkey",
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
		EchoTest(t, tc)

		tc = TestCases{
			Name:                 "list sshkey id",
			EchoFunc:             "RestGetAllResources",
			HttpMethod:           http.MethodGet,
			WhenURL:              "/tumblebug/ns/:nsId/resources/sshKey",
			GivenQueryParams:     "?option=id",
			GivenParaNames:       []string{"nsId"},
			GivenParaVals:        []string{"ns-unit-01"},
			GivenPostData:        "",
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"idList":["mock-unit-config01-dev"]}`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:                 "get sshkey",
			EchoFunc:             "RestGetResource",
			HttpMethod:           http.MethodGet,
			WhenURL:              "/tumblebug/ns/:nsId/resources/sshKey/:resourceId",
			GivenQueryParams:     "",
			GivenParaNames:       []string{"nsId", "resourceId"},
			GivenParaVals:        []string{"ns-unit-01", "mock-unit-config01-dev"},
			GivenPostData:        "",
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"id":"mock-unit-config01-dev"`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:             "inspect sshkey",
			EchoFunc:         "RestInspectResources",
			HttpMethod:       http.MethodPost,
			WhenURL:          "/tumblebug/inspectResources",
			GivenQueryParams: "",
			GivenParaNames:   nil,
			GivenParaVals:    nil,
			GivenPostData: `{
				"connectionName": "mock-unit-config01",
				"type": "sshKey"				
			}`,
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"resourcesOnCsp":[{"id":"mock-unit-config01-dev"`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:             "list lookup image",
			EchoFunc:         "RestLookupImageList",
			HttpMethod:       http.MethodGet,
			WhenURL:          "/tumblebug/lookupImages",
			GivenQueryParams: "",
			GivenParaNames:   nil,
			GivenParaVals:    nil,
			GivenPostData: `{
				"connectionName": "mock-unit-config01"
			}`,
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"image":[{"Name":"","IId":{"NameId":"mock-vmimage-01"`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:             "lookup image",
			EchoFunc:         "RestLookupImage",
			HttpMethod:       http.MethodGet,
			WhenURL:          "/tumblebug/lookupImage",
			GivenQueryParams: "",
			GivenParaNames:   nil,
			GivenParaVals:    nil,
			GivenPostData: `{
				"connectionName": "mock-unit-config01",
				"cspImageId": "mock-vmimage-01"		
			}`,
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"Name":"","IId":{"NameId":"mock-vmimage-01"`,
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
			Name:                 "list image",
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
		EchoTest(t, tc)

		tc = TestCases{
			Name:                 "list image id",
			EchoFunc:             "RestGetAllResources",
			HttpMethod:           http.MethodGet,
			WhenURL:              "/tumblebug/ns/:nsId/resources/image",
			GivenQueryParams:     "?option=id",
			GivenParaNames:       []string{"nsId"},
			GivenParaVals:        []string{"ns-unit-01"},
			GivenPostData:        "",
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"idList":["mock-unit-config01-dev"]}`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:                 "get image",
			EchoFunc:             "RestGetResource",
			HttpMethod:           http.MethodGet,
			WhenURL:              "/tumblebug/ns/:nsId/resources/image/:resourceId",
			GivenQueryParams:     "",
			GivenParaNames:       []string{"nsId", "resourceId"},
			GivenParaVals:        []string{"ns-unit-01", "mock-unit-config01-dev"},
			GivenPostData:        "",
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"namespace":"ns-unit-01","id":"mock-unit-config01-dev"`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:             "search image",
			EchoFunc:         "RestSearchImage",
			HttpMethod:       http.MethodPost,
			WhenURL:          "/tumblebug/ns/:nsId/resources/searchImage",
			GivenQueryParams: "",
			GivenParaNames:   []string{"nsId"},
			GivenParaVals:    []string{"ns-unit-01"},
			GivenPostData: `{
				"keywords": [
						"mock"
					]
			}`,
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"image":[{"namespace":"ns-unit-01","id":"mock-unit-config01-dev"`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:                 "fetch images",
			EchoFunc:             "RestFetchImages",
			HttpMethod:           http.MethodPost,
			WhenURL:              "/tumblebug/ns/:nsId/resources/fetchImages",
			GivenQueryParams:     "",
			GivenParaNames:       []string{"nsId"},
			GivenParaVals:        []string{"ns-unit-01"},
			GivenPostData:        "",
			ExpectStatus:         http.StatusCreated,
			ExpectBodyStartsWith: `{"message":"Fetched 5 images (from 1 connConfigs)"}`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:             "list lookup spec",
			EchoFunc:         "RestLookupSpecList",
			HttpMethod:       http.MethodGet,
			WhenURL:          "/tumblebug/lookupSpecs",
			GivenQueryParams: "",
			GivenParaNames:   nil,
			GivenParaVals:    nil,
			GivenPostData: `{
				"connectionName": "mock-unit-config01"
			}`,
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"vmspec":[{"Region":"default","Name":"mock-vmspec-01"`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:             "lookup spec",
			EchoFunc:         "RestLookupSpec",
			HttpMethod:       http.MethodGet,
			WhenURL:          "/tumblebug/lookupSpec",
			GivenQueryParams: "",
			GivenParaNames:   nil,
			GivenParaVals:    nil,
			GivenPostData: `{
				"connectionName": "mock-unit-config01",
				"cspSpecName": "mock-vmspec-01"				
			}`,
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"Region":"default","Name":"mock-vmspec-01"`,
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
			Name:                 "list spec",
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
		EchoTest(t, tc)

		tc = TestCases{
			Name:                 "list spec id",
			EchoFunc:             "RestGetAllResources",
			HttpMethod:           http.MethodGet,
			WhenURL:              "/tumblebug/ns/:nsId/resources/spec",
			GivenQueryParams:     "?option=id",
			GivenParaNames:       []string{"nsId"},
			GivenParaVals:        []string{"ns-unit-01"},
			GivenPostData:        "",
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"idList":["mock-unit-config01-dev"]}`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:                 "get spec",
			EchoFunc:             "RestGetResource",
			HttpMethod:           http.MethodGet,
			WhenURL:              "/tumblebug/ns/:nsId/resources/spec/:resourceId",
			GivenQueryParams:     "",
			GivenParaNames:       []string{"nsId", "resourceId"},
			GivenParaVals:        []string{"ns-unit-01", "mock-unit-config01-dev"},
			GivenPostData:        "",
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"namespace":"ns-unit-01","id":"mock-unit-config01-dev"`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:             "update spec",
			EchoFunc:         "RestPutSpec",
			HttpMethod:       http.MethodPut,
			WhenURL:          "/tumblebug/ns/:nsId/resources/spec/:specId",
			GivenQueryParams: "",
			GivenParaNames:   []string{"nsId", "specId"},
			GivenParaVals:    []string{"ns-unit-01", "mock-unit-config01-dev"},
			GivenPostData: `{
				"id": "mock-unit-config01-dev", 
				"description": "UpdateSpec() test"
			}`,
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"namespace":"ns-unit-01","id":"mock-unit-config01-dev"`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:             "filter spec",
			EchoFunc:         "RestFilterSpecs",
			HttpMethod:       http.MethodPost,
			WhenURL:          "/tumblebug/ns/:nsId/resources/filterSpecs",
			GivenQueryParams: "",
			GivenParaNames:   []string{"nsId"},
			GivenParaVals:    []string{"ns-unit-01"},
			GivenPostData: `{
		    "num_vCPU": 4, 
		    "mem_GiB": 32
			}`,
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"spec":[{"namespace":"ns-unit-01","id":"mock-unit-config01-dev"`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:             "range filter spec",
			EchoFunc:         "RestFilterSpecsByRange",
			HttpMethod:       http.MethodPost,
			WhenURL:          "/tumblebug/ns/:nsId/resources/filterSpecsByRange",
			GivenQueryParams: "",
			GivenParaNames:   []string{"nsId"},
			GivenParaVals:    []string{"ns-unit-01"},
			GivenPostData: `{
		    "mem_GiB": {
			    "min": 4
		    }
			}`,
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"spec":[{"namespace":"ns-unit-01","id":"mock-unit-config01-dev"`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:                 "fetch specs",
			EchoFunc:             "RestFetchSpecs",
			HttpMethod:           http.MethodPost,
			WhenURL:              "/tumblebug/ns/:nsId/resources/fetchSpecs",
			GivenQueryParams:     "",
			GivenParaNames:       []string{"nsId"},
			GivenParaVals:        []string{"ns-unit-01"},
			GivenPostData:        "",
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
			Name:                 "list mcis",
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
		EchoTest(t, tc)

		tc = TestCases{
			Name:                 "list mcis id",
			EchoFunc:             "RestGetAllMcis",
			HttpMethod:           http.MethodGet,
			WhenURL:              "/tumblebug/ns/:nsId/mcis",
			GivenQueryParams:     "?option=id",
			GivenParaNames:       []string{"nsId"},
			GivenParaVals:        []string{"ns-unit-01"},
			GivenPostData:        "",
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"idList":["mock-unit-config01-dev"]}`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:                 "list mcis status",
			EchoFunc:             "RestGetAllMcis",
			HttpMethod:           http.MethodGet,
			WhenURL:              "/tumblebug/ns/:nsId/mcis",
			GivenQueryParams:     "?option=status",
			GivenParaNames:       []string{"nsId"},
			GivenParaVals:        []string{"ns-unit-01"},
			GivenPostData:        "",
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"mcis":[{"id":"mock-unit-config01-dev"`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:                 "get mcis",
			EchoFunc:             "RestGetMcis",
			HttpMethod:           http.MethodGet,
			WhenURL:              "/tumblebug/ns/:nsId/mcis/:mcisId",
			GivenQueryParams:     "",
			GivenParaNames:       []string{"nsId", "mcisId"},
			GivenParaVals:        []string{"ns-unit-01", "mock-unit-config01-dev"},
			GivenPostData:        "",
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"id":"mock-unit-config01-dev"`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:                 "get mcis status",
			EchoFunc:             "RestGetMcis",
			HttpMethod:           http.MethodGet,
			WhenURL:              "/tumblebug/ns/:nsId/mcis/:mcisId",
			GivenQueryParams:     "?action=status",
			GivenParaNames:       []string{"nsId", "mcisId"},
			GivenParaVals:        []string{"ns-unit-01", "mock-unit-config01-dev"},
			GivenPostData:        "",
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"status":{"id":"mock-unit-config01-dev"`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:                 "suspend mcis",
			EchoFunc:             "RestGetMcis",
			HttpMethod:           http.MethodGet,
			WhenURL:              "/tumblebug/ns/:nsId/mcis/:mcisId",
			GivenQueryParams:     "?action=suspend",
			GivenParaNames:       []string{"nsId", "mcisId"},
			GivenParaVals:        []string{"ns-unit-01", "mock-unit-config01-dev"},
			GivenPostData:        "",
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"message":"Suspending the MCIS"}`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:                 "resume mcis",
			EchoFunc:             "RestGetMcis",
			HttpMethod:           http.MethodGet,
			WhenURL:              "/tumblebug/ns/:nsId/mcis/:mcisId",
			GivenQueryParams:     "?action=resume",
			GivenParaNames:       []string{"nsId", "mcisId"},
			GivenParaVals:        []string{"ns-unit-01", "mock-unit-config01-dev"},
			GivenPostData:        "",
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"message":"Resuming the MCIS"}`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:                 "reboot mcis",
			EchoFunc:             "RestGetMcis",
			HttpMethod:           http.MethodGet,
			WhenURL:              "/tumblebug/ns/:nsId/mcis/:mcisId",
			GivenQueryParams:     "?action=reboot",
			GivenParaNames:       []string{"nsId", "mcisId"},
			GivenParaVals:        []string{"ns-unit-01", "mock-unit-config01-dev"},
			GivenPostData:        "",
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"message":"Rebooting the MCIS"}`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:                 "refine mcis",
			EchoFunc:             "RestGetMcis",
			HttpMethod:           http.MethodGet,
			WhenURL:              "/tumblebug/ns/:nsId/mcis/:mcisId",
			GivenQueryParams:     "?action=refine",
			GivenParaNames:       []string{"nsId", "mcisId"},
			GivenParaVals:        []string{"ns-unit-01", "mock-unit-config01-dev"},
			GivenPostData:        "",
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"message":"Refined the MCIS"}`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:                 "list vm id",
			EchoFunc:             "RestGetMcis",
			HttpMethod:           http.MethodGet,
			WhenURL:              "/tumblebug/ns/:nsId/mcis/:mcisId",
			GivenQueryParams:     "?option=id",
			GivenParaNames:       []string{"nsId", "mcisId"},
			GivenParaVals:        []string{"ns-unit-01", "mock-unit-config01-dev"},
			GivenPostData:        "",
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"idList":["mock-unit-config01-dev","mock-unit-config01-dev-0"`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:             "inspect vm",
			EchoFunc:         "RestInspectResources",
			HttpMethod:       http.MethodPost,
			WhenURL:          "/tumblebug/inspectResources",
			GivenQueryParams: "",
			GivenParaNames:   nil,
			GivenParaVals:    nil,
			GivenPostData: `{
				"connectionName": "mock-unit-config01",
				"type": "vm"				
			}`,
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"resourcesOnCsp":[{"id":"ns-unit-01-mock-unit-config01-dev-mock-unit-config01-dev"`,
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
			Name:                 "list mcis policy",
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
		EchoTest(t, tc)

		tc = TestCases{
			Name:                 "get mcis policy",
			EchoFunc:             "RestGetMcisPolicy",
			HttpMethod:           http.MethodGet,
			WhenURL:              "/tumblebug/ns/:nsId/policy/mcis/:mcisId",
			GivenQueryParams:     "",
			GivenParaNames:       []string{"nsId", "mcisId"},
			GivenParaVals:        []string{"ns-unit-01", "mock-unit-config01-dev"},
			GivenPostData:        "",
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"Name":"mock-unit-config01-dev"`,
		}
		EchoTest(t, tc)

		tc = TestCases{
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
		EchoTest(t, tc)

		tc = TestCases{
			Name:                 "list config",
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
		EchoTest(t, tc)

		tc = TestCases{
			Name:                 "get config",
			EchoFunc:             "RestGetConfig",
			HttpMethod:           http.MethodGet,
			WhenURL:              "/tumblebug/config/:configId",
			GivenQueryParams:     "",
			GivenParaNames:       []string{"configId"},
			GivenParaVals:        []string{"key01"},
			GivenPostData:        "",
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"id":"key01","name":"key01","value":"value01"}`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:             "install agent",
			EchoFunc:         "RestPostInstallMonitorAgentToMcis",
			HttpMethod:       http.MethodPost,
			WhenURL:          "/tumblebug/ns/:nsId/monitoring/install/mcis/:mcisId",
			GivenQueryParams: "",
			GivenParaNames:   []string{"nsId", "mcisId"},
			GivenParaVals:    []string{"ns-unit-01", "mock-unit-config01-dev"},
			GivenPostData: `{
				"command": "echo -n [CMD] Works! [Public IP: ; curl https://api.ipify.org ; echo -n ], [HostName: ; hostname ; echo -n ]"
			}`,
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"result_array":[{"mcisId":"mock-unit-config01-dev"`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:                 "get monitoring data",
			EchoFunc:             "RestGetMonitorData",
			HttpMethod:           http.MethodPost,
			WhenURL:              "/tumblebug/ns/:nsId/monitoring/mcis/:mcisId/metric/:metric",
			GivenQueryParams:     "",
			GivenParaNames:       []string{"nsId", "mcisId", "metric"},
			GivenParaVals:        []string{"ns-unit-01", "mock-unit-config01-dev", "cpu"},
			GivenPostData:        "",
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"nsId":"ns-unit-01","mcisId":"mock-unit-config01-dev","mcisMonitoring":[{"metric"`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:                 "list connection config",
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
		EchoTest(t, tc)

		tc = TestCases{
			Name:                 "get connection config",
			EchoFunc:             "RestGetConnConfig",
			HttpMethod:           http.MethodGet,
			WhenURL:              "/tumblebug/connConfig/:connConfigName",
			GivenQueryParams:     "",
			GivenParaNames:       []string{"connConfigName"},
			GivenParaVals:        []string{"mock-unit-config01"},
			GivenPostData:        "",
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"ConfigName":"mock-unit-config01","ProviderName":"MOCK"`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:                 "list region",
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
		EchoTest(t, tc)

		tc = TestCases{
			Name:                 "get region",
			EchoFunc:             "RestGetRegion",
			HttpMethod:           http.MethodGet,
			WhenURL:              "/tumblebug/region/:regionName",
			GivenQueryParams:     "",
			GivenParaNames:       []string{"regionName"},
			GivenParaVals:        []string{"mock-unit-region01"},
			GivenPostData:        "",
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"RegionName":"mock-unit-region01","ProviderName":"MOCK"`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:             "command mcis",
			EchoFunc:         "RestPostCmdMcis",
			HttpMethod:       http.MethodPost,
			WhenURL:          "/tumblebug/ns/:nsId/cmd/mcis/:mcisId",
			GivenQueryParams: "",
			GivenParaNames:   []string{"nsId", "mcisId"},
			GivenParaVals:    []string{"ns-unit-01", "mock-unit-config01-dev"},
			GivenPostData: `{
				"command": "echo -n [CMD] Works! [Public IP: ; curl https://api.ipify.org ; echo -n ], [HostName: ; hostname ; echo -n ]"
			}`,
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"result_array":[{"mcisId":"mock-unit-config01-dev"`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:             "command mcis vm",
			EchoFunc:         "RestPostCmdMcisVm",
			HttpMethod:       http.MethodPost,
			WhenURL:          "/tumblebug/ns/:nsId/cmd/mcis/:mcisId/vm/:vmId",
			GivenQueryParams: "",
			GivenParaNames:   []string{"nsId", "mcisId", "vmId"},
			GivenParaVals:    []string{"ns-unit-01", "mock-unit-config01-dev", "mock-unit-config01-dev"},
			GivenPostData: `{
				"command": "echo -n [CMD] Works! [Public IP: ; curl https://api.ipify.org ; echo -n ], [HostName: ; hostname ; echo -n ]"
			}`,
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"result":"echo -n [CMD] Works`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:                 "list object",
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
		EchoTest(t, tc)

		tc = TestCases{
			Name:                 "get object",
			EchoFunc:             "RestGetObject",
			HttpMethod:           http.MethodGet,
			WhenURL:              "/tumblebug/object",
			GivenQueryParams:     "?key=/config/key01",
			GivenParaNames:       nil,
			GivenParaVals:        nil,
			GivenPostData:        "",
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"id":"key01","name":"key01","value":"value01"}`,
		}
		EchoTest(t, tc)

		//
		// Delete Resources
		//

		tc = TestCases{
			Name:                 "delete object",
			EchoFunc:             "RestDeleteObject",
			HttpMethod:           http.MethodDelete,
			WhenURL:              "/tumblebug/object",
			GivenQueryParams:     "?key=/config/key01",
			GivenParaNames:       nil,
			GivenParaVals:        nil,
			GivenPostData:        "",
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"message":"The object has been deleted"}`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:                 "delete all sub object",
			EchoFunc:             "RestDeleteObjects",
			HttpMethod:           http.MethodDelete,
			WhenURL:              "/tumblebug/objects",
			GivenQueryParams:     "?key=/config",
			GivenParaNames:       nil,
			GivenParaVals:        nil,
			GivenPostData:        "",
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"message":"Objects have been deleted"}`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:                 "init config",
			EchoFunc:             "RestInitConfig",
			HttpMethod:           http.MethodDelete,
			WhenURL:              "/tumblebug/config/:configId",
			GivenQueryParams:     "",
			GivenParaNames:       []string{"configId"},
			GivenParaVals:        []string{"key01"},
			GivenPostData:        "",
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"message":"The config key01 has been initialized."}`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:                 "init all config",
			EchoFunc:             "RestInitAllConfig",
			HttpMethod:           http.MethodDelete,
			WhenURL:              "/tumblebug/config",
			GivenQueryParams:     "",
			GivenParaNames:       nil,
			GivenParaVals:        nil,
			GivenPostData:        "",
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"message":"All configs has been initialized."}`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:                 "delete mcis policy",
			EchoFunc:             "RestDelMcisPolicy",
			HttpMethod:           http.MethodDelete,
			WhenURL:              "/tumblebug/ns/:nsId/policy/mcis/:mcisId",
			GivenQueryParams:     "",
			GivenParaNames:       []string{"nsId", "mcisId"},
			GivenParaVals:        []string{"ns-unit-01", "mock-unit-config01-dev"},
			GivenPostData:        "",
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"message":"Deleting the MCIS Policy info"}`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:                 "delete all mcis policy",
			EchoFunc:             "RestDelAllMcisPolicy",
			HttpMethod:           http.MethodDelete,
			WhenURL:              "/tumblebug/ns/:nsId/policy/mcis",
			GivenQueryParams:     "",
			GivenParaNames:       []string{"nsId"},
			GivenParaVals:        []string{"ns-unit-01"},
			GivenPostData:        "",
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"message":"No MCIS Policy to delete"}`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:                 "terminate mcis",
			EchoFunc:             "RestGetMcis",
			HttpMethod:           http.MethodGet,
			WhenURL:              "/tumblebug/ns/:nsId/mcis/:mcisId",
			GivenQueryParams:     "?action=terminate",
			GivenParaNames:       []string{"nsId", "mcisId"},
			GivenParaVals:        []string{"ns-unit-01", "mock-unit-config01-dev"},
			GivenPostData:        "",
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"message":"Terminating the MCIS"}`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:                 "delete mcis",
			EchoFunc:             "RestDelMcis",
			HttpMethod:           http.MethodDelete,
			WhenURL:              "/tumblebug/ns/:nsId/mcis/:mcisId",
			GivenQueryParams:     "",
			GivenParaNames:       []string{"nsId", "mcisId"},
			GivenParaVals:        []string{"ns-unit-01", "mock-unit-config01-dev"},
			GivenPostData:        "",
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"message":"Deleting the MCIS mock-unit-config01-dev"}`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:                 "delete all mcis",
			EchoFunc:             "RestDelAllMcis",
			HttpMethod:           http.MethodDelete,
			WhenURL:              "/tumblebug/ns/:nsId/mcis",
			GivenQueryParams:     "",
			GivenParaNames:       []string{"nsId"},
			GivenParaVals:        []string{"ns-unit-01"},
			GivenPostData:        "",
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"message":"No MCIS to delete"}`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:                 "delete spec",
			EchoFunc:             "RestDelResource",
			HttpMethod:           http.MethodDelete,
			WhenURL:              "/tumblebug/ns/:nsId/resources/spec/:resourceId",
			GivenQueryParams:     "",
			GivenParaNames:       []string{"nsId", "resourceId"},
			GivenParaVals:        []string{"ns-unit-01", "mock-unit-config01-dev"},
			GivenPostData:        "",
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"message":"The spec mock-unit-config01-dev has been deleted"}`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:                 "delete all spec",
			EchoFunc:             "RestDelAllResources",
			HttpMethod:           http.MethodDelete,
			WhenURL:              "/tumblebug/ns/:nsId/resources/spec",
			GivenQueryParams:     "",
			GivenParaNames:       []string{"nsId"},
			GivenParaVals:        []string{"ns-unit-01"},
			GivenPostData:        "",
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"message":"All specs has been deleted"}`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:                 "delete image",
			EchoFunc:             "RestDelResource",
			HttpMethod:           http.MethodDelete,
			WhenURL:              "/tumblebug/ns/:nsId/resources/image/:resourceId",
			GivenQueryParams:     "",
			GivenParaNames:       []string{"nsId", "resourceId"},
			GivenParaVals:        []string{"ns-unit-01", "mock-unit-config01-dev"},
			GivenPostData:        "",
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"message":"The image mock-unit-config01-dev has been deleted"}`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:                 "delete all image",
			EchoFunc:             "RestDelAllResources",
			HttpMethod:           http.MethodDelete,
			WhenURL:              "/tumblebug/ns/:nsId/resources/image",
			GivenQueryParams:     "",
			GivenParaNames:       []string{"nsId"},
			GivenParaVals:        []string{"ns-unit-01"},
			GivenPostData:        "",
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"message":"All images has been deleted"}`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:                 "delete sshkey",
			EchoFunc:             "RestDelResource",
			HttpMethod:           http.MethodDelete,
			WhenURL:              "/tumblebug/ns/:nsId/resources/sshKey/:resourceId",
			GivenQueryParams:     "",
			GivenParaNames:       []string{"nsId", "resourceId"},
			GivenParaVals:        []string{"ns-unit-01", "mock-unit-config01-dev"},
			GivenPostData:        "",
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"message":"The sshKey mock-unit-config01-dev has been deleted"}`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:                 "delete all sshkey",
			EchoFunc:             "RestDelAllResources",
			HttpMethod:           http.MethodDelete,
			WhenURL:              "/tumblebug/ns/:nsId/resources/sshKey",
			GivenQueryParams:     "",
			GivenParaNames:       []string{"nsId"},
			GivenParaVals:        []string{"ns-unit-01"},
			GivenPostData:        "",
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"message":"All sshKeys has been deleted"}`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:                 "delete security",
			EchoFunc:             "RestDelResource",
			HttpMethod:           http.MethodDelete,
			WhenURL:              "/tumblebug/ns/:nsId/resources/securityGroup/:resourceId",
			GivenQueryParams:     "",
			GivenParaNames:       []string{"nsId", "resourceId"},
			GivenParaVals:        []string{"ns-unit-01", "mock-unit-config01-dev"},
			GivenPostData:        "",
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"message":"The securityGroup mock-unit-config01-dev has been deleted"}`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:                 "delete all security",
			EchoFunc:             "RestDelAllResources",
			HttpMethod:           http.MethodDelete,
			WhenURL:              "/tumblebug/ns/:nsId/resources/securityGroup",
			GivenQueryParams:     "",
			GivenParaNames:       []string{"nsId"},
			GivenParaVals:        []string{"ns-unit-01"},
			GivenPostData:        "",
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"message":"All securityGroups has been deleted"}`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:                 "delete vnet",
			EchoFunc:             "RestDelResource",
			HttpMethod:           http.MethodDelete,
			WhenURL:              "/tumblebug/ns/:nsId/resources/vNet/:resourceId",
			GivenQueryParams:     "",
			GivenParaNames:       []string{"nsId", "resourceId"},
			GivenParaVals:        []string{"ns-unit-01", "mock-unit-config01-dev"},
			GivenPostData:        "",
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"message":"The vNet mock-unit-config01-dev has been deleted"}`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:                 "delete all vnet",
			EchoFunc:             "RestDelAllResources",
			HttpMethod:           http.MethodDelete,
			WhenURL:              "/tumblebug/ns/:nsId/resources/vNet",
			GivenQueryParams:     "",
			GivenParaNames:       []string{"nsId"},
			GivenParaVals:        []string{"ns-unit-01"},
			GivenPostData:        "",
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"message":"All vNets has been deleted"}`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:                 "delete namespace",
			EchoFunc:             "RestDelNs",
			HttpMethod:           http.MethodDelete,
			WhenURL:              "/tumblebug/ns/:nsId",
			GivenQueryParams:     "",
			GivenParaNames:       []string{"nsId"},
			GivenParaVals:        []string{"ns-unit-01"},
			GivenPostData:        "",
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"message":"The ns ns-unit-01 has been deleted"}`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:                 "delete all namespace",
			EchoFunc:             "RestDelAllNs",
			HttpMethod:           http.MethodDelete,
			WhenURL:              "/tumblebug/ns",
			GivenQueryParams:     "",
			GivenParaNames:       nil,
			GivenParaVals:        nil,
			GivenPostData:        "",
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"message":"All namespaces has been deleted"}`,
		}
		EchoTest(t, tc)

		TearDownForRest()
	})

}
