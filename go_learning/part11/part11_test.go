package part11

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

//功能测试
func TestAdd(t *testing.T) {

	re := add(1, 3)
	if re != 4 {
		t.Errorf("except %d,actual %d ", 4, re)
	}

}

//过滤耗时多的测试用例
// go test -short
func TestAdd2(t *testing.T) {
	if testing.Short() {
		t.Skip("short 模式下跳过")
	}

	re := add(1, 3)
	if re != 4 {
		t.Errorf("except %d,actual %d ", 4, re)
	}
	fmt.Println("TestAdd2 test over ")
}

// 表格驱动管理数据
func TestAdd3(t *testing.T) {
	data := []struct {
		a   int
		b   int
		out int
	}{
		{1, 2, 3,},
		{23, 45, 68,},
	}

	for _, v := range data {
		if re := add(v.a, v.b); re != v.out {
			t.Errorf("except %d,actual %d ", v.out, re)
		}
	}
}

func BenchmarkAdd(bb *testing.B) {
	a := 123
	b := 456
	c := 579
	for i := 0; i < bb.N; i++ {
		if re := add(a, b); re != c {
			fmt.Printf("%d + %d,except  %d ,actual: %d", a, b, c, re)
		}
	}
}

// 运行所有benchmark 比较
// get test -bench =".*"
/*
  BenchmarkAdd-4          1000000000               0.7500 ns/op
  BenchmarkSprintf-4             1        15770902000 ns/op
  BenchmarkSprintf2-4            1        6410366600 ns/op
  BenchmarkBuild-4             212          20864400 ns/op

*/

const COUNT = 100000

func BenchmarkSprintf(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var temp string
		for j := 0; j <= COUNT; j++ {
			temp = fmt.Sprintf("%s%d", temp, j)
		}
	}
	b.StopTimer()
}

func BenchmarkSprintf2(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var temp string
		for j := 0; j <= COUNT; j++ {
			temp = temp + strconv.Itoa(j)
		}
	}
	b.StopTimer()
}

func BenchmarkBuild(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var build strings.Builder
		for j := 0; j <= COUNT; j++ {
			build.WriteString(strconv.Itoa(j))
		}
		_ = build.String()
	}
	b.StopTimer()
}
