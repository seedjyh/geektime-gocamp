package dao

import "context"

type Student struct {
	Id   int
	Name string
	Age  int
}

type StudentDAO interface {
	// FindByID 根据ID查询一个学生
	// 如果没出错，且找到了数据，返回 (TheStudent, nil)
	// 如果没找到或出错了，返回(nil, TheError)，根据TheError的类型判断是没找到还是出错了。
	FindByID(ctx context.Context, id int) (*Student, error)
}
