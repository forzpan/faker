package faker

import (
	"fmt"
	"math/rand"
	"strconv"
)

// IntMatrix 整数母体
type IntMatrix struct {
	min int        //最小值
	max int        //最大值
	rnd *rand.Rand //随机数
}

// NewIntMatrix 构造一个 IntMatrix
func NewIntMatrix(min, max int) (*IntMatrix, error) {
	if min > max {
		return nil, fmt.Errorf("min[%d] > max[%d]，请反转它们！", min, max)
	}
	if min == max {
		return nil, fmt.Errorf("min[%d] == max[%d]，请使用固定值构造！", min, max)
	}
	return &IntMatrix{
		min,
		max,
		getrand(),
	}, nil
}

func (m *IntMatrix) breed() int {
	return m.min + m.rnd.Intn(m.max-m.min+1)
}

// Breed 实现接口matrix
func (m *IntMatrix) Breed() string {
	return strconv.Itoa(m.breed())
}
