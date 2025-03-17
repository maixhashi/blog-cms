package layout_test

import (
	"net/http/httptest"
	"strings"

	"github.com/labstack/echo/v4"
)

// Helper function to setup the Echo context
func setupContext(method, url string, body string) (echo.Context, *httptest.ResponseRecorder) {
	e := echo.New()
	req := httptest.NewRequest(method, url, strings.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	return c, rec
}
