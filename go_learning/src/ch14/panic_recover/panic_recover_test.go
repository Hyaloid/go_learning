package panic_recover_test

import (
	"fmt"
	"os"
	"testing"
)

func TestPanicVxExit(t *testing.T) {
	//defer func() {
	//	fmt.Println("Finally!")
	//}()
	//defer func() {
	//	if err := recover(); err != nil {
	//		fmt.Println("recover from ", err)
	//	}
	//}()
	fmt.Println("Start")
	//panic(errors.New("Something wrong!"))
	os.Exit(-1)
	fmt.Println("End")
}

