package creational

import (
	"patterns-in-go/tools"
	"reflect"
	"testing"
)

func TestSumGenerics(t *testing.T) {
	input := map[int]float64{
		1: 12.5,
		2: 2.4,
		3: -1.3,
	}
	expected := 13.6
	got := tools.Sum(input)
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("Sum expected %f but got %f", expected, got)
	}
}

func BenchmarkWithoutGenerics(b *testing.B) {
	input := map[int]float64{
		1: 12.5123,
		2: 2.41231,
		3: -1.3123,
		4: -123123123.77,
		5: 1231232.5544,
		6: 123123123.2,
	}
	for i := 0; i < 1; i++ {
		alloc := make([]byte, 100000)
		tools.Sum(input)
		alloc[0] = 0
	}
}

func BenchmarkGenerics(b *testing.B) {
	input := map[int]float64{
		1: 2,
		2: 3,
	}
	for i := 0; i < 1; i++ {
		alloc := make([]byte, 100000)
		tools.SumGenerics(input)
		alloc[0] = 0
	}
}
