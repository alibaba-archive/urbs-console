package tpl

import (
	"github.com/teambition/gear"
)

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
