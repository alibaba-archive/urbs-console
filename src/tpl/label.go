package tpl

import (
	"time"

	"github.com/teambition/gear"
)

// LabelBody ...
type LabelBody struct {
	Name string   `json:"name"`
	Desc string   `json:"desc"`
	Uids []string `json:"uids"`
}

// Validate 实现 gear.BodyTemplate。
func (t *LabelBody) Validate() error {
	if !validLabelReg.MatchString(t.Name) {
		return gear.ErrBadRequest.WithMsgf("invalid label: %s", t.Name)
	}
	if len(t.Desc) > 1022 {
		return gear.ErrBadRequest.WithMsgf("desc too long: %d (<= 1022)", len(t.Desc))
	}
	return nil
}

// LabelUpdateBody ...
type LabelUpdateBody struct {
	Desc     *string   `json:"desc"`
	Channels *[]string `json:"channels"`
	Clients  *[]string `json:"clients"`
	Uids     []string  `json:"uids"`
}

// Validate 实现 gear.BodyTemplate。
func (t *LabelUpdateBody) Validate() error {
	if t.Desc == nil && t.Channels == nil && t.Clients == nil {
		return gear.ErrBadRequest.WithMsgf("desc or channels or clients required")
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
	return nil
}

// LabelInfo ...
type LabelInfo struct {
	ID        int64      `json:"-"`
	HID       string     `json:"hid"`
	Product   string     `json:"product"`
	Name      string     `json:"name"`
	Desc      string     `json:"desc"`
	Channels  []string   `json:"channels"`
	Clients   []string   `json:"clients"`
	Status    int64      `json:"status"`
	Release   int64      `json:"release"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	OfflineAt *time.Time `json:"offlineAt"`
	Users     []*User    `json:"users"`
}

// LabelInfoRes ...
type LabelInfoRes struct {
	Result LabelInfo `json:"result"`
}

// LabelsInfoRes ...
type LabelsInfoRes struct {
	SuccessResponseType
	Result []*LabelInfo `json:"result"`
}

// LabelGroupsRes ...
type LabelGroupsRes struct {
	SuccessResponseType
	Result []*Group `json:"result"`
}

// LabelUsersRes ...
type LabelUsersRes struct {
	SuccessResponseType
	Result []*User `json:"result"`
}

// LabelReleaseInfo ...
type LabelReleaseInfo struct {
	Release int64    `json:"release"`
	Users   []string `json:"users"`
	Groups  []string `json:"groups"`
}

// LabelReleaseInfoRes ...
type LabelReleaseInfoRes struct {
	SuccessResponseType
	Result LabelReleaseInfo `json:"result"` // 空数组也保留
}
