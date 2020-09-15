package main

import "fmt"

type mytype struct{
	pi float32
	num int
}



func main() {


	mm := make( map[string]mytype)

	mm["heman"] = mytype{
		3.14 ,20 }
	mm["himan"] = mytype{
		3.44 ,30 }

	for k,v :=  range mm {
		fmt.Printf("Key= %s -Value1= %f -Value2=%d \n",k,v.pi,v.num)
	}

}
