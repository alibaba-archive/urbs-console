package bll

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/teambition/urbs-console/src/tpl"
)

func testProductCreate(tt *TestTools, name string) (*tpl.ProductRes, error) {
	req := &tpl.NameDescBody{
		Name: name,
	}
	res, err := blls.Product.Create(context.Background(), req)
	tt.Require.Nil(err)
	tt.Require.Equal(name, res.Result.Name)
	return res, nil
}

func TestProduct(t *testing.T) {
	require := require.New(t)
	tt := SetUpTestTools(require)

	uid := tpl.RandUID()
	testAddUrbsAcUser(tt, uid)

	req := &tpl.NameDescBody{
		Name: tpl.RandName(),
		Uids: []string{uid},
	}
	res, err := blls.Product.Create(context.Background(), req)
	require.Nil(err)
	require.Equal(req.Name, res.Result.Name)
	require.True(len(res.Result.Users) > 0, req.Uids[0])
	require.Equal(req.Uids[0], res.Result.Users[0].Uid)

	uid2 := tpl.RandUID()
	testAddUrbsAcUser(tt, uid2)

	req1 := &tpl.ProductUpdateBody{
		Uids: []string{uid2},
		Desc: &uid2,
	}
	res, err = blls.Product.Update(context.Background(), req.Name, req1)
	require.Nil(err)
	require.Equal(req.Name, res.Result.Name)
	require.True(len(res.Result.Users) == 1, req.Uids[0])
	require.Equal(uid2, res.Result.Users[0].Uid)
}
