package tpl

import (
	"github.com/teambition/gear"
	"github.com/teambition/urbs-console/src/schema"
)

// ModuleUpdateBody ...
type ModuleUpdateBody struct {
	Desc *string  `json:"desc"`
	Uids []string `json:"uids"`
}

// Validate 实现 gear.BodyTemplate。
func (t *ModuleUpdateBody) Validate() error {
	if t.Desc != nil && len(*t.Desc) > 1022 {
		return gear.ErrBadRequest.WithMsgf("desc too long: %d", len(*t.Desc))
	}
	return nil
}

// ModuleInfo ...
type ModuleInfo struct {
	schema.Module
	Users []*User `json:"users"`
}

// ModuleInfoRes ...
type ModuleInfoRes struct {
	SuccessResponseType
	Result ModuleInfo `json:"result"`
}

// ModulesInfoRes ...
type ModulesInfoRes struct {
	SuccessResponseType
	Result []*ModuleInfo `json:"result"`
}
