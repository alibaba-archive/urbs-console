package tpl

import (
	"time"

	"github.com/teambition/gear"
)

// ProductURL ...
type ProductURL struct {
	Product string `json:"product" param:"product"`
}

// Validate 实现 gear.BodyTemplate。
func (t *ProductURL) Validate() error {
	if !validNameReg.MatchString(t.Product) {
		return gear.ErrBadRequest.WithMsgf("invalid product name: %s", t.Product)
	}
	return nil
}

// ProductUpdateBody ...
type ProductUpdateBody struct {
	Desc *string   `json:"desc"`
	Uids *[]string `json:"uids"`
}

// Validate 实现 gear.BodyTemplate。
func (t *ProductUpdateBody) Validate() error {
	if t.Desc != nil && len(*t.Desc) > 1022 {
		return gear.ErrBadRequest.WithMsgf("desc too long: %d", len(*t.Desc))
	}
	if t.Uids != nil {
		if len(*t.Uids) > 9 {
			return gear.ErrBadRequest.WithMsgf("uids length should 0 < %d < 10", len(*t.Uids))
		}
		if !SortStringsAndCheck(*t.Uids) {
			return gear.ErrBadRequest.WithMsgf("invalid uids: %v", *t.Uids)
		}
	}
	return nil
}

// ProductPaginationURL ...
type ProductPaginationURL struct {
	Pagination
	Product string `json:"product" param:"product"`
}

// Validate 实现 gear.BodyTemplate。
func (t *ProductPaginationURL) Validate() error {
	if !validNameReg.MatchString(t.Product) {
		return gear.ErrBadRequest.WithMsgf("invalid product name: %s", t.Product)
	}
	if err := t.Pagination.Validate(); err != nil {
		return err
	}
	return nil
}

// ProductLabelURL ...
type ProductLabelURL struct {
	ProductPaginationURL
	Label string `json:"label" param:"label"`
}

// Validate 实现 gear.BodyTemplate。
func (t *ProductLabelURL) Validate() error {
	if err := t.ProductPaginationURL.Validate(); err != nil {
		return err
	}
	if !validLabelReg.MatchString(t.Label) {
		return gear.ErrBadRequest.WithMsgf("invalid label: %s", t.Label)
	}
	return nil
}

// UIDProductURL ...
type UIDProductURL struct {
	Pagination
	UID     string `json:"uid" param:"uid"`
	Product string `json:"product" query:"product"`
}

// Validate 实现 gear.BodyTemplate。
func (t *UIDProductURL) Validate() error {
	if !validIDReg.MatchString(t.UID) {
		return gear.ErrBadRequest.WithMsgf("invalid uid: %s", t.UID)
	}
	if !validNameReg.MatchString(t.Product) {
		return gear.ErrBadRequest.WithMsgf("invalid product name: %s", t.Product)
	}

	if err := t.Pagination.Validate(); err != nil {
		return err
	}
	return nil
}

// ProductModuleURL ...
type ProductModuleURL struct {
	Pagination
	ProductURL
	Module string `json:"module" param:"module"`
}

// Validate 实现 gear.BodyTemplate。
func (t *ProductModuleURL) Validate() error {
	if !validNameReg.MatchString(t.Module) {
		return gear.ErrBadRequest.WithMsgf("invalid module name: %s", t.Module)
	}
	if err := t.ProductURL.Validate(); err != nil {
		return err
	}
	if err := t.Pagination.Validate(); err != nil {
		return err
	}
	return nil
}

// ProductModuleSettingURL ...
type ProductModuleSettingURL struct {
	ProductModuleURL
	Setting string `json:"setting" param:"setting"`
}

// Validate 实现 gear.BodyTemplate。
func (t *ProductModuleSettingURL) Validate() error {
	if err := t.ProductModuleURL.Validate(); err != nil {
		return err
	}
	if !validNameReg.MatchString(t.Setting) {
		return gear.ErrBadRequest.WithMsgf("invalid setting name: %s", t.Setting)
	}
	return nil
}

// Product ...
type Product struct {
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
	OfflineAt *time.Time `json:"offlineAt"`
	Name      string     `json:"name"`
	Desc      string     `json:"desc"`
	Status    int64      `json:"status"`
	Users     []*User    `json:"users"`
}

// ProductRes ...
type ProductRes struct {
	SuccessResponseType
	Result Product `json:"result"`
}

// ProductsRes ...
type ProductsRes struct {
	SuccessResponseType
	Result []*Product `json:"result"`
}

// User ...
type User struct {
	Uid  string `json:"uid"`
	Name string `json:"name"`
}

// ProductLabelPaginationURL ...
type ProductLabelPaginationURL struct {
	Pagination
	ProductURL
	Label string `json:"label" param:"label"`
}

// Validate 实现 gear.BodyTemplate。
func (t *ProductLabelPaginationURL) Validate() error {
	if err := t.ProductURL.Validate(); err != nil {
		return err
	}
	if !validLabelReg.MatchString(t.Label) {
		return gear.ErrBadRequest.WithMsgf("invalid label: %s", t.Label)
	}
	if err := t.Pagination.Validate(); err != nil {
		return err
	}
	return nil
}

// ProductStatisticsRes ...
type ProductStatisticsRes struct {
	Result ProductStatistics `json:"result"`
}

// ProductStatistics ...
type ProductStatistics struct {
	Labels   int64 `json:"labels"`
	Modules  int64 `json:"modules"`
	Settings int64 `json:"settings"`
	Release  int64 `json:"release"`
	Status   int64 `json:"status"`
}

// ProductModuleSettingHIDURL ...
type ProductModuleSettingHIDURL struct {
	ProductModuleSettingURL
	HID string `json:"hid" param:"hid"`
}

// Validate 实现 gear.BodyTemplate。
func (t *ProductModuleSettingHIDURL) Validate() error {
	if !validHIDReg.MatchString(t.HID) {
		return gear.ErrBadRequest.WithMsgf("invalid hid: %s", t.HID)
	}
	if err := t.ProductModuleSettingURL.Validate(); err != nil {
		return err
	}
	return nil
}

// ProductLabelHIDURL ...
type ProductLabelHIDURL struct {
	ProductLabelURL
	HID string `json:"hid" param:"hid"`
}

// Validate 实现 gear.BodyTemplate。
func (t *ProductLabelHIDURL) Validate() error {
	if !validHIDReg.MatchString(t.HID) {
		return gear.ErrBadRequest.WithMsgf("invalid hid: %s", t.HID)
	}
	if err := t.ProductLabelURL.Validate(); err != nil {
		return err
	}
	return nil
}

// ProductLabelUIDURL ...
type ProductLabelUIDURL struct {
	ProductLabelURL
	UID string `json:"uid" param:"uid"`
}

// Validate 实现 gear.BodyTemplate。
func (t *ProductLabelUIDURL) Validate() error {
	if !validIDReg.MatchString(t.UID) {
		return gear.ErrBadRequest.WithMsgf("invalid uid: %s", t.UID)
	}
	if err := t.ProductLabelURL.Validate(); err != nil {
		return err
	}
	return nil
}

// ProductModuleSettingUIDURL ...
type ProductModuleSettingUIDURL struct {
	ProductModuleSettingURL
	UID string `json:"uid" param:"uid"`
}

// Validate 实现 gear.BodyTemplate。
func (t *ProductModuleSettingUIDURL) Validate() error {
	if !validIDReg.MatchString(t.UID) {
		return gear.ErrBadRequest.WithMsgf("invalid uid: %s", t.UID)
	}
	if err := t.ProductModuleSettingURL.Validate(); err != nil {
		return err
	}
	return nil
}
