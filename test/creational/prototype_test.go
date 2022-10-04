package creational

import (
	"patterns-in-go/creational-patterns/prototype"
	"reflect"
	"testing"
)

func TestPrototypePattern(t *testing.T) {
	person := prototype.Person{
		Name:      "Olzhas",
		Age:       190,
		Nicknames: []string{"Oljik", "heyoung"},
		Educations: []prototype.Education{
			{
				"KBTU bachelor",
				[]string{"OOP", "algorithm and data structure"},
			},
			{
				"KBTU master",
				[]string{"testing and bug fixing", "computer architecture"},
			},
		},
	}
	clonePerson := person.Copy()

	if !reflect.DeepEqual(person, clonePerson) {
		t.Fatalf("copied wrong was")
	}

}
