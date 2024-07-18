package pkg4

type Griffen interface {
	Get() string
	Set(string)
	SetMore(C)
}

type C interface {
	Get() string
	Set(data string)
}
