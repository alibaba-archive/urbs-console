package service

import (
	"net/url"
	"strconv"
	"strings"

	"github.com/teambition/urbs-console/src/conf"
	"github.com/teambition/urbs-console/src/util/request"
)

// GroupMember ...
type GroupMember struct {
}

// List ...
func (a *GroupMember) List(groupId string, skip int, pageSize int) (*ListGroupMembersResp, error) {
	groupUrl := strings.Replace(conf.Config.Thrid.GroupMember.URL, "{groupId}", groupId, -1)
	httpUrl, err := url.Parse(groupUrl)
	if err != nil {
		return nil, err
	}
	q := httpUrl.Query()
	q.Add("pageSize", strconv.Itoa(pageSize))
	q.Add("skip", strconv.Itoa(skip))
	httpUrl.RawQuery = q.Encode()

	result := new(ListGroupMembersResp)
	resp, err := request.Get(httpUrl.String()).Header(genThridHeader()).Result(result).Do()
	if err := HanderResponse(resp, err); err != nil {
		return nil, err
	}
	return result, nil
}

// ListGroupMembersResp ...
type ListGroupMembersResp struct {
	Members []*Member `json:"result"`
}

// Member ....
type Member struct {
	UID string `json:"uid"`
}
