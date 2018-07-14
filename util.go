package faker

import (
	"math/rand"
	"time"
)

// 由于随机函数默认不是并发安全的，所以需要大量制造
func getrand() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}

// 定界符号集
func delimiter() string {
	return string([]byte{'\r', '\n', Delimiter})
}
