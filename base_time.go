package faker

import (
	"fmt"
	"math/rand"
	"time"
)

// TimeMatrix 时间母体
type TimeMatrix struct {
	min    time.Time  //最小值
	max    time.Time  //最大值
	layout string     //输出格式
	rnd    *rand.Rand //随机数
}

// NewTimeMatrix 构造一个 TimeMatrix
func NewTimeMatrix(min, max time.Time, layout string) (*TimeMatrix, error) {
	if !min.Before(max) {
		return nil, fmt.Errorf("时间范围不合法！")
	}
	return &TimeMatrix{
		min,
		max,
		layout,
		NewRand(),
	}, nil
}

// Time 生成一个时间
func (m *TimeMatrix) Time() time.Time {
	d := int(m.max.Sub(m.min))
	rd := time.Duration(m.rnd.Intn(int(d)))
	return m.min.Add(rd)
}

// Breed 实现接口matrix
func (m *TimeMatrix) Breed() string {
	return m.Time().Format(m.layout)
}
