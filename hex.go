package faker

import (
	"strconv"
)

// HexMatrix 十六进制整数母体
type HexMatrix IntMatrix

// NewHexMatrix 构造一个 HexMatrix
func NewHexMatrix(min, max int) (*HexMatrix, error) {
	rv, err := NewIntMatrix(min, max)
	if err != nil {
		return nil, err
	}
	return (*HexMatrix)(rv), nil
}

// Breed 实现接口matrix
func (m *HexMatrix) Breed() string {
	return strconv.FormatInt(int64((*IntMatrix)(m).breed()), 16)
}
