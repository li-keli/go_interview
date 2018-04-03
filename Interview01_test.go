package test

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

type People1 interface {
	Show()
}

type Student struct{}

func (stu *Student) Show() {

}

func live() People1 {
	var stu *Student
	return stu
}

func Test_interface(t *testing.T) {
	Convey("比较interface和nil", t, func() {
		// 解答：https://studygolang.com/articles/1908
		So(live(), ShouldBeTrue)
	})
}
