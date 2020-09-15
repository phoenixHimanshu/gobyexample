package main

import (
	"fmt"
	"math"
)

type rect struct {
	width, height float32
}

type cir struct {
	radius float32

}

type generic interface {
	area() float32
	perimeter() float32
}

func measure(g generic) {
	fmt.Printf("Area==>%f\n", g.area())
	fmt.Printf("Perimeter==>%f\n", g.perimeter())
}

func (ptr *rect) area() float32 {
	return ptr.height * ptr.width
}

func (ptr *rect) perimeter() float32 {

	return (ptr.height * 2) + (ptr.width * 2)
}

func (ptr *cir) area() float32 {
	return math.Pi * ptr.radius * ptr.radius
}

func (ptr *cir) perimeter() float32 {

	return 2 * ptr.radius * math.Pi
}

func main() {

	r := rect{width: 5, height: 4}
	c := cir{radius: 5.5}
	measure(&c)
	measure(&r)

	//fmt.Printf("Area ==>%d \n", r.area())
	//fmt.Printf("Perimeter ==>%d \n", r.perimeter())

	//rptr := &r
	//fmt.Printf("Area ==>%d ", rptr.area())
	//fmt.Printf("Perimeter ==>%d", rptr.perimeter())
	//fmt.Println(r.height)
	//fmt.Println(r.width)

}
