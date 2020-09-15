//https://tour.golang.org/methods/12
/*
 If the concrete value inside the interface itself is nil,
the method will be called with a nil receiver.

In some languages this would trigger a null pointer exception,
but in Go it is common to write methods that gracefully handle being called with a
nil receiver (as with the method M in this example.)

Note that an interface value that holds a nil concrete value is itself non-nil.

*/
package main

import "fmt"

type Intf interface {
	M()
}

type Tstrct struct {
	S string
}

func (t *Tstrct) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func main() {
	var i Intf

	var t *Tstrct
	i = t
	described(i)
	i.M()

	i = &Tstrct{"hello"}
	described(i)
	i.M()
}

func described(i Intf) {
	fmt.Printf("(%v, %T)\n", i, i)
}
