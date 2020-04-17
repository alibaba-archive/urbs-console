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
		parseLogContent(log.Content, item)
		items[i] = item
	}
	res := &tpl.OperationLogListRes{Result: items}
	if len(res.Result) > req.PageSize {
		res.NextPageToken = req.GetNextPageToken()
		res.Result = items[:req.PageSize]
	}
	return res, nil
}

// AddSettingAssignLog ...
func (a *OperationLog) AddSettingAssignLog(ctx context.Context, args *tpl.ProductModuleSettingURL, body *tpl.UsersGroupsBody) error {
	log := &schema.OperationLog{
		Operator: util.GetUid(ctx),
		Object:   args.Product + args.Module + args.Setting,
		Action:   "create",
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
	if body.Percentage > 0 {
		content += "04" + strconv.Itoa(body.Percentage)
	}
	return content
}

func parseLogContent(content string, log *tpl.OperationLogListItem) {

	items := strings.Split(content, "\r\n")
	for _, item := range items {
		kind := item[2:4]
		content := item[4:]
		switch kind {
		case "01": // users
			log.Users = strings.Split(content, ",")
		case "02": // groups
			log.Groups = strings.Split(content, ",")
		case "03":
			log.Value = content
		case "04": // percentage
			log.Percentage, _ = strconv.Atoi(content)
		}
	}
}
