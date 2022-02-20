package mysql

import (
	"context"
	"fmt"
	"geektime-gocamp/week4/homework/internal/code"
	"geektime-gocamp/week4/homework/internal/student"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/pkg/errors"
	"xorm.io/builder"
)

type repository struct {
	engine *xorm.Engine
}

func New(engine *xorm.Engine) *repository {
	return &repository{engine: engine}
}

func (r *repository) Add(ctx context.Context, do *student.StudentDO) error {
	po := new(studentPO).initFromStudentDO(do)
	if cnt, err := r.engine.InsertOne(po); err != nil {
		return err
	} else if cnt != 1 {
		return fmt.Errorf("some error happened")
	} else {
		return nil
	}
}

func (r *repository) DeleteByUID(ctx context.Context, uid student.UID) error {
	if cnt, err := r.engine.Where(builder.Eq{"uid": uid.String()}).
		Cols("is_deleted").Update(&studentPO{IsDeleted: true}); err != nil {
		return err
	} else if cnt == 0 {
		return errors.Wrap(code.ErrNotFound, fmt.Sprintf("no such student, uid=[%+v]", uid))
	} else {
		return nil
	}
}
