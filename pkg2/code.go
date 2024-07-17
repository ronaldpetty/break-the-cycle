package pkg2

import "bs.com/pkg1"

type Bar struct {
  data string
}

func (Bar) DoSomething() {
  print("boo bar")
  
  f := pkg1.Foo{}
  f.DoSomething()
}
