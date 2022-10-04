package abstract_factory

import "patterns-in-go/creational-patterns/abstract-factory/interfaces"

func NewRocketFactory(companyName string) interfaces.IRocketFactory {
	switch companyName {
	case NasaFactory:
		return NewNasa()
	case SpaceXFactory:
		return NewSpaceX()
	case BlueOriginFactory:
		return NewBlueOrigin()
	}
	return nil
}
