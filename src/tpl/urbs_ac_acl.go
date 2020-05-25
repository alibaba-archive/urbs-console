package tpl

import (
	"github.com/asaskevich/govalidator"
	"github.com/teambition/gear"
)

// UrbsAcAclURL ...
type UrbsAcAclURL struct {
	Uid string `json:"uid" param:"uid"`
}

// Validate 实现 gear.BodyTemplate。
func (t *UrbsAcAclURL) Validate() error {
	return nil
}

// UrbsAcAclAddBody ...
type UrbsAcAclAddBody struct {
	Product string `json:"product" valid:"required"`
	Label   string `json:"label"`
	Module  string `json:"module"`
	Setting string `json:"setting"`

	Permission string `json:"permission"`
}

// Validate 实现 gear.BodyTemplate。
func (t *UrbsAcAclAddBody) Validate() error {
	if _, err := govalidator.ValidateStruct(t); err != nil {
		return gear.ErrBadRequest.WithMsg(err.Error())
	}
	return nil
}

// UrbsAcAclCheckBody ...
type UrbsAcAclCheckBody struct {
	Product string `json:"product"`
	Label   string `json:"label"`
	Module  string `json:"module"`
	Setting string `json:"setting"`
}

// Validate 实现 gear.BodyTemplate。
func (t *UrbsAcAclCheckBody) Validate() error {
	return nil
}
