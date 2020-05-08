package util

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/mushroomsir/request"
	"github.com/teambition/gear"
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
	// HeaderAuthorize ...
	HeaderAuthorize = "Authorization"
	// HeaderAuthorizeBearer ...
	HeaderAuthorizeBearer = "Bearer "
	// HeaderAuthorizeOAuth2 ...
	HeaderAuthorizeOAuth2 = "OAuth2 "
)

// TokenExtractor ...
func TokenExtractor(ctx *gear.Context) (token string) {
	if val := ctx.Get(HeaderAuthorize); strings.HasPrefix(val, HeaderAuthorizeBearer) {
		token = val[7:]
	} else if val := ctx.Get(HeaderAuthorize); strings.HasPrefix(val, HeaderAuthorizeOAuth2) {
		token = val[7:]
	}
	return
}

// StringInSlice ...
func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
