package user

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type bindingHandler struct {
}

func (h *bindingHandler) Bind(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

func (h *bindingHandler) Unbind(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}
