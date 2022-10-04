package abstract_factory

import (
	"patterns-in-go/creational-patterns/abstract-factory/interfaces"
	"patterns-in-go/creational-patterns/abstract-factory/model"
)

const BlueOriginFactory = "BlueOriginFactory"

type BlueOrigin struct {
	CompanyName string
}

func NewBlueOrigin() *BlueOrigin {
	return &BlueOrigin{
		CompanyName: "BlueOrigin",
	}
}

func (blueO *BlueOrigin) MakeRocket() interfaces.IRocket {
	return &model.BlueOriginRocket{
		Rocket: model.Rocket{
			Name: "New Glen",
		},
	}
}
