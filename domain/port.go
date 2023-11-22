package domain

type ID string

type Name string

type Coordinates []float64

type Location struct {
	City     string
	Province string
	Country  string
}

type Aliases []string

type Regions []string

type Timezone string

type Code string

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
