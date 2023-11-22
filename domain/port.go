package domain

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var (
	idRegex       = regexp.MustCompile("^[A-Z0-9]{5}$")
	timezoneRegex = regexp.MustCompile("^[A-Za-z_-]+/[A-Za-z_-]+$")
)

type ID string

func NewID(id string) (*ID, error) {
	if !idRegex.Match([]byte(id)) {
		return nil, Error(ErrInvalidModel, fmt.Sprintf("id %q does not match regex %q", id, idRegex))
	}

	newID := ID(id)
	return &newID, nil
}

type Name string

func NewName(name string) (*Name, error) {
	if len(strings.TrimSpace(name)) == 0 {
		return nil, Error(ErrInvalidModel, "name cannot be empty")
	}

	newName := Name(name)
	return &newName, nil
}

type Coordinates []float64

func NewCoordinates(coordinates []float64) (*Coordinates, error) {
	if len(coordinates) != 2 {
		return nil, Error(ErrInvalidModel, "coordinates length must be two")
	}

	newCoordinates := Coordinates(coordinates)
	return &newCoordinates, nil
}

type Location struct {
	City     string
	Province string
	Country  string
}

func NewLocation(city, province, country string) (*Location, error) {
	if len(strings.TrimSpace(city)) == 0 {
		return nil, Error(ErrInvalidModel, "city cannot be empty")
	}

	if len(strings.TrimSpace(province)) == 0 {
		return nil, Error(ErrInvalidModel, "province cannot be empty")
	}

	if len(strings.TrimSpace(country)) == 0 {
		return nil, Error(ErrInvalidModel, "country cannot be empty")
	}

	return &Location{
		City:     city,
		Province: province,
		Country:  country,
	}, nil
}

type Aliases []string

func NewAliases(aliases []string) (*Aliases, error) {
	if len(aliases) > 0 {
		for _, alias := range aliases {
			if len(strings.TrimSpace(alias)) == 0 {
				return nil, Error(ErrInvalidModel, "alias from aliases cannot be empty")
			}
		}
	}

	newAliases := Aliases(aliases)
	return &newAliases, nil
}

type Regions []string

func NewRegions(regions []string) (*Regions, error) {
	if len(regions) > 0 {
		for _, region := range regions {
			if len(strings.TrimSpace(region)) == 0 {
				return nil, Error(ErrInvalidModel, "region from regions cannot be empty")
			}
		}
	}

	newRegions := Regions(regions)
	return &newRegions, nil
}

type Timezone string

func NewTimezone(timezone string) (*Timezone, error) {
	if !timezoneRegex.Match([]byte(timezone)) {
		return nil, Error(ErrInvalidModel, fmt.Sprintf("timezone %q does not match regex %q", timezone, timezoneRegex))
	}

	newTimezone := Timezone(timezone)
	return &newTimezone, nil
}

type Code string

func NewCode(code string) (*Code, error) {
	if _, err := strconv.Atoi(code); err != nil {
		return nil, Error(ErrInvalidModel, fmt.Sprintf("code %q does not contain a valid number", code))
	}

	newCode := Code(code)
	return &newCode, nil
}

type Port struct {
	ID          ID
	Name        Name
	Coordinates Coordinates
	Location    Location
	Alias       Aliases
	Regions     Regions
	Timezone    Timezone
	Unlocs      []ID
	Code        Code
}

type Repository interface {
	Save(Port) error
}
