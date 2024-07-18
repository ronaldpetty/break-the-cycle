// AKA package A
package pkg1

import (
	"bs.com/pkg2"
	"bs.com/pkg3"
	"bs.com/pkg4"
)

// get interface instance object from B
func DoSomething() {
	var g pkg4.Griffen
	g = pkg2.New()

	//print(g.Get()) // should be empty

	pkg3.Update(g)

	print(g.Get())
}
