package main

import (
	"context"
	"fmt"
	sd "geektime-gocamp/week2/homework/dao/student"
)

func main() {
	dsn := "root:123456@tcp(localhost:3306)/geektime"
	studentDAO, err := sd.NewStudentDAO(dsn)
	if err != nil {
		fmt.Printf("Failed! error=[%+v]", err)
	}
	// id := 1
	id := 2
	if s, found, err := studentDAO.FindByID(context.Background(), id); err != nil {
		fmt.Printf("Failed! %+v", err)
	} else if !found {
		fmt.Printf("No such student. ID=[%+v]", id)
	} else {
		fmt.Printf("Found the student. student=[%+v]", s)
	}
}
