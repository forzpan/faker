package faker

import (
	"fmt"
	"math/rand"
)

// FixedLenStringMatrix 定长字串母体
type FixedLenStringMatrix struct {
	chars  string     //基础字符库
	length int        //字串长度
	rnd    *rand.Rand //随机数
}

// NewFixedLenStringMatrix 构造一个 FixedLenStringMatrix
func NewFixedLenStringMatrix(length int, chars string) (*FixedLenStringMatrix, error) {
	if 0 == len(chars) {
		return nil, fmt.Errorf("基础字符库不能为空！")
	}
	if 0 == length {
		return nil, fmt.Errorf("len[%d]，试图构造空字串，请使用固定值构造！", length)
	}
	return &FixedLenStringMatrix{
		chars,
		length,
		getrand(),
	}, nil
}

// Breed 实现接口matrix
func (m *FixedLenStringMatrix) Breed() string {
	rv := make([]byte, m.length)
	for i := range rv {
		rv[i] = m.chars[m.rnd.Intn(len(m.chars))]
	}
	return string(rv)
}

// RandLenStringMatrix 变长字串母体
type RandLenStringMatrix struct {
	chars string //基础字符库
	*IntMatrix
}

// NewRandLenStringMatrix 构造一个 RandLenStringMatrix
func NewRandLenStringMatrix(min, max int, chars string) (*RandLenStringMatrix, error) {
	if 0 == len(chars) {
		return nil, fmt.Errorf("基础字符库不能为空！")
	}
	im, err := NewIntMatrix(min, max)
	if err != nil {
		return nil, err
	}
	return &RandLenStringMatrix{
		chars,
		im,
	}, nil
}

// Breed 实现接口matrix
func (m *RandLenStringMatrix) Breed() string {
	rv := make([]byte, m.IntMatrix.breed())
	for i := range rv {
		rv[i] = m.chars[m.rnd.Intn(len(m.chars))]
	}
	return string(rv)
}
