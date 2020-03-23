package util

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"runtime"
	"time"
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
}
