package bll

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/teambition/urbs-console/src/constant"
	"github.com/teambition/urbs-console/src/dto"
	"github.com/teambition/urbs-console/src/tpl"
)

func TestOperationLog(t *testing.T) {
	tt := SetUpTestTools(require.New(t))

	uid := tpl.RandUID()
	ctx := getUidContext(uid)
	object := tpl.RandName()

	body := new(tpl.UsersGroupsBody)
	body.Users = []string{tpl.RandUID()}
	body.Groups = []string{tpl.RandUID()}
	body.Desc = tpl.RandName()
	body.Value = "true"

	testAddUrbsAcUser(tt, uid)

	t.Run("add operationLog", func(t *testing.T) {
		percent := 2

		logContent := &dto.OperationLogContent{
			Users:   body.Users,
			Groups:  body.Groups,
			Desc:    body.Desc,
			Value:   body.Value,
			Percent: &percent,
		}
		err := testBlls.OperationLog.Add(ctx, object, constant.OperationCreate, logContent)
		require.Nil(t, err)
	})

	t.Run("get operationLog", func(t *testing.T) {
		require := require.New(t)
		// 获取操作日志
		page := &tpl.ConsolePagination{}
		page.Validate()
		res, err := testBlls.OperationLog.List(ctx, object, page)
		require.Nil(err)

		require.Equal(body.Users, res.Result[0].Users)
		require.Equal(body.Groups, res.Result[0].Groups)
		require.Equal(body.Desc, res.Result[0].Desc)
		require.Equal(body.Value, res.Result[0].Value)
		require.Equal(2, *res.Result[0].Percent)

		totalSize, err := testDaos.OperationLog.CountByObject(ctx, object)
		require.Nil(err)
		require.True(totalSize > 0)
	})

	t.Run("operationLog content", func(t *testing.T) {
		percent := 2

		require := require.New(t)
		object := tpl.RandName()
		logContent := &dto.OperationLogContent{
			Users:   body.Users,
			Groups:  body.Groups,
			Desc:    tpl.RandName(),
			Value:   body.Value,
			Percent: &percent,
			Release: 111,
		}
		err := testBlls.OperationLog.Add(ctx, object, constant.OperationCreate, logContent)
		require.Nil(err)

		log, err := testDaos.OperationLog.FindOneByObject(ctx, object)
		require.Nil(err)
		require.Equal(logContent.Desc, log.Desc)

		// 2
		release := getRelease(log.Content)
		require.Equal(int64(111), release)

		// 3
		item := &tpl.OperationLogListItem{}
		parseLogContent(log.Content, item)

		require.Equal(body.Users, item.Users)
		require.Equal(body.Groups, item.Groups)
		require.Equal(body.Value, item.Value)
		require.Equal("userPercent", item.Kind)
		require.Equal(2, *item.Percent)
	})

	t.Run("FindAll", func(t *testing.T) {
		require := require.New(t)
		logs, err := testDaos.OperationLog.FindAll(nil)
		require.Nil(err)
		require.True(len(logs) > 0, len(logs))
	})
}
