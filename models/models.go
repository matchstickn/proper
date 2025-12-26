package models

type Bind struct {
	Keybind string
	Output  []string
}

type Macro struct {
	Binds     []Bind
	Activated bool
	Name      string
}

func (m *Macro) toggleActivated() {
	if m.Activated {
		m.Activated = false
	} else {
		m.Activated = true
	}
}
