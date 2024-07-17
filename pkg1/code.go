package pkg1

import "bs.com/pkg2"

type Foo struct {
  data string
}

func (Foo) DoSomething() {
  print("boo foo")
  b := pkg2.Bar{}
  b.DoSomething()
}
