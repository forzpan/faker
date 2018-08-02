package faker

import (
	"math/rand"
	"time"
)

// NewRand 创建一个随机函数源
// 由于随机函数源不是并发安全的，全局随机函数源又是带锁的
func NewRand() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}

func Delimiter() string {
	return "\r\n"
}
