package demo

// Demo服务的key
const Key = "hade:demo"

// Demo服务的接口
type Service interface {
	GetFoo() Foo
}

type Foo struct {
	Name string
}
