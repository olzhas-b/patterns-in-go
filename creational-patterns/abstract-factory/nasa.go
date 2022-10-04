package abstract_factory

import (
	"patterns-in-go/creational-patterns/abstract-factory/interfaces"
	"patterns-in-go/creational-patterns/abstract-factory/model"
)

const NasaFactory = "NasaFactory"

type Nasa struct {
	CompanyName string
}

func NewNasa() *Nasa {
	return &Nasa{
		CompanyName: "Nasa",
	}
}

func (nasa *Nasa) MakeRocket() interfaces.IRocket {
	return &model.NasaRocket{
		Rocket: model.Rocket{
			Name: "Space Shuttle Atlantis",
		},
	}
}
