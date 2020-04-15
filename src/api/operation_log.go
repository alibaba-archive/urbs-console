package api

import (
	"github.com/teambition/gear"
	"github.com/teambition/urbs-console/src/bll"
	"github.com/teambition/urbs-console/src/tpl"
)

// OperationLog ..
type OperationLog struct {
	blls *bll.Blls
}

// Add ...
func (a *OperationLog) Add(ctx *gear.Context) error {
	return nil
}

// List 返回操作日志列表
func (a *OperationLog) List(ctx *gear.Context) error {
	req := &tpl.OperationLogListReq{}
	if err := ctx.ParseURL(req); err != nil {
		return err
	}
	res, err := a.blls.OperationLog.List(ctx, req)
	if err != nil {
		return err
	}
	return ctx.OkJSON(res)
}
