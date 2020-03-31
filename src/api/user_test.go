package api

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/teambition/urbs-console/src/dto/urbssetting"
	"github.com/teambition/urbs-console/src/tpl"
	"github.com/teambition/urbs-console/src/util/request"
)

func testUserBatchCreate(tt *TestTools, uid string) {
	req := &tpl.UsersBody{
		Users: []string{uid},
	}
	result := &tpl.BoolRes{}

	res, err := request.Post(fmt.Sprintf("%s/v1/users:batch", tt.Host)).Body(req).Result(result).Do()

	tt.Require.Nil(err)
	tt.Require.True(res.OK())
}

func testUserListLables(tt *TestTools, uid string, count int) {
	result := &urbssetting.LabelsInfoRes{}

	res, err := request.Get(fmt.Sprintf("%s/v1/users/%s/labels", tt.Host, uid)).Result(result).Do()

	tt.Require.Nil(err)
	tt.Require.True(res.OK())
	tt.Require.Equal(count, len(result.Result))
}

func testUserSettingAll(tt *TestTools, product string, uid string, count int) {
	result := &urbssetting.LabelsInfoRes{}

	res, err := request.Get(fmt.Sprintf("%s/v1/users/%s/settings:unionAll?product=%s", tt.Host, uid, product)).Result(result).Do()

	tt.Require.Nil(err)
	tt.Require.True(res.OK())
	tt.Require.Equal(count, len(result.Result))
}

func TestUserBatchAdd(t *testing.T) {
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

		assert.Equal(200, resp.StatusCode)

		resp, err = request.Post(fmt.Sprintf("%s/v1/users:batch", tt.Host)).Body(req).Result(result).Do()
		assert.Nil(err)

		assert.Equal(200, resp.StatusCode)
	})
}
