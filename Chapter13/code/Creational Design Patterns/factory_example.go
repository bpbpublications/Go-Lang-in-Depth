package main

import (
	"fmt"
)

const (
	RuggedModel   = 1
	StandardModel = 2
)

type Truck interface {
	NumOfTyres() int
	GetModelName() string
}

func GetTruck(truck int) (Truck, error) {
	switch truck {
	case RuggedModel:
		return new(RuggedModelType), nil
	case StandardModel:
		return new(StandardModelType), nil
	default:
		return nil, fmt.Errorf("Not a Known Truck Model")
	}
	return nil, fmt.Errorf("Not implemented yet")
}

type RuggedModelType struct{}

func (f *RuggedModelType) NumOfTyres() int {
	return 6
}

func (f *RuggedModelType) GetModelName() string {
	return "Rugged"
}

type StandardModelType struct {
}

func (f *StandardModelType) NumOfTyres() int {
	return 6
}

func (f *StandardModelType) GetModelName() string {
	return "Standard"
}
func main() {

	truck, err := GetTruck(RuggedModel)
	if err != nil {
		fmt.Errorf("This model does not exist")
	}

	fmt.Printf("%v Truck has %d tyres\n", truck.GetModelName(), truck.NumOfTyres())

}
