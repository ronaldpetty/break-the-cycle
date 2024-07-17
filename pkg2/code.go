package pkg2

import "bs.com/pkg3"

type Bar struct {
  data string
  other pkg3.Griffen
}

func (b Bar) DoSomething() {
  print("boo bar", b.data)
  
  b.other.DoSomething()
}

func New(otherThing pkg3.Griffen) pkg3.Griffen {
  return &Bar{"pkg2 data", otherThing}
}
