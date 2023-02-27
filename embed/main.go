package main

/*
//#include <stdlib.h>
*/
import (
	"C"
	"fmt"
	"os"

	"github.com/coolcoder613eb/go.vm/cpu"
)

// Embedding
//
//export RunCode
func RunCode(code []byte) (success bool) { //code *C.char
	//fmt.Printf(code)
	//arr := C.GoBytes(unsafe.Pointer(&code), 19)
	arr := code
	//fmt.Println(len(arr))
	// for x := 0; x < len(arr); x++ {
	// 	fmt.Print(arr[x], ",")
	// }
	// fmt.Println("")
	fmt.Println(arr)
	c := cpu.NewCPU()
	oout := os.Stdout
	os.Stdout, _ = os.Open(os.DevNull)
	//print("b")
	c.LoadBytes(arr)
	os.Stdout = oout
	//print("a")
	err := c.Run()
	if err != nil {
		fmt.Printf("Error running code: %s\n", err)
		return false
	}
	return true
}

func main() {

}
