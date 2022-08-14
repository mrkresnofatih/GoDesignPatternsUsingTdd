package main

import "fmt"

// Decorator is a structural design pattern that lets you attach new behaviors to objects by placing these objects inside special wrapper objects that contain the behaviors.

type IIceCream interface {
	GetIceCream() string
}

type VanillaIceCream struct{}

func (v *VanillaIceCream) GetIceCream() string {
	return "Vanilla IceCream"
}

type ChocolateFrostingDecorator struct {
	IceCream IIceCream
}

func (c *ChocolateFrostingDecorator) GetIceCream() string {
	return c.IceCream.GetIceCream() + " with Chocolate Frosting"
}

type CaramelSauceDecorator struct {
	IceCream IIceCream
}

func (c *CaramelSauceDecorator) GetIceCream() string {
	return c.IceCream.GetIceCream() + " with Caramel Sauce"
}

func main() {
	iceCreamOne := &VanillaIceCream{}
	iceCreamOneWithChocolateFrosting := &ChocolateFrostingDecorator{IceCream: iceCreamOne}
	iceCreamOneWithChocolateFrostingWithCaramelSauce := &CaramelSauceDecorator{IceCream: iceCreamOneWithChocolateFrosting}
	fmt.Println("iceCreamOne: " + iceCreamOneWithChocolateFrostingWithCaramelSauce.GetIceCream())

	iceCreamTwo := &VanillaIceCream{}
	iceCreamTwoWithCaramelFrosting := &ChocolateFrostingDecorator{IceCream: iceCreamTwo}
	fmt.Println("iceCreamTwo: " + iceCreamTwoWithCaramelFrosting.GetIceCream())
}
