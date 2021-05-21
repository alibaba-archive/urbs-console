package middleware

import (
	"context"
	"time"

	otgo "github.com/open-trust/ot-go-lib"
	"github.com/teambition/gear"
	auth "github.com/teambition/gear-auth"
	authjwt "github.com/teambition/gear-auth/jwt"
	"github.com/teambition/gear/logging"
	"github.com/teambition/urbs-console/src/bll"
	"github.com/teambition/urbs-console/src/conf"
	"github.com/teambition/urbs-console/src/dto/thrid"
	"github.com/teambition/urbs-console/src/logger"
	"github.com/teambition/urbs-console/src/service"
	"github.com/teambition/urbs-console/src/util"
)

func init() {
	keys := conf.Config.AuthKeys
	if len(keys) > 0 {
		Auther = auth.New(authjwt.StrToKeys(keys...)...)
		Auther.JWT().SetExpiresIn(time.Minute * 10)
	}
	otConf := conf.Config.OpenTrust
	if otConf.OTID != "" {
		otVerifier = newVerifier(otConf.OTID, otConf.DomainPublicKeys)
		if otConf.LegacyOTID != "" {
			otLegacyVerifier = newVerifier(otConf.LegacyOTID, otConf.DomainPublicKeys)
		}
	}
	if otVerifier == nil {
		logger.Default.Warningf("`open_trust` is empty, Auth middleware will not be executed.")
	}
	if Auther == nil {
		logger.Default.Warningf("`auth_keys` is empty, Auth middleware will not be executed.")
	}
}

func newVerifier(otidStr string, publicKeys []string) *otgo.Verifier {
	otid, err := otgo.ParseOTID(otidStr)
	if err != nil {
		logger.Default.Panicf("Parse Open Trust config failed: %s", err)
	}

	otVerifier, err := otgo.NewVerifier(conf.Config.GlobalCtx, otid, false, publicKeys...)
	if err != nil {
		logger.Default.Panicf("Parse Open Trust config failed: %s", err)
	}
	return otVerifier
}

// Auther 是基于 JWT 的身份验证，当 config.auth_keys 配置了才会启用
var Auther *auth.Auth
var otVerifier *otgo.Verifier
var otLegacyVerifier *otgo.Verifier

// Verify ...
func Verify(services *service.Services) func(ctx *gear.Context) error {
	return func(ctx *gear.Context) error {
		var isOpenTrust bool
		var uid string
		xToken := util.XAuthExtractor(ctx)
		if xToken != "" && Auther != nil { // 旧版本实现，兼容
			claims, err := Auther.JWT().Verify(xToken)
			if err != nil {
				return gear.ErrUnauthorized.WithMsg(err.Error())
			}
			uid = claims.Get("uid").(string)
		} else {
			token := util.ExtractBearerToken(ctx)
			if otVerifier != nil && len(token) > 0 {
				vid, err := otVerifier.ParseOTVID(token)
				if err != nil {
					vid, err = otLegacyVerifier.ParseOTVID(token)
				}
				if err == nil {
					isOpenTrust = true
					if len(conf.Config.SuperAdmins) > 0 { // open-trust 不需要从用户身份中提取用 UID，默认拥有超级管理员权限
						uid = conf.Config.SuperAdmins[0]
					}
					logging.SetTo(ctx, "otSub", vid.ID.String())
				} else {
					logger.Warning(ctx, err.Error())
				}
			}
			token = util.ExtractToken(ctx)
			if !isOpenTrust { // open-trust 验证失败，走用户验证
				body := &thrid.UserVerifyReq{}
				body.Cookie, _ = ctx.Cookies.Get(conf.Config.Thrid.UserAuth.CookieKey)
				body.Singed, _ = ctx.Cookies.Get(conf.Config.Thrid.UserAuth.CookieKey + ".sig")
				body.Token = token
				if body.Cookie == "" && body.Token == "" {
					return gear.ErrUnauthorized.WithMsg("invalid authorization")
				}
				var err error
				uid, err = services.UserAuth.Verify(ctx, body)
				if err != nil {
					return gear.ErrUnauthorized.WithMsg(err.Error())
				}
			}
		}
		if uid == "" {
			return gear.ErrUnauthorized.WithMsg("invalid uid")
		}
		_ctx := context.WithValue(ctx.Context(), util.UidKey{}, uid)
		ctx.WithContext(_ctx)
		return nil
	}
}

// VerifyService ...
func VerifyService(services *service.Services) func(ctx *gear.Context) error {
	return func(ctx *gear.Context) error {
		token := util.ExtractBearerToken(ctx)
		if token == "" {
			token = util.XAuthExtractor(ctx)
		}
		_, err := Auther.JWT().Verify(token)
		if err != nil {
			return gear.ErrUnauthorized.WithMsg(err.Error())
		}
		return nil
	}
}

// CheckSuperAdmin ...
func CheckSuperAdmin(blls *bll.Blls) func(ctx *gear.Context) error {
	return func(ctx *gear.Context) error {
		err := blls.UrbsAcAcl.CheckSuperAdmin(ctx)
		if err != nil {
			return err
		}
		return nil
	}
}

// CheckViewer ...
func CheckViewer(blls *bll.Blls) func(ctx *gear.Context) error {
	return func(ctx *gear.Context) error {
		err := blls.UrbsAcAcl.CheckViewer(ctx)
		if err != nil {
			return err
		}
		return nil
	}
}
