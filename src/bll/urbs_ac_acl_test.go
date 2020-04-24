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

	args := &tpl.UrbsAcAclURL{
		Uid: tpl.RandUID(),
	}
	body := &tpl.UrbsAcAclAddReq{}
	body.Product = tpl.RandName()
	body.Label = tpl.RandLabel()
	body.Permission = constant.PermissionAll

	err := blls.UrbsAcAcl.AddByReq(context.Background(), args, body)
	require.Nil(err)

	userBody := &tpl.UrbsAcUsersBody{
		Users: []*tpl.UrbsAcUserBody{
			{
				Uid:  args.Uid,
				Name: args.Uid,
			},
		},
	}
	err = blls.UrbsAcUser.Add(context.Background(), userBody)
	require.Nil(err)

	err = blls.UrbsAcAcl.CheckViewer(getUidContext(args.Uid))
	require.Nil(err)

	object := body.Product + body.Label

	err = blls.UrbsAcAcl.CheckAdmin(getUidContext(args.Uid), object)
	require.Nil(err)

	subjects := []string{tpl.RandUID()}
	userBody = &tpl.UrbsAcUsersBody{
		Users: []*tpl.UrbsAcUserBody{
			{
				Uid:  subjects[0],
				Name: subjects[0],
			},
		},
	}
	err = blls.UrbsAcUser.Add(context.Background(), userBody)
	require.Nil(err)

	err = blls.UrbsAcAcl.Update(context.Background(), subjects, object)
	require.Nil(err)

	err = blls.UrbsAcAcl.CheckAdmin(getUidContext(subjects[0]), object)
	require.Nil(err, object)

	err = blls.UrbsAcAcl.CheckAdmin(getUidContext(args.Uid), object)
	require.NotNil(err)

	users, err := blls.UrbsAcAcl.FindUsersByObject(context.Background(), object)
	require.Nil(err)
	require.Equal(1, len(users))
	require.Equal(subjects[0], users[0].Uid)
	require.Equal(subjects[0], users[0].Name)
}
