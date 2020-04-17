package api

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/teambition/urbs-console/src/dto/urbssetting"
	"github.com/teambition/urbs-console/src/tpl"
	"github.com/teambition/urbs-console/src/util/request"
)

func TestSetting(t *testing.T) {
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

	settingName := tpl.RandName()
	testSettingCreate(tt, productName, moduleName, settingName)

	testSettingGet(tt, productName, moduleName, settingName)

	testSettingList(tt, productName, moduleName)

	testSettingUpdate(tt, productName, moduleName, settingName)

	// create user
	uid := tpl.RandUID()
	testUserBatchCreate(tt, uid)

	testSettingAssign(tt, productName, moduleName, settingName, uid)

	testUserSettingAll(tt, productName, uid, 1)

	testSettingOffline(tt, productName, moduleName, settingName)

	testUserSettingAll(tt, productName, uid, 0)
}

func testSettingCreate(tt *TestTools, product, module, setting string) (*urbssetting.SettingInfoRes, error) {

	body := &tpl.NameDescBody{
		Name: setting,
	}

	result := &urbssetting.SettingInfoRes{}

	res, err := request.Post(fmt.Sprintf("%s/api/v1/products/%s/modules/%s/settings", tt.Host, product, module)).Body(body).Result(result).Do()
	tt.Require.Nil(err)
	tt.Require.True(res.OK())
	tt.Require.Equal(body.Name, result.Result.Name)
	return result, nil
}

func testSettingList(tt *TestTools, product, module string) (*urbssetting.SettingsInfoRes, error) {

	result := &urbssetting.SettingsInfoRes{}

	res, err := request.Get(fmt.Sprintf("%s/api/v1/products/%s/modules/%s/settings", tt.Host, product, module)).Result(result).Do()

	tt.Require.Nil(err)
	tt.Require.True(res.OK())
	tt.Require.Equal(1, len(result.Result))
	return result, nil
}

func testSettingGet(tt *TestTools, product, module, setting string) (*urbssetting.SettingInfoRes, error) {

	result := &urbssetting.SettingInfoRes{}

	res, err := request.Get(fmt.Sprintf("%s/api/v1/products/%s/modules/%s/settings/%s", tt.Host, product, module, setting)).Result(result).Do()

	tt.Require.Nil(err)
	tt.Require.True(res.OK())
	tt.Require.Equal(setting, result.Result.Name)
	return result, nil
}

func testSettingUpdate(tt *TestTools, product, module, setting string) (*urbssetting.SettingInfoRes, error) {

	body := &tpl.SettingUpdateBody{
		Desc:     &setting,
		Channels: &[]string{"beta"},
		Clients:  &[]string{"ios"},
	}

	result := &urbssetting.SettingInfoRes{}

	res, err := request.Put(fmt.Sprintf("%s/api/v1/products/%s/modules/%s/settings/%s", tt.Host, product, module, setting)).Body(body).Result(result).Do()
	tt.Require.Nil(err)
	tt.Require.True(res.OK())

	tt.Require.Equal(*body.Desc, result.Result.Desc)
	tt.Require.Equal(*body.Channels, result.Result.Channels)
	tt.Require.Equal(*body.Clients, result.Result.Clients)

	return result, nil
}

func testSettingOffline(tt *TestTools, product, module, setting string) (*tpl.BoolRes, error) {

	result := &tpl.BoolRes{}

	res, err := request.Put(fmt.Sprintf("%s/api/v1/products/%s/modules/%s/settings/%s:offline", tt.Host, product, module, setting)).Result(result).Do()
	tt.Require.Nil(err)
	tt.Require.True(res.OK())

	return result, nil
}

func testSettingAssign(tt *TestTools, product, module, setting string, uid string) (*tpl.BoolRes, error) {

	body := &tpl.UsersGroupsBody{
		Users: []string{uid},
	}

	result := &tpl.BoolRes{}

	res, err := request.Post(fmt.Sprintf("%s/api/v1/products/%s/modules/%s/settings/%s:assign", tt.Host, product, module, setting)).Body(body).Result(result).Do()
	tt.Require.Nil(err)
	tt.Require.True(res.OK())

	return result, nil
}
