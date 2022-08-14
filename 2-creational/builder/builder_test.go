package builder

import "testing"

func TestBuildCar(t *testing.T) {
	director := LandVehicleBuilderDirector{}

	carBuilder := CarBuilder{}

	director.SetBuilder(&carBuilder)
	director.Construct()

	newCar := carBuilder.GetLandVehicle()

	if newCar.Name != "Car" || newCar.Seats != 6 || newCar.Wheels != 4 {
		t.Error("Name or seats or wheels are incorrect for car!")
		return
	}
}

func TestBuildBike(t *testing.T) {
	director := LandVehicleBuilderDirector{}

	bikeBuilder := BikeBuilder{}

	director.SetBuilder(&bikeBuilder)
	director.Construct()

	newBike := bikeBuilder.GetLandVehicle()

	if newBike.Name != "Bike" || newBike.Seats != 1 || newBike.Wheels != 2 {
		t.Error("Name or seats or wheels are incorrect for bike!")
		return
	}
}
