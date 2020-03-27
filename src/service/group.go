package service

import (
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/teambition/gear-auth/jwt"
	"github.com/teambition/urbs-console/src/conf"
	"github.com/teambition/urbs-console/src/util/request"
)

// GroupMember ...
type GroupMember struct {
}

// List ...
func (a *GroupMember) List(groupId string, skip int, pageSize int) (*ListGroupMembersResp, error) {
	j := jwt.New([]byte(conf.Config.GroupMember.Key))
	token, err := j.Sign(conf.Config.GroupMember.TokenKV, time.Hour)
	if err != nil {
		return nil, err
	}
	header := http.Header{}
	header.Set("Authorization", "Bearer "+token)

	groupUrl := strings.Replace(conf.Config.GroupMember.URL, "{groupId}", groupId, -1)
	httpUrl, err := url.Parse(groupUrl)
	if err != nil {
		return nil, err
	}
	q := httpUrl.Query()
	q.Add("pageSize", strconv.Itoa(pageSize))
	q.Add("skip", strconv.Itoa(skip))
	for k, v := range conf.Config.GroupMember.BodyKK {
		switch k {
		case "groupId":
			q.Add(v, groupId)
		}
	}
	httpUrl.RawQuery = q.Encode()

	result := new(ListGroupMembersResp)
	resp, err := request.Get(httpUrl.String()).Header(header).Result(result).Do()
	if err := HanderResponse(resp, err); err != nil {
		return nil, err
	}
	return result, nil
}

// ListGroupMembersResp ...
type ListGroupMembersResp struct {
	Members   []*Member `json:"result"`
	TotalSize int       `json:"totalSize"`
}

// Member ....
type Member struct {
	UserID string `json:"_userId"` // 兼容 Teambition，后面要去掉
	UID    string `json:"uid"`
}

// GetUID ...
func (a *Member) GetUID() string {
	if a.UID != "" {
		return a.UID
	}
	return a.UserID
}
