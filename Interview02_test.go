package test

import (
	"fmt"
	"testing"
)

// 写出下面代码输出内容。
func Test_Call(t *testing.T) {
	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()

	panic("触发异常")
}
