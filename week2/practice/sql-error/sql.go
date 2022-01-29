package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

func main() {
	dsn := "root:123456@tcp(localhost:3306)/geektime"
	if err := run(dsn, 2); err != nil {
		fmt.Printf("%v\n", err)
		fmt.Println("=============")
		fmt.Printf("%+v\n", err)
		fmt.Println("=============")
	} else {
		fmt.Println("done!")
	}
}

func run(dsn string, id int) error {
	ctx := context.Background()
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return errors.Wrapf(err, "define mysql connection, dsn=%+v", dsn)
	}
	if err := db.PingContext(ctx); err != nil {
		return errors.Wrapf(err, "build connection too mysql failed")
	}
	// 1) try QueryContext, Next, Scan
	//rows, err := db.QueryContext(ctx, "SELECT name, age FROM student WHERE id = ?", id)
	//if err != nil {
	//	return errors.Wrapf(err, "query failed")
	//}
	//defer rows.Close()
	//var name string
	//var age int
	//for rows.Next() {
	//	var name string
	//	var age int
	//	if err := rows.Scan(&name, &age); err != nil {
	//		return errors.Wrapf(err, "parse row failed")
	//	}
	//	fmt.Printf("Row> name=%s, age=%d\n", name, age)
	//}
	//if rows.Err() != nil {
	//	return errors.Wrapf(err, "failed while scanning rows")
	//}
	// 2) QueryRows, Scan
	var name string
	var age int
	if err := db.QueryRow("SELECT name, age FROM student WHERE id = ?", id).Scan(&name, &age); err != nil {
		fmt.Println("check err 1:", errors.Is(err, sql.ErrNoRows))
		fmt.Println("check err 2:", errors.Cause(err) == sql.ErrNoRows)
		return errors.Wrap(err, "QueryRow and Scan failed")
	} else {
		fmt.Printf("Row> name=%s, age=%d\n", name, age)
	}
	return nil
}
