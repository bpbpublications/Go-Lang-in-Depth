package main

import (
	"fmt"
)

type FlightCloner interface {
	GetClone(s int) (ItemInfoGetter, error)
}

const (
	Boeing     = 1
	Bombardier = 2
	Embraer    = 3
)

func GetFlightsCloner() FlightCloner {
	return new(FlightsCache)
}

type FlightsCache struct{}

func (s *FlightsCache) GetClone(flight int) (ItemInfoGetter, error) {
	switch flight {
	case Boeing:
		newItem := *flightPrototype
		return &newItem, nil
	}
	return nil, fmt.Errorf("Not implemented yet")
}

type ItemInfoGetter interface {
	GetInfo() string
}

type FlightEngine byte

type Flight struct {
	Price  float32
	Model  string
	Engine FlightEngine
}

func (s *Flight) GetInfo() string {
	return "" + s.Model
}

var flightPrototype *Flight = &Flight{
	Price:  15.00,
	Model:  "777",
	Engine: Boeing,
}

func (i *Flight) GetPrice() float32 {
	return i.Price
}
func main() {

	flightCache := GetFlightsCloner()
	if flightCache == nil {
		fmt.Errorf("The current cache is not valid")
	}

	firstFlight, err := flightCache.GetClone(Boeing)

	fmt.Printf("Flight cloned is of Model type : %s\n", firstFlight.GetInfo())
	if err != nil {
		fmt.Println(err)
	}

	if firstFlight == flightPrototype {
		fmt.Errorf("firstitem cannot be equal to the white prototype")
	}

}
