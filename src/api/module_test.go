package api

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/teambition/urbs-console/src/dto/urbssetting"
	"github.com/teambition/urbs-console/src/tpl"
	"github.com/teambition/urbs-console/src/util/request"
)

func TestModule(t *testing.T) {
	require := require.New(t)
	tt, cleanup := SetUpTestTools()
	tt.Require = require
	defer cleanup()

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

func testModuleCreate(tt *TestTools, product, module string) (*urbssetting.ModuleRes, error) {
	req := &tpl.NameDescBody{
		Name: module,
	}
	result := &urbssetting.ModuleRes{}

	res, err := request.Post(fmt.Sprintf("%s/v1/products/%s/modules", tt.Host, product)).Body(req).Result(result).Do()

	tt.Require.Nil(err)
	tt.Require.True(res.OK())
	tt.Require.Equal(module, result.Result.Name)
	return result, nil
}

func testModuleGet(tt *TestTools, product string) (*urbssetting.ModulesRes, error) {

	result := &urbssetting.ModulesRes{}

	res, err := request.Get(fmt.Sprintf("%s/v1/products/%s/modules", tt.Host, product)).Result(result).Do()

	tt.Require.Nil(err)
	tt.Require.True(res.OK())
	tt.Require.Equal(1, len(result.Result))
	return result, nil
}

func testModuleUpdate(tt *TestTools, product, module string) (*urbssetting.ModuleRes, error) {
	desc := tpl.RandName()
	req := &tpl.ModuleUpdateBody{
		Desc: &desc,
	}
	result := &urbssetting.ModuleRes{}

	res, err := request.Put(fmt.Sprintf("%s/v1/products/%s/modules/%s", tt.Host, product, module)).Body(req).Result(result).Do()

	tt.Require.Nil(err)
	tt.Require.True(res.OK())
	tt.Require.Equal(*req.Desc, result.Result.Desc)
	return result, nil
}

func testModuleOffline(tt *TestTools, product, module string) (*tpl.BoolRes, error) {

	result := &tpl.BoolRes{}

	res, err := request.Put(fmt.Sprintf("%s/v1/products/%s/modules/%s:offline", tt.Host, product, module)).Result(result).Do()

	tt.Require.Nil(err)
	tt.Require.True(res.OK())
	return result, nil
}
