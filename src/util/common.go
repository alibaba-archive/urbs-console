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
		Timeout:   time.Second * 10,
		Transport: tr,
	}

	request.SetHttpClient(httpClient)
	request.SetUserAgent(UserAgent)
}

const (
	// HeaderXAuth ...
	HeaderXAuth = "X-Auth"
	// HeaderAuthorize ...
	HeaderAuthorize = "Authorization"
	// HeaderAuthorizeBearer ...
	HeaderAuthorizeBearer = "Bearer "
	// HeaderAuthorizeOAuth2 ...
	HeaderAuthorizeOAuth2 = "OAuth2 "
)

// ExtractBearerToken ...
func ExtractBearerToken(ctx *gear.Context) (token string) {
	if val := ctx.Get(HeaderAuthorize); strings.HasPrefix(val, HeaderAuthorizeBearer) {
		token = val[7:]
	}
	return
}

// ExtractToken ...
func ExtractToken(ctx *gear.Context) (token string) {
	if val := ctx.Get(HeaderAuthorize); strings.HasPrefix(val, HeaderAuthorizeBearer) {
		token = val[7:]
	} else if val := ctx.Get(HeaderAuthorize); strings.HasPrefix(val, HeaderAuthorizeOAuth2) {
		token = val[7:]
	}
	return
}

// XAuthExtractor ...
func XAuthExtractor(ctx *gear.Context) (token string) {
	if val := ctx.Get(HeaderXAuth); strings.HasPrefix(val, HeaderAuthorizeBearer) {
		token = val[7:]
	}
	return
}

// StringSliceHas ...
func StringSliceHas(list []string, a string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// PathExists returns whether the given file or directory exists
func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return true
}
