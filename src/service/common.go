package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"net/http"

	"github.com/mushroomsir/request"
	otgo "github.com/open-trust/ot-go-lib"
	"github.com/teambition/gear"
	"github.com/teambition/urbs-console/src/conf"
	"github.com/teambition/urbs-console/src/logger"
	"github.com/teambition/urbs-console/src/util"
)

var thirdOTID otgo.OTID
var urbsSettingOTID otgo.OTID
var otHolder *otgo.Holder

func init() {
	util.DigProvide(NewServices)

	otConf := conf.Config.OpenTrust
	otHolder = genHolder(otConf)

	addOTVIDTokens(otHolder, conf.Config.UrbsSetting.OTVIDs)
	addOTVIDTokens(otHolder, conf.Config.Thrid.OTVIDs)

	urbsSettingOTID = parseOTID(conf.Config.UrbsSetting.OTID)
	thirdOTID = parseOTID(conf.Config.Thrid.OTID)
}

func parseOTID(otidStr string) otgo.OTID {
	otid, err := otgo.ParseOTID(otidStr)
	if err != nil {
		panic(fmt.Errorf("Parse Urbs Setting config otid failed: %s", err))
	}
	return otid
}

func genHolder(otConf conf.OpenTrust) *otgo.Holder {
	otid := parseOTID(otConf.OTID)
	otHolder, err := otgo.NewHolder(context.Background(), otid, otConf.PrivateKeys...)
	if err != nil {
		panic(fmt.Errorf("Parse Open Trust config failed: %s", err))
	}
	return otHolder
}

func addOTVIDTokens(otHolder *otgo.Holder, otvids []string) {
	err := otHolder.AddOTVIDTokens(otvids...)
	if err != nil {
		panic(fmt.Errorf("Parse Open Trust config otvids failed: %s", err))
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
	token, err := otHolder.GetOTVIDToken(thirdOTID)
	if err != nil {
		panic(err)
	}
	otgo.AddTokenToHeader(header, token)
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
