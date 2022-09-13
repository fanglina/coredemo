package demo

const Key = "demo"

type IService interface {
	GetALlStudent() []Student
}

type Student struct {
	ID   int
	Name string
}

