package tpl

import "github.com/teambition/gear"

// OperationLogListReq ...
type OperationLogListReq struct {
	Pagination
	ProductURL
	Label   string `json:"label" query:"label"`
	Module  string `json:"module" query:"module"`
	Setting string `json:"setting" query:"setting"`
}

// Validate 实现 gear.BodyTemplate。
func (t *OperationLogListReq) Validate() error {
	if err := t.Pagination.Validate(); err != nil {
		return err
	}
	if err := t.ProductURL.Validate(); err != nil {
		return err
	}
	if t.Label != "" {
		if !validLabelReg.MatchString(t.Label) {
			return gear.ErrBadRequest.WithMsgf("invalid label: %s", t.Label)
		}
	} else {
		if !validNameReg.MatchString(t.Module) {
			return gear.ErrBadRequest.WithMsgf("invalid module name: %s", t.Module)
		}
		if !validNameReg.MatchString(t.Setting) {
			return gear.ErrBadRequest.WithMsgf("invalid setting name: %s", t.Setting)
		}
	}
	return nil
}

// OperationLogListRes ...
type OperationLogListRes struct {
	SuccessResponseType
	Result []*OperationLogListItem `json:"result"` // 空数组也保留
}

// OperationLogListItem ...
type OperationLogListItem struct {
	Operator     string `json:"operator"`     // 操作人
	OperatorName string `json:"operatorName"` // 操作人
	Action       string `json:"action"`       // 操作行为
	Desc         string `json:"description"`  // 操作说明

	Groups     []string `json:"groups,omitempty"`     // 群组
	Users      []string `json:"users,omitempty"`      // 用户
	Value      string   `json:"value,omitempty"`      // 灰度百分比
	Percentage int      `json:"percentage,omitempty"` // 灰度百分比
}
