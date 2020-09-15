package main

import "fmt"
//callme function takes no parameter and returns func() which returns integer
func callme() func() int{
	i :=0
	return func() int {
		i++
		return i
	}
}

func factorial(n int )  int  {
	if n == 0 {
		return  1

	}
		return n * factorial(n-1)
}


func main() {

	incr := callme()

	//How to know variable type
	fmt.Printf("%T\n", incr)

	fmt.Println(incr())
	fmt.Println(incr())
	fmt.Println(incr())

	incr2 := callme()
	fmt.Println(incr2())

	//recursion
	fmt.Println(factorial(7))

}

