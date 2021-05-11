package alg

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestStringSliceEqual(t *testing.T) {
	Convey("TestStringSliceEqual should return when a != nil && b != nil", t, func() {
		a := []string{"hello", "goconvey"}
		b := []string{"hello", "goconvey"}
		So(StringSliceEqual(a, b), ShouldBeTrue)
	})

	Convey("TestStringSliceEqual should return ture when a == nil && b == nil", t, func() {
		So(StringSliceEqual(nil, nil), ShouldBeTrue)
	})

	Convey("TestStringSliceEqual should return false when a == nil && b != nil", t, func() {
		a := []string(nil)
		b := []string{}
		So(StringSliceEqual(a, b), ShouldBeFalse)
	})

	Convey("TestStringSliceEqual should return false when a != nil && b != nil", t, func() {
		a := []string{"hello", "world"}
		b := []string{"hello", "goconvey"}
		So(StringSliceEqual(a, b), ShouldBeFalse)
	})

	Convey("TestStringSliceEqual should return false when len(a) != len(b)", t, func() {
		a := []string{"hello"}
		b := []string{"hello", "world"}
		So(StringSliceEqual(a, b), ShouldBeFalse)
	})
}

func TestStringSliceEqual2(t *testing.T) {
	Convey("TestStringSliceEqual,", t, func() {
		Convey("should return when a != nil && b != nil", func() {
			a := []string{"hello", "goconvey"}
			b := []string{"hello", "goconvey"}
			So(StringSliceEqual(a, b), ShouldBeTrue)
		})

		Convey("should return ture when a == nil && b == nil", func() {
			So(StringSliceEqual(nil, nil), ShouldBeTrue)
		})

		Convey("should return false when a == nil && b != nil", func() {
			a := []string(nil)
			b := []string{}
			So(StringSliceEqual(a, b), ShouldBeFalse)
		})

		Convey("should return false when a != nil && b != nil", func() {
			a := []string{"hello", "world"}
			b := []string{"hello", "goconvey"}
			So(StringSliceEqual(a, b), ShouldBeFalse)
		})

		Convey("should return false when len(a) != len(b)", func() {
			a := []string{"hello"}
			b := []string{"hello", "world"}
			So(StringSliceEqual(a, b), ShouldBeFalse)
		})
	})
}

func TestShouldSummerBeComing(t *testing.T) {
	Convey("TestSummer", t, func() {
		So("summer", ShouldSummerBeComing, "coming")
		So("winter", ShouldSummerBeComing, "coming")
	})
}
