package faker

import (
	"math/rand"
)

// CompositeMatrix 混合母体，控制各母体的比例
type CompositeMatrix struct {
	matrixs []Matrix   //组件母体
	rnd     *rand.Rand //随机数
}

// NewCompositeMatrix 构造一个 CompositeMatrix
func NewCompositeMatrix(ms ...Matrix) *CompositeMatrix {
	matrixs := make([]Matrix, len(ms))
	for i, m := range ms {
		matrixs[i] = m
	}
	return &CompositeMatrix{
		matrixs,
		getrand(),
	}
}

// Add 往CompositeMatrix里增加Matrix
func (m *CompositeMatrix) Add(ms ...Matrix) {
	if len(ms) > 0 {
		m.matrixs = append(m.matrixs, ms...)
	}
}

// Breed 实现接口matrix
func (m *CompositeMatrix) Breed() string {
	return m.matrixs[m.rnd.Intn(len(m.matrixs))].Breed()
}
