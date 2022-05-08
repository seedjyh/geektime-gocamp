package user

import (
	"anh/internal/pkg/mylog"
	"bytes"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"io/ioutil"
	"strings"
	"time"
)
import "anh/internal/pkg/uuid"

var sessionIDGenerator = uuid.NewUUID32Generator()

const (
	echoKeySessionID = "echo-key-session-id"
)

// appendSessionID 返回一个MiddlewareFunc，能在echo.Context里Set进一个SessionID。
func appendSessionID() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set(echoKeySessionID, sessionIDGenerator.Next())
			return next(c)
		}
	}
}

func getSessionID(c echo.Context) string {
	if v, ok := c.Get(echoKeySessionID).(string); ok {
		return v
	} else {
		return "not-defined"
	}
}

// recordLatency 返回一个MiddlewareFunc，能计算处理请求的耗时，并打印出来。
func recordLatency() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			startTime := time.Now()
			if err := next(c); err != nil {
				c.Error(err)
			}
			duration := time.Now().Sub(startTime)
			// 打印日志
			sessionID := c.Get(echoKeySessionID).(string)
			mylog.CloneLogger().
				WithTag("session_id", sessionID).
				WithFields(mylog.String("latency", duration.String())).
				Info("request done")
			return nil
		}
	}
}

// dumpRequest 返回一个MiddlewareFunc，在处理请求之前打印请求信息（包括请求头和请求体）
func dumpRequest() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			// Request
			var reqBody []byte
			if c.Request().Body != nil { // Read
				reqBody, _ = ioutil.ReadAll(c.Request().Body)
			}
			c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(reqBody)) // Reset
			sessionID := c.Get(echoKeySessionID).(string)
			mylog.CloneLogger().
				WithTag("session_id", sessionID).
				WithFields(mylog.String("method", req.Method)).
				WithFields(mylog.String("url", req.URL.String())).
				WithFields(mylog.String("headers", fmt.Sprintf("%+v", req.Header))).
				WithFields(mylog.String("body", escapeInvisible(string(reqBody)))).
				Info("received HTTP request")
			return next(c)
		}
	}
}

func escapeInvisible(raw string) string {
	working := raw
	working = strings.ReplaceAll(working, "\r", "\\r")
	working = strings.ReplaceAll(working, "\n", "\\n")
	working = strings.ReplaceAll(working, "\t", "\\t")
	return working
}

// dumpResponse 返回一个MiddlewareFunc，在处理请求之后打印响应信息（包括响应码和响应体）。
func dumpResponse() echo.MiddlewareFunc {
	return middleware.BodyDump(
		func(c echo.Context, reqBody, resBody []byte) {
			rec := c.Response()
			sessionID := c.Get(echoKeySessionID).(string)
			mylog.CloneLogger().
				WithTag("session_id", sessionID).
				WithFields(mylog.Int("status", rec.Status)).
				WithFields(mylog.String("body", escapeInvisible(string(resBody)))).
				Info("sending HTTP response")
		},
	)
}
