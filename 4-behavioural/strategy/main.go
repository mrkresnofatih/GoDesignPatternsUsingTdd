package main

import "fmt"

type ISendPackageStrategy interface {
	SendPackage(string, string)
}

type IPackageSenderClient interface {
	SetSendPackageStrategy(ISendPackageStrategy)
	ExecuteSendPackage()
}

type AmazonPackageSenderClient struct {
	SendPackageStrategy ISendPackageStrategy
	PackageToSend       string
	Destination         string
}

func (f *AmazonPackageSenderClient) SetSendPackageStrategy(strategy ISendPackageStrategy) {
	f.SendPackageStrategy = strategy
}

func (f *AmazonPackageSenderClient) SetPackageToSend(pkg string) {
	f.PackageToSend = pkg
}

func (f *AmazonPackageSenderClient) SetDestination(destination string) {
	f.Destination = destination
}

func (f *AmazonPackageSenderClient) ExecuteSendPackage() {
	f.SendPackageStrategy.SendPackage(f.PackageToSend, f.Destination)
}

type PlanePackageSenderStrategy struct{}

func (f *PlanePackageSenderStrategy) SendPackage(packageToSend string, destination string) {
	fmt.Printf("Sending %s by plane strategy to %s\n", packageToSend, destination)
}

type ShipPackageSenderStrategy struct{}

func (f *ShipPackageSenderStrategy) SendPackage(packageToSend string, destination string) {
	fmt.Printf("Sending %s by ship strategyto %s\n", packageToSend, destination)
}

func main() {
	amazonSenderClient := &AmazonPackageSenderClient{}

	shoes := "Shoes"
	addressOne := "Qatar"
	planeStrategy := &PlanePackageSenderStrategy{}
	amazonSenderClient.SetDestination(addressOne)
	amazonSenderClient.SetSendPackageStrategy(planeStrategy)
	amazonSenderClient.SetPackageToSend(shoes)
	amazonSenderClient.ExecuteSendPackage()

	books := "Books"
	addressTwo := "Singapore"
	shipStrategy := &ShipPackageSenderStrategy{}
	amazonSenderClient.SetDestination(addressTwo)
	amazonSenderClient.SetSendPackageStrategy(shipStrategy)
	amazonSenderClient.SetPackageToSend(books)
	amazonSenderClient.ExecuteSendPackage()

	// OUTPUT
	// ======
	// Sending Shoes by plane strategy
	// Sending Books by ship strategy

	// EXPLANATION
	// ===========
	// A client of an application should not care how a functionality works, better off keep those under the hood
	// A client should provide flexibility of changing the strategy of a certain task, this is done by separating the strategies into separate objects
	// Very similar to bridge pattern, the client provides APIs for user to set the object & destination to send & choose the type of strategy & tell the strategy to execute
	// Each strategy provides an api for the client to execute/do something i.e send the package
}
