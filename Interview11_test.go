package test

import (
	"fmt"
	"testing"
)

type People03 interface {
	Speak(string) string
}

type Stduent03 struct{}

func (stu *Stduent03) Speak(think string) (talk string) {
	if think == "bitch" {
		talk = "You are a good boy"
	} else {
		talk = "hi"
	}
	return
}

// 以下代码能编译过去吗？为什么？
func Test11(t *testing.T) {
	var peo People03 = Stduent03{}
	think := "bitch"
	fmt.Println(peo.Speak(think))
}