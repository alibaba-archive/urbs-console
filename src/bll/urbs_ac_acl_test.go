package bll

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/teambition/urbs-console/src/constant"
	"github.com/teambition/urbs-console/src/tpl"
)

func TestUrbsAcACL(t *testing.T) {
	require := require.New(t)

	body := &tpl.UrbsAcAclAddReq{}
	body.Uid = tpl.RandUID()
	body.Product = tpl.RandName()
	body.Label = tpl.RandLabel()
	body.Permission = constant.PermissionAll

	// 添加权限
	err := blls.UrbsAcAcl.AddByReq(context.Background(), body)
	require.Nil(err)

	// 添加用户
	userBody := &tpl.UrbsAcUsersBody{
		Users: []*tpl.UrbsAcUserBody{
			{
				Uid:  body.Uid,
				Name: body.Uid,
			},
		},
	}

	err = blls.UrbsAcUser.Add(context.Background(), userBody)
	require.Nil(err)

	// 检查浏览者权限
	err = blls.UrbsAcAcl.CheckViewer(getUidContext(body.Uid))
	require.Nil(err)

	// 检查管理者权限
	err = blls.UrbsAcAcl.CheckAdmin(getUidContext(body.Uid), body.Product+body.Label)
	require.Nil(err)
}
