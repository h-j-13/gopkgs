package add

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	// 最基础的方式
	got := Add(0, 1)
	if got != 1 {
		t.Errorf("Abs(-1) = %d; want 1", got)
	}

	// assert
	assert.Equal(t, 3, Add(1, 2), "The two words should be the same.")

	// convey
	convey.Convey("use convey for ut", t, func() {
		convey.So(Add(1, 1), convey.ShouldEqual, 2)
		convey.So(Add(1, 2), convey.ShouldNotEqual, 2)
		convey.So(Add(1, 2), convey.ShouldEqual, 3)
	})
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(i, i+1)
	}
	// 基准测试不能直接go test允许, 需要 go test --bench=Add

	// goos: windows
	// goarch: amd64
	// pkg: github.com/h-j-13/gopkgs/ut/add
	// cpu: 13th Gen Intel(R) Core(TM) i7-13700KF
	// BenchmarkAdd         执行次数              每次时间
	// BenchmarkAdd-24    	1000000000	         0.09749 ns/op
	// PASS

	// go test -bench=Add -benchmem
	// goos: windows
	// goarch: amd64
	// pkg: github.com/h-j-13/gopkgs/ut/add
	// cpu: 13th Gen Intel(R) Core(TM) i7-13700KF
	//                                                                        每次执行内存分配大小 每次执行内存分配次数
	// BenchmarkAdd-24         1000000000               0.09905 ns/op         0 B/op          0 allocs/op
	// PASS
	// ok      github.com/h-j-13/gopkgs/ut/add 0.259s

	// -benchtime=5s 执行时间
	// -count=1 执行次数
}
