
package main

import (

	d "github.com/synesissoftware/diagnosticism.Go"
)


func SomeFunction(x, y int, order string) {
	d.Trace("SomeFunction",
		d.Trarg("x", x),
		d.Trarg("y", y),
		d.Trarg("order", order),
	)
	//. . . impl. of SomeFunc()
}

func main() {

	d.EnableTracing(true)

	SomeFunction(1, 2, "first")
}
