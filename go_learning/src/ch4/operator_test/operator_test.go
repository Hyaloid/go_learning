package operator_test

import "testing"

func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 3, 4, 5}
	//c := [...]int{1, 2, 3, 4, 5}
	d := [...]int{1, 2, 3, 4}
	t.Log(a==b)
	//a 与 c比较会出现编译错误，因为它们的维度不同
	//t.Log(a==c)
	t.Log(a==d)
}
