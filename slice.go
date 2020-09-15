package main

import "fmt"

func main() {

	s := make([]string, 3)
	//fmt.Println("emp",s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	s = append(s, "d")
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println(c)
	decSlice := []string{"e","f","g"}
	fmt.Println(decSlice)
	decArr := [5]string{"A", "R", "R", "A", "Y"}

	fmt.Println("Values in Array ==>")
	for _,a := range decArr {

		fmt.Printf(" %s ",a)
	}


	}
