package service

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/teambition/gear"
	authjwt "github.com/teambition/gear-auth/jwt"
	"github.com/teambition/urbs-console/src/conf"
	"github.com/teambition/urbs-console/src/logger"
	"github.com/teambition/urbs-console/src/util"
	"github.com/teambition/urbs-console/src/util/request"
)

var urbsSettingJwt *authjwt.JWT
var thirdJwt *authjwt.JWT

func init() {
	util.DigProvide(NewServices)
	urbsSettingJwt = authjwt.New([]byte(conf.Config.UrbsSetting.Key))
	thirdJwt = authjwt.New([]byte(conf.Config.Thrid.Key))
}

// Services ...
type Services struct {
	UrbsSetting UrbsSettingInterface
	UserAuth    UserAuthInterface
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
	var requestId string
	header := http.Header{}
	if gearCtx, ok := ctx.(*gear.Context); ok {
		requestId = gearCtx.GetHeader(gear.HeaderXRequestID)
		if requestId == "" {
			requestId = gearCtx.Res.Header().Get(gear.HeaderXRequestID)
		}
		canary := gearCtx.GetHeader("X-Canary")
		if canary != "" {
			header.Set("X-Canary", canary)
		}
	}
	if requestId == "" {
		requestId, _ = ctx.Value(gear.HeaderXRequestID).(string)
	}
	if requestId != "" {
		header.Set(gear.HeaderXRequestID, requestId)
	}
	header.Set(util.HeaderAuthorize, "Bearer "+genToken(urbsSettingJwt))
	return header
}

// HanderResponse ...
func HanderResponse(response *request.Response, err error) error {
	if err != nil {
		return gear.ErrBadRequest.WithMsg(err.Error())
	}
	if !response.OK() {
		resRequestID := response.Request.Header.Get(gear.HeaderXRequestID)
		logger.Default.Err("error", string(response.Content), "status", response.StatusCode, "xRequestId", resRequestID, "url", response.Request.URL.String())

		gearErr := new(gear.Error)
		json.Unmarshal(response.Content, gearErr)
		if gearErr.Err != "" {
			return gearErr.WithCode(response.StatusCode)
		}
		return gear.ErrBadRequest.WithCode(response.StatusCode).WithMsg(response.String())
	}
	return nil
}

func genThridHeader(ctx context.Context) http.Header {
	header := http.Header{}
	header.Set(util.HeaderAuthorize, "Bearer "+genToken(thirdJwt))
	addRequestId(ctx, header)
	return header
}

func addRequestId(ctx context.Context, header http.Header) {
	if ctx == nil {
		return
	}
	requestId, _ := ctx.Value(gear.HeaderXRequestID).(string)
	if requestId == "" {
		if gearCtx, ok := ctx.(*gear.Context); ok {
			requestId = gearCtx.GetHeader(gear.HeaderXRequestID)
			if requestId == "" {
				requestId = gearCtx.Res.Header().Get(gear.HeaderXRequestID)
			}
		}
	}
	if requestId != "" {
		header.Set(gear.HeaderXRequestID, requestId)
	}
}

func genToken(j *authjwt.JWT) string {
	m := make(map[string]interface{})
	m["name"] = "urbs-console"
	token, err := j.Sign(m, time.Hour)
	if err != nil {
		panic(err)
	}
	return token
}
