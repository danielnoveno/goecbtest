// +build linux

package gpio

import (
	"fmt"
	"periph.io/x/periph/conn/gpio"
	"periph.io/x/periph/conn/gpio/gpioreg"
	"periph.io/x/periph/host"
)

type RealPin struct {
	PinName string
}

func NewRealPin(pinName string) *RealPin {
	host.Init()
	return &RealPin{PinName: pinName}
}

func (p *RealPin) Read() bool {
	pin := gpioreg.ByName(p.PinName)
	if pin == nil {
		fmt.Println("Pin not found:", p.PinName)
		return false
	}
	return pin.Read() == gpio.High
}

func (p *RealPin) Write(value bool) error {
	pin := gpioreg.ByName(p.PinName)
	if pin == nil {
		return fmt.Errorf("pin not found: %s", p.PinName)
	}
	if value {
		pin.Out(gpio.High)
	} else {
		pin.Out(gpio.Low)
	}
	return nil
}
