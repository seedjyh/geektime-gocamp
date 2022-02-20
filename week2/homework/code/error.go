package code

import "github.com/pkg/errors"

var (
	NotFound = errors.New("Not Found") // 某个数据没找到
	Internal = errors.New("Internal")  // 依赖库的内部错误
)
