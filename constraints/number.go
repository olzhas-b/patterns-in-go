package constraints

type Number interface {
	int32 | int64 | float64 | float32 | int8 | byte
}
