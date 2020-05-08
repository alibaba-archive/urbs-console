package api

import (
	"fmt"
	"testing"

	"github.com/mushroomsir/request"
	"github.com/stretchr/testify/require"
)

func TestApp(t *testing.T) {
	t.Run(`app should work`, func(t *testing.T) {
		require := require.New(t)
		json := map[string]string{}
		res, err := request.Get(testHost).Result(&json).Do()
		require.Nil(err)
		require.Equal(200, res.StatusCode)
		require.Equal("urbs-console", json["name"])
		require.NotEqual("", json["version"])
		require.NotEqual("", json["gitSHA1"])
		require.NotEqual("", json["buildTime"])
	})

	t.Run(`"GET /version" should work`, func(t *testing.T) {
		require := require.New(t)

		json := map[string]string{}
		res, err := request.Get(fmt.Sprintf("%s/version", testHost)).Result(&json).Do()
		require.Nil(err)
		require.Equal(200, res.StatusCode)
		require.Equal("urbs-console", json["name"])
		require.NotEqual("", json["version"])
		require.NotEqual("", json["gitSHA1"])
		require.NotEqual("", json["buildTime"])
	})
}
