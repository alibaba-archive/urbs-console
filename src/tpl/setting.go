package tpl

import (
	"github.com/teambition/gear"
	"github.com/teambition/urbs-console/src/schema"
)

// MySettingsQueryURL ...
type MySettingsQueryURL struct {
	Pagination
	UID     string `json:"uid" param:"uid"`
	Product string `json:"product" query:"product"`
	Channel string `json:"channel" query:"channel"`
	Client  string `json:"client" query:"client"`
}

// Validate 实现 gear.BodyTemplate。
func (t *MySettingsQueryURL) Validate() error {
	if !validIDReg.MatchString(t.UID) {
		return gear.ErrBadRequest.WithMsgf("invalid user: %s", t.UID)
	}
	if !validNameReg.MatchString(t.Product) {
		return gear.ErrBadRequest.WithMsgf("invalid product name: %s", t.Product)
	}

	if err := t.Pagination.Validate(); err != nil {
		return err
	}
	return nil
}

// SettingUpdateBody ...
type SettingUpdateBody struct {
	Desc     *string   `json:"desc"`
	Channels *[]string `json:"channels"`
	Clients  *[]string `json:"clients"`
	Values   *[]string `json:"values"`
	Uids     []string  `json:"uids"`
}

// Validate 实现 gear.BodyTemplate。
func (t *SettingUpdateBody) Validate() error {
	if t.Desc == nil && t.Channels == nil && t.Clients == nil && t.Values == nil {
		return gear.ErrBadRequest.WithMsgf("desc or channels or clients or values required")
	}
	if t.Desc != nil && len(*t.Desc) > 1022 {
		return gear.ErrBadRequest.WithMsgf("desc too long: %d", len(*t.Desc))
	}
	if t.Channels != nil {
		if len(*t.Channels) > 5 {
			return gear.ErrBadRequest.WithMsgf("too many channels: %d", len(*t.Channels))
		}
		if !SortStringsAndCheck(*t.Channels) {
			return gear.ErrBadRequest.WithMsgf("invalid channels: %v", *t.Channels)
		}
	}
	if t.Clients != nil {
		if len(*t.Clients) > 10 {
			return gear.ErrBadRequest.WithMsgf("too many clients: %d", len(*t.Clients))
		}
		if !SortStringsAndCheck(*t.Clients) {
			return gear.ErrBadRequest.WithMsgf("invalid clients: %v", *t.Clients)
		}

	}
	if t.Values != nil {
		if len(*t.Values) > 10 {
			return gear.ErrBadRequest.WithMsgf("too many values: %d", len(*t.Clients))
		}
		if !SortStringsAndCheck(*t.Values) {
			return gear.ErrBadRequest.WithMsgf("invalid values: %v", *t.Values)
		}
	}
	return nil
}

// MySettingsQueryURLClient ...
type MySettingsQueryURLClient struct {
	Pagination
	Product string `json:"product" query:"product"`
	Channel string `json:"channel" query:"channel"`
	Client  string `json:"client" query:"client"`
}

// Validate 实现 gear.BodyTemplate。
func (t *MySettingsQueryURLClient) Validate() error {
	if !validNameReg.MatchString(t.Product) {
		return gear.ErrBadRequest.WithMsgf("invalid product name: %s", t.Product)
	}

	if err := t.Pagination.Validate(); err != nil {
		return err
	}
	return nil
}

// SettingInfo ...
type SettingInfo struct {
	schema.Setting
	Users []*User `json:"users"`
}

// SettingInfoRes ...
type SettingInfoRes struct {
	Result SettingInfo `json:"result"`
}

// SettingsInfoRes ...
type SettingsInfoRes struct {
	SuccessResponseType
	Result []*SettingInfo `json:"result"`
}
