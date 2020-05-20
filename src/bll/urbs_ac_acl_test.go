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
	tt := SetUpTestTools(require)

	t.Run("check viewer", func(t *testing.T) {
		uid := tpl.RandUID()

		testAddUrbsAcUser(tt, uid)
		testAddUrbsAcAcl(tt, uid)

		err := blls.UrbsAcAcl.CheckViewer(getUidContext(uid))
		require.Nil(err)
	})

	t.Run("check admin", func(t *testing.T) {
		uid := tpl.RandUID()

		testAddUrbsAcUser(tt, uid)
		res := testAddUrbsAcAcl(tt, uid)
		object := res.Product + res.Label

		err := blls.UrbsAcAcl.CheckAdmin(getUidContext(uid), object)
		require.Nil(err)
	})

	t.Run("check update permission", func(t *testing.T) {
		uid := tpl.RandUID()

		testAddUrbsAcUser(tt, uid)
		res := testAddUrbsAcAcl(tt, uid)
		object := res.Product + res.Label

		uid2 := tpl.RandUID()

		err := blls.UrbsAcAcl.Update(context.Background(), &[]string{uid2}, object)
		require.Nil(err)

		err = blls.UrbsAcAcl.CheckAdmin(getUidContext(uid), object)
		require.NotNil(err)
	})

	t.Run("find ysers by object", func(t *testing.T) {
		uid := tpl.RandUID()

		testAddUrbsAcUser(tt, uid)
		res := testAddUrbsAcAcl(tt, uid)
		object := res.Product + res.Label

		users, err := blls.UrbsAcAcl.FindUsersByObject(context.Background(), object)
		require.Nil(err)
		require.Equal(1, len(users))
		require.Equal(uid, users[0].Uid)
		require.Equal(uid, users[0].Name)
	})

	t.Run("check empty admin", func(t *testing.T) {
		uid := tpl.RandUID()
		object := ""
		err := blls.UrbsAcAcl.CheckAdmin(getUidContext(uid), object)
		require.NotNil(err)
	})
}

func testAddUrbsAcAcl(tt *TestTools, uid string) *tpl.UrbsAcAclAddBody {
	args := &tpl.UrbsAcAclURL{
		Uid: uid,
	}

	body := &tpl.UrbsAcAclAddBody{}
	body.Product = tpl.RandName()
	body.Label = tpl.RandLabel()
	body.Permission = constant.PermissionAll

	err := blls.UrbsAcAcl.AddByReq(context.Background(), args, body)
	tt.Require.Nil(err)
	return body
}
