package main

import "fmt"

type IHasTrainWidth interface {
	GetTrainWidth() int
}

type Railroad struct {
	Width int
}

func (r *Railroad) IsCorrectWidthForTrain(t IHasTrainWidth) bool {
	return t.GetTrainWidth() == r.Width
}

type Train struct {
	TrainWidth int
}

func (t *Train) GetTrainWidth() int {
	return t.TrainWidth
}

func main() {
	railroad := Railroad{Width: 10}

	passengerTrain := Train{TrainWidth: 10}
	bulletTrain := Train{TrainWidth: 12}

	canPassengerTrainPass := railroad.IsCorrectWidthForTrain(&passengerTrain)
	canBulletTrainPass := railroad.IsCorrectWidthForTrain(&bulletTrain)

	fmt.Println(canPassengerTrainPass)
	fmt.Println(canBulletTrainPass)
}
