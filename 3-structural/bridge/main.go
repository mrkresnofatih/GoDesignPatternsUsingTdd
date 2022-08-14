package main

import "fmt"

// Bridge Example
// 1. a car provides the user with one api for "moving the car" (gas pedal, interface for the user)
// 2. a car should not have to care about what brand of wheels
//    it's got to move the car, so long as it implements a standard api to communicate with the car (spin wheels, interface for the car)
// 3. a wheel should not care about what kind of car it is to start spinning

// Interfaces

type ICar interface {
	Gas()
	SetWheels(IWheel)
}

type IWheel interface {
	Spin()
}

// Cars

type ToyotaCar struct {
	Wheel IWheel
}

func (car *ToyotaCar) Gas() {
	fmt.Println("Toyota Car attempting: Gas")
	car.Wheel.Spin()
}

func (car *ToyotaCar) SetWheels(wheels IWheel) {
	car.Wheel = wheels
}

type TeslaCar struct {
	Wheel IWheel
}

func (car *TeslaCar) Gas() {
	fmt.Println("Tesla Car attempting: Gas")
	car.Wheel.Spin()
}

func (car *TeslaCar) SetWheels(wheels IWheel) {
	car.Wheel = wheels
}

// Wheels

type PremiumWheel struct{}

func (wheel *PremiumWheel) Spin() {
	fmt.Println("Premium wheel spinning!")
}

type StandardWheel struct{}

func (wheel *StandardWheel) Spin() {
	fmt.Println("Standard wheel spinning!")
}

func main() {
	toyotaCar := &ToyotaCar{}
	premiumWheels := &PremiumWheel{}
	toyotaCar.SetWheels(premiumWheels)
	toyotaCar.Gas()

	// toyotacar wants to downgrade its wheels
	standardWheels := &StandardWheel{}
	toyotaCar.SetWheels(standardWheels)
	toyotaCar.Gas()

	newStandardWheels := &StandardWheel{}
	teslaCar := &TeslaCar{}
	teslaCar.SetWheels(newStandardWheels)
	teslaCar.Gas()
}
