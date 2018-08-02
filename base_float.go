package faker

import (
	"fmt"
	"math/rand"
	"strconv"
)

// FloatMatrix 浮点数母体
type FloatMatrix struct {
	min  float64    //最小值
	max  float64    //最大值
	prec int        //精度
	rnd  *rand.Rand //随机数

}

// NewFloatMatrix 构造一个 FloatMatrix [min,max)
func NewFloatMatrix(min, max float64, prec int) (*FloatMatrix, error) {
	if !(min < max) {
		return nil, fmt.Errorf("范围不合法！")
	}
	return &FloatMatrix{
		min,
		max,
		prec,
		NewRand(),
	}, nil
}

// Float64 产生一个 float64
func (m *FloatMatrix) Float64() float64 {
	return m.min + m.rnd.Float64()*(m.max-m.min)
}

// Breed 实现接口matrix
func (m *FloatMatrix) Breed() string {
	return strconv.FormatFloat(m.Float64(), 'f', m.prec, 64)
}
