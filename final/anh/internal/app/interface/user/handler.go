package user

import (
	"context"
	"github.com/labstack/echo/v4"
	"net/http"
)

type bindingHandler struct {
	xbrClient *xbrClient
}

func newBindingHandler(client *xbrClient) *bindingHandler {
	return &bindingHandler{xbrClient: client}
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

func internalError(c echo.Context, data interface{}) error {
	return c.JSON(http.StatusInternalServerError, &ResponseBody{
		Code:    3,
		Message: "internal error",
		Data:    data,
	})
}

func (h *bindingHandler) Bind(c echo.Context) error {
	body := new(BindRequestBody)
	if err := c.Bind(body); err != nil {
		return invalidRequestBody(c, nil)
	}
	if bindId, err := h.xbrClient.Bind(context.Background(), &BindParameter{
		TelA: Number(body.TelA),
		TelX: Number(body.TelX),
		TelB: Number(body.TelB),
	}); err != nil {
		return internalError(c, map[string]string{"error": err.Error()})
	} else {
		return success(c, &BindResponseData{BindID: bindId.String()})
	}
}

func (h *bindingHandler) Unbind(c echo.Context) error {
	bindID := c.Param("bind_id")
	if err := h.xbrClient.Unbind(context.Background(), BindId(bindID)); err != nil {
		return internalError(c, map[string]string{"error": err.Error()})
	} else {
		return success(c, nil)
	}
}
