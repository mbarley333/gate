package main

import (
	"fmt"
	"gate"
)

func main() {

	inputA := 1
	inputB := 1

	result := gate.AND(inputA, inputB)
	fmt.Println(result)

}
