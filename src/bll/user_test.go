package bll

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/teambition/urbs-console/src/service"
	"github.com/teambition/urbs-console/src/service/mock_service"
	"github.com/teambition/urbs-console/src/tpl"
	"github.com/teambition/urbs-setting/src/schema"
)

func TestUser(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	usMock := mock_service.NewMockUrbsSettingInterface(ctrl)

	user := &User{services: service.NewServices(testDB)}
	user.services.UrbsSetting = usMock

	t.Run(`BatchAdd`, func(t *testing.T) {
		require := require.New(t)

		mockReturn := new(tpl.BoolRes)
		mockReturn.Result = true
		usMock.EXPECT().UserBatchAdd(nil, []string{"123"}).Return(mockReturn, nil)

		_, err := user.BatchAdd(nil, []string{"123"})
		require.Nil(err)
	})

	t.Run(`UserListSettingsUnionAll should work`, func(t *testing.T) {
		assert := assert.New(t)

		uid := tpl.RandUID()
		product := tpl.RandName()
		label := tpl.RandLabel()

		module := tpl.RandName()
		setting := tpl.RandName()
		HID := tpl.RandUID()

		value := tpl.RandName()

		args := &tpl.MySettingsQueryURL{
			UID:     uid,
			Product: product,
			Client:  "ios",
		}
		mockReturn := &tpl.MySettingsRes{
			Result: []*tpl.MySetting{{
				HID:     HID,
				Product: product,
				Module:  module,
				Name:    setting,
				Value:   value,
			}},
		}
		usMock.EXPECT().UserListSettingsUnionAll(nil, args).Return(mockReturn, nil)

		mockCacheLabelsInfoRes := &tpl.CacheLabelsInfoRes{
			Result: []schema.UserCacheLabel{
				{
					Label:   label,
					Clients: []string{"ios"},
				},
			},
		}
		usMock.EXPECT().LabelsCache(nil, product, uid).Return(mockCacheLabelsInfoRes, nil)

		res, err := user.ListSettingsUnionAll(nil, args)
		assert.Nil(err)
		assert.Equal(product, res.Result[0].Product)
		assert.Equal("urbs", res.Result[0].Module)
		assert.Equal("", res.Result[0].HID)
		assert.Equal(label, res.Result[0].Name)
		assert.Equal("true", res.Result[0].Value)

		assert.Equal(product, res.Result[1].Product)
		assert.Equal(module, res.Result[1].Module)
		assert.Equal(HID, res.Result[1].HID)
		assert.Equal(setting, res.Result[1].Name)
		assert.Equal(value, res.Result[1].Value)
	})
}

func testUserBatchCreate(tt *TestTools, uid string) {
	_, err := testBlls.User.BatchAdd(context.Background(), []string{uid})

	tt.Require.Nil(err)
}

func testUserListLables(tt *TestTools, uid string, count int) {
	args := &tpl.UIDPaginationURL{
		UID: uid,
	}
	res, err := testBlls.User.ListLables(getUidContext(), args)

	tt.Require.Nil(err)
	tt.Require.Equal(count, len(res.Result))
}
