package abstract_factory

import (
	"patterns-in-go/creational-patterns/abstract-factory/interfaces"
	"patterns-in-go/creational-patterns/abstract-factory/model"
)

const SpaceXFactory = "SpaceXFactory"

type SpaceX struct {
	CompanyName string
}

func NewSpaceX() *SpaceX {
	return &SpaceX{
		CompanyName: "SpaceX",
	}
}

func (spx *SpaceX) MakeRocket() interfaces.IRocket {
	return &model.SpaceXRocket{
		Rocket: model.Rocket{
			Name: "Falcon9",
		},
	}
}
