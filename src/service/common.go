package service

import (
	"context"
	"encoding/json"
	"net"
	"net/http"
	"time"

	"github.com/mushroomsir/request"
	"github.com/teambition/gear"
	authjwt "github.com/teambition/gear-auth/jwt"
	"github.com/teambition/urbs-console/src/conf"
	"github.com/teambition/urbs-console/src/logger"
	"github.com/teambition/urbs-console/src/util"
)

var urbsSettingJwt *authjwt.JWT
var thirdJwt *authjwt.JWT

func init() {
	util.DigProvide(NewServices)
	urbsSettingJwt = authjwt.New([]byte(conf.Config.UrbsSetting.Key))
	thirdJwt = authjwt.New([]byte(conf.Config.Thrid.Key))
}

// UrbsSettingHeader ...
func UrbsSettingHeader(ctx context.Context) http.Header {
	header := http.Header{}
	addRequestId(ctx, header)
	header.Set(util.HeaderAuthorize, util.HeaderAuthorizeBearer+genToken(urbsSettingJwt))
	return header
}

// ThridHeader ...
func ThridHeader(ctx context.Context) http.Header {
	header := http.Header{}
	addRequestId(ctx, header)
	header.Set(util.HeaderAuthorize, util.HeaderAuthorizeBearer+genToken(thirdJwt))
	return header
}

// HanderResponse ...
func HanderResponse(response *request.Response, err error) error {
	if err, ok := err.(net.Error); ok && err.Timeout() {
		return gear.ErrGatewayTimeout.WithMsg(err.Error())
	}
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

func addRequestId(ctx context.Context, header http.Header) {
	if ctx == nil {
		return
	}
	requestId, _ := ctx.Value(gear.HeaderXRequestID).(string)
	if requestId == "" {
		if gearCtx, ok := ctx.(*gear.Context); ok && gearCtx != nil {
			requestId = gearCtx.GetHeader(gear.HeaderXRequestID)
			if requestId == "" {
				requestId = gearCtx.Res.Header().Get(gear.HeaderXRequestID)
			}
			canary := gearCtx.GetHeader("X-Canary")
			if canary != "" {
				header.Set("X-Canary", canary)
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
