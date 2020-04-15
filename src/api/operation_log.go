package api

import (
	"github.com/teambition/gear"
	"github.com/teambition/urbs-console/src/bll"
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
	return nil
}
