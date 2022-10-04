package factory

type Man struct {
	name string
	age  int64
}

func NewMan(name string, age int64) *Man {
	return &Man{
		name: name,
		age:  age,
	}
}

func (m Man) GetName() string {
	return m.name
}

func (m Man) GetAge() int64 {
	return m.age
}
