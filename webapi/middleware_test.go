package webapi_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/chaipawn/assessment/webapi"
	"github.com/labstack/echo/v4"
)

func TestAuthorizeMiddlewareFail(t *testing.T) {
	handler := func(c echo.Context) error {
		return nil
	}
	req := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(""))
	rec := httptest.NewRecorder()
	app := echo.New()
	c := app.NewContext(req, rec)

	webapi.Authorize(handler)(c)

	if rec.Code != http.StatusUnauthorized {
		t.Errorf("Authorized middleware status code expect %d, but got %d", http.StatusUnauthorized, rec.Code)
	}
}

func TestAuthorizeMiddlewareSuccess(t *testing.T) {
	handler := func(c echo.Context) error {
		return c.String(http.StatusOK, "")
	}
	req := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(""))
	req.Header.Add("Authorization", "November 10, 2009")
	rec := httptest.NewRecorder()
	app := echo.New()
	c := app.NewContext(req, rec)

	webapi.Authorize(handler)(c)

	if rec.Code != http.StatusOK {
		t.Errorf("Authorized middleware status code expect %d, but got %d", http.StatusOK, rec.Code)
	}
}
