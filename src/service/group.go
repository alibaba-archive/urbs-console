package service

import (
	"context"
	"net/http"
	"time"

	"github.com/teambition/gear-auth/jwt"
	"github.com/teambition/urbs-console/src/conf"
	"github.com/teambition/urbs-console/src/util"
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
	resp := new(ListGroupMembersResp)
	err = util.RequestPost(context.Background(), conf.Config.GroupMember.URL, header, body, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// ListGroupMembersResp ...
type ListGroupMembersResp struct {
	Members   []*Member `json:"result"`
	TotalSize int       `json:"totalSize"`
}

// Member ....
type Member struct {
	UserID string `json:"_userId"`
}
