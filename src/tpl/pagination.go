package tpl

import (
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"

	"github.com/teambition/gear"
)

// Search 搜索
type Search struct {
	Q string `json:"q" query:"q"`
}

// Validate escape and build MySQL LIKE pattern
func (s *Search) Validate() error {
	if s.Q != "" {
		s.Q = strings.ReplaceAll(s.Q, `\`, "-")
		s.Q = strings.ReplaceAll(s.Q, "%", `\%`)
		s.Q = strings.ReplaceAll(s.Q, "_", `\_`)
	}
	return nil
}

// Pagination 分页
type Pagination struct {
	Search
	PageToken string `json:"pageToken" query:"pageToken"`
	Skip      int    `json:"skip,omitempty" query:"skip"`
	PageSize  int    `json:"pageSize,omitempty" query:"pageSize"`
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

	if err := pg.Search.Validate(); err != nil {
		return err
	}

	return nil
}

// ConsolePagination 分页
type ConsolePagination struct {
	Search
	PageToken string `json:"pageToken" query:"pageToken"`
	Skip      int    `json:"skip,omitempty" query:"skip"`
	PageSize  int    `json:"pageSize,omitempty" query:"pageSize"`
}

// Validate ...
func (pg *ConsolePagination) Validate() error {
	if pg.PageToken != "" && !strings.HasPrefix(pg.PageToken, "hid.") {
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

	if err := pg.Search.Validate(); err != nil {
		return err
	}

	return nil
}

// GetNextPageToken ...
func (pg *ConsolePagination) GetNextPageToken() string {
	pg.Skip += pg.PageSize
	pageInfo := fmt.Sprintf("%d,%d", pg.Skip, pg.PageSize)
	return base64.URLEncoding.EncodeToString([]byte(pageInfo))
}
