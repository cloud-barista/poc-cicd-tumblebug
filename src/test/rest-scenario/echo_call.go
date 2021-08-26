package restscenario

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"strings"
	"testing"

	rest_common "github.com/cloud-barista/poc-cicd-tumblebug/src/api/rest/server/common"
	rest_mcir "github.com/cloud-barista/poc-cicd-tumblebug/src/api/rest/server/mcir"
	rest_mcis "github.com/cloud-barista/poc-cicd-tumblebug/src/api/rest/server/mcis"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var funcs = map[string]interface{}{
	"RestGetHealth":                     rest_common.RestGetHealth,
	"RestCheckNs":                       rest_common.RestCheckNs,
	"RestGetConnConfigList":             rest_common.RestGetConnConfigList,
	"RestGetConnConfig":                 rest_common.RestGetConnConfig,
	"RestGetRegionList":                 rest_common.RestGetRegionList,
	"RestGetRegion":                     rest_common.RestGetRegion,
	"RestLookupSpecList":                rest_mcir.RestLookupSpecList,
	"RestLookupSpec":                    rest_mcir.RestLookupSpec,
	"RestLookupImageList":               rest_mcir.RestLookupImageList,
	"RestLookupImage":                   rest_mcir.RestLookupImage,
	"RestInspectResources":              rest_common.RestInspectResources,
	"RestPostConfig":                    rest_common.RestPostConfig,
	"RestGetConfig":                     rest_common.RestGetConfig,
	"RestGetAllConfig":                  rest_common.RestGetAllConfig,
	"RestInitConfig":                    rest_common.RestInitConfig,
	"RestInitAllConfig":                 rest_common.RestInitAllConfig,
	"RestGetObject":                     rest_common.RestGetObject,
	"RestGetObjects":                    rest_common.RestGetObjects,
	"RestDeleteObject":                  rest_common.RestDeleteObject,
	"RestDeleteObjects":                 rest_common.RestDeleteObjects,
	"RestPostNs":                        rest_common.RestPostNs,
	"RestGetNs":                         rest_common.RestGetNs,
	"RestGetAllNs":                      rest_common.RestGetAllNs,
	"RestPutNs":                         rest_common.RestPutNs,
	"RestDelNs":                         rest_common.RestDelNs,
	"RestDelAllNs":                      rest_common.RestDelAllNs,
	"RestPostMcis":                      rest_mcis.RestPostMcis,
	"RestGetMcis":                       rest_mcis.RestGetMcis,
	"RestGetAllMcis":                    rest_mcis.RestGetAllMcis,
	"RestPutMcis":                       rest_mcis.RestPutMcis,
	"RestDelMcis":                       rest_mcis.RestDelMcis,
	"RestDelAllMcis":                    rest_mcis.RestDelAllMcis,
	"RestPostMcisVm":                    rest_mcis.RestPostMcisVm,
	"RestPostMcisVmGroup":               rest_mcis.RestPostMcisVmGroup,
	"RestGetMcisVm":                     rest_mcis.RestGetMcisVm,
	"RestDelMcisVm":                     rest_mcis.RestDelMcisVm,
	"RestPostMcisRecommend":             rest_mcis.RestPostMcisRecommend,
	"RestRecommendVm":                   rest_mcis.RestRecommendVm,
	"RestPostCmdMcis":                   rest_mcis.RestPostCmdMcis,
	"RestPostCmdMcisVm":                 rest_mcis.RestPostCmdMcisVm,
	"RestPostInstallAgentToMcis":        rest_mcis.RestPostInstallAgentToMcis,
	"RestGetBenchmark":                  rest_mcis.RestGetBenchmark,
	"RestGetAllBenchmark":               rest_mcis.RestGetAllBenchmark,
	"RestPostMcisPolicy":                rest_mcis.RestPostMcisPolicy,
	"RestGetMcisPolicy":                 rest_mcis.RestGetMcisPolicy,
	"RestGetAllMcisPolicy":              rest_mcis.RestGetAllMcisPolicy,
	"RestPutMcisPolicy":                 rest_mcis.RestPutMcisPolicy,
	"RestDelMcisPolicy":                 rest_mcis.RestDelMcisPolicy,
	"RestDelAllMcisPolicy":              rest_mcis.RestDelAllMcisPolicy,
	"RestPostInstallMonitorAgentToMcis": rest_mcis.RestPostInstallMonitorAgentToMcis,
	"RestGetMonitorData":                rest_mcis.RestGetMonitorData,
	"RestPostImage":                     rest_mcir.RestPostImage,
	"RestGetResource":                   rest_mcir.RestGetResource,
	"RestGetAllResources":               rest_mcir.RestGetAllResources,
	"RestPutImage":                      rest_mcir.RestPutImage,
	"RestDelResource":                   rest_mcir.RestDelResource,
	"RestDelAllResources":               rest_mcir.RestDelAllResources,
	"RestPostSshKey":                    rest_mcir.RestPostSshKey,
	"RestPutSshKey":                     rest_mcir.RestPutSshKey,
	"RestPostSpec":                      rest_mcir.RestPostSpec,
	"RestPutSpec":                       rest_mcir.RestPutSpec,
	"RestFetchSpecs":                    rest_mcir.RestFetchSpecs,
	"RestFilterSpecs":                   rest_mcir.RestFilterSpecs,
	"RestFilterSpecsByRange":            rest_mcir.RestFilterSpecsByRange,
	"RestTestSortSpecs":                 rest_mcir.RestTestSortSpecs,
	"RestFetchImages":                   rest_mcir.RestFetchImages,
	"RestSearchImage":                   rest_mcir.RestSearchImage,
	"RestPostSecurityGroup":             rest_mcir.RestPostSecurityGroup,
	"RestPutSecurityGroup":              rest_mcir.RestPutSecurityGroup,
	"RestPostVNet":                      rest_mcir.RestPostVNet,
	"RestPutVNet":                       rest_mcir.RestPutVNet,
	"RestCheckResource":                 rest_mcir.RestCheckResource,
	"RestCheckMcis":                     rest_mcis.RestCheckMcis,
	"RestCheckVm":                       rest_mcis.RestCheckVm,
	"RestTestAddObjectAssociation":      rest_mcir.RestTestAddObjectAssociation,
	"RestTestDeleteObjectAssociation":   rest_mcir.RestTestDeleteObjectAssociation,
	"RestTestGetAssociatedObjectCount":  rest_mcir.RestTestGetAssociatedObjectCount,
}

