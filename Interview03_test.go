package test

import "testing"

type student struct {
	Name string
	Age  int
}

// 以下代码有什么问题，说明原因
func Test_student(t *testing.T) {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	for _, stu := range stus {
		m[stu.Name] = &stu
	}
}
