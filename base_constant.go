package faker

import (
	"errors"
	"strings"
)

// ConstantMatrix 固定值母体
type ConstantMatrix struct {
	val string
}

// NewConstantMatrix 构造一个 ConstantMatrix
func NewConstantMatrix(val string) (*ConstantMatrix, error) {
	if strings.ContainsAny(val, Delimiter()) {
		return nil, errors.New("不能包含分隔符和换行符")
	}
	return &ConstantMatrix{val}, nil
}

// Breed 实现接口matrix
func (m *ConstantMatrix) Breed() string {
	return m.val
}