func Call(m map[string]interface{}, name string, params ...interface{}) (result []reflect.Value, err error) {
	f := reflect.ValueOf(m[name])
	if len(params) != f.Type().NumIn() {
		err = errors.New("The number of params is not adapted.")
		return
	}

	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}
	result = f.Call(in)
	return
}

func EchoTest(t *testing.T, tc TestCases) (string, error) {

	var (
		body string = ""
		err  error  = nil
	)

	t.Run(tc.Name, func(t *testing.T) {
		e := echo.New()
		var req *http.Request = nil
		if tc.GivenPostData != "" {
			req = httptest.NewRequest(tc.HttpMethod, "/"+tc.GivenQueryParams, bytes.NewBuffer([]byte(tc.GivenPostData)))
		} else {
			req = httptest.NewRequest(tc.HttpMethod, "/"+tc.GivenQueryParams, nil)
		}
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath(tc.WhenURL)
		if tc.GivenParaNames != nil {
			c.SetParamNames(tc.GivenParaNames...)
			c.SetParamValues(tc.GivenParaVals...)
		}

		res, err := Call(funcs, tc.EchoFunc, c)
		if assert.NoError(t, err) {
			if res != nil && !res[0].IsNil() {
				he, ok := res[0].Interface().(*echo.HTTPError)
				if ok { // echo.NewHTTPError() 로 에러를 리턴했을 경우
					assert.Equal(t, tc.ExpectStatus, he.Code)
					body = fmt.Sprintf("%v", he.Message)
				} else { // err 로 에러를 리턴했을 경우
					body = fmt.Sprintf("%v", res[0])
				}
				if tc.ExpectBodyStartsWith != "" {
					if !assert.True(t, strings.HasPrefix(body, tc.ExpectBodyStartsWith)) {
						fmt.Fprintf(os.Stderr, "\n                Not Equal(echo.NewHTTPError): \n"+
							"                  Expected Start With: %s\n"+
							"                  Actual  : %s\n", tc.ExpectBodyStartsWith, body)
					}
				}
				if tc.ExpectBodyContains != "" {
					if !assert.True(t, strings.Contains(body, tc.ExpectBodyContains)) {
						fmt.Fprintf(os.Stderr, "\n                Not Equal(echo.NewHTTPError): \n"+
							"                  Expected Contains: %s\n"+
							"                  Actual  : %s\n", tc.ExpectBodyContains, body)
					}
				}
				if tc.ExpectBodyStartsWith == "" && tc.ExpectBodyContains == "" {
					if !assert.True(t, "" == body) {
						fmt.Fprintf(os.Stderr, "\n                Not Equal(echo.NewHTTPError): \n"+
							"      Expected StartWith/Contains: %s\n"+
							"      Actual  : %s\n", tc.ExpectBodyStartsWith, body)
					}
				}
			} else {
				assert.Equal(t, tc.ExpectStatus, rec.Code)
				body = rec.Body.String()
				if tc.ExpectBodyStartsWith != "" {
					if !assert.True(t, strings.HasPrefix(body, tc.ExpectBodyStartsWith)) {
						fmt.Fprintf(os.Stderr, "\n                Not Equal(echo.Context): \n"+
							"                  Expected Start With: %s\n"+
							"                  Actual  : %s\n", tc.ExpectBodyStartsWith, body)
					}
				}
				if tc.ExpectBodyContains != "" {
					if !assert.True(t, strings.Contains(body, tc.ExpectBodyContains)) {
						fmt.Fprintf(os.Stderr, "\n                Not Equal(echo.Context): \n"+
							"                  Expected Contains: %s\n"+
							"                  Actual  : %s\n", tc.ExpectBodyContains, body)
					}
				}
				if tc.ExpectBodyStartsWith == "" && tc.ExpectBodyContains == "" {
					if !assert.True(t, "" == body) {
						fmt.Fprintf(os.Stderr, "\n                Not Equal(echo.Context): \n"+
							"      Expected StartWith/Contains: %s\n"+
							"      Actual  : %s\n", tc.ExpectBodyStartsWith, body)
					}
				}
			}
		}
	})

	return body, err
}
