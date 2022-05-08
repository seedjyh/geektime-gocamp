package user

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type bindingHandler struct {
}

type BindRequestBody struct {
	TelA string `json:"tel_a"`
	TelX string `json:"tel_x"`
	TelB string `json:"tel_b"`
}

type ResponseBody struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type BindResponseData struct {
	BindID string
}

func success(c echo.Context, data interface{}) error {
	return c.JSON(http.StatusOK, &ResponseBody{
		Code:    0,
		Message: "success",
		Data:    data,
	})
}

func invalidRequestBody(c echo.Context, data interface{}) error {
	return c.JSON(http.StatusBadRequest, &ResponseBody{
		Code:    1,
		Message: "invalid request body",
		Data:    data,
	})
}

func invalidQueryParameter(c echo.Context, data interface{}) error {
	return c.JSON(http.StatusBadRequest, &ResponseBody{
		Code:    2,
		Message: "invalid query parameter",
		Data:    data,
	})
}

func (h *bindingHandler) Bind(c echo.Context) error {
	body := new(BindRequestBody)
	if err := c.Bind(body); err != nil {
		return invalidRequestBody(c, nil)
	}
	bindID := "-new-bind-id-"
	return success(c, &BindResponseData{BindID: bindID})
}

func (h *bindingHandler) Unbind(c echo.Context) error {
	bindID := c.Param("bind_id")
	if len(bindID) < 10 {
		return invalidQueryParameter(c, map[string]interface{}{"bind_id": bindID})
	}
	return success(c, nil)
}
