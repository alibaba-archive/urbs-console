package api

import (
	"fmt"
	"testing"

	"github.com/mushroomsir/request"
	"github.com/stretchr/testify/require"
)

func TestCanary(t *testing.T) {

	t.Run(`canary should work`, func(t *testing.T) {
		require := require.New(t)

		res, err := request.Get(fmt.Sprintf("%s/v1/canary", testHost)).Do()

		require.Nil(err)
		require.Equal(200, res.StatusCode)
	})
}
