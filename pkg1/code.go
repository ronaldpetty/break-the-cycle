package pkg1

import (
  "bs.com/pkg2"
)

type Foo struct {
  data string
}

func (Foo) DoSomething() {
  print("boo foo")
}

func DoSomething() {
  g := pkg2.New(Foo{"pkg1 data"})
  g.DoSomething() 
}
