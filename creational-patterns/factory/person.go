package factory

type Person interface {
	GetName() string
	GetAge() int64
}

const Female string = "Female"
const Male string = "Male"

// #pattern factory
// pattern factory return interface such as Person
// other struct which returned must be implemented Person functions
func NewPerson(name string, age int64, gender string) Person {
	switch gender {
	case Male:
		return NewMan(name, age)
	case Female:
		return NewWomen(name, age)
	}
	return nil
}
