package faker

// BoolMatrix 布尔母体
type BoolMatrix = EnumMatrix

// NewBoolMatrix 构造一个 BoolMatrix
func NewBoolMatrix() *BoolMatrix {
	vals := []string{"true", "false"}
	matrix, _ := NewEnumMatrix(vals)
	return matrix
}
