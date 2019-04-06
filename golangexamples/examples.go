package golangexamples

import (
	"fmt"

	"github.com/ehteshamz/greetings"
)

//Slice concataion
func ConcatSlice(sliceToConcat []byte) string {
	var str string
	for i := 0; i < len(sliceToConcat)-1; i++ {
		str = str + string(sliceToConcat[i]) + "-"
	}
	str = str + string(sliceToConcat[len(sliceToConcat)-1])
	return str
}

// Encrypt
func Encrypt(sliceToConcat []byte, ceaserCount int) string {

	// var dummy string
	return "HELLO"
}

//	Greetings
func EZGreetings(name string) string {
	fmt.Printf(greetings.PrintGreetings(name))
	return greetings.PrintGreetings(name)

}
