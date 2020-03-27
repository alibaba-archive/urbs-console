package tpl

import "github.com/teambition/gear"

// ModuleUpdateBody ...
type ModuleUpdateBody struct {
	Desc *string `json:"desc"`
}

// Validate 实现 gear.BodyTemplate。
func (t *ModuleUpdateBody) Validate() error {
	if t.Desc == nil {
		return gear.ErrBadRequest.WithMsgf("desc required")
	}

	if len(*t.Desc) > 1022 {
		return gear.ErrBadRequest.WithMsgf("desc too long: %d", len(*t.Desc))
	}
	return nil
}
