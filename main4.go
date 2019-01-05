package main

import (
	"fmt"
)

func main3() {
	fmt.Println("Boom")
	x := "superfox"
	trySwitch(x)
}

func trySwitch(input string) {
	switch {
	case (input == "testing"), (len(input) == 8):
		fmt.Println("Hey this is less than 2")
	case (len(input) < 5):
		fmt.Println("Negative value detected")
	case (len(input) > 2):
		fmt.Println("hey this is larger than 2")
		fallthrough
	case (input == input):
		fmt.Println("Just say true")
	default:
		fmt.Println(`I don't know what you mean`)
	}
}
