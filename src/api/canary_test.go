package api

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/teambition/urbs-console/src/util/request"
)

func TestCanary(t *testing.T) {

	t.Run(`canary should work`, func(t *testing.T) {
		require := require.New(t)

		json := map[string]string{}
		res, err := request.Get(fmt.Sprintf("%s/v1/canary", testHost)).Result(&json).Do()

		require.Nil(err)
		require.Equal(200, res.StatusCode)
		require.Equal("urbs-console", json["name"])
		require.NotEqual("", json["version"])
		require.NotEqual("", json["gitSHA1"])
		require.NotEqual("", json["buildTime"])
	})
}
