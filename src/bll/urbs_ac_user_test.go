package bll

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/teambition/urbs-console/src/tpl"
)

func TestUrbsAcUser(t *testing.T) {
	require := require.New(t)
	tt := SetUpTestTools(require)

	t.Run("add urbsAcUser", func(t *testing.T) {
		for i := 0; i < 5; i++ {
			testAddUrbsAcUser(tt, tpl.RandUID())
		}
	})

	t.Run("urbsAcUser pagination", func(t *testing.T) {
		args := new(tpl.ConsolePagination)
		args.PageSize = 1
		res, err := blls.UrbsAcUser.List(context.Background(), args)
		require.Nil(err)
		require.Equal(1, len(res.Result))
		require.NotEmpty(res.NextPageToken)

		args2 := new(tpl.ConsolePagination)
		args2.PageToken = res.NextPageToken
		args2.Validate()
		require.Equal(1, args2.Skip)

		res2, err := blls.UrbsAcUser.List(context.Background(), args2)
		require.Nil(err)
		require.Equal(1, len(res2.Result))
		require.NotEmpty(res2.NextPageToken)
		require.NotEqual(res.Result[0].UID, res2.Result[0].UID)
		require.True(res2.TotalSize > 0)
	})
	t.Run("search", func(t *testing.T) {
		uid := tpl.RandUID()
		testAddUrbsAcUser(tt, uid)

		res3, err := blls.UrbsAcUser.Search(context.Background(), uid)
		require.Nil(err)
		require.Equal(uid, res3.Result[0].Name)
	})

	t.Run("delete", func(t *testing.T) {
		uid := tpl.RandUID()
		testAddUrbsAcUser(tt, uid)

		res, err := blls.UrbsAcUser.Search(context.Background(), uid)
		require.Nil(err)
		require.Equal(uid, res.Result[0].Name)

		err = blls.UrbsAcUser.DeleteByUID(context.Background(), uid)
		require.Nil(err)

		res, err = blls.UrbsAcUser.Search(context.Background(), uid)
		require.Nil(err)
		require.Equal(0, len(res.Result))
	})

	t.Run("update", func(t *testing.T) {
		uid := tpl.RandUID()
		testAddUrbsAcUser(tt, uid)

		name := tpl.RandName()
		err := blls.UrbsAcUser.UpdateByUID(context.Background(), name, uid)
		require.Nil(err)

		res, err := blls.UrbsAcUser.Search(context.Background(), uid)
		require.Nil(err)
		require.Equal(name, res.Result[0].Name)
	})
}

func testAddUrbsAcUser(tt *TestTools, uid string) {
	userBody := &tpl.UrbsAcUsersBody{
		Users: []*tpl.UrbsAcUserBody{
			{
				Uid:  uid,
				Name: uid,
			},
		},
	}
	err := blls.UrbsAcUser.Add(context.Background(), userBody)
	tt.Require.Nil(err)
}
