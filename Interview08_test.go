package test

import (
	"testing"
	"fmt"
)

// 请写出以下输入内容
func Test08(t *testing.T) {
	s := make([]int, 5)
	s = append(s, 1, 2, 3)
	fmt.Println(s)
}
