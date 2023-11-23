package domain_test

import (
	"reflect"
	"testing"

	"github.com/acerquetti/ports/domain"
)

func TestNewPortFromRaw(t *testing.T) {
	id, err := domain.NewID("ABC12")
	mustNotError(t, err)

	name, err := domain.NewName("hello")
	mustNotError(t, err)

	coords, err := domain.NewCoordinates([]float64{1, 1})
	mustNotError(t, err)

	location, err := domain.NewLocation("Madrid", "Madrid", "Spain")
	mustNotError(t, err)

	aliases, err := domain.NewAliases([]string{})
	mustNotError(t, err)

	regions, err := domain.NewRegions([]string{})
	mustNotError(t, err)

	timezone, err := domain.NewTimezone("first/second")
	mustNotError(t, err)

	code, err := domain.NewCode("2387123")
	mustNotError(t, err)

	unlocs := []domain.ID{*id}

	tests := map[string]struct {
		input  domain.PortRaw
		output *domain.Port
		err    error
	}{
		"Should successfully build a Port": {
			input: domain.PortRaw{
				ID:          "ABC12",
				Name:        "hello",
				Coordinates: []float64{1, 1},
				City:        "Madrid",
				Province:    "Madrid",
				Country:     "Spain",
				Aliases:     []string{},
				Regions:     []string{},
				Timezone:    "first/second",
				Unlocs:      []string{"ABC56"},
				Code:        "1232193",
			},
			output: &domain.Port{
				ID:          *id,
				Name:        *name,
				Coordinates: *coords,
				Location:    *location,
				Aliases:     *aliases,
				Regions:     *regions,
				Timezone:    *timezone,
				Unlocs:      unlocs,
				Code:        *code,
			},
		},
	}

	for tname, test := range tests {
		gotPort, err := domain.NewPortFromRaw(test.input)
		if err != nil && test.err == nil {
			t.Error("unexpected error:", err)
		} else if err == nil && test.err != nil {
			t.Errorf("expected error %v but was nil", test.err)
		}

		if reflect.DeepEqual(test.output, gotPort) {
			t.Errorf("test %q: exp: %v, got: %v", tname, test.output, gotPort)
		}
	}
}

func mustNotError(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}
