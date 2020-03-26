package urbssetting

import (
	"regexp"

	"github.com/teambition/gear"
)

var validIDReg = regexp.MustCompile(`^[0-9A-Za-z._=-]{3,63}$`)

// SuccessResponseType 定义了标准的 API 接口成功时返回数据模型
type SuccessResponseType struct {
	TotalSize     int         `json:"totalSize,omitempty"`
	NextPageToken string      `json:"nextPageToken,omitempty"`
	Result        interface{} `json:"result"`
}

// BoolRes ...
type BoolRes struct {
	SuccessResponseType
	Result bool `json:"result"`
}

// UsersGroupsBody ...
type UsersGroupsBody struct {
	Users  []string `json:"users"`
	Groups []string `json:"groups"`
	Value  string   `json:"value"`
}

// Pagination 分页
type Pagination struct {
	PageToken string `json:"pageToken" query:"pageToken"`
	PageSize  int    `json:"pageSize,omitempty" query:"pageSize"`
	Skip      int    `json:"skip,omitempty" query:"skip"`
}

// Validate ...
func (pg *Pagination) Validate() error {
	if pg.Skip < 0 {
		pg.Skip = 0
	}

	if pg.PageSize > 1000 {
		return gear.ErrBadRequest.WithMsgf("pageSize(%v) should not great than 1000", pg.PageSize)
	}

	if pg.PageSize <= 0 {
		pg.PageSize = 10
	}

	return nil
}

// UIDPaginationURL ...
type UIDPaginationURL struct {
	Pagination
	UID string `json:"uid" param:"uid"`
}

// Validate 实现 gear.BodyTemplate。
func (t *UIDPaginationURL) Validate() error {
	if !validIDReg.MatchString(t.UID) {
		return gear.ErrBadRequest.WithMsgf("invalid uid: %s", t.UID)
	}
	if err := t.Pagination.Validate(); err != nil {
		return err
	}
	return nil
}

// UIDURL ...
type UIDURL struct {
	UID string `json:"uid" param:"uid"`
}

// UIDHIDURL ...
type UIDHIDURL struct {
	UIDURL
	HID string `json:"hid" param:"hid"`
}
