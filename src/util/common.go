package util

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/teambition/gear"
	"github.com/teambition/urbs-console/src/util/request"
)

var (
	// UserAgent ...
	UserAgent string

	httpClient *http.Client
)

func init() {
	hostname, _ := os.Hostname()
	UserAgent = fmt.Sprintf("golang/%v hostname/%s version/%s", runtime.Version(), hostname, Version)

	tr := &http.Transport{
		MaxIdleConnsPerHost: 500,
		IdleConnTimeout:     2 * time.Minute,
		TLSHandshakeTimeout: 500 * time.Millisecond,
		TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
	}
	httpClient = &http.Client{
		Timeout:   time.Second * 3,
		Transport: tr,
	}

	request.SetHttpClient(httpClient)
	request.SetUserAgent(UserAgent)
}

const (
	headerAuthorize = "Authorization"
	schemaBearer    = "Bearer "
	schemaOAuth2    = "OAuth2 "
)

// TokenExtractor ...
func TokenExtractor(ctx *gear.Context) (token string) {
	if val := ctx.Get(headerAuthorize); strings.HasPrefix(val, schemaBearer) {
		token = val[7:]
	} else if val := ctx.Get(headerAuthorize); strings.HasPrefix(val, schemaOAuth2) {
		token = val[7:]
	}
	return
}
