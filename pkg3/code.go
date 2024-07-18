// AKA package C
package pkg3

import (
	"bs.com/pkg4"
)

type CThing struct {
	data string
}

func (c CThing) Get() string {
	return c.data
}

func (c CThing) Set(data string) {
	c.data = data
}

func Update(i pkg4.Griffen) {
	c := CThing{"cthing"}
	i.Set("pkg 3 data")
	i.SetMore(c)
}
