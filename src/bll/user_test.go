package bll

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/teambition/urbs-console/src/dto/urbssetting"
	"github.com/teambition/urbs-console/src/service"
	"github.com/teambition/urbs-console/src/service/mock_service"
)

func TestUser(t *testing.T) {
	require := require.New(t)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	usMock := mock_service.NewMockUrbsSettingInterface(ctrl)

	user := &User{services: service.NewServices()}
	user.services.UrbsSetting = usMock

	mockReturn := new(urbssetting.BoolRes)
	mockReturn.Result = true
	usMock.EXPECT().UserBatchAdd(nil, []string{"123"}).Return(mockReturn, nil)

	err := user.BatchAdd(nil, []string{"123"})
	require.Nil(err)
}
