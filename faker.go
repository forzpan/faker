package faker

var Delimiter byte = '\t'

const Chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// Matrix 母体接口
type Matrix interface {
	Breed() string //生产字串
}
