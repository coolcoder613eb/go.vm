package main

import (
	"C"
	"fmt"

	"github.com/coolcoder613eb/go.vm/cpu"
)

// Embedding
//
//export RunCode
func RunCode(code []byte) (success bool) {
	c := cpu.NewCPU()
	c.LoadBytes(code)
	err := c.Run()
	if err != nil {
		fmt.Printf("Error running code: %s\n", err)
		return false
	}
	return true
}

func main() {

}
