package faker

// SexMatrix 性别母体
type SexMatrix = EnumMatrix

// NewSexMatrix 构造一个 SexMatrix
func NewSexMatrix() *SexMatrix {
	vals := []string{"男", "女"}
	matrix, _ := NewEnumMatrix(vals)
	return matrix
}
