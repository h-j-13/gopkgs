package add_test

import (
	"github.com/h-j-13/gopkgs/ut/add"
	"github.com/smartystreets/goconvey/convey"
	"testing"

	"github.com/stretchr/testify/assert"
)

// 黑盒测试方式

func TestAdd(t *testing.T) {
	// 最基础的方式
	got := add.Add(0, 1)
	if got != 1 {
		t.Errorf("Abs(-1) = %d; want 1", got)
	}

	// assert
	assert.Equal(t, 3, add.Add(1, 2), "The two words should be the same.")

	// convey
	convey.Convey("use convey for ut", t, func() {
		convey.So(add.Add(1, 1), convey.ShouldEqual, 2)
		convey.So(add.Add(1, 2), convey.ShouldNotEqual, 2)
		convey.So(add.Add(1, 2), convey.ShouldEqual, 3)
	})
}
