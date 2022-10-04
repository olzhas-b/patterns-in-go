package creational

import (
	abstract_factory "patterns-in-go/creational-patterns/abstract-factory"
	"patterns-in-go/creational-patterns/abstract-factory/interfaces"
	"testing"
)

func TestAbstractFactory(t *testing.T) {
	spaceX := abstract_factory.NewRocketFactory(abstract_factory.SpaceXFactory)
	falcon9 := spaceX.MakeRocket()

	blueOrigin := abstract_factory.NewRocketFactory(abstract_factory.BlueOriginFactory)
	glenn := blueOrigin.MakeRocket()

	nasa := abstract_factory.NewRocketFactory(abstract_factory.NasaFactory)
	shuttle := nasa.MakeRocket()

	tests := []struct {
		name string
		interfaces.IRocket
		expected string
	}{
		{
			"testing falcon9",
			falcon9,
			"Falcon9",
		},
		{
			"testing glenn",
			glenn,
			"New Glen",
		},
		{
			"testing shuttle",
			shuttle,
			"Space Shuttle Atlantis",
		},
	}

	for _, test := range tests {
		got := test.GetName()
		t.Run(test.name, func(t *testing.T) {
			if test.expected != got {
				t.Fatalf("exptected %s, but got %s", test.expected, got)
			}
		})
	}

}
