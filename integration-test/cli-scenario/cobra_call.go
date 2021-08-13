package cliscenario

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/cloud-barista/poc-cicd-tumblebug/src/api/grpc/cbadm/cmd"
	"github.com/stretchr/testify/assert"
)

func TumblebugCmdTest(t *testing.T, tc TestCases) (string, error) {

	var (
		res string = ""
		err error  = nil
	)

	t.Run(tc.Name, func(t *testing.T) {

		tumblebugCmd := cmd.NewRootCmd()
		b := bytes.NewBufferString("")
		e := bytes.NewBufferString("")
		tumblebugCmd.SetOut(b)
		tumblebugCmd.SetErr(e)
		tumblebugCmd.SetArgs(tc.CmdArgs)
		tumblebugCmd.Execute()

		out, err := ioutil.ReadAll(e)
		if assert.NoError(t, err) && string(out) == "" {
			out, err = ioutil.ReadAll(b)
		}

		if assert.NoError(t, err) {
			if strings.HasPrefix(string(out), "{") {
				dst := new(bytes.Buffer)
				err = json.Compact(dst, out)
				if assert.NoError(t, err) {
					res = dst.String()
				}
			} else {
				res = string(out)
			}

			if tc.ExpectResStartsWith != "" {
				if !assert.True(t, strings.HasPrefix(res, tc.ExpectResStartsWith)) {
					fmt.Fprintf(os.Stderr, "\n                Not Equal: \n"+
						"                  Expected Start With: %s\n"+
						"                  Actual  : %s\n", tc.ExpectResStartsWith, res)
				}
			}
			if tc.ExpectResContains != "" {
				if !assert.True(t, strings.Contains(res, tc.ExpectResContains)) {
					fmt.Fprintf(os.Stderr, "\n                Not Equal: \n"+
						"                  Expected Contains: %s\n"+
						"                  Actual  : %s\n", tc.ExpectResContains, res)
				}
			}
			if tc.ExpectResStartsWith == "" && tc.ExpectResContains == "" {
				if !assert.True(t, "" == res) {
					fmt.Fprintf(os.Stderr, "\n                Not Equal: \n"+
						"      Expected StartWith/Contains: %s\n"+
						"      Actual  : %s\n", tc.ExpectResStartsWith, res)
				}
			}
		}

	})

	return res, err
}
