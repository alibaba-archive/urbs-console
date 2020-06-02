package api

import (
	"net/http"
	"testing"

	"github.com/mushroomsir/request"
	"github.com/stretchr/testify/require"
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
		{Method: http.MethodGet, URL: "/api/v1/products/product/statistics"},
		{Method: http.MethodPost, URL: "/api/v1/products"},
		{Method: http.MethodPut, URL: "/api/v1/products/product", Body: `{"uids":["123"]}`},
		{Method: http.MethodPut, URL: "/api/v1/products/product:offline", Body: "{}"},
		{Method: http.MethodDelete, URL: "/api/v1/products/product"},
		// ***** label ******
		{Method: http.MethodGet, URL: "/api/v1/products/product/labels"},
		{Method: http.MethodGet, URL: "/api/v1/products/product/labels/label/logs"},
		{Method: http.MethodGet, URL: "/api/v1/products/product/labels/label/groups"},
		{Method: http.MethodDelete, URL: "/api/v1/products/product/labels/label/groups/123"},
		{Method: http.MethodGet, URL: "/api/v1/products/product/labels/label/users"},
		{Method: http.MethodDelete, URL: "/api/v1/products/product/labels/label/users/123"},
		{Method: http.MethodPost, URL: "/api/v1/products/product/labels", Body: `{"name":"name","uids":["123"]}`},
		{Method: http.MethodPut, URL: "/api/v1/products/product/labels/label", Body: `{"desc":"name"}`},
		// {Method: http.MethodDelete, URL: "/api/v1/products/product/labels/label"},
		{Method: http.MethodPut, URL: "/api/v1/products/product/labels/label:offline"},
		{Method: http.MethodPost, URL: "/api/v1/products/product/labels/label:assign", Body: `{"users":["123"]}`},
		{Method: http.MethodPost, URL: "/api/v1/products/product/labels/label:recall", Body: `{"hid":"123"}`},
		{Method: http.MethodPost, URL: "/api/v1/products/product/labels/label/rules", Body: `{"kind":"userPercent"}`},
		{Method: http.MethodGet, URL: "/api/v1/products/product/labels/label/rules", Body: `{"kind":"userPercent"}`},
		{Method: http.MethodPut, URL: "/api/v1/products/product/labels/label/rules/AwAAAAAAAAB25V_QnbhCuRwF", Body: `{"kind":"userPercent"}`},
		{Method: http.MethodDelete, URL: "/api/v1/products/product/labels/label/rules/AwAAAAAAAAB25V_QnbhCuRwF"},
		// ***** module ******
		{Method: http.MethodGet, URL: "/api/v1/products/product/modules"},
		{Method: http.MethodPost, URL: "/api/v1/products/product/modules", Body: `{"name":"xcccc","uids":["123"]}`},
		{Method: http.MethodPut, URL: "/api/v1/products/product/modules/module", Body: `{"name":"xcccc","uids":["123"]}`},
		{Method: http.MethodPut, URL: "/api/v1/products/product/modules/module:offline"},
		// ***** setting ******
		{Method: http.MethodGet, URL: "/api/v1/products/product/modules/module/settings"},
		{Method: http.MethodGet, URL: "/api/v1/products/product/modules/module/settings/setting/groups"},
		{Method: http.MethodDelete, URL: "/api/v1/products/product/modules/module/settings/setting/groups/123"},
		{Method: http.MethodPut, URL: "/api/v1/products/product/modules/module/settings/setting/groups/123:rollback"},

		{Method: http.MethodGet, URL: "/api/v1/products/product/modules/module/settings/setting/users"},
		{Method: http.MethodDelete, URL: "/api/v1/products/product/modules/module/settings/setting/users/123"},
		{Method: http.MethodPut, URL: "/api/v1/products/product/modules/module/settings/setting/users/123:rollback"},

		{Method: http.MethodGet, URL: "/api/v1/products/product/modules/module/settings/setting/logs"},
		{Method: http.MethodPost, URL: "/api/v1/products/product/modules/module/settings", Body: `{"name":"xxxx","uids":["123"]}`},
		{Method: http.MethodPut, URL: "/api/v1/products/product/modules/module/settings/setting", Body: `{"desc":"xxxx","uids":["123"]}`},
		{Method: http.MethodPut, URL: "/api/v1/products/product/modules/module/settings/setting:offline"},
		{Method: http.MethodPost, URL: "/api/v1/products/product/modules/module/settings/setting:assign", Body: `{"users":["xxxx"]}`},
		{Method: http.MethodPost, URL: "/api/v1/products/product/modules/module/settings/setting:recall", Body: `{"hid":"123"}`},
		{Method: http.MethodPost, URL: "/api/v1/products/product/modules/module/settings/setting/rules", Body: `{"kind":"userPercent"}`},
		{Method: http.MethodGet, URL: "/api/v1/products/product/modules/module/settings/setting/rules", Body: `{"kind":"userPercent"}`},
		{Method: http.MethodPut, URL: "/api/v1/products/product/modules/module/settings/setting/rules/AwAAAAAAAAB25V_QnbhCuRwF", Body: `{"kind":"userPercent"}`},
		{Method: http.MethodDelete, URL: "/api/v1/products/product/modules/module/settings/setting/rules/AwAAAAAAAAB25V_QnbhCuRwF"},
		// ***** user ******
		{Method: http.MethodGet, URL: "/api/v1/users"},
		{Method: http.MethodGet, URL: "/api/v1/users/uid/labels"},
		{Method: http.MethodGet, URL: "/api/v1/users/uid/settings"},
		{Method: http.MethodPut, URL: "/api/v1/users/uid/labels:cache"},
		{Method: http.MethodPost, URL: "/api/v1/users:batch"},
		// ***** group ******
		{Method: http.MethodGet, URL: "/api/v1/groups"},
		{Method: http.MethodGet, URL: "/api/v1/groups/uid/labels"},
		{Method: http.MethodGet, URL: "/api/v1/groups/uid/settings"},
		{Method: http.MethodGet, URL: "/api/v1/groups/uid/members"},
		{Method: http.MethodPost, URL: "/api/v1/groups:batch"},
		{Method: http.MethodPut, URL: "/api/v1/groups/uid"},
		{Method: http.MethodDelete, URL: "/api/v1/groups/uid"},
		{Method: http.MethodDelete, URL: "/api/v1/groups/uid/members"},
		// ***** UrbsAc ******
		{Method: http.MethodPost, URL: "/api/v1/ac/users"},
		{Method: http.MethodDelete, URL: "/api/v1/ac/users/users"},
		{Method: http.MethodGet, URL: "/api/v1/ac/users"},
		{Method: http.MethodGet, URL: "/api/v1/ac/users:search"},
		{Method: http.MethodPost, URL: "/api/v1/ac/users/:uid/permissions"},
		{Method: http.MethodDelete, URL: "/api/v1/ac/users/:uid/permissions"},
	}

	for _, req := range reqArgs {
		// 验证 401
		res, err := request.Method(req.Method).Url(testHost + req.URL).RawBody(req.Body).Do()
		require.Nil(err, err)
		require.Equal(http.StatusUnauthorized, res.StatusCode, req.Method+":"+req.URL)

		// 验证 403
		res, err = request.Method(req.Method).Url(testHost + req.URL).RawBody(req.Body).Header(userTokenHeader()).Do()
		require.Nil(err)
		require.Equal(http.StatusForbidden, res.StatusCode, req.Method+":"+req.URL)
	}
}
