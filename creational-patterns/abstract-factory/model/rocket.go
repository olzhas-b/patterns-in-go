package model

import "time"

type Rocket struct {
	Name     string
	maxSpeed int64
	cost     float64
	diameter float64
	created  time.Time
}

func (rocket *Rocket) GetName() string {
	return rocket.Name
}

func (rocket *Rocket) SetCost(cost float64) {
	rocket.cost = cost
}

func (rocket *Rocket) GetCost() float64 {
	return rocket.cost
}

func (rocket *Rocket) GetCreatedDate() time.Time {
	return rocket.created
}
