package errutils

import (
	"testing"

	"errors"
	"fmt"

	. "github.com/smartystreets/goconvey/convey"
)

func TestJoinErrsStr(t *testing.T) {
	Convey("No delimiter with no errors", t, func() {
		expected := ""
		actual := JoinErrsStr(", ")

		So(actual, ShouldEqual, expected)
	})

	Convey("No delimiter with errors", t, func() {
		err1 := errors.New("err1")
		err2 := errors.New("err2")
		err3 := errors.New("err3")
		expected := fmt.Sprint(err1.Error(), err2.Error(), err3.Error())
		actual := JoinErrsStr("", err1, err2, err3)

		So(actual, ShouldEqual, expected)
	})

	Convey("With delimiter and with errors", t, func() {
		err1 := errors.New("err1")
		err2 := errors.New("err2")
		err3 := errors.New("err3")
		expected := fmt.Sprintf("%s, %s, %s", err1.Error(), err2.Error(), err3.Error())
		actual := JoinErrsStr(", ", err1, err2, err3)

		So(actual, ShouldEqual, expected)
	})
}

func TestJoinErrs(t *testing.T) {
	Convey("No delimiter with no errors", t, func() {
		actual := JoinErrs(", ")

		So(actual, ShouldBeNil)
	})

	Convey("No delimiter with errors", t, func() {
		err1 := errors.New("err1")
		err2 := errors.New("err2")
		err3 := errors.New("err3")
		expected := fmt.Errorf("%s%s%s", err1.Error(), err2.Error(), err3.Error())
		actual := JoinErrs("", err1, err2, err3)

		So(actual, ShouldResemble, expected)
	})

	Convey("With delimiter and with errors", t, func() {
		err1 := errors.New("err1")
		err2 := errors.New("err2")
		err3 := errors.New("err3")
		expected := fmt.Errorf("%s, %s, %s", err1.Error(), err2.Error(), err3.Error())
		actual := JoinErrs(", ", err1, err2, err3)

		So(actual, ShouldResemble, expected)
	})
}
