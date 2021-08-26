package goscenario

import (
	"testing"

	api "github.com/cloud-barista/poc-cicd-tumblebug/src/api/grpc/request"
	core_common "github.com/cloud-barista/poc-cicd-tumblebug/src/core/common"
	core_mcir "github.com/cloud-barista/poc-cicd-tumblebug/src/core/mcir"
	core_mcis "github.com/cloud-barista/poc-cicd-tumblebug/src/core/mcis"
)

func TestGoFullParamArg(t *testing.T) {
	t.Run("go api full test for mock driver by parameter args style", func(t *testing.T) {
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
			Method:   "GetNSByParam",
			Args: []interface{}{
				"ns-unit-01",
			},
			ExpectResStartsWith: `{"id":"ns-unit-01"`,
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
			Name:     "list vnet",
			Instance: McirApi,
			Method:   "ListVNetByParam",
			Args: []interface{}{
				"ns-unit-01",
			},
			ExpectResStartsWith: `{"vNet":[{"id":"mock-unit-config01-dev"`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "list vnet id",
			Instance: McirApi,
			Method:   "ListVNetIdByParam",
			Args: []interface{}{
				"ns-unit-01",
			},
			ExpectResStartsWith: `{"idList":["mock-unit-config01-dev"]}`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "get vnet",
			Instance: McirApi,
			Method:   "GetVNetByParam",
			Args: []interface{}{
				"ns-unit-01",
				"mock-unit-config01-dev",
			},
			ExpectResStartsWith: `{"id":"mock-unit-config01-dev"`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "inspect vnet",
			Instance: TbutilApi,
			Method:   "InspectMcirResourcesByParam",
			Args: []interface{}{
				"mock-unit-config01",
				"vNet",
			},
			ExpectResStartsWith: `{"resourcesOnCsp":[{"id":"mock-unit-config01-dev"`,
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
			Name:     "list security",
			Instance: McirApi,
			Method:   "ListSecurityGroupByParam",
			Args: []interface{}{
				"ns-unit-01",
			},
			ExpectResStartsWith: `{"securityGroup":[{"id":"mock-unit-config01-dev"`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "list security id",
			Instance: McirApi,
			Method:   "ListSecurityGroupIdByParam",
			Args: []interface{}{
				"ns-unit-01",
			},
			ExpectResStartsWith: `{"idList":["mock-unit-config01-dev"]}`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "get security",
			Instance: McirApi,
			Method:   "GetSecurityGroupByParam",
			Args: []interface{}{
				"ns-unit-01",
				"mock-unit-config01-dev",
			},
			ExpectResStartsWith: `{"id":"mock-unit-config01-dev"`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "inspect security",
			Instance: TbutilApi,
			Method:   "InspectMcirResourcesByParam",
			Args: []interface{}{
				"mock-unit-config01",
				"securityGroup",
			},
			ExpectResStartsWith: `{"resourcesOnCsp":[{"id":"mock-unit-config01-dev"`,
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
			Name:     "list sshkey",
			Instance: McirApi,
			Method:   "ListSshKeyByParam",
			Args: []interface{}{
				"ns-unit-01",
			},
			ExpectResStartsWith: `{"sshKey":[{"id":"mock-unit-config01-dev"`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "list sshkey id",
			Instance: McirApi,
			Method:   "ListSshKeyIdByParam",
			Args: []interface{}{
				"ns-unit-01",
			},
			ExpectResStartsWith: `{"idList":["mock-unit-config01-dev"]}`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "get sshkey",
			Instance: McirApi,
			Method:   "GetSshKeyByParam",
			Args: []interface{}{
				"ns-unit-01",
				"mock-unit-config01-dev",
			},
			ExpectResStartsWith: `{"id":"mock-unit-config01-dev"`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "inspect sshkey",
			Instance: TbutilApi,
			Method:   "InspectMcirResourcesByParam",
			Args: []interface{}{
				"mock-unit-config01",
				"sshKey",
			},
			ExpectResStartsWith: `{"resourcesOnCsp":[{"id":"mock-unit-config01-dev"`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "list lookup image",
			Instance: McirApi,
			Method:   "ListLookupImageByParam",
			Args: []interface{}{
				"mock-unit-config01",
			},
			ExpectResStartsWith: `{"image":[{"IId":{"NameId":"mock-vmimage-01"`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "lookup image",
			Instance: McirApi,
			Method:   "GetLookupImageByParam",
			Args: []interface{}{
				"mock-unit-config01",
				"mock-vmimage-01",
			},
			ExpectResStartsWith: `{"IId":{"NameId":"mock-vmimage-01"`,
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
			Name:     "list image",
			Instance: McirApi,
			Method:   "ListImageByParam",
			Args: []interface{}{
				"ns-unit-01",
			},
			ExpectResStartsWith: `{"image":[{"namespace":"ns-unit-01","id":"mock-unit-config01-dev"`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "list image id",
			Instance: McirApi,
			Method:   "ListImageIdByParam",
			Args: []interface{}{
				"ns-unit-01",
			},
			ExpectResStartsWith: `{"idList":["mock-unit-config01-dev"]}`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "get image",
			Instance: McirApi,
			Method:   "GetImageByParam",
			Args: []interface{}{
				"ns-unit-01",
				"mock-unit-config01-dev",
			},
			ExpectResStartsWith: `{"namespace":"ns-unit-01","id":"mock-unit-config01-dev"`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "search image",
			Instance: McirApi,
			Method:   "SearchImageByParam",
			Args: []interface{}{
				&api.SearchImageQryRequest{
					NsId: "ns-unit-01",
					Keywords: []string{
						"mock",
					},
				},
			},
			ExpectResStartsWith: `{"image":[{"namespace":"ns-unit-01","id":"mock-unit-config01-dev"`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "fetch images",
			Instance: McirApi,
			Method:   "FetchImageByParam",
			Args: []interface{}{
				"!all",
				"ns-unit-01",
			},
			ExpectResStartsWith: `{"message":"Fetched 5 images (from 1 connConfigs)"}`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "list lookup spec",
			Instance: McirApi,
			Method:   "ListLookupSpecByParam",
			Args: []interface{}{
				"mock-unit-config01",
			},
			ExpectResStartsWith: `{"vmspec":[{"Region":"default","Name":"mock-vmspec-01"`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "lookup spec",
			Instance: McirApi,
			Method:   "GetLookupSpecByParam",
			Args: []interface{}{
				"mock-unit-config01",
				"mock-vmspec-01",
			},
			ExpectResStartsWith: `{"Region":"default","Name":"mock-vmspec-01"`,
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
			Name:     "list spec",
			Instance: McirApi,
			Method:   "ListSpecByParam",
			Args: []interface{}{
				"ns-unit-01",
			},
			ExpectResStartsWith: `{"spec":[{"namespace":"ns-unit-01","id":"mock-unit-config01-dev"`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "list spec id",
			Instance: McirApi,
			Method:   "ListSpecIdByParam",
			Args: []interface{}{
				"ns-unit-01",
			},
			ExpectResStartsWith: `{"idList":["mock-unit-config01-dev"]}`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "get spec",
			Instance: McirApi,
			Method:   "GetSpecByParam",
			Args: []interface{}{
				"ns-unit-01",
				"mock-unit-config01-dev",
			},
			ExpectResStartsWith: `{"namespace":"ns-unit-01","id":"mock-unit-config01-dev"`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "fetch specs",
			Instance: McirApi,
			Method:   "FetchSpecByParam",
			Args: []interface{}{
				"!all",
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
			Name:     "add vm to mcis",
			Instance: McisApi,
			Method:   "CreateMcisVMByParam",
			Args: []interface{}{
				&api.TbVmCreateRequest{
					NsId:   "ns-unit-01",
					McisId: "mock-unit-config01-dev",
					Item: api.TbVmInfo{
						Name:           "mock-unit-config01-dev",
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
			ExpectResStartsWith: `{"id":"mock-unit-config01-dev"`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "add vmgroup to mcis",
			Instance: McisApi,
			Method:   "CreateMcisVMGroupByParam",
			Args: []interface{}{
				&api.TbVmGroupCreateRequest{
					NsId:   "ns-unit-01",
					McisId: "mock-unit-config01-dev",
					Item: core_mcis.TbVmReq{
						VmGroupSize:    "3",
						Name:           "mock-unit-config01-dev",
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
			ExpectResStartsWith: `{"id":"mock-unit-config01-dev"`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "list mcis",
			Instance: McisApi,
			Method:   "ListMcisByParam",
			Args: []interface{}{
				"ns-unit-01",
			},
			ExpectResStartsWith: `{"mcis":[{"id":"mock-unit-config01-dev"`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "list mcis id",
			Instance: McisApi,
			Method:   "ListMcisIdByParam",
			Args: []interface{}{
				"ns-unit-01",
			},
			ExpectResStartsWith: `{"idList":["mock-unit-config01-dev"]}`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "list mcis status",
			Instance: McisApi,
			Method:   "ListMcisStatusByParam",
			Args: []interface{}{
				"ns-unit-01",
			},
			ExpectResStartsWith: `{"mcis":[{"id":"mock-unit-config01-dev"`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "get mcis",
			Instance: McisApi,
			Method:   "GetMcisInfoByParam",
			Args: []interface{}{
				"ns-unit-01",
				"mock-unit-config01-dev",
			},
			ExpectResStartsWith: `{"id":"mock-unit-config01-dev"`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "get mcis status",
			Instance: McisApi,
			Method:   "GetMcisStatusByParam",
			Args: []interface{}{
				"ns-unit-01",
				"mock-unit-config01-dev",
			},
			ExpectResStartsWith: `{"status":{"id":"mock-unit-config01-dev"`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "suspend mcis",
			Instance: McisApi,
			Method:   "ControlMcisByParam",
			Args: []interface{}{
				"ns-unit-01",
				"mock-unit-config01-dev",
				"suspend",
			},
			ExpectResStartsWith: `{"message":"Suspending the MCIS"}`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "resume mcis",
			Instance: McisApi,
			Method:   "ControlMcisByParam",
			Args: []interface{}{
				"ns-unit-01",
				"mock-unit-config01-dev",
				"resume",
			},
			ExpectResStartsWith: `{"message":"Resuming the MCIS"}`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "reboot mcis",
			Instance: McisApi,
			Method:   "ControlMcisByParam",
			Args: []interface{}{
				"ns-unit-01",
				"mock-unit-config01-dev",
				"reboot",
			},
			ExpectResStartsWith: `{"message":"Rebooting the MCIS"}`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "refine mcis",
			Instance: McisApi,
			Method:   "ControlMcisByParam",
			Args: []interface{}{
				"ns-unit-01",
				"mock-unit-config01-dev",
				"refine",
			},
			ExpectResStartsWith: `{"message":"Refined the MCIS"}`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "list vm id",
			Instance: McisApi,
			Method:   "ListMcisVmIdByParam",
			Args: []interface{}{
				"ns-unit-01",
				"mock-unit-config01-dev",
			},
			ExpectResStartsWith: `{"idList":["mock-unit-config01-dev","mock-unit-config01-dev-0"`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "inspect vm",
			Instance: TbutilApi,
			Method:   "InspectVmResourcesByParam",
			Args: []interface{}{
				"mock-unit-config01",
			},
			ExpectResStartsWith: `{"resourcesOnCsp":[{"id":"ns-unit-01-mock-unit-config01-dev-mock-unit-config01-dev"`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "create mcis policy",
			Instance: McisApi,
			Method:   "CreateMcisPolicyByParam",
			Args: []interface{}{
				&api.McisPolicyCreateRequest{
					NsId:   "ns-unit-01",
					McisId: "mock-unit-config01-dev",
					Item: core_mcis.McisPolicyInfo{
						Description: "Tumblebug Auto Control Demo",
						Policy: []core_mcis.Policy{
							core_mcis.Policy{
								AutoCondition: core_mcis.AutoCondition{
									Metric:           "cpu",
									Operator:         ">=",
									Operand:          "80",
									EvaluationPeriod: "10",
								},
								AutoAction: core_mcis.AutoAction{
									ActionType: "ScaleOut",
									Vm: core_mcis.TbVmInfo{
										Name: "AutoGen",
									},
									PostCommand: core_mcis.McisCmdReq{
										Command: "wget https://raw.githubusercontent.com/cloud-barista/cb-tumblebug/master/assets/scripts/setweb.sh -O ~/setweb.sh; chmod +x ~/setweb.sh; sudo ~/setweb.sh; wget https://raw.githubusercontent.com/cloud-barista/cb-tumblebug/master/assets/scripts/runLoadMaker.sh -O ~/runLoadMaker.sh; chmod +x ~/runLoadMaker.sh; sudo ~/runLoadMaker.sh",
									},
									PlacementAlgo: "random",
								},
							},
							core_mcis.Policy{
								AutoCondition: core_mcis.AutoCondition{
									Metric:           "cpu",
									Operator:         "<=",
									Operand:          "60",
									EvaluationPeriod: "10",
								},
								AutoAction: core_mcis.AutoAction{
									ActionType: "ScaleIn",
								},
							},
						},
					},
				},
			},
			ExpectResStartsWith: `{"Name":"mock-unit-config01-dev"`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "list mcis policy",
			Instance: McisApi,
			Method:   "ListMcisPolicyByParam",
			Args: []interface{}{
				"ns-unit-01",
			},
			ExpectResStartsWith: `{"mcisPolicy":[{"Name":"mock-unit-config01-dev"`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "get mcis policy",
			Instance: McisApi,
			Method:   "GetMcisPolicyByParam",
			Args: []interface{}{
				"ns-unit-01",
				"mock-unit-config01-dev",
			},
			ExpectResStartsWith: `{"Name":"mock-unit-config01-dev"`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "update config",
			Instance: TbutilApi,
			Method:   "CreateConfigByParam",
			Args: []interface{}{
				&core_common.ConfigReq{
					Name:  "key01",
					Value: "value01",
				},
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
			Method:   "GetConfigByParam",
			Args: []interface{}{
				"key01",
			},
			ExpectResStartsWith: `{"id":"key01","name":"key01","value":"value01"}`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "install agent",
			Instance: McisApi,
			Method:   "InstallMonitorAgentToMcisByParam",
			Args: []interface{}{
				&api.McisCmdCreateRequest{
					NsId:   "ns-unit-01",
					McisId: "mock-unit-config01-dev",
					Item: core_mcis.McisCmdReq{
						Command: "echo -n [CMD] Works! [Public IP: ; curl https://api.ipify.org ; echo -n ], [HostName: ; hostname ; echo -n ]",
					},
				},
			},
			ExpectResStartsWith: `{"result_array":[{"mcisId":"mock-unit-config01-dev"`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "get monitoring data",
			Instance: McisApi,
			Method:   "GetMonitorDataByParam",
			Args: []interface{}{
				"ns-unit-01",
				"mock-unit-config01-dev",
				"cpu",
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
			Method:   "GetConnConfigByParam",
			Args: []interface{}{
				"mock-unit-config01",
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
			Method:   "GetRegionByParam",
			Args: []interface{}{
				"mock-unit-region01",
			},
			ExpectResStartsWith: `{"RegionName":"mock-unit-region01","ProviderName":"MOCK"`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "command mcis",
			Instance: McisApi,
			Method:   "CmdMcisByParam",
			Args: []interface{}{
				&api.McisCmdCreateRequest{
					NsId:   "ns-unit-01",
					McisId: "mock-unit-config01-dev",
					Item: core_mcis.McisCmdReq{
						Command: "echo -n [CMD] Works! [Public IP: ; curl https://api.ipify.org ; echo -n ], [HostName: ; hostname ; echo -n ]",
					},
				},
			},
			ExpectResStartsWith: `{"result_array":[{"mcisId":"mock-unit-config01-dev"`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "command mcis vm",
			Instance: McisApi,
			Method:   "CmdMcisVmByParam",
			Args: []interface{}{
				&api.McisCmdVmCreateRequest{
					NsId:   "ns-unit-01",
					McisId: "mock-unit-config01-dev",
					VmId:   "mock-unit-config01-dev",
					Item: core_mcis.McisCmdReq{
						Command: "echo -n [CMD] Works! [Public IP: ; curl https://api.ipify.org ; echo -n ], [HostName: ; hostname ; echo -n ]",
					},
				},
			},
			ExpectResStartsWith: `{"Result":"echo -n [CMD] Works`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "list object",
			Instance: TbutilApi,
			Method:   "ListObjectByParam",
			Args: []interface{}{
				"",
			},
			ExpectResStartsWith: `{"object":["/config/key01"`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "list sub object",
			Instance: TbutilApi,
			Method:   "ListObjectByParam",
			Args: []interface{}{
				"/config",
			},
			ExpectResStartsWith: `{"object":["/config/key01"]}`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "get object",
			Instance: TbutilApi,
			Method:   "GetObjectByParam",
			Args: []interface{}{
				"/config/key01",
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
			Method:   "DeleteObjectByParam",
			Args: []interface{}{
				"/config/key01",
			},
			ExpectResStartsWith: `{"message":"The object has been deleted"}`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "delete all sub object",
			Instance: TbutilApi,
			Method:   "DeleteAllObjectByParam",
			Args: []interface{}{
				"/config",
			},
			ExpectResStartsWith: `{"message":"Objects have been deleted"}`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "init config",
			Instance: TbutilApi,
			Method:   "InitConfigByParam",
			Args: []interface{}{
				"key01",
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
			Method:   "DeleteMcisPolicyByParam",
			Args: []interface{}{
				"ns-unit-01",
				"mock-unit-config01-dev",
			},
			ExpectResStartsWith: `{"message":"Deleting the MCIS Policy info"}`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "delete all mcis policy",
			Instance: McisApi,
			Method:   "DeleteAllMcisPolicyByParam",
			Args: []interface{}{
				"ns-unit-01",
			},
			ExpectResStartsWith: `{"message":"No MCIS Policy to delete"}`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "terminate mcis",
			Instance: McisApi,
			Method:   "ControlMcisByParam",
			Args: []interface{}{
				"ns-unit-01",
				"mock-unit-config01-dev",
				"terminate",
			},
			ExpectResStartsWith: `{"message":"Terminating the MCIS"}`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "delete mcis",
			Instance: McisApi,
			Method:   "DeleteMcisByParam",
			Args: []interface{}{
				"ns-unit-01",
				"mock-unit-config01-dev",
			},
			ExpectResStartsWith: `{"message":"Deleting the MCIS mock-unit-config01-dev"}`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "delete all mcis",
			Instance: McisApi,
			Method:   "DeleteAllMcisByParam",
			Args: []interface{}{
				"ns-unit-01",
			},
			ExpectResStartsWith: `{"message":"No MCIS to delete"}`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "delete spec",
			Instance: McirApi,
			Method:   "DeleteSpecByParam",
			Args: []interface{}{
				"ns-unit-01",
				"mock-unit-config01-dev",
				"false",
			},
			ExpectResStartsWith: `{"message":"The spec mock-unit-config01-dev has been deleted"}`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "delete all spec",
			Instance: McirApi,
			Method:   "DeleteAllSpecByParam",
			Args: []interface{}{
				"ns-unit-01",
				"false",
			},
			ExpectResStartsWith: `{"message":"All specs has been deleted"}`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "delete image",
			Instance: McirApi,
			Method:   "DeleteImageByParam",
			Args: []interface{}{
				"ns-unit-01",
				"mock-unit-config01-dev",
				"false",
			},
			ExpectResStartsWith: `{"message":"The image mock-unit-config01-dev has been deleted"}`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "delete all image",
			Instance: McirApi,
			Method:   "DeleteAllImageByParam",
			Args: []interface{}{
				"ns-unit-01",
				"false",
			},
			ExpectResStartsWith: `{"message":"All images has been deleted"}`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "delete sshkey",
			Instance: McirApi,
			Method:   "DeleteSshKeyByParam",
			Args: []interface{}{
				"ns-unit-01",
				"mock-unit-config01-dev",
				"false",
			},
			ExpectResStartsWith: `{"message":"The sshKey mock-unit-config01-dev has been deleted"}`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "delete all sshkey",
			Instance: McirApi,
			Method:   "DeleteAllSshKeyByParam",
			Args: []interface{}{
				"ns-unit-01",
				"false",
			},
			ExpectResStartsWith: `{"message":"All sshKeys has been deleted"}`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "delete security",
			Instance: McirApi,
			Method:   "DeleteSecurityGroupByParam",
			Args: []interface{}{
				"ns-unit-01",
				"mock-unit-config01-dev",
				"false",
			},
			ExpectResStartsWith: `{"message":"The securityGroup mock-unit-config01-dev has been deleted"}`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "delete all security",
			Instance: McirApi,
			Method:   "DeleteAllSecurityGroupByParam",
			Args: []interface{}{
				"ns-unit-01",
				"false",
			},
			ExpectResStartsWith: `{"message":"All securityGroups has been deleted"}`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "delete vnet",
			Instance: McirApi,
			Method:   "DeleteVNetByParam",
			Args: []interface{}{
				"ns-unit-01",
				"mock-unit-config01-dev",
				"false",
			},
			ExpectResStartsWith: `{"message":"The vNet mock-unit-config01-dev has been deleted"}`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "delete all vnet",
			Instance: McirApi,
			Method:   "DeleteAllVNetByParam",
			Args: []interface{}{
				"ns-unit-01",
				"false",
			},
			ExpectResStartsWith: `{"message":"All vNets has been deleted"}`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "delete namespace",
			Instance: NsApi,
			Method:   "DeleteNSByParam",
			Args: []interface{}{
				"ns-unit-01",
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

		TearDownForGrpc()
	})

}
