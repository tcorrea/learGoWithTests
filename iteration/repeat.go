package iteration

import (
	"strings"
)

func Repeat(character string, times int) string {
	// var repeated string
	//
	// for i := 0; i < times; i++ {
	// 	repeated += character
	// }
	//
	// return repeated
	return strings.Repeat(character, times)
}
