package module

import (
	"github.com/clevyr/scaffold/iexec"
)

type Module struct {
	Name    string         `json:"name,omitempty"`
	Dev     bool           `json:"dev,omitempty"`
	Enabled bool           `json:"enabled,omitempty"`
	Hidden  bool           `json:"hidden,omitempty"`
	Version string         `json:"version,omitempty"`
	Then    []ActionsUnion `json:"then,omitempty"`
}

func (module *Module) WriteAnswer(name string, value interface{}) error {
	module.Enabled = value.(bool)
	return nil
}

// ActionsUnion Union class to hold different actions
type ActionsUnion struct {
	Run *RunAction `json:"run,omitempty"`
}

func (then ActionsUnion) Activate() (err error) {
	if then.Run != nil {
		err = then.Run.Activate()
		if err != nil {
			return err
		}
	}
	return nil
}

// RunAction Runs an action when activated
type RunAction []string

func (run RunAction) Activate() error {
	return iexec.Command(run[0], run[1:]...)
}
