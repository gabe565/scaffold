package module

import (
	"github.com/clevyr/scaffold/internal/iexec"
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

func (module *Module) WriteAnswer(name string, value any) error {
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
		if err = then.Copy.Activate(); err != nil {
			return err
		}
	}
	if then.Run != nil {
		if err = then.Run.Activate(); err != nil {
			return err
		}
	}
	return nil
}

// RunAction Runs an action when activated
type RunAction []string

func (run RunAction) Activate() error {
	return iexec.NewBuilder(run...).Run()
}

// CopyAction Copies a file or directory when activated
type CopyAction struct {
	Src string `json:"src"`
	Dst string `json:"dst"`
}

func (copy CopyAction) Activate() error {
	dir := filepath.Dir(copy.Dst)
	if dir != "." {
		err := iexec.NewBuilder("mkdir", "-p", filepath.Dir(copy.Dst)).Run()
		if err != nil {
			return err
		}
	}
	return iexec.NewBuilder("cp", "-a", copy.Src, copy.Dst).Run()
}
