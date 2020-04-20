package bll

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/teambition/urbs-console/src/service"
	"github.com/teambition/urbs-console/src/service/mock_service"
	"github.com/teambition/urbs-console/src/tpl"
)

func TestUser(t *testing.T) {
	require := require.New(t)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	usMock := mock_service.NewMockUrbsSettingInterface(ctrl)

	user := &User{services: service.NewServices()}
	user.services.UrbsSetting = usMock

	mockReturn := new(tpl.BoolRes)
	mockReturn.Result = true
	usMock.EXPECT().UserBatchAdd(nil, []string{"123"}).Return(mockReturn, nil)

	_, err := user.BatchAdd(nil, []string{"123"})
	require.Nil(err)
}

func testUserBatchCreate(tt *TestTools, uid string) {
	_, err := blls.User.BatchAdd(context.Background(), []string{uid})

	tt.Require.Nil(err)
}

func testUserListLables(tt *TestTools, uid string, count int) {
	args := &tpl.UIDPaginationURL{
		UID: uid,
	}
	res, err := blls.User.ListLables(getUidContext(), args)

	tt.Require.Nil(err)
	tt.Require.Equal(count, len(res.Result))
}
