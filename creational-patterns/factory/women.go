package factory

type Women struct {
	Name string
	Age  int64
}

func NewWomen(name string, age int64) *Women {
	return &Women{
		Name: name,
		Age:  age,
	}
}

func (w Women) GetName() string {
	return w.Name
}

func (w Women) GetAge() int64 {
	return w.Age
}
