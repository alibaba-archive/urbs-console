package service

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/teambition/gear"
	"github.com/teambition/gear-auth/jwt"
	"github.com/teambition/urbs-console/src/conf"
	"github.com/teambition/urbs-console/src/util"
	"github.com/teambition/urbs-console/src/util/request"
)

func init() {
	util.DigProvide(NewServices)
}

// Services ...
type Services struct {
	UrbsSetting UrbsSettingInterface
	UserAuth    *UserAuth
	GroupMember *GroupMember
}

// NewServices ...
func NewServices() *Services {
	s := &Services{
		GroupMember: &GroupMember{},
		UrbsSetting: &UrbsSetting{},
		UserAuth:    &UserAuth{},
	}
	return s
}

// UrbsSettingHeader ...
func UrbsSettingHeader(ctx context.Context) http.Header {
	header := http.Header{}
	requestId, _ := ctx.Value(gear.HeaderXRequestID).(string)
	if requestId == "" {
		if gearCtx, ok := ctx.(*gear.Context); ok {
			requestId = gearCtx.GetHeader(gear.HeaderXRequestID)
			if requestId == "" {
				requestId = gearCtx.Res.Header().Get(gear.HeaderXRequestID)
			}
			header.Set("X-Canary", gearCtx.GetHeader("X-Canary"))
		}
	}
	if requestId != "" {
		header.Set(gear.HeaderXRequestID, requestId)
	}
	header.Set("Authorization", "Bearer "+genToken(conf.Config.UrbsSetting.Key))
	return header
}

// HanderResponse ...
func HanderResponse(response *request.Response, err error) error {
	if err != nil {
		return gear.ErrBadRequest.WithMsg(err.Error())
	}
	if !response.OK() {
		gearErr := new(gear.Error)
		json.Unmarshal(response.Content, gearErr)
		if gearErr.Err != "" {
			return gearErr.WithCode(response.StatusCode)
		}
		return gear.ErrBadRequest.WithCode(response.StatusCode).WithMsg(response.String())
	}
	return nil
}

func genThridHeader() http.Header {
	header := http.Header{}
	header.Set("Authorization", "Bearer "+genToken(conf.Config.Thrid.Key))
	return header
}

func genToken(key string) string {
	j := jwt.New([]byte(key))
	m := make(map[string]interface{})
	m["name"] = "urbs-console"
	token, err := j.Sign(m, time.Hour)
	if err != nil {
		panic(err)
	}
	return token
}
