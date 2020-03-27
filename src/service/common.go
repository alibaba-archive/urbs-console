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

	UserAuth    UserAuth
	GroupMember *GroupMember
}

// NewServices ...
func NewServices() *Services {
	s := &Services{
		GroupMember: &GroupMember{},
		UrbsSetting: &UrbsSetting{},
	}
	if conf.Config.UserAuth.UserAuthThrid.URL == "" {
		s.UserAuth = &UserAuthLocal{}
	} else {
		s.UserAuth = &UserAuthThrid{}
	}
	return s
}

// UrbsSettingHeader ...
func UrbsSettingHeader(ctx context.Context) http.Header {
	header := http.Header{}
	requestId, _ := ctx.Value("X-Request-ID").(string)
	if requestId == "" {
		if gearCtx, ok := ctx.(*gear.Context); ok {
			requestId = gearCtx.GetHeader(gear.HeaderXRequestID)
			if requestId == "" {
				requestId = gearCtx.Res.Header().Get(gear.HeaderXRequestID)
			}
		}
	}
	if requestId != "" {
		header.Set("X-Request-ID", requestId)
	}
	header.Set("Authorization", "Bearer "+genToken())
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
			gearErr.Msg = "urbs-setting: " + gearErr.Msg
			return gearErr.WithCode(response.StatusCode)
		}
		return gear.ErrBadRequest.WithCode(response.StatusCode).WithMsg("urbs-setting: " + response.String())
	}
	return nil
}

func genToken() string {
	j := jwt.New([]byte(conf.Config.UrbsSetting.AuthKeys[0]))
	m := make(map[string]interface{})
	m["name"] = "urbs-console"
	token, err := j.Sign(m, time.Hour)
	if err != nil {
		panic(err)
	}
	return token
}
