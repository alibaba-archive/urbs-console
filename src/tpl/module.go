package tpl

import (
	"time"

	"github.com/teambition/gear"
)

// ModuleUpdateBody ...
type ModuleUpdateBody struct {
	Desc *string   `json:"desc"`
	Uids *[]string `json:"uids"`
}

// Validate 实现 gear.BodyTemplate。
func (t *ModuleUpdateBody) Validate() error {
	if t.Desc != nil && len(*t.Desc) > 1022 {
		return gear.ErrBadRequest.WithMsgf("desc too long: %d", len(*t.Desc))
	}
	if t.Uids != nil && len(*t.Uids) > 9 {
		return gear.ErrBadRequest.WithMsgf("uids length should 0 < %d < 10", len(*t.Uids))
	}
	return nil
}

// Module ...
type Module struct {
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	OfflineAt *time.Time `json:"offlineAt"`
	Name      string     `json:"name"`
	Desc      string     `json:"desc"`
	Status    int64      `json:"status"`
}

// ModuleInfo ...
type ModuleInfo struct {
	Module
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
