package faker

import (
	"fmt"
)

// ZipcodeMatrix 邮编母体
type ZipcodeMatrix struct {
	*IntMatrix
}

// NewZipcodeMatrix 构造一个 ZipcodeMatrix
func NewZipcodeMatrix() *ZipcodeMatrix {
	matrix, _ := NewIntMatrix(10000, 870000, 10)
	return &ZipcodeMatrix{
		matrix,
	}
}

// Breed 实现接口matrix
func (m *ZipcodeMatrix) Breed() string {
	return fmt.Sprintf("%06d", m.Int64())
}
