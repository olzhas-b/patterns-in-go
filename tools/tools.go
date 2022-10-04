package tools

import "patterns-in-go/constraints"

func Max[T int32 | float64 | float32](a, b T) T {
	if a < b {
		return b
	}
	return a
}

func SumGenerics[K comparable, num constraints.Number](nums map[K]num) num {
	var result num
	for _, value := range nums {
		result += value
	}
	return result
}

func Sum(nums map[int]float64) float64 {
	var result float64
	for _, num := range nums {
		result += num
	}
	return result
}
