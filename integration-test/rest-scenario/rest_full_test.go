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

		TearDownForRest()
	})

}
