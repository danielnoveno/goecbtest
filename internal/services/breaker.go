package services

import (
	"time"

	"goecbtest/internal/gpio"
)

type BreakerService struct {
	Sensor gpio.Pin
	Relay  gpio.Pin
	Logs   chan string // channel buat kirim pesan ke UI
}

func NewBreaker(sensor gpio.Pin, relay gpio.Pin) *BreakerService {
	return &BreakerService{
		Sensor: sensor,
		Relay:  relay,
		Logs:   make(chan string, 100),
	}
}

func (b *BreakerService) Monitor() {
	for {
		if b.Sensor.Read() {
			b.Logs <- "[ALERT] ⚠️ Leakage detected! Cutting off power..."
			b.Relay.Write(false)
		} else {
			b.Relay.Write(true)
		}
		time.Sleep(1 * time.Second)
	}
}
