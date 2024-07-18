// AKA package B
package pkg2

import (
	"bs.com/pkg4"
)

type G struct {
	data     string
	moreData pkg4.C
}

func (g G) Get() string {
	return g.data + g.moreData.Get()
}

func (g G) Set(data string) {
	g.data = data
}

func (g *G) SetMore(c pkg4.C) {
	g.moreData = c
}

func New() pkg4.Griffen {
	return &G{}
}
