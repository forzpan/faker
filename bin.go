package faker

import (
	"strconv"
)

// BinMatrix 二进制整数母体
type BinMatrix IntMatrix

// NewBinMatrix 构造一个 BinMatrix
func NewBinMatrix(min, max int) (*BinMatrix, error) {
	rv, err := NewIntMatrix(min, max)
	if err != nil {
		return nil, err
	}
	return (*BinMatrix)(rv), nil
}

// Breed 实现接口matrix
func (m *BinMatrix) Breed() string {
	return strconv.FormatInt(int64((*IntMatrix)(m).breed()), 2)
}
