package util

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/teambition/gear"
)

// RequestGet ...
func RequestGet(ctx context.Context, url string, header http.Header, out interface{}) (err error) {
	return Request(ctx, url, http.MethodGet, header, nil, out)
}

// RequestPost ...
func RequestPost(ctx context.Context, url string, header http.Header, body interface{}, out interface{}) (err error) {
	return Request(ctx, url, http.MethodPost, header, body, out)
}

// Request ...
func Request(ctx context.Context, url string, method string, header http.Header, body interface{}, out interface{}) error {
	var err error
	bs := []byte{}
	if body != nil {
		bs, err = json.Marshal(body)
		if err != nil {
			return err
		}
	}
	if header == nil {
		header = http.Header{}
	}
	header.Set("Content-Type", "application/json")
	requestId, _ := ctx.Value("X-Request-ID").(string)
	if requestId == "" {
		if gearCtx, ok := ctx.(*gear.Context); ok {
			requestId = gearCtx.GetHeader(gear.HeaderXRequestID)
		}
	}
	header.Set("X-Request-ID", requestId)
	header.Set("User-Agent", UserAgent)

	req, err := http.NewRequest(method, url, bytes.NewReader(bs))
	if err != nil {
		return err
	}
	req.Header = header
	resp, err := httpClient.Do(req)
	if err != nil {
		return gear.ErrBadRequest.WithMsg(err.Error())
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return gear.ErrBadRequest.WithCode(resp.StatusCode).WithMsg(err.Error())
	}
	if resp.StatusCode >= 300 {
		return gear.ErrBadRequest.WithCode(resp.StatusCode).WithMsg(err.Error())
	}
	if out != nil {
		err = json.Unmarshal(respBody, out)
	} else {
		io.Copy(ioutil.Discard, resp.Body)
	}
	return err
}
