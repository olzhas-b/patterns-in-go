package builder

import "fmt"

type conditions []condition

type condition string

func NewCondition(cond string) condition {
	return condition(cond)
}

func (c conditions) ToString() string {
	if len(c) == 0 {
		return " "
	}
	result := "WHERE ("
	for i := 0; i < len(c)-1; i++ {
		result += fmt.Sprintf(string(c[i]) + " and ")
	}
	result += string(c[len(c)-1]) + ") "
	return result
}
