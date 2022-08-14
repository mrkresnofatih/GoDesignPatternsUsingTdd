package main

import "fmt"

type LightningClient struct{}

type Computer interface {
	ReceiveInLightningPort()
}

func (c *LightningClient) InsertLightningConnectorIntoComputer(computer Computer) {
	fmt.Println("Lightning client insert into computer")
	computer.ReceiveInLightningPort()
}

type Mac struct{}

func (m *Mac) ReceiveInLightningPort() {
	fmt.Println("Mac receives in lightning port")
}

type Windows struct{}

func (w *Windows) ReceiveInUsbPort() {
	fmt.Println("Windows receives in usb port")
}

type WindowsAdapter struct {
	windowsMachine *Windows
}

func (w *WindowsAdapter) ReceiveInLightningPort() {
	fmt.Println("Added windows adapter for lightning ports")
	w.windowsMachine.ReceiveInUsbPort()
}

func main() {
	lightningClient := &LightningClient{}

	macComputer := &Mac{}

	lightningClient.InsertLightningConnectorIntoComputer(macComputer)

	windowsComputer := &Windows{}

	windowsComputerAdapter := &WindowsAdapter{windowsMachine: windowsComputer}

	lightningClient.InsertLightningConnectorIntoComputer(windowsComputerAdapter)

	// Lightning client insert into computer
	// Mac receives in lightning port
	// Lightning client insert into computer
	// Added windows adapter for lightning ports
	// Windows receives in usb port
}
