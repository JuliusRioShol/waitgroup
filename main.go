// Filename: main.go
// Topic : Wait group
package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	//set the flags:
	msg := flag.String("msg", "Howdy Stranger!", "The message to display")
	num := flag.Int("num", 1, "How many times to print the message?")

	// Print all in Caps:
	caps := flag.Bool("caps", false, "should the string be in all uppercase?")
	flag.Parse()
	// Print in all Caps:

	// Make the message print in uppercase:

	//check to see if it should be in upper case
	// if *caps {
	// 	for i := 0; i < *num; i++ {
	// 		fmt.Println(strings.ToUpper(*msg))
	// 	}
	// } else if !(*caps) {
	// 	for i := 0; i < *num; i++ {
	// 		fmt.Println((*msg))
	// 	}
	// }

	// The beter approach:
	if *caps {
		*msg = strings.ToUpper(*msg)
	}
	for i := 0; i < *num; i++ {
		fmt.Println(*msg)
	}
}
