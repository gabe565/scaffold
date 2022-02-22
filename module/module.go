package module

import (
	"github.com/clevyr/scaffold/iexec"
	"path/filepath"
)

type Module struct {
	Name     string         `json:"-"`
	Dev      bool           `json:"dev,omitempty"`
	Enabled  bool           `json:"enabled,omitempty"`
	Hidden   bool           `json:"hidden,omitempty"`
	Version  string         `json:"version,omitempty"`
	Priority int8           `json:"priority,omitempty"`
	Then     []ActionsUnion `json:"then,omitempty"`
}

func (module *Module) WriteAnswer(name string, value interface{}) error {
	module.Enabled = value.(bool)
	return nil
}

// ActionsUnion Union class to hold different actions
type ActionsUnion struct {
	Copy *CopyAction `json:"copy,omitempty"`
	Run  *RunAction  `json:"run,omitempty"`
}

func (then ActionsUnion) Activate() (err error) {
	if then.Copy != nil {
		err = then.Copy.Activate()
		if err != nil {
			return err
		}
	}
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

// CopyAction Copies a file or directory when activated
type CopyAction struct {
	Src string `json:"src"`
	Dst string `json:"dst"`
}

func (copy CopyAction) Activate() error {
	dir := filepath.Dir(copy.Dst)
	if dir != "." {
		err := iexec.Command("mkdir", "-p", filepath.Dir(copy.Dst))
		if err != nil {
			return err
		}
	}
	return iexec.Command("cp", "-a", copy.Src, copy.Dst)
}
