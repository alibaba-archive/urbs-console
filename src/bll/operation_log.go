package bll

import (
	"context"
	"strings"

	"github.com/teambition/urbs-console/src/dao"
	"github.com/teambition/urbs-console/src/schema"
	"github.com/teambition/urbs-console/src/tpl"
)

// OperationLog table `operation_log`
type OperationLog struct {
	daos *dao.Daos
}

// Add ...
func (a *OperationLog) Add(ctx context.Context, obj *schema.OperationLog) error {
	return a.daos.OperationLog.Add(ctx, obj)
}

// List 返回操作日志列表
func (a *OperationLog) List(ctx context.Context, req *tpl.OperationLogListReq) (*tpl.OperationLogListRes, error) {
	objectLog := req.Product + req.Label
	if req.Label == "" {
		objectLog = req.Product + req.Module + req.Setting
	}
	logs, err := a.daos.OperationLog.FindByObject(ctx, objectLog, &req.Pagination)
	if err != nil {
		return nil, err
	}
	items := make([]*tpl.OperationLogListItem, len(logs))
	for i, log := range logs {
		item := &tpl.OperationLogListItem{
			Operator:     log.Operator,
			OperatorName: log.Name,
			Action:       log.Action,
			Desc:         log.Desc,
		}
		items[i] = item
	}
	res := &tpl.OperationLogListRes{Result: items}
	if len(res.Result) > req.PageSize {
		res.NextPageToken = req.GetNextPageToken()
		res.Result = items[:req.PageSize]
	}
	return res, nil
}

func parseLogContent(content string) ([]string, []string, int) {
	var users []string
	var groups []string
	var percentage int

	items := strings.Split(content, "\r\n")
	for _, item := range items {
		kind := item[2:4]
		content := item[4:]
		switch kind {
		case "01": // users
			users = strings.Split(content, ",")
		case "02": // groups
			groups = strings.Split(content, ",")
		case "03": // percentage

		}
	}
	return users, groups, percentage
}
