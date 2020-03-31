package api

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/teambition/urbs-console/src/dto/urbssetting"
	"github.com/teambition/urbs-console/src/tpl"
	"github.com/teambition/urbs-console/src/util/request"
)

func TestAssign(t *testing.T) {
	require := require.New(t)
	tt, cleanup := SetUpTestTools()
	tt.Require = require
	defer cleanup()

	// create product
	productName := tpl.RandName()
	testProductCreate(tt, productName)

	// create label
	labelName := tpl.RandName()
	label, err := testLabelCreate(tt, productName, labelName)
	require.Nil(err)
	require.Equal(labelName, label.Result.Name)

	// update label
	testLabelUpdate(tt, productName, labelName)

	// label list
	testLabelList(tt, productName, 1)

	// create user
	uid := tpl.RandUID()
	testUserBatchCreate(tt, uid)

	// assign label
	testLabelAssign(tt, productName, labelName, uid)

	// users labels
	testUserListLables(tt, uid, 1)

	// offline
	testLabelOffline(tt, productName, labelName)

	// users labels
	testUserListLables(tt, uid, 0)

	testLabelList(tt, productName, 1)

	// delete
	testLabelDelete(tt, productName, labelName)

	testLabelList(tt, productName, 0)
}

func testLabelCreate(tt *TestTools, product, labelName string) (*urbssetting.LabelInfoRes, error) {

	result := &urbssetting.LabelInfoRes{}

	body := &tpl.LabelBody{
		Name: labelName,
	}

	res, err := request.Post(fmt.Sprintf("%s/v1/products/%s/labels", tt.Host, product)).Body(body).Result(result).Do()
	tt.Require.Nil(err)
	tt.Require.True(res.OK())
	return result, nil
}

func testLabelUpdate(tt *TestTools, product, labelName string) (*urbssetting.LabelInfoRes, error) {
	body := &tpl.LabelUpdateBody{
		Desc:     &labelName,
		Channels: &[]string{"beta"},
		Clients:  &[]string{"ios"},
	}
	result := &urbssetting.LabelInfoRes{}

	res, err := request.Put(fmt.Sprintf("%s/v1/products/%s/labels/%s", tt.Host, product, labelName)).Body(body).Result(result).Do()
	tt.Require.Nil(err)
	tt.Require.True(res.OK())

	tt.Require.Equal(*body.Desc, result.Result.Desc)
	tt.Require.Equal(*body.Channels, result.Result.Channels)
	tt.Require.Equal(*body.Clients, result.Result.Clients)
	return result, nil
}

func testLabelList(tt *TestTools, product string, count int) {

	result := &urbssetting.LabelsInfoRes{}

	res, err := request.Get(fmt.Sprintf("%s/v1/products/%s/labels", tt.Host, product)).Result(result).Do()

	tt.Require.Nil(err)
	tt.Require.True(res.OK())
	tt.Require.Equal(count, len(result.Result))
}

func testLabelAssign(tt *TestTools, product, labelName, uid string) {
	body := &tpl.UsersGroupsBody{
		Users: []string{uid},
	}
	result := &tpl.BoolRes{}

	res, err := request.Post(fmt.Sprintf("%s/v1/products/%s/labels/%s:assign", tt.Host, product, labelName)).Body(body).Result(result).Do()

	tt.Require.Nil(err)
	tt.Require.True(res.OK())
}

func testLabelOffline(tt *TestTools, product, labelName string) {

	result := &tpl.BoolRes{}

	res, err := request.Put(fmt.Sprintf("%s/v1/products/%s/labels/%s:offline", tt.Host, product, labelName)).Result(result).Do()

	tt.Require.Nil(err)
	tt.Require.True(res.OK())
}

func testLabelDelete(tt *TestTools, product, labelName string) {

	result := &tpl.BoolRes{}

	res, err := request.Delete(fmt.Sprintf("%s/v1/products/%s/labels/%s", tt.Host, product, labelName)).Result(result).Do()

	tt.Require.Nil(err)
	tt.Require.True(res.OK())
}
