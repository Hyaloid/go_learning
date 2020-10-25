package string

import "testing"

func TestString(t *testing.T) {
	var s string
	t.Log(s) //初始化为默认零值
	s = "hello"
	t.Log(len(s))

	// string是一个不可变的byte切片
	s = "\xE4\xB8\xA5"
	t.Log(s)
	t.Log(len(s))

	s = "中"
	t.Log(len(s))
	// Unicode 和 UTF8
	c := []rune(s)
	t.Log(len(c))
	t.Logf("中 unicode %x", c[0])
	t.Logf("中 UTF8 %x", s)
}

func TestStringToRune(t *testing.T) {
	s := "中华人民共和国"
	for _, c := range s {
		t.Logf("%[1]c %[1]d", c)
	}
}
