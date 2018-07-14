package main

import (
	"fmt"
	"unsafe"
)

// var baseMatrixs = make(map[string]faker.Matrix, 0)    // 基础母体，即使root用户也不可以删除
// var defaultMatrixs = make(map[string]faker.Matrix, 0) // 默认的一些母体，root用户定义的所有母体都是默认母体，普通用户可以向root用户推荐
// var userMatrixs = make(map[string]faker.Matrix, 0)    // 用户定义的一些母体，可以向其他用户推荐

// 测试git
func main() {
	s := [3]byte{1, 2, 3}

	a := (*[20]byte)(unsafe.Pointer(&s))

	fmt.Println(a)

}
