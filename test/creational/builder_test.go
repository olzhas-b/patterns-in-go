package creational

import (
	"patterns-in-go/creational-patterns/builder"
	"patterns-in-go/creational-patterns/factory"
	"testing"
)

func TestBuilderPattern(t *testing.T) {
	expected := "SELECT age, name, gender FROM  person WHERE (age < 18 and name = 'olzhas' and gender = 'Male') ;"

	got := builder.New().
		Table("person").
		Select("age", "name", "gender").
		Where("age < %d", 18).
		Where("name = '%s'", "olzhas").
		Where("gender = '%s'", factory.Male).
		BuildSQL()

	if got != expected {
		t.Fatalf("TestBuilderPattern got %s but expected %s", got, expected)
	}
}
