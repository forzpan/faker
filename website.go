package faker

// WebMatrix 网站母体
type WebMatrix struct {
	matrixs []Matrix
}

// NewWebMatrix 构造一个 WebMatrix
func NewWebMatrix() *WebMatrix {
	head, _ := NewEnumMatrix([]string{"http://", "https://", ""})
	body, _ := NewStringMatrix(5, 12, []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"))
	tail, _ := NewEnumMatrix([]string{".com", ".com.cn", ".gov.cn", ".edu.cn", ".org", ".net", ".info", ".io"})

	matrixs := []Matrix{head, body, tail}

	return &WebMatrix{
		matrixs,
	}
}

// Breed 实现接口matrix
func (m *WebMatrix) Breed() string {
	return m.matrixs[0].Breed() + "www." + m.matrixs[1].Breed() + m.matrixs[2].Breed()
}
