package main

import (
	"fizz-buzz/comparator"
	"fmt"
)
func main() {
	for i:=1; i <= 100; i++ {
		fmt.Println(comparator.FizzBuzz(i))
	}
}