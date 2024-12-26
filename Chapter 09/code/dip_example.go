package main

import "fmt"

type vehicle interface {
	getSpeed() float64
	getDistanceTravelled() float64
	start()
}

type Engine struct {
	engineType string
}

type car struct {
	brand              string
	manufacturing_year int
	engine             *Engine
}

func NewCar(brand string, year int, engine *Engine) *car {
	return &car{
		brand:              brand,
		manufacturing_year: year,
		engine:             engine,
	}
}

func (Car car) getSpeed() float64 {

	return 120
}

func (Car car) getDistanceTravelled() float64 {

	return 4000
}

func (Car car) start() {

	fmt.Printf(">>> starting the [%s] car ... \n", Car.brand)
}

func main() {

	var veh vehicle
	var engine = &Engine{engineType: "V3"}
	veh = NewCar("Toyota Tercel", 1997, engine)
	fmt.Println("Car's Speed of the Vehicle :", veh.getSpeed())
	fmt.Println("Car's Distance Travelled by the Vehicle:", veh.getDistanceTravelled())

	veh.start()

}
