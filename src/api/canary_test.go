package api

import (
	"fmt"
	"testing"

	"github.com/DavidCai1993/request"
	"github.com/stretchr/testify/assert"
)

func TestCanary(t *testing.T) {
	tt, cleanup := SetUpTestTools()
	defer cleanup()

	t.Run(`canary should work`, func(t *testing.T) {
		assert := assert.New(t)

		res, err := request.Get(fmt.Sprintf("%s/v1/canary", tt.Host)).End()
		assert.Nil(err)

		json := map[string]string{}
		res.JSON(&json)

		assert.Nil(err)
		assert.Equal(200, res.StatusCode)
		assert.Equal("urbs-console", json["name"])
		assert.NotEqual("", json["version"])
		assert.NotEqual("", json["gitSHA1"])
		assert.NotEqual("", json["buildTime"])
	})
}