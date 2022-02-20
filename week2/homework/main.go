package main

import (
	"context"
	"fmt"
	"geektime-gocamp/week2/homework/code"
	sd "geektime-gocamp/week2/homework/dao/student"
	"github.com/pkg/errors"
)

func main() {
	dsn := "root:123456@tcp(localhost:3306)/geektime"
	studentDAO, err := sd.NewStudentDAO(dsn)
	if err != nil {
		fmt.Printf("Failed! error=[%+v]", err)
	}
	// id := 1
	id := 2
	if s, err := studentDAO.FindByID(context.Background(), id); err != nil {
		if errors.Is(err, code.NotFound) {
			fmt.Printf("No such student. ID=[%+v], error=[%+v]", id, err)
		} else {
			fmt.Printf("Failed! error=[%+v]", err)
		}
	} else {
		fmt.Printf("Found the student. student=[%+v]", s)
	}
}
