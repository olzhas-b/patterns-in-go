package creational

import (
	"fmt"
	"patterns-in-go/creational-patterns/factory"
	"testing"
)

func TestFactoryPattern(t *testing.T) {
	p1 := factory.NewPerson("Olzhas", 23, factory.Male)
	p2 := factory.NewPerson("Aigerym", 25, factory.Female)

	persons := []factory.Person{p1, p2}
	for _, person := range persons {
		fmt.Println(person.GetAge())
	}
}
