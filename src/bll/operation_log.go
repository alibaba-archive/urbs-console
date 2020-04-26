package bll

import (
	"context"
	"strconv"
	"strings"

	"github.com/teambition/urbs-console/src/dao"
	"github.com/teambition/urbs-console/src/schema"
	"github.com/teambition/urbs-console/src/tpl"
	"github.com/teambition/urbs-console/src/util"
)

// OperationLog table `operation_log`
type OperationLog struct {
	daos *dao.Daos
}

var (
	actionCreate = "create"
	actionDelete = "delete"
)

// List 返回操作日志列表
func (a *OperationLog) List(ctx context.Context, object string, req *tpl.Pagination) (*tpl.OperationLogListRes, error) {
	logs, err := a.daos.OperationLog.FindByObject(ctx, object, req)
	if err != nil {
		return nil, err
	}
	totalSize, err := a.daos.OperationLog.CountByObject(ctx, object)
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
		parseLogContent(log.Content, item)
		items[i] = item
	}
	res := &tpl.OperationLogListRes{Result: items}
	res.TotalSize = totalSize
	if len(res.Result) > req.PageSize {
		res.NextPageToken = req.GetNextPageToken()
		res.Result = items[:req.PageSize]
	}
	return res, nil
}

// Add ...
func (a *OperationLog) Add(ctx context.Context, object string, action string, body *tpl.UsersGroupsBody) error {
	log := &schema.OperationLog{
		Operator: util.GetUid(ctx),
		Object:   object,
		Action:   action,
		Content:  genContent(body),
		Desc:     body.Desc,
	}
	return a.daos.OperationLog.Add(ctx, log)
}

func genContent(body *tpl.UsersGroupsBody) string {
	content := "01"
	if len(body.Users) > 0 {
		content += "01" + strings.Join(body.Users, ",") + "\r\n"
	}
	if len(body.Groups) > 0 {
		content += "02" + strings.Join(body.Groups, ",") + "\r\n"
	}
	if body.Value != "" {
		content += "03" + body.Value + "\r\n"
	}
	if body.Percent > 0 {
		content += "04" + strconv.Itoa(body.Percent)
	}
	return content
}

func parseLogContent(content string, log *tpl.OperationLogListItem) {
	content = content[2:]

	items := strings.Split(content, "\r\n")
	for _, item := range items {
		kind := item[0:2]
		content := item[2:]
		switch kind {
		case "01": // users
			log.Users = strings.Split(content, ",")
		case "02": // groups
			log.Groups = strings.Split(content, ",")
		case "03":
			log.Value = content
		case "04": // percent
			log.Percent, _ = strconv.Atoi(content)
		}
	}
}
