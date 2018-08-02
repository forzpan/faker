package faker

import (
	"errors"
)

// StringMatrix 字串母体
type StringMatrix struct {
	chars      []rune //候选字符集
	length     uint32 //定长字串，长度
	*IntMatrix        //如果是变长字串
}

// NewStringMatrix 构造一个变长 StringMatrix
func NewStringMatrix(min, max uint32, chars []rune) (*StringMatrix, error) {
	if 0 == len(chars) {
		return nil, errors.New("候选字符集不能为空！")
	}
	im, err := NewIntMatrix(int64(min), int64(max), 10)
	if err != nil {
		return nil, err
	}
	return &StringMatrix{
		chars,
		0,
		im,
	}, nil
}

// NewLenStringMatrix 构造一个定长 StringMatrix
func NewLenStringMatrix(length uint32, chars []rune) (*StringMatrix, error) {
	if length == 0 {
		return nil, errors.New("字串长度不能为0！")
	}
	if 0 == len(chars) {
		return nil, errors.New("候选字符集不能为空！")
	}
	return &StringMatrix{
		chars,
		length,
		nil,
	}, nil
}

// Breed 实现接口matrix
func (m *StringMatrix) Breed() string {
	var rv []rune
	if m.IntMatrix == nil {
		rv = make([]rune, m.length)
	} else {
		rv = make([]rune, m.IntMatrix.Int64())
	}
	for i := range rv {
		rv[i] = m.chars[m.rnd.Intn(len(m.chars))]
	}
	return string(rv)
}
