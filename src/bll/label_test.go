package bll

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/teambition/urbs-console/src/tpl"
)

func TestLabelAssign(t *testing.T) {
	require := require.New(t)
	tt := SetUpTestTools(require)

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

func testLabelCreate(tt *TestTools, product, labelName string) (*tpl.LabelInfoRes, error) {
	body := &tpl.LabelBody{
		Name: labelName,
	}
	res, err := blls.Label.Create(context.Background(), product, body)
	tt.Require.Nil(err)
	return res, nil
}

func testLabelUpdate(tt *TestTools, product, labelName string) (*tpl.LabelInfoRes, error) {
	body := &tpl.LabelUpdateBody{
		Desc:     &labelName,
		Channels: &[]string{"beta"},
		Clients:  &[]string{"ios"},
	}
	result, err := blls.Label.Update(context.Background(), product, labelName, body)
	tt.Require.Nil(err)

	tt.Require.Equal(*body.Desc, result.Result.Desc)
	tt.Require.Equal(*body.Channels, result.Result.Channels)
	tt.Require.Equal(*body.Clients, result.Result.Clients)
	return result, nil
}

func testLabelList(tt *TestTools, product string, count int) {
	args := &tpl.ProductPaginationURL{
		Product: product,
	}
	result, err := blls.Label.List(context.Background(), args)
	tt.Require.Nil(err)
	tt.Require.Equal(count, len(result.Result))
}

func testLabelAssign(tt *TestTools, product, label, uid string) {
	body := &tpl.UsersGroupsBody{
		Users: []string{uid},
	}

	_, err := blls.Label.Assign(getUidContext(), product, label, body)

	tt.Require.Nil(err)
}

func testLabelOffline(tt *TestTools, product, label string) {
	_, err := blls.Label.Offline(getUidContext(), product, label)
	tt.Require.Nil(err)
}

func testLabelDelete(tt *TestTools, product, label string) {
	_, err := blls.Label.Delete(getUidContext(), product, label)
	tt.Require.Nil(err)
}
