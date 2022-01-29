package dao

import "context"

type Student struct {
	Id int
	Name string
	Age int
}

type StudentDAO interface {
	// FindByID 根据ID查询一个学生
	// 如果出错了，返回 (nil, false, TheError)
	// 如果没出错，但没找到数据，返回 (nil, false, nil)
	// 如果没出错，且找到了数据，返回 (TheStudent, true, nil)
	FindByID(ctx context.Context, id int) (*Student, bool, error)
}
