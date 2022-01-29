// package student is the repository via MySQL.
//mysql table definition
//
//mysql> desc student;
//+-------+-------------+------+-----+---------+----------------+
//| Field | Type        | Null | Key | Default | Extra          |
//+-------+-------------+------+-----+---------+----------------+
//| id    | bigint      | NO   | PRI | NULL    | auto_increment |
//| name  | varchar(32) | YES  |     | NULL    |                |
//| age   | int         | YES  |     | NULL    |                |
//+-------+-------------+------+-----+---------+----------------+
//
//mysql> select * from student;
//+----+---------+------+
//| id | name    | age  |
//+----+---------+------+
//|  1 | seedjyh |   38 |
//+----+---------+------+

package student

import (
	"context"
	"database/sql"
	"geektime-gocamp/week2/homework/dao"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

type studentDAO struct {
	db *sql.DB
}

func NewStudentDAO(dsn string) (*studentDAO, error) {
	if db, err := sql.Open("mysql", dsn); err != nil {
		return nil, err
	} else {
		return &studentDAO{
			db: db,
		}, nil
	}
}

func (s *studentDAO) FindByID(ctx context.Context, id int) (*dao.Student, bool, error) {
	student := new(dao.Student)
	if err := s.db.QueryRow("SELECT id, name, age FROM student WHERE id = ?", id).Scan(&student.Id, &student.Name, &student.Age); err != nil {
		if err == sql.ErrNoRows {
			return nil, false, nil
		} else {
			return nil, false, errors.Wrapf(err, "Query student by ID failed, id=[%+v]", id)
		}
	} else {
		return student, true, nil
	}
}
