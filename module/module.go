package module

type Module struct {
	Name            string     `json:",omitempty"`
	Dev             bool       `json:"-"`
	Enabled         bool       `json:",omitempty"`
	Hidden          bool       `json:"-"`
	Version         string     `json:",omitempty"`
	PostInstallCmds [][]string `json:",omitempty"`
}

func (module *Module) WriteAnswer(name string, value interface{}) error {
	module.Enabled = value.(bool)
	return nil
}
