package bll

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/teambition/urbs-console/src/tpl"
)

func TestUrbsAcUser(t *testing.T) {
	require := require.New(t)

	for i := 0; i < 5; i++ {
		args := &tpl.UrbsAcAclURL{
			Uid: tpl.RandUID(),
		}
		userBody := &tpl.UrbsAcUsersBody{
			Users: []*tpl.UrbsAcUserBody{
				{
					Uid:  args.Uid,
					Name: args.Uid,
				},
			},
		}
		err := blls.UrbsAcUser.Add(context.Background(), userBody)
		require.Nil(err)
	}

	args := new(tpl.Pagination)
	args.PageSize = 1
	res, err := blls.UrbsAcUser.List(context.Background(), args)
	require.Nil(err)
	require.Equal(1, len(res.Result))
	require.NotEmpty(res.NextPageToken)

	args2 := new(tpl.Pagination)
	args2.PageToken = res.NextPageToken
	args2.Validate()
	require.Equal(1, args2.Skip)

	res2, err := blls.UrbsAcUser.List(context.Background(), args2)
	require.Nil(err)
	require.Equal(1, len(res2.Result))
	require.NotEmpty(res2.NextPageToken)
	require.NotEqual(res.Result[0].UID, res2.Result[0].UID)

	require.True(res2.TotalSize > 0)

	res3, err := blls.UrbsAcUser.Search(context.Background(), res.Result[0].Name)
	require.Nil(err)
	require.Equal(res.Result[0].Name, res3.Result[0].Name)
}
