package faker

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
)

// IntMatrix 整数母体
type IntMatrix struct {
	min  int64      //最小值
	max  int64      //最大值
	base int        //进制[2,36]
	rnd  *rand.Rand //随机数
}

// NewIntMatrix 构造一个 IntMatrix [min,max)
func NewIntMatrix(min, max int64, base int) (*IntMatrix, error) {
	if !(min < max) {
		return nil, errors.New("范围不合法！")
	}
	if base < 2 || base > 36 {
		return nil, fmt.Errorf("不支持表示为%d进制！", base)
	}
	return &IntMatrix{
		min,
		max,
		base,
		NewRand(),
	}, nil
}

// Int64 产生一个 int64
func (m *IntMatrix) Int64() int64 {
	return m.min + int64(m.rnd.Intn(int(m.max-m.min)))
}

// Breed 实现接口matrix
func (m *IntMatrix) Breed() string {
	return strconv.FormatInt(m.Int64(), m.base)
}
