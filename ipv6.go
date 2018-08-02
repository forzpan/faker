package faker

import "strings"

// IPV6Matrix IPV6母体
type IPV6Matrix struct {
	im *IntMatrix
}

// NewIPV6Matrix 构造一个 SexMatrix
func NewIPV6Matrix() *IPV6Matrix {
	matrix, _ := NewIntMatrix(0, 65536, 16)
	return &IPV6Matrix{
		im: matrix,
	}
}

// Breed 实现接口matrix
func (m *IPV6Matrix) Breed() string {
	return strings.Join([]string{m.im.Breed(), m.im.Breed(), m.im.Breed(), m.im.Breed(), m.im.Breed(), m.im.Breed(), m.im.Breed(), m.im.Breed()}, ":")
}
