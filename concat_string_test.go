package ch40

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"testing"
)

const numbers = 100

func BenchmarkSprintf(b *testing.B) {
	b.ResetTimer()
	for idx := 0; idx < b.N; idx++ {
		var s string
		for i := 0; i < numbers; i++ {
			s = fmt.Sprintf("%v%v", s, i)
		}
	}
	b.StopTimer()
}

// 推荐使用1
func BenchmarkStringBuilder(b *testing.B) {
	b.ResetTimer()
	for idx := 0; idx < b.N; idx++ {
		var builder strings.Builder
		for i := 0; i < numbers; i++ {
			builder.WriteString(strconv.Itoa(i))
		}
		_ = builder.String()
	}
	b.StopTimer()
}

// 推荐使用2
func BenchmarkBytesBuf(b *testing.B) {
	b.ResetTimer()
	for idx := 0; idx < b.N; idx++ {
		var buf bytes.Buffer
		for i := 0; i < numbers; i++ {
			buf.WriteString(strconv.Itoa(i))
		}
		_ = buf.String()
	}
	b.StopTimer()
}

func BenchmarkStringAdd(b *testing.B) {
	b.ResetTimer()
	for idx := 0; idx < b.N; idx++ {
		var s string
		for i := 0; i < numbers; i++ {
			s += strconv.Itoa(i)
		}
	}
	b.StopTimer()
}

func TestBook(t *testing.T) {
	//s := make([]int, 0, 10)
	//s = append(s, 10)
	//t.Log(s)

	data := [3]int{10, 20, 30}
	//for i, x := range data {
	//	if i == 0 {
	//		data[0] += 100
	//		data[1] += 200
	//		data[2] += 300
	//	}
	//	t.Logf("x: %d, data: %d\n", x, data[i])
	//}
	for i, x := range data[:] {
		if i == 0 {
			data[0] += 100
			data[1] += 200
			data[2] += 300
		}
		t.Logf("x: %d, data: %d\n", x, data[i])
	}
}

func add(x, y int) (z int) {
	{   // 不是说 { 不能单独占用一行吗？为什么这一行不会报错呢
		z := x + y
		return z
	}
}

func div(x, y int) (z int, err error) {
	if y == 0 {
		err = errors.New("Division by zero.")
		return
	}
	z = x / y
	return
}