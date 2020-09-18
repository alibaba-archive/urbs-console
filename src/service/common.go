package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/mushroomsir/request"
	otgo "github.com/open-trust/ot-go-lib"
	"github.com/teambition/gear"
	authjwt "github.com/teambition/gear-auth/jwt"
	"github.com/teambition/urbs-console/src/conf"
	"github.com/teambition/urbs-console/src/logger"
	"github.com/teambition/urbs-console/src/util"
)

var thirdJwt *authjwt.JWT
var otHolder *otgo.Holder
var urbsSettingOTID otgo.OTID

func init() {
	util.DigProvide(NewServices)

	thirdJwt = authjwt.New([]byte(conf.Config.Thrid.Key))

	otConf := conf.Config.OpenTrust
	otid, err := otgo.ParseOTID(otConf.OTID)
	if err != nil {
		panic(fmt.Errorf("Parse Open Trust config otid failed: %s", err))
	}

	urbsSettingOTID, err = otgo.ParseOTID(conf.Config.UrbsSetting.OTID)
	if err != nil {
		panic(fmt.Errorf("Parse Urbs Setting config otid failed: %s", err))
	}

	otHolder, err = otgo.NewHolder(conf.Config.GlobalCtx, otid, otConf.PrivateKeys...)
	if err != nil {
		panic(fmt.Errorf("Parse Open Trust config failed: %s", err))
	}
	if len(otConf.OTVIDs) > 0 {
		if err = otHolder.AddOTVIDTokens(otConf.OTVIDs...); err != nil {
			panic(fmt.Errorf("Parse Open Trust config otvids failed: %s", err))
		}
	}
}

// UrbsSettingHeader ...
func UrbsSettingHeader(ctx context.Context) http.Header {
	header := http.Header{}
	addRequestId(ctx, header)
	token, err := otHolder.GetOTVIDToken(urbsSettingOTID)
	if err != nil {
		panic(err)
	}
	otgo.AddTokenToHeader(header, token)
	return header
}

// ThridHeader ...
func ThridHeader(ctx context.Context) http.Header {
	header := http.Header{}
	addRequestId(ctx, header)
	otgo.AddTokenToHeader(header, genToken(thirdJwt))
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
