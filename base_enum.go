package faker

import (
	"errors"
	"math/rand"
	"strings"
)

// EnumMatrix 枚举母体
// 从枚举集合里随机选一个value输出
type EnumMatrix struct {
	vals []string   //枚举值
	rnd  *rand.Rand //随机数
}

// NewEnumMatrix 构造一个 EnumMatrix
func NewEnumMatrix(strs []string) (*EnumMatrix, error) {
	vals := make([]string, len(strs))
	for i, str := range strs {
		if strings.ContainsAny(str, Delimiter()) {
			return nil, errors.New("不能包含分隔符和换行符")
		}
		vals[i] = str
	}
	return &EnumMatrix{
		vals,
		NewRand(),
	}, nil
}

// Breed 实现接口Matrix
func (m *EnumMatrix) Breed() string {
	return m.vals[m.rnd.Intn(len(m.vals))]
}
