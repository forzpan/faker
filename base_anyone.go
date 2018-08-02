package faker

import (
	"math/rand"
)

// SliceMatrix 切片组合母体
// 实现各个子母体按比例输出场景
// 子母体全为ConstantMatrix，且均衡比例输出的情况下，不建议使用SliceMatrix
type SliceMatrix struct {
	matrixs []Matrix   //组件母体
	rnd     *rand.Rand //随机数
}

// NewSliceMatrix 构造一个 SliceMatrix
func NewSliceMatrix(ms ...Matrix) *SliceMatrix {
	matrixs := make([]Matrix, len(ms))
	for i, m := range ms {
		matrixs[i] = m
	}
	return &SliceMatrix{
		matrixs,
		NewRand(),
	}
}

// Add 往SliceMatrix里增加Matrix
func (m *SliceMatrix) Add(ms ...Matrix) {
	if len(ms) > 0 {
		m.matrixs = append(m.matrixs, ms...)
	}
}

// Breed 实现接口Matrix,每次随机一个子母体执行
func (m *SliceMatrix) Breed() string {
	return m.matrixs[m.rnd.Intn(len(m.matrixs))].Breed()
}
