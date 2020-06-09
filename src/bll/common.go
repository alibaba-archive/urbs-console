package bll

import (
	"github.com/teambition/urbs-console/src/dto"
	"github.com/teambition/urbs-console/src/tpl"
)

type operationLogAdd struct {
	Object  string                   `json:"object"`
	Content *dto.OperationLogContent `json:"content"`
	Action  string                   `json:"Action"`
}

type settingRecallReq struct {
	Args *tpl.ProductModuleSettingURL `json:"args"`
	Body *tpl.RecallBody              `json:"body"`
}
