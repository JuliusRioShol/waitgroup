// Filename: main.go
// Topic : Wait group
package main

import (
	"fmt"
	"sync"
)

func one(wg *sync.WaitGroup) {
	defer wg.Done() //Done is decrementing , defer - is allways  the last that get run within a function
	fmt.Println("hola")
}
func two(wg *sync.WaitGroup) {
	defer wg.Done() //Done is decrementing , defer - is allways  the last that get run within a function
	fmt.Println("ni hao")
}

func three(wg *sync.WaitGroup) {
	defer wg.Done() //Done is decrementing , defer - is allways  the last that get run within a function
	fmt.Println("hello")
}
func main() {

	var wg sync.WaitGroup
	wg.Add(3)
	go one(&wg)
	go two(&wg)
	go three(&wg)

	// let's delay
	wg.Wait() //Wait means: wait until wg holds the value 0.
}
