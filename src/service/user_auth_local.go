package service

import (
	"time"

	"github.com/teambition/gear"
	auth "github.com/teambition/gear-auth"
	authjwt "github.com/teambition/gear-auth/jwt"
	"github.com/teambition/gear/logging"
	"github.com/teambition/urbs-console/src/conf"
	"github.com/teambition/urbs-console/src/dto/thrid"
	"github.com/teambition/urbs-console/src/logger"
)

func init() {
	key := conf.Config.Thrid.Key
	if key != "" {
		Auther = auth.New(authjwt.StrToKeys(key)...)
		Auther.JWT().SetExpiresIn(time.Minute * 10)
	} else {
		logger.Default.Warningf("`user_auth.keys` is empty, Auth middleware will not be executed.")
	}
}

// UserAuthLocal ...
type UserAuthLocal struct {
}

// Auther ...
var Auther *auth.Auth

// Verify ...
func (a *UserAuthLocal) Verify(ctx *gear.Context, body *thrid.UserVerifyReq) error {
	if Auther != nil {
		claims, err := Auther.FromCtx(ctx)
		if err != nil {
			return err
		}
		if sub, ok := claims.Subject(); ok {
			logging.SetTo(ctx, "jwt_sub", sub)
		}
		if jti, ok := claims.JWTID(); ok {
			logging.SetTo(ctx, "jwt_id", jti)
		}
	}
	return nil
}
