package goscenario

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGoMisc(t *testing.T) {
	t.Run("go api misc test", func(t *testing.T) {
		SetUpForGrpc()

		// NsApi Set/Get Testing
		NsApi.SetServerAddr("127.0.0.1")
		res, _ := NsApi.GetServerAddr()
		assert.True(t, "127.0.0.1" == res)

		NsApi.SetTLSCA("./tls.ca")
		res, _ = NsApi.GetTLSCA()
		assert.True(t, "./tls.ca" == res)

		sec, _ := time.ParseDuration("10s")
		NsApi.SetTimeout(sec)
		tt, _ := NsApi.GetTimeout()
		assert.True(t, sec == tt)

		NsApi.SetJWTToken("abcdefg")
		res, _ = NsApi.GetJWTToken()
		assert.True(t, "abcdefg" == res)

		NsApi.SetInType("json")
		res, _ = NsApi.GetInType()
		assert.True(t, "json" == res)

		NsApi.SetInType("yaml")
		res, _ = NsApi.GetInType()
		assert.True(t, "yaml" == res)

		err := NsApi.SetInType("text")
		assert.True(t, err != nil)

		NsApi.SetOutType("json")
		res, _ = NsApi.GetOutType()
		assert.True(t, "json" == res)

		NsApi.SetOutType("yaml")
		res, _ = NsApi.GetOutType()
		assert.True(t, "yaml" == res)

		err = NsApi.SetOutType("text")
		assert.True(t, err != nil)

		// McirApi Set/Get Testing
		McirApi.SetServerAddr("127.0.0.1")
		res, _ = McirApi.GetServerAddr()
		assert.True(t, "127.0.0.1" == res)

		McirApi.SetTLSCA("./tls.ca")
		res, _ = McirApi.GetTLSCA()
		assert.True(t, "./tls.ca" == res)

		sec, _ = time.ParseDuration("10s")
		McirApi.SetTimeout(sec)
		tt, _ = McirApi.GetTimeout()
		assert.True(t, sec == tt)

		McirApi.SetJWTToken("abcdefg")
		res, _ = McirApi.GetJWTToken()
		assert.True(t, "abcdefg" == res)

		McirApi.SetInType("json")
		res, _ = McirApi.GetInType()
		assert.True(t, "json" == res)

		McirApi.SetInType("yaml")
		res, _ = McirApi.GetInType()
		assert.True(t, "yaml" == res)

		err = McirApi.SetInType("text")
		assert.True(t, err != nil)

		McirApi.SetOutType("json")
		res, _ = McirApi.GetOutType()
		assert.True(t, "json" == res)

		McirApi.SetOutType("yaml")
		res, _ = McirApi.GetOutType()
		assert.True(t, "yaml" == res)

		err = McirApi.SetOutType("text")
		assert.True(t, err != nil)

		// McisApi Set/Get Testing
		McisApi.SetServerAddr("127.0.0.1")
		res, _ = McisApi.GetServerAddr()
		assert.True(t, "127.0.0.1" == res)

		McisApi.SetTLSCA("./tls.ca")
		res, _ = McisApi.GetTLSCA()
		assert.True(t, "./tls.ca" == res)

		sec, _ = time.ParseDuration("10s")
		McisApi.SetTimeout(sec)
		tt, _ = McisApi.GetTimeout()
		assert.True(t, sec == tt)

		McisApi.SetJWTToken("abcdefg")
		res, _ = McisApi.GetJWTToken()
		assert.True(t, "abcdefg" == res)

		McisApi.SetInType("json")
		res, _ = McisApi.GetInType()
		assert.True(t, "json" == res)

		McisApi.SetInType("yaml")
		res, _ = McisApi.GetInType()
		assert.True(t, "yaml" == res)

		err = McisApi.SetInType("text")
		assert.True(t, err != nil)

		McisApi.SetOutType("json")
		res, _ = McisApi.GetOutType()
		assert.True(t, "json" == res)

		McisApi.SetOutType("yaml")
		res, _ = McisApi.GetOutType()
		assert.True(t, "yaml" == res)

		err = McisApi.SetOutType("text")
		assert.True(t, err != nil)

		// TbutilApi Set/Get Testing
		TbutilApi.SetServerAddr("127.0.0.1")
		res, _ = TbutilApi.GetServerAddr()
		assert.True(t, "127.0.0.1" == res)

		TbutilApi.SetTLSCA("./tls.ca")
		res, _ = TbutilApi.GetTLSCA()
		assert.True(t, "./tls.ca" == res)

		sec, _ = time.ParseDuration("10s")
		TbutilApi.SetTimeout(sec)
		tt, _ = TbutilApi.GetTimeout()
		assert.True(t, sec == tt)

		TbutilApi.SetJWTToken("abcdefg")
		res, _ = TbutilApi.GetJWTToken()
		assert.True(t, "abcdefg" == res)

		TbutilApi.SetInType("json")
		res, _ = TbutilApi.GetInType()
		assert.True(t, "json" == res)

		TbutilApi.SetInType("yaml")
		res, _ = TbutilApi.GetInType()
		assert.True(t, "yaml" == res)

		err = TbutilApi.SetInType("text")
		assert.True(t, err != nil)

		TbutilApi.SetOutType("json")
		res, _ = TbutilApi.GetOutType()
		assert.True(t, "json" == res)

		TbutilApi.SetOutType("yaml")
		res, _ = TbutilApi.GetOutType()
		assert.True(t, "yaml" == res)

		err = TbutilApi.SetOutType("text")
		assert.True(t, err != nil)

		TearDownForGrpc()
	})
}
