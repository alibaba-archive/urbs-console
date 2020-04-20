package bll

import (
	"context"

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
