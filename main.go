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

	//Flag to reverse a string:
	reverse := flag.Bool("r", false, "reverse the string")
	// check if we should reverse the string:
	flag.Parse()
	if *reverse {
		//reverse string
		var result string //1. Create an empty string.

		//2. Iterate over the string that we want to reverse:
		for _, value := range *msg {
			result = string(value) + result
		}
		*msg = result
	}

	// The beter approach:
	if *caps {
		*msg = strings.ToUpper(*msg)
	}
	for i := 0; i < *num; i++ {
		fmt.Println(*msg)
	}
}
