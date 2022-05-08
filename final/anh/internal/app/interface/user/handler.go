package user

import (
	"anh/internal/pkg/mylog"
	"context"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
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

func withSessionID(c echo.Context) (context.Context, *mylog.Logger) {
	sessionID := getSessionID(c)
	ctx := context.WithValue(context.Background(), "session_id", sessionID)
	logger := mylog.CloneLogger().WithTag("session_id", sessionID)
	return ctx, logger
}

func (h *bindingHandler) Bind(c echo.Context) error {
	ctx, logger := withSessionID(c)
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	body := new(BindRequestBody)
	if err := c.Bind(body); err != nil {
		_ = logger.CloneLogger().WithFields(mylog.Error("error", err)).Error("parse request body failed")
		return invalidRequestBody(c, nil)
	}
	parameter := &BindParameter{
		TelA: Number(body.TelA),
		TelX: Number(body.TelX),
		TelB: Number(body.TelB),
	}
	if err := parameter.AssertValid(); err != nil {
		_ = logger.CloneLogger().
			WithFields(mylog.String("parameter", parameter.String())).
			WithFields(mylog.Error("error", err)).
			Error("parse request body failed")
		return invalidQueryParameter(c, map[string]string{"error": err.Error()})
	}
	if bindId, err := h.xbrClient.Bind(ctx, parameter); err != nil {
		_ = logger.WithFields(mylog.Error("error", err)).Error("xbr bind failed")
		return internalError(c, map[string]string{"error": err.Error()})
	} else {
		return success(c, &BindResponseData{BindID: bindId.String()})
	}
}

func (h *bindingHandler) Unbind(c echo.Context) error {
	ctx, logger := withSessionID(c)
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	bindID := c.Param("bind_id")
	if err := h.xbrClient.Unbind(ctx, BindId(bindID)); err != nil {
		_ = logger.WithFields(mylog.Error("error", err)).Error("xbr unbind failed")
		return internalError(c, map[string]string{"error": err.Error()})
	} else {
		return success(c, nil)
	}
}
