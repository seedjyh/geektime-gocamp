package echo

import (
	"context"
	students2 "geektime-gocamp/week4/homework/internal/students"
	"github.com/labstack/echo/v4"
	"net/http"
)

type register struct {
	service *students2.Service
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
	do := &students2.StudentDO{
		UID:      students2.UID(req.UID),
		RealName: students2.RealName(req.RealName),
	}
	if err := r.service.Register(context.TODO(), do); err != nil {
		return c.JSON(http.StatusBadRequest, &ResponseBody{
			Code:    1,
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, defaultSuccessResponse)
}
