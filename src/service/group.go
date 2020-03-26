package service

import (
	"net/http"
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
	j := jwt.New(conf.Config.GroupMember.Keys)
	token, err := j.Sign(conf.Config.GroupMember.TokenKV, time.Hour)
	if err != nil {
		return nil, err
	}
	header := http.Header{}
	header.Set("Authorization", "Bearer "+token)

	body := make(map[string]interface{})
	for k, v := range conf.Config.GroupMember.BodyKK {
		switch v {
		case "groupId":
			body[k] = groupId
		case "pageSize":
			body[k] = pageSize
		case "skip":
			body[k] = skip
		}
	}
	result := new(ListGroupMembersResp)
	resp, err := request.Post(conf.Config.GroupMember.URL).Header(header).Body(body).Result(result).Do()
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
