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
}
