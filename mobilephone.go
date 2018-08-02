package faker

// MobilephoneMatrix 手机号母体
type MobilephoneMatrix struct {
	matrixs []Matrix
}

// NewMobilephoneMatrix 构造一个 MobilephoneMatrix
func NewMobilephoneMatrix() *MobilephoneMatrix {
	head, _ := NewEnumMatrix([]string{"130", "131", "132", "133", "134", "135", "136", "137", "138", "139", "145", "146",
		"147", "149", "150", "151", "152", "153", "155", "156", "157", "158", "159", "166", "170", "171", "172", "173",
		"175", "176", "177", "178", "180", "181", "182", "183", "184", "185", "186", "187", "188", "189", "198", "199"})
	code, _ := NewIntMatrix(10000000, 100000000, 10)

	matrixs := []Matrix{head, code}

	return &MobilephoneMatrix{
		matrixs,
	}
}

// Breed 实现接口matrix
func (m *MobilephoneMatrix) Breed() string {
	return m.matrixs[0].Breed() + m.matrixs[1].Breed()
}
