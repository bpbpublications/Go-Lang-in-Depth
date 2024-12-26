package main

import (
	"fmt"
)

const (
	CivilType = 1
)

const (
	RuggedModel   = 1
	StandardModel = 2
)

type Truck interface {
	NumOfTyres() int
	GetModelName() string
}

type TruckFactory interface {
	Create(t int) (Truck, error)
}

type CivilTruckFactory struct{}

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

func (i *CivilTruckFactory) Create(t int) (Truck, error) {
	switch t {
	case RuggedModel:
		return new(RuggedModelType), nil
	case StandardModel:
		return new(StandardModelType), nil
	}
	return nil, fmt.Errorf("There are no trucks of type %d\n", t)
}

func CreateAbstractFactory(c int) (TruckFactory, error) {
	switch c {
	case CivilType:
		return new(CivilTruckFactory), nil
	default:
		return nil, fmt.Errorf("There is no factory with id %d\n", c)
	}
}

func main() {

	civilT, _ := CreateAbstractFactory(CivilType)
	truckM, _ := civilT.Create(StandardModel)

	truck, ok := truckM.(Truck)
	if !ok {
		fmt.Errorf("Invalid model")
	}

	fmt.Printf("%v Truck has %d tyres\n", truck.GetModelName(), truck.NumOfTyres())

}
