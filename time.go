package faker

import (
	"fmt"
	"math/rand"
	"strconv"
)

// TimeMatrix 时间母体
type TimeMatrix struct {
	min uint       //最小时间
	max uint       //最大时间
	rnd *rand.Rand //随机数
}

// NewTimeMatrix 构造一个 TimeMatrix
func NewTimeMatrix(min, max uint，fmt string) (*TimeMatrix, error) {
	if min > max {
		return nil, fmt.Errorf("min[%d] > max[%d]，请反转它们！", min, max)
	}
	if min == max {
		return nil, fmt.Errorf("min[%d] == max[%d]，请使用固定值构造！", min, max)
	}
	return &TimeMatrix{
		min,
		max,
		getrand(),
	}, nil
}

func (m *TimeMatrix) breed() int {
	return m.min + m.rnd.Intn(m.max-m.min+1)
}

// Breed 实现接口matrix
func (m *TimeMatrix) Breed() string {
	return strconv.Itoa(m.breed())
}
