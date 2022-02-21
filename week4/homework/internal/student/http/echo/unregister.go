package echo

import (
	"context"
	"errors"
	"geektime-gocamp/week4/homework/internal/student"
	"github.com/labstack/echo/v4"
	"net/http"
)

type unregister struct {
	service *student.Service
}

func (ur *unregister) Handle(c echo.Context) error {
	uid := c.Param("uid")
	if len(uid) == 0 {
		return c.JSON(http.StatusBadRequest, &ResponseBody{
			Code:    2,
			Message: "uid shouldn't be empty",
		})
	}
	if err := ur.service.Unregister(context.TODO(), student.UID(uid)); err != nil {
		if errors.Is(err, student.ErrNotFound) {
			return c.JSON(http.StatusNotFound, &ResponseBody{
				Code:    2,
				Message: err.Error(),
			})
		}
		return c.JSON(http.StatusBadRequest, &ResponseBody{
			Code:    2,
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, defaultSuccessResponse)
}
