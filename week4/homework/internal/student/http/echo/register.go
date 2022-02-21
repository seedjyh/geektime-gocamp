package echo

import (
	"context"
	"errors"
	"geektime-gocamp/week4/homework/internal/pkg/code"
	"geektime-gocamp/week4/homework/internal/student"
	"github.com/labstack/echo/v4"
	"net/http"
)

type register struct {
	service *student.Service
}

type registerRequest struct {
	UID      string `json:"uid"`
	RealName string `json:"realName"`
}

func (r *register) Handle(c echo.Context) error {
	req := new(registerRequest)
	if err := c.Bind(req); err != nil {
		return err
	}
	do := &student.StudentDO{
		UID:      student.UID(req.UID),
		RealName: student.RealName(req.RealName),
	}
	if err := r.service.Register(context.TODO(), do); err != nil {
		var codeErr *code.Error
		if errors.As(err, &codeErr) {
			return c.JSON(http.StatusBadRequest, &ResponseBody{
				Code:    codeErr.Code,
				Message: err.Error(),
			})
		} else {
			return c.JSON(http.StatusBadRequest, &ResponseBody{
				Code:    -1,
				Message: err.Error(),
			})
		}
	}
	return c.JSON(http.StatusOK, defaultSuccessResponse)
}
