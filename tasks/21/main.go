package main

import "fmt"

type USB_C interface {
	PlugIn()
}

type USB_A interface {
	Plug(device USB_C)
}

type micron_USB_C struct{}

func (m *micron_USB_C) PlugIn() {
	fmt.Println("Device connected")
}

type hp_USB_A struct{}

func (h *hp_USB_A) Plug(device USB_C) {
	fmt.Println("Connected")
	device.PlugIn()
	fmt.Println("Disconnected")
}

type C_to_A struct {
	device USB_C
}

func (c *C_to_A) PlugIn() {
	c.device.PlugIn()
}

func main() {
	device := micron_USB_C{}

	socket := hp_USB_A{}

	adapter := &C_to_A{device: &device}

	socket.Plug(adapter)
}
