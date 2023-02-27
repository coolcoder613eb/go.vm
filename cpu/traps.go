//
// This file contains the callbacks that the virtual machine
// can implement via the `int` operation.
//
//

package cpu

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// TrapFunction is the signature for a function that is available
// as a trap.
type TrapFunction func(c *CPU, num int) error

// TRAPS is an array of our trap-functions.
var TRAPS [0xffff]TrapFunction

//
// Trap Functions now follow
//

// TrapNOP is the default trap-function for any trap IDs that haven't
// explicitly been setup.
func TrapNOP(c *CPU, num int) error {
	return fmt.Errorf("trap function not defined: 0x%04X", num)
}

// StrLenTrap returns the length of a string.
//
// Input:
//
//	The string to measure in register 0.
//
// Output:
//
//	Sets register 0 with the length
func StrLenTrap(c *CPU, num int) error {
	str, err := c.regs[0].GetString()
	if err != nil {
		return err
	}
	c.regs[0].SetInt(len(str))
	return nil
}

// ReadStringTrap reads a string from the console
//
// Input: None
//
// Output:
//
//	Sets register 0 with the user-provided string
func ReadStringTrap(c *CPU, num int) error {
	text, err := c.STDIN.ReadString('\n')
	if err != nil {
		return err
	}
	c.regs[0].SetString(text)
	return nil
}

// LoadProgTrap loads a program into memory
//
// Input: Memory address in register 0.
//
// Output: None
func LoadProgTrap(c *CPU, num int) error {
	path, err := c.regs[0].GetString()
	if err != nil {
		return err
	}
	addr, err := c.regs[1].GetInt()
	if err != nil {
		return err
	}
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return fmt.Errorf("failed to read file: %s - %s", path, err.Error())
	}

	if (len(data) + addr) >= 0xFFFF {
		return fmt.Errorf("program too large for RAM %d", len(data))
	}

	// Copy contents of file to our memory region.
	// NOTE: This calls `Reset` too :)
	for i := 0; i < len(data); i++ {
		// Addition to fix a linter warning suggesting the used
		// of `copy`.
		c.mem[i+0] = data[i+0]
	}
	return nil
}

// RemoveNewLineTrap removes any trailing newline from the string in #0
//
// Input:
//
//	The string operate upon in #0.
//
// Output:
//
//	Sets register #0 with the updated string
func RemoveNewLineTrap(c *CPU, num int) error {
	str, err := c.regs[0].GetString()
	if err != nil {
		return err
	}
	c.regs[0].SetString(strings.TrimSpace(str))
	return nil
}

// init configures our registered traps.
func init() {

	// Default to all traps being "empty", i.e. configured to
	// to hold a reference to a function that just reports an
	// error.
	for i := 0; i < 0xFFFF; i++ {
		TRAPS[i] = TrapNOP
	}

	//
	// Now setup the actual traps we implement.
	//
	TRAPS[0] = StrLenTrap
	TRAPS[1] = ReadStringTrap
	TRAPS[2] = RemoveNewLineTrap
	TRAPS[3] = LoadProgTrap

}
