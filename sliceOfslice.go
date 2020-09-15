package main

import (
	"fmt"
	"strings"
)

//func tictacto() {

//}

func main() {
	arr := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	arr[0][0] = "X"
	arr[1][1] = "X"
	arr[2][2] = "X"
	arr[0][2] = "0"
	arr[2][0] = "0"

	for i := 0; i < len(arr); i++ {
		fmt.Printf("%s\n", strings.Join(arr[i], " "))

	}

}
