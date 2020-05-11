package tpl

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
	Product string `json:"product"`
	Label   string `json:"label"`
	Module  string `json:"module"`
	Setting string `json:"setting"`

	Permission string `json:"permission"`
}

// Validate 实现 gear.BodyTemplate。
func (t *UrbsAcAclAddBody) Validate() error {
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
