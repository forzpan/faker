package faker

// BoolMatrix 布尔母体
type BoolMatrix = CompositeMatrix

// NewBoolMatrix 构造一个 BoolMatrix
func NewBoolMatrix() *BoolMatrix {
	return NewCompositeMatrix(&ConstantMatrix{"true"}, &ConstantMatrix{"false"})
}
