package faker

// TelephoneMatrix 固定电话母体
type TelephoneMatrix struct {
	matrixs []Matrix
}

// NewTelephoneMatrix 构造一个 TelephoneMatrix
func NewTelephoneMatrix() *TelephoneMatrix {
	area1, _ := NewEnumMatrix([]string{"10", "20", "21", "22", "23", "24", "25", "26", "27", "28", "29"})
	area2, _ := NewIntMatrix(310, 1000, 10)
	area := NewSliceMatrix(area1, area2, area2, area2, area2, area2, area2, area2, area2, area2)
	code, _ := NewIntMatrix(8000000, 100000000, 10)

	matrixs := []Matrix{area, code}

	return &TelephoneMatrix{
		matrixs,
	}
}

// Breed 实现接口matrix
func (m *TelephoneMatrix) Breed() string {
	return "0" + m.matrixs[0].Breed() + "-" + m.matrixs[1].Breed()
}
