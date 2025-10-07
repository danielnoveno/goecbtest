package gpio

import "fmt"

type MockPin struct {
	State bool
	Name  string
}

func NewMockPin(name string) *MockPin {
	return &MockPin{Name: name}
}

func (p *MockPin) Read() bool {
	return p.State
}

func (p *MockPin) Write(value bool) error {
	p.State = value
	fmt.Printf("[MOCK] %s set to %v\n", p.Name, value)
	return nil
}
