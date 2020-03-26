package api

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/teambition/urbs-console/src/tpl"
	"github.com/teambition/urbs-console/src/util/request"
)

func TestUser(t *testing.T) {
	tt, cleanup := SetUpTestTools()
	defer cleanup()

	t.Run(`BatchAdd`, func(t *testing.T) {
		assert := assert.New(t)

		req := &tpl.UsersBody{
			Users: []string{"123"},
		}
		result := &tpl.BoolRes{}

		resp, err := request.Post(fmt.Sprintf("%s/v1/users:batch", tt.Host)).Header(genHeader()).Body(req).Result(result).Do()
		assert.Nil(err)

		assert.Equal(400, resp.StatusCode)
		//log.Println(resp.String())

		resp, err = request.Post(fmt.Sprintf("%s/v1/users:batch", tt.Host)).Body(req).Result(result).Do()
		assert.Nil(err)

		assert.Equal(401, resp.StatusCode)
	})
}
