package bll

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/teambition/urbs-console/src/dto"
	"github.com/teambition/urbs-console/src/service"
	"github.com/teambition/urbs-console/src/service/mock_service"
	"github.com/teambition/urbs-console/src/tpl"
)

func TestSetting(t *testing.T) {
	require := require.New(t)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	usMock := mock_service.NewMockUrbsSettingInterface(ctrl)

	uid := tpl.RandUID()
	object := tpl.RandUID()
	logContent := &dto.OperationLogContent{
		Users:   []string{tpl.RandUID()},
		Groups:  []string{tpl.RandUID()},
		Desc:    "desc",
		Value:   "true",
		Release: 1,
	}
	err := blls.OperationLog.Add(getUidContext(uid), object, actionCreate, logContent)
	require.Nil(err)

	t.Run("recall", func(t *testing.T) {
		setting := &Setting{services: service.NewServices()}

		// 1
		args := &tpl.ProductModuleSettingURL{}
		body := &tpl.RecallBody{
			HID: service.IDToHID(1000000, "log"),
		}
		_, err = setting.Recall(getUidContext(uid), args, body)
		require.NotNil(err)

		// 2
		log1, err := daos.OperationLog.FindOneByObject(nil, object)
		require.Nil(err)

		body = &tpl.RecallBody{
			HID: service.IDToHID(log1.ID, "log"),
		}

		_, err = setting.Recall(getUidContext(uid), args, body)
		require.NotNil(err)
	})

	t.Run("recall", func(t *testing.T) {
		setting := &Setting{services: service.NewServices()}
		setting.services.UrbsSetting = usMock

		args := &tpl.ProductModuleSettingURL{}
		log1, err := daos.OperationLog.FindOneByObject(nil, object)
		require.Nil(err)

		body := &tpl.RecallBody{
			HID:     service.IDToHID(log1.ID, "log"),
			Release: log1.ID,
		}

		usMock.EXPECT().SettingRecall(getUidContext(uid), args, body).Return(&tpl.BoolRes{}, nil)

		_, err = setting.Recall(getUidContext(uid), args, body)
		require.Nil(err)

		_, err = daos.OperationLog.FindOneByObject(nil, object)
		require.NotNil(err)
	})

}
