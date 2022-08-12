package builder

import "fmt"

// A Builder design pattern tries to:
// Abstract complex creations so that object creation is separated from the object user
// Create an object step by step filling its fields
// Reuse the object creation algorithm between many objects

type ILandVehicleBuilder interface {
	SetWheels() ILandVehicleBuilder
	SetSeats() ILandVehicleBuilder
	SetName() ILandVehicleBuilder
	GetLandVehicle() LandVehicle
}

type LandVehicle struct {
	Wheels int
	Seats  int
	Name   string
}

type LandVehicleBuilderDirector struct {
	builder ILandVehicleBuilder
}

func (l *LandVehicleBuilderDirector) SetBuilder(b ILandVehicleBuilder) {
	l.builder = b
}

func (l *LandVehicleBuilderDirector) Construct() {
	l.builder.SetWheels().SetSeats().SetName()
}

type CarBuilder struct {
	Car LandVehicle
}

func (c *CarBuilder) SetWheels() ILandVehicleBuilder {
	c.Car.Wheels = 4
	return c
}

func (c *CarBuilder) SetSeats() ILandVehicleBuilder {
	c.Car.Seats = 6
	return c
}

func (c *CarBuilder) SetName() ILandVehicleBuilder {
	c.Car.Name = "Car"
	return c
}

func (c *CarBuilder) GetLandVehicle() LandVehicle {
	return c.Car
}

type BikeBuilder struct {
	Bike LandVehicle
}

func (b *BikeBuilder) SetWheels() ILandVehicleBuilder {
	b.Bike.Wheels = 2
	return b
}

func (b *BikeBuilder) SetSeats() ILandVehicleBuilder {
	b.Bike.Seats = 1
	return b
}

func (b *BikeBuilder) SetName() ILandVehicleBuilder {
	b.Bike.Name = "Bike"
	return b
}

func (b *BikeBuilder) GetLandVehicle() LandVehicle {
	return b.Bike
}

func main() {
	director := LandVehicleBuilderDirector{}

	carBuilder := CarBuilder{}

	director.SetBuilder(&carBuilder)
	director.Construct()

	newCar := carBuilder.GetLandVehicle()

	bikeBuilder := BikeBuilder{}

	director.SetBuilder(&bikeBuilder)
	director.Construct()

	newBike := bikeBuilder.GetLandVehicle()

	fmt.Println(newCar.Name)
	fmt.Println(newCar.Seats)
	fmt.Println(newCar.Wheels)

	fmt.Println(newBike.Name)
	fmt.Println(newBike.Seats)
	fmt.Println(newBike.Wheels)
}
