package bll

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/teambition/urbs-console/src/constant"
	"github.com/teambition/urbs-console/src/dto"
	"github.com/teambition/urbs-console/src/tpl"
)

func TestOperationLog(t *testing.T) {
	require := require.New(t)
	tt := SetUpTestTools(require)

	uid := tpl.RandUID()
	ctx := getUidContext(uid)
	object := tpl.RandName()

	body := new(tpl.UsersGroupsBody)
	body.Users = []string{tpl.RandUID()}
	body.Groups = []string{tpl.RandUID()}
	body.Desc = tpl.RandName()
	body.Value = "true"

	testAddUrbsAcUser(tt, uid)

	t.Run("add operationLog", func(t *testing.T) {
		logContent := &dto.OperationLogContent{
			Users:   body.Users,
			Groups:  body.Groups,
			Desc:    body.Desc,
			Value:   body.Value,
			Percent: 2,
		}
		err := blls.OperationLog.Add(ctx, object, constant.OperationCreate, logContent)
		require.Nil(err)
	})

	t.Run("get operationLog", func(t *testing.T) {
		// 获取操作日志
		page := &tpl.Pagination{}
		page.Validate()
		res, err := blls.OperationLog.List(ctx, object, page)
		require.Nil(err)

		require.Equal(body.Users, res.Result[0].Users)
		require.Equal(body.Groups, res.Result[0].Groups)
		require.Equal(body.Desc, res.Result[0].Desc)
		require.Equal(body.Value, res.Result[0].Value)
		require.Equal(2, res.Result[0].Percent)

		totalSize, err := daos.OperationLog.CountByObject(ctx, object)
		require.Nil(err)
		require.True(totalSize > 0)
	})
}
