package test

import (
	"fmt"
	"testing"
)

type People2 struct{}

func (p *People2) ShowA() {
	fmt.Println("showA")
	p.ShowB()
}
func (p *People2) ShowB() {
	fmt.Println("showB")
}

type Teacher2 struct {
	People2
}

func (t *Teacher2) ShowB() {
	fmt.Println("teacher showB")
}

// 下面代码会输出什么？
func Test05(t *testing.T) {
	//t := Teacher2{}
	//t.ShowA()
}
