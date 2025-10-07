package gpio

type MockPin struct {
	State bool
	Name  string
	Logs  chan string
}

func NewMockPin(name string, logs chan string) *MockPin {
	return &MockPin{Name: name, Logs: logs}
}

func (p *MockPin) Read() bool {
	return p.State
}

func (p *MockPin) Write(value bool) error {
	p.State = value
	if p.Logs != nil {
		p.Logs <- "[MOCK] " + p.Name + " set to " + boolToText(value)
	}
	return nil
}

func boolToText(v bool) string {
	if v {
		return "ON"
	}
	return "OFF"
}
