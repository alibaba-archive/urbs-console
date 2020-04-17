package api

import (
	"fmt"

	"github.com/teambition/urbs-console/src/dto/urbssetting"
	"github.com/teambition/urbs-console/src/tpl"
	"github.com/teambition/urbs-console/src/util/request"
)

func testProductCreate(tt *TestTools, name string) (*urbssetting.ProductRes, error) {
	req := &tpl.NameDescBody{
		Name: name,
	}
	result := &urbssetting.ProductRes{}

	res, err := request.Post(fmt.Sprintf("%s/api/v1/products", tt.Host)).Body(req).Result(result).Do()
	tt.Require.Nil(err)
	tt.Require.True(res.OK())
	tt.Require.Equal(name, result.Result.Name)
	return result, nil
}
