package bll

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/teambition/urbs-console/src/constant"
	"github.com/teambition/urbs-console/src/tpl"
)

func TestOperationLog(t *testing.T) {
	require := require.New(t)

	uid := tpl.RandUID()
	ctx := getUidContext(uid)
	object := tpl.RandName()

	body := new(tpl.UsersGroupsBody)
	body.Users = []string{tpl.RandUID()}
	body.Groups = []string{tpl.RandUID()}
	body.Desc = tpl.RandName()
	body.Percentage = 2
	body.Value = "true"

	// 添加用户
	userBody := &tpl.UrbsAcUsersBody{
		Users: []*tpl.UrbsAcUserBody{
			{
				Uid:  uid,
				Name: tpl.RandUID(),
			},
		},
	}

	err := blls.UrbsAcUser.Add(context.Background(), userBody)
	require.Nil(err)

	// 添加操作日志
	err = blls.OperationLog.Add(ctx, object, constant.OperationCreate, body)
	require.Nil(err)

	// 获取操作日志
	page := &tpl.Pagination{}
	page.Validate()
	res, err := blls.OperationLog.List(ctx, object, page)
	require.Nil(err)

	require.Equal(body.Users, res.Result[0].Users)
	require.Equal(body.Groups, res.Result[0].Groups)
	require.Equal(body.Desc, res.Result[0].Desc)
	require.Equal(body.Percentage, res.Result[0].Percentage)
	require.Equal(body.Value, res.Result[0].Value)
}
