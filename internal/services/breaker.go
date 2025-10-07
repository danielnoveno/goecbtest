package services

import (
	"fmt"
	"time"

	"ecb-system/internal/gpio"
)

type BreakerService struct {
	Sensor gpio.Pin
	Relay  gpio.Pin
}

func NewBreaker(sensor gpio.Pin, relay gpio.Pin) *BreakerService {
	return &BreakerService{
		Sensor: sensor,
		Relay:  relay,
	}
}

func (b *BreakerService) Monitor() {
	for {
		if b.Sensor.Read() {
			fmt.Println("[ALERT] Leakage detected! Cutting off power...")
			b.Relay.Write(false) // matikan relay
		} else {
			b.Relay.Write(true) // hidupkan relay
		}
		time.Sleep(1 * time.Second)
	}
}
