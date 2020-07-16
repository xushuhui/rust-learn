package test

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"skrshop-cms-api/core"
	"skrshop-cms-api/request"
	"skrshop-cms-api/route"
	"testing"
)

var resp core.JsonResponse

func TestPostBrand(t *testing.T) {
	router := route.InitRouter()
	w := httptest.NewRecorder()
	json.Marshal(request.Brand{Name: "brand", Desc: "desc", LogoUrl: "logourl", CreateBy: 1})
	req, _ := http.NewRequest(http.MethodPost, "/brand/", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	json.Unmarshal(w.Body.Bytes(), &resp)
	assert.Equal(t, 0, resp.Code)
}
func TestGetBrand(t *testing.T) {
	router := route.InitRouter()
	w := httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodPost, "/brand/5", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	json.Unmarshal(w.Body.Bytes(), &resp)
	assert.Equal(t, 0, resp.Code)
}
