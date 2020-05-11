package tpl

import (
	"time"

	"github.com/teambition/gear"
)

// MySettingsQueryURL ...
type MySettingsQueryURL struct {
	Pagination
	UidsBody
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
	if err := t.UidsBody.Validate(); err != nil {
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
	Uids     *[]string `json:"uids"`
}

// Validate 实现 gear.BodyTemplate。
func (t *SettingUpdateBody) Validate() error {
	if t.Desc == nil && t.Channels == nil && t.Clients == nil && t.Values == nil && t.Uids == nil {
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
	if t.Uids != nil && len(*t.Uids) > 9 {
		return gear.ErrBadRequest.WithMsgf("uids length should 0 < %d < 10", len(*t.Uids))
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
	HID       string     `json:"hid"`
	Product   string     `json:"product"`
	Module    string     `json:"module"`
	Name      string     `json:"name"`
	Desc      string     `json:"desc"`
	Channels  []string   `json:"channels"`
	Clients   []string   `json:"clients"`
	Values    []string   `json:"values"`
	Status    int64      `json:"status"`
	Release   int64      `json:"release"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	OfflineAt *time.Time `json:"offlineAt"`
	Users     []*User    `json:"users"`
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

// SettingReleaseInfo ...
type SettingReleaseInfo struct {
	Release int64    `json:"release"`
	Users   []string `json:"users"`
	Groups  []string `json:"groups"`
	Value   string   `json:"value"`
}

// SettingReleaseInfoRes ...
type SettingReleaseInfoRes struct {
	SuccessResponseType
	Result SettingReleaseInfo `json:"result"` // 空数组也保留
}

// MySettingsRes ...
type MySettingsRes struct {
	SuccessResponseType
	Result []*MySetting `json:"result"` // 空数组也保留
}

// MySetting ...
type MySetting struct {
	ID         int64     `json:"-"`
	HID        string    `json:"hid"`
	Product    string    `json:"product"`
	Module     string    `json:"module"`
	Name       string    `json:"name"`
	Desc       string    `json:"desc"`
	Value      string    `json:"value"`
	LastValue  string    `json:"lastValue"`
	Release    int64     `json:"release"`
	AssignedAt time.Time `json:"assignedAt"`

	UpdatedAt time.Time `json:"updated_at"` //兼容
}
