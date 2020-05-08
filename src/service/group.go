package service

import (
	"context"
	"net/url"
	"strconv"
	"strings"

	"github.com/mushroomsir/request"
	"github.com/teambition/urbs-console/src/conf"
)

// GroupMember ...
type GroupMember struct {
}

// List ...
func (a *GroupMember) List(ctx context.Context, groupId string, pageToken string, pageSize int) (*ListGroupMembersResp, error) {
	groupUrl := strings.Replace(conf.Config.Thrid.GroupMember.URL, "{groupId}", groupId, -1)
	httpUrl, err := url.Parse(groupUrl)
	if err != nil {
		return nil, err
	}
	q := httpUrl.Query()
	q.Add("pageSize", strconv.Itoa(pageSize))
	q.Add("pageToken", pageToken)
	httpUrl.RawQuery = q.Encode()

	result := new(ListGroupMembersResp)
	resp, err := request.Get(httpUrl.String()).Header(ThridHeader(ctx)).Result(result).Do()
	if err := HanderResponse(resp, err); err != nil {
		return nil, err
	}
	return result, nil
}

// ListGroupMembersResp ...
type ListGroupMembersResp struct {
	Members       []*Member `json:"result"`
	NextPageToken string    `json:"nextPageToken"`
}

// Member ....
type Member struct {
	UID string `json:"uid"`
}
