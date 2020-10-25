package main

import (
	"fmt"
	"os"
)

//func test() []func() {
//	var s []func()
//
//	for i := 0; i < 2; i++ {
//		x := i
//		s = append(s, func() {
//			println(&x, x)
//		})
//	}
//	return s
//}
//
//func main() {
//	for _, f := range test() {
//		f()
//	}
//}

func main() {
	//var m sync.Mutex
	//m.Lock(); {
	//	m.Lock()
	//	m.Unlock()
	//}
	//m.Unlock()

	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = "*"
	}
	fmt.Println(s)
}
