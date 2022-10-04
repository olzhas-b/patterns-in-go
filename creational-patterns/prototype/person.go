package prototype

import "fmt"

type Person struct {
	Name       string
	Age        int
	Nicknames  []string
	Educations []Education
}

func (person *Person) Copy() Person {
	cloneNicknames := make([]string, len(person.Nicknames))
	copy(cloneNicknames, person.Nicknames)

	cloneEducations := make([]Education, len(person.Educations))
	for i := 0; i < len(cloneEducations); i++ {
		cloneEducations[i] = person.Educations[i].Copy()
	}

	newPerson := Person{
		Name:       person.Name,
		Age:        person.Age,
		Nicknames:  cloneNicknames,
		Educations: cloneEducations,
	}
	return newPerson
}

func (person *Person) Print() string {
	return fmt.Sprintf("%+v", person)
}
