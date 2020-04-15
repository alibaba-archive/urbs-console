package tpl

import (
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"

	"github.com/teambition/gear"
)

// Pagination 分页
type Pagination struct {
	PageToken string `json:"pageToken" query:"pageToken"`
	Skip      int    `json:"skip,omitempty" query:"skip"`
	PageSize  int    `json:"pageSize,omitempty" query:"pageSize"`
}

// Validate ...
func (pg *Pagination) Validate() error {
	if pg.PageToken != "" {
		pageToken, err := base64.URLEncoding.DecodeString(pg.PageToken)
		if err != nil {
			return gear.ErrBadRequest.WithMsgf("invalid PageToken: %v", pg.PageToken)
		}
		pageInfo := strings.Split(string(pageToken), ",")
		if len(pageInfo) == 2 {
			pg.Skip, _ = strconv.Atoi(pageInfo[0])
			pg.PageSize, _ = strconv.Atoi(pageInfo[1])
		}
	}

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

// GetNextPageToken ...
func (pg *Pagination) GetNextPageToken() string {
	pg.Skip += pg.PageSize
	pageInfo := fmt.Sprintf("%d,%d", pg.Skip, pg.PageSize)
	return base64.URLEncoding.EncodeToString([]byte(pageInfo))
}
