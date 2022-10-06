package structural

import (
	"patterns-in-go/structural-patterns/facade"
	"testing"
)

func Test_SteamFacade(t *testing.T) {
	steamFacade := facade.NewSteamFacade(facade.NewCreditCard())
	tests := []struct {
		game          facade.Game
		expected      error
		expectedError bool
	}{
		{
			game: facade.Game{
				Name: "GTA",
				Cost: 5,
				ID:   123,
			},
			expected:      nil,
			expectedError: false,
		},
		{
			game: facade.Game{
				Name: "GTA",
				Cost: 5,
				ID:   123,
			},
			expected:      nil,
			expectedError: true,
		},
		{
			game: facade.Game{
				Name: "Call out Duty",
				Cost: 51,
				ID:   12,
			},
			expected:      nil,
			expectedError: false,
		},
	}
	for _, test := range tests {
		err := steamFacade.BuyGame(test.game)
		if test.expectedError && err == nil {
			t.Fatalf("expected error but got nil")
		} else if !test.expectedError && err != nil {
			t.Fatalf("got error %v", err)
		}
	}
}
