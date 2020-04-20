package api

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/teambition/urbs-console/src/util/request"
)

func TestAPIsAuth(t *testing.T) {
	require := require.New(t)

	type reqArg struct {
		Method string
		URL    string
		Body   string
	}
	reqArgs := []reqArg{
		// ***** product ******
		{Method: http.MethodGet, URL: "/api/v1/products"},
		{Method: http.MethodPost, URL: "/api/v1/products"},
		{Method: http.MethodPut, URL: "/api/v1/products/product", Body: "{}"},
		{Method: http.MethodPut, URL: "/api/v1/products/product:offline", Body: "{}"},
		{Method: http.MethodDelete, URL: "/api/v1/products/product"},
		// ***** label ******
		{Method: http.MethodGet, URL: "/api/v1/products/product/labels"},
		{Method: http.MethodGet, URL: "/api/v1/products/product/labels/label/logs"},
		{Method: http.MethodGet, URL: "/api/v1/products/product/labels/label/groups"},
		{Method: http.MethodGet, URL: "/api/v1/products/product/labels/label/users"},
		{Method: http.MethodPost, URL: "/api/v1/products/product/labels", Body: `{"name":"name"}`},
		{Method: http.MethodPut, URL: "/api/v1/products/product/labels/label", Body: `{"desc":"name"}`},
		{Method: http.MethodDelete, URL: "/api/v1/products/product/labels/label"},
		{Method: http.MethodPut, URL: "/api/v1/products/product/labels/label:offline"},
		{Method: http.MethodPost, URL: "/api/v1/products/product/labels/label:assign", Body: `{"users":["123"]}`},
		{Method: http.MethodPost, URL: "/api/v1/products/product/labels/label:recall"},
		// ***** module ******
		{Method: http.MethodGet, URL: "/api/v1/products/product/modules"},
		{Method: http.MethodPost, URL: "/api/v1/products/product/modules", Body: `{"name":"xcccc"}`},
		{Method: http.MethodPut, URL: "/api/v1/products/product/modules/module", Body: `{"name":"xcccc"}`},
		{Method: http.MethodPut, URL: "/api/v1/products/product/modules/module:offline"},
		// ***** setting ******
		{Method: http.MethodGet, URL: "/api/v1/products/product/modules/module/settings"},
		{Method: http.MethodGet, URL: "/api/v1/products/product/modules/module/settings/setting/logs"},
		{Method: http.MethodPost, URL: "/api/v1/products/product/modules/module/settings", Body: `{"name":"xxxx"}`},
		{Method: http.MethodPut, URL: "/api/v1/products/product/modules/module/settings/setting", Body: `{"desc":"xxxx"}`},
		{Method: http.MethodPut, URL: "/api/v1/products/product/modules/module/settings/setting:offline"},
		{Method: http.MethodPost, URL: "/api/v1/products/product/modules/module/settings/setting:assign", Body: `{"users":["xxxx"]}`},
		{Method: http.MethodPost, URL: "/api/v1/products/product/modules/module/settings/setting:recall"},
		// ***** user ******
		{Method: http.MethodGet, URL: "/api/v1/users"},
		{Method: http.MethodGet, URL: "/api/v1/users/uid/labels"},
		{Method: http.MethodGet, URL: "/api/v1/users/uid/settings"},
		{Method: http.MethodPut, URL: "/api/v1/users/uid/labels:cache"},
		{Method: http.MethodDelete, URL: "/api/v1/users/uid/labels/hid"},
		{Method: http.MethodPut, URL: "/api/v1/users/uid/settings/hid:rollback"},
		{Method: http.MethodDelete, URL: "/api/v1/users/uid/settings/hid"},
		{Method: http.MethodPost, URL: "/api/v1/users:batch"},
		// ***** group ******
		{Method: http.MethodGet, URL: "/api/v1/groups"},
		{Method: http.MethodGet, URL: "/api/v1/groups/uid/labels"},
		{Method: http.MethodGet, URL: "/api/v1/groups/uid/settings"},
		{Method: http.MethodGet, URL: "/api/v1/groups/uid/members"},
		{Method: http.MethodPost, URL: "/api/v1/groups:batch"},
		{Method: http.MethodPut, URL: "/api/v1/groups/uid"},
		{Method: http.MethodDelete, URL: "/api/v1/groups/uid"},
		{Method: http.MethodDelete, URL: "/api/v1/groups/uid/labels/hid"},
		{Method: http.MethodPut, URL: "/api/v1/groups/uid/settings/hid:rollback"},
		{Method: http.MethodDelete, URL: "/api/v1/groups/uid/settings/hid"},
		{Method: http.MethodDelete, URL: "/api/v1/groups/uid/members"},
		// ***** UrbsAc ******
		{Method: http.MethodPost, URL: "/api/v1/ac/users"},
		{Method: http.MethodPost, URL: "/api/v1/ac/users/:uid/permissions"},
	}

	for _, req := range reqArgs {
		// 验证 401
		res, err := request.Method(req.Method).Url(testHost + req.URL).RawBody(req.Body).Do()
		require.Nil(err)
		require.Equal(http.StatusUnauthorized, res.StatusCode)

		// 验证 403
		res, err = request.Method(req.Method).Url(testHost + req.URL).RawBody(req.Body).Header(userTokenHeader()).Do()
		require.Nil(err)
		require.Equal(http.StatusForbidden, res.StatusCode)
	}
}
