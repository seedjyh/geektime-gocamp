package student

import "geektime-gocamp/week4/homework/internal/pkg/code"

var (
	ErrNotFound = code.NewError(1, "Not found")
	ErrInternal = code.NewError(2, "Internal error")
)
