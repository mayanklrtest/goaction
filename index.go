// Package goaction implements methods for returning str.
//
// Experience has shown that a small number of procedures can prove
// helpful when attempting to save the world.
package goaction

import "fmt"

//ReturnStr - Rturn Str
func ReturnStr(str string) string {
	fmt.Println("Print", str)
	return str
}
