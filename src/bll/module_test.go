package bll

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/teambition/urbs-console/src/tpl"
)

func TestModule(t *testing.T) {
	require := require.New(t)
	tt := SetUpTestTools(require)

	// create product
	productName := tpl.RandName()
	testProductCreate(tt, productName)

	// create module
	moduleName := tpl.RandName()
	testModuleCreate(tt, productName, moduleName)

	testModuleGet(tt, productName)

	testModuleUpdate(tt, productName, moduleName)

	testModuleOffline(tt, productName, moduleName)
}

func testModuleCreate(tt *TestTools, product, module string) (*tpl.ModuleInfoRes, error) {
	req := &tpl.NameDescBody{
		Name: module,
	}
	result, err := blls.Module.Create(context.Background(), product, req)

	tt.Require.Nil(err)
	tt.Require.Equal(module, result.Result.Name)
	tt.Require.True(result.Result.CreatedAt.Before(time.Now()))
	return result, nil
}

func testModuleGet(tt *TestTools, product string) (*tpl.ModulesInfoRes, error) {
	args := &tpl.ProductPaginationURL{
		Product: product,
	}
	result, err := blls.Module.List(context.Background(), args)
	tt.Require.Nil(err)
	tt.Require.Equal(1, len(result.Result))
	tt.Require.True(result.Result[0].CreatedAt.Before(time.Now()))
	return result, nil
}

func testModuleUpdate(tt *TestTools, product, module string) (*tpl.ModuleInfoRes, error) {
	desc := tpl.RandName()
	req := &tpl.ModuleUpdateBody{
		Desc: &desc,
	}
	result, err := blls.Module.Update(context.Background(), product, module, req)

	tt.Require.Nil(err)
	tt.Require.Equal(*req.Desc, result.Result.Desc)
	tt.Require.True(result.Result.CreatedAt.Before(time.Now()))
	return result, nil
}

func testModuleOffline(tt *TestTools, product, module string) (*tpl.BoolRes, error) {

	res, err := blls.Module.Offline(context.Background(), product, module)

	tt.Require.Nil(err)
	return res, nil
}