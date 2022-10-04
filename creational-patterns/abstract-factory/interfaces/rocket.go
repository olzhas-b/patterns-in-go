package interfaces

import "time"

type IRocket interface {
	GetName() string
	GetCreatedDate() time.Time
	SetCost(cost float64)
	GetCost() float64
}
