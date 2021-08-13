package cliscenario

import (
	"testing"
)

func TestCliMiscJson(t *testing.T) {
	t.Run("command misc json in/out test for mock driver", func(t *testing.T) {
		SetUpForCli()

		tc := TestCases{
			Name: "create namespace",
			CmdArgs: []string{"namespace", "create", "--config", "../conf/grpc_conf.yaml", "-i", "json", "-o", "json", "-d", `{
				"name": "ns-unit-01",
				"description": "NameSpace for General Testing"
			}`},
			ExpectResStartsWith: `{"id":"ns-unit-01"`,
		}
		TumblebugCmdTest(t, tc)

		tc = TestCases{
			Name:                "list namespace id",
			CmdArgs:             []string{"namespace", "list-id", "--config", "../conf/grpc_conf.yaml", "-i", "json", "-o", "json"},
			ExpectResStartsWith: `{"idList":["ns-unit-01"]}`,
		}
		TumblebugCmdTest(t, tc)

		tc = TestCases{
			Name: "create vnet",
			CmdArgs: []string{"network", "create", "--config", "../conf/grpc_conf.yaml", "-i", "json", "-o", "json", "-d", `{
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
			}`},
			ExpectResStartsWith: `{"id":"mock-unit-config01-dev"`,
		}
		TumblebugCmdTest(t, tc)

		tc = TestCases{
			Name: "create security",
			CmdArgs: []string{"securitygroup", "create", "--config", "../conf/grpc_conf.yaml", "-i", "json", "-o", "json", "-d", `{
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
			}`},
			ExpectResStartsWith: `{"id":"mock-unit-config01-dev"`,
		}
		TumblebugCmdTest(t, tc)

		tc = TestCases{
			Name: "create sshkey",
			CmdArgs: []string{"keypair", "create", "--config", "../conf/grpc_conf.yaml", "-i", "json", "-o", "json", "-d", `{
				"nsId":  "ns-unit-01",
				"sshKey": {
					"name": "mock-unit-config01-dev",
					"connectionName": "mock-unit-config01",
					"description": ""	
				}	
			}`},
			ExpectResStartsWith: `{"id":"mock-unit-config01-dev"`,
		}
		TumblebugCmdTest(t, tc)

		tc = TestCases{
			Name: "register image with id",
			CmdArgs: []string{"image", "create-id", "--config", "../conf/grpc_conf.yaml", "-i", "json", "-o", "json", "-d", `{
				"nsId":  "ns-unit-01",
				"image": {
					"connectionName": "mock-unit-config01",
					"name": "mock-unit-config01-dev",
					"cspImageId": "mock-vmimage-01",
					"description": "Canonical, Ubuntu, 18.04 LTS, amd64 bionic"
				}
			}`},
			ExpectResStartsWith: `{"namespace":"ns-unit-01","id":"mock-unit-config01-dev"`,
		}
		TumblebugCmdTest(t, tc)

		tc = TestCases{
			Name:                "connection fetch images",
			CmdArgs:             []string{"image", "fetch", "--config", "../conf/grpc_conf.yaml", "-i", "json", "-o", "json", "--cc", "mock-unit-config01", "--ns", "ns-unit-01"},
			ExpectResStartsWith: `{"message":"Fetched 5 images (from 1 connConfigs)"}`,
		}
		TumblebugCmdTest(t, tc)

		tc = TestCases{
			Name: "register spec",
			CmdArgs: []string{"spec", "create-id", "--config", "../conf/grpc_conf.yaml", "-i", "json", "-o", "json", "-d", `{
				"nsId":  "ns-unit-01",
				"spec": {
					"connectionName": "mock-unit-config01",
					"name": "mock-unit-config01-dev",
					"cspSpecName": "mock-vmspec-01"		
				}	
			}`},
			ExpectResStartsWith: `{"namespace":"ns-unit-01","id":"mock-unit-config01-dev"`,
		}
		TumblebugCmdTest(t, tc)

		tc = TestCases{
			Name:                "connection fetch specs",
			CmdArgs:             []string{"spec", "fetch", "--config", "../conf/grpc_conf.yaml", "-i", "json", "-o", "json", "--cc", "mock-unit-config01", "--ns", "ns-unit-01"},
			ExpectResStartsWith: `{"message":"Fetched 4 specs (from 1 connConfigs)"}`,
		}
		TumblebugCmdTest(t, tc)

		tc = TestCases{
			Name: "create mcis",
			CmdArgs: []string{"mcis", "create", "--config", "../conf/grpc_conf.yaml", "-i", "json", "-o", "json", "-d", `{
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
			}`},
			ExpectResStartsWith: `{"id":"mock-unit-config01-dev"`,
		}
		TumblebugCmdTest(t, tc)

		tc = TestCases{
			Name:                "get mcis vm",
			CmdArgs:             []string{"mcis", "get-vm", "--config", "../conf/grpc_conf.yaml", "-i", "json", "-o", "json", "--ns", "ns-unit-01", "--mcis", "mock-unit-config01-dev", "--vm", "mock-unit-config01-dev-01"},
			ExpectResStartsWith: `{"id":"mock-unit-config01-dev-01"`,
		}
		TumblebugCmdTest(t, tc)

		tc = TestCases{
			Name:                "get mcis vm status",
			CmdArgs:             []string{"mcis", "status-vm", "--config", "../conf/grpc_conf.yaml", "-i", "json", "-o", "json", "--ns", "ns-unit-01", "--mcis", "mock-unit-config01-dev", "--vm", "mock-unit-config01-dev-01"},
			ExpectResStartsWith: `{"id":"mock-unit-config01-dev-01"`,
		}
		TumblebugCmdTest(t, tc)

		tc = TestCases{
			Name:                "suspend mcis vm",
			CmdArgs:             []string{"mcis", "suspend-vm", "--config", "../conf/grpc_conf.yaml", "-i", "json", "-o", "json", "--ns", "ns-unit-01", "--mcis", "mock-unit-config01-dev", "--vm", "mock-unit-config01-dev-01"},
			ExpectResStartsWith: `{"message":"Suspending the VM"}`,
		}
		TumblebugCmdTest(t, tc)

		tc = TestCases{
			Name:                "resume mcis vm",
			CmdArgs:             []string{"mcis", "resume-vm", "--config", "../conf/grpc_conf.yaml", "-i", "json", "-o", "json", "--ns", "ns-unit-01", "--mcis", "mock-unit-config01-dev", "--vm", "mock-unit-config01-dev-01"},
			ExpectResStartsWith: `{"message":"Resuming the VM"}`,
		}
		TumblebugCmdTest(t, tc)

		tc = TestCases{
			Name:                "reboot mcis vm",
			CmdArgs:             []string{"mcis", "reboot-vm", "--config", "../conf/grpc_conf.yaml", "-i", "json", "-o", "json", "--ns", "ns-unit-01", "--mcis", "mock-unit-config01-dev", "--vm", "mock-unit-config01-dev-01"},
			ExpectResStartsWith: `{"message":"Rebooting the VM"}`,
		}
		TumblebugCmdTest(t, tc)

		tc = TestCases{
			Name:                "all benchmark",
			CmdArgs:             []string{"mcis", "benchmark", "--config", "../conf/grpc_conf.yaml", "-i", "json", "-o", "json", "--ns", "ns-unit-01", "--mcis", "mock-unit-config01-dev", "--host", "localhost"},
			ExpectResStartsWith: `{"resultarray":[{"result":"1.0","unit":"unit","desc":"mrtt complete"`,
		}
		TumblebugCmdTest(t, tc)

		tc = TestCases{
			Name:                "benchmark",
			CmdArgs:             []string{"mcis", "benchmark", "--config", "../conf/grpc_conf.yaml", "-i", "json", "-o", "json", "--ns", "ns-unit-01", "--mcis", "mock-unit-config01-dev", "--host", "localhost", "--action", "cpus"},
			ExpectResStartsWith: `{"resultarray":[{"result":"1.0","unit":"unit","desc":"cpus complete"`,
		}
		TumblebugCmdTest(t, tc)

		tc = TestCases{
			Name:                "terminate mcis vm",
			CmdArgs:             []string{"mcis", "terminate-vm", "--config", "../conf/grpc_conf.yaml", "-i", "json", "-o", "json", "--ns", "ns-unit-01", "--mcis", "mock-unit-config01-dev", "--vm", "mock-unit-config01-dev-01"},
			ExpectResStartsWith: `{"message":"Terminating the VM"}`,
		}
		TumblebugCmdTest(t, tc)

		tc = TestCases{
			Name:                "delete mcis vm",
			CmdArgs:             []string{"mcis", "del-vm", "--config", "../conf/grpc_conf.yaml", "-i", "json", "-o", "json", "--ns", "ns-unit-01", "--mcis", "mock-unit-config01-dev", "--vm", "mock-unit-config01-dev-01"},
			ExpectResStartsWith: `{"message":"Deleting the VM info"}`,
		}
		TumblebugCmdTest(t, tc)

		tc = TestCases{
			Name: "create mcis(installMonAgent:yes,vmGroupSize>0)",
			CmdArgs: []string{"mcis", "create", "--config", "../conf/grpc_conf.yaml", "-i", "json", "-o", "json", "-d", `{
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
			}`},
			ExpectResStartsWith: `{"id":"mock-unit-config01-dev2"`,
		}
		TumblebugCmdTest(t, tc)

		TearDownForCli()
	})
}
