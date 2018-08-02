package faker

import "strings"

// IPV4Matrix IPV4母体
type IPV4Matrix struct {
	im *IntMatrix
}

// NewIPV4Matrix 构造一个 SexMatrix
func NewIPV4Matrix() *IPV4Matrix {
	matrix, _ := NewIntMatrix(0, 256, 10)
	return &IPV4Matrix{
		im: matrix,
	}
}

// Breed 实现接口matrix
func (m *IPV4Matrix) Breed() string {
	return strings.Join([]string{m.im.Breed(), m.im.Breed(), m.im.Breed(), m.im.Breed()}, ".")
}
