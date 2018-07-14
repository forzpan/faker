package faker

import (
	"fmt"
	"strings"
)

// ConstantMatrix 固定值母体
type ConstantMatrix struct {
	val string
}

// NewConstantMatrix 构造一个 ConstantMatrix
func NewConstantMatrix(val string) (*ConstantMatrix, error) {
	dchars := delimiter()
	if strings.ContainsAny(val, dchars) {
		return nil, fmt.Errorf("数据中不允许包含定界符号集[%s]中的字符，请检查！", dchars)
	}
	return &ConstantMatrix{val}, nil
}

// Breed 实现接口matrix
func (m *ConstantMatrix) Breed() string {
	return m.val
}
