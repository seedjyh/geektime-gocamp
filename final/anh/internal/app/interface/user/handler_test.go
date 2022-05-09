package user

import (
	"bytes"
	"encoding/json"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBindingHandler_Bind(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	xbrClient := NewMockXBRClient(mockCtrl)
	handler := newBindingHandler(xbrClient)

	const (
		telA   = "13700001111"
		telX   = "18622223333"
		telB   = "19944445555"
		bindID = "this-is-bind-id"
	)

	var reqBuf *bytes.Buffer
	if buf, err := json.Marshal(&BindRequestBody{
		TelA: telA,
		TelX: telX,
		TelB: telB,
	}); assert.NoError(t, err) {
		reqBuf = bytes.NewBuffer(buf)
	}
	req := httptest.NewRequest(http.MethodPost, "/", reqBuf)
	req.Header.Add(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	rec := httptest.NewRecorder()
	echoCtx := echo.New().NewContext(req, rec)
	xbrClient.EXPECT().Bind(gomock.Any(), &BindParameter{
		TelA: telA,
		TelX: telX,
		TelB: telB,
	}).Return(BindId(bindID), nil).Times(1)
	assert.NoError(t, handler.Bind(echoCtx))
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestBindingHandler_Unbind(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	xbrClient := NewMockXBRClient(mockCtrl)
	handler := newBindingHandler(xbrClient)

	const (
		bindID = "this-is-bind-id"
	)

	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	rec := httptest.NewRecorder()
	echoCtx := echo.New().NewContext(req, rec)
	echoCtx.SetParamNames("bind_id")
	echoCtx.SetParamValues(bindID)
	xbrClient.EXPECT().Unbind(gomock.Any(), BindId(bindID)).Return(nil).Times(1)
	assert.NoError(t, handler.Unbind(echoCtx))
	assert.Equal(t, http.StatusOK, rec.Code)
}
