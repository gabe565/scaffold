package module

import (
	"os"
	"path/filepath"
	"testing"
)

func TestModule_WriteAnswer(t *testing.T) {
	t.Parallel()

	tests := []bool{true, false}

	module := Module{}

	for _, test := range tests {
		if err := module.WriteAnswer("", test); err != nil {
			t.Error(err)
		}
		if module.Enabled != test {
			t.Errorf(`invalid value for enabled. expected %t, got %t"`, test, module.Enabled)
		}
	}
}

func TestRunAction_Activate(t *testing.T) {
	t.Parallel()

	action := RunAction{"echo", "success"}
	if err := action.Activate(); err != nil {
		t.Error(err)
	}
}

func TestCopyAction_Activate(t *testing.T) {
	t.Parallel()

	tmpDir := os.TempDir()

	// Create a temp file
	f, err := os.CreateTemp(tmpDir, "copy_action_test")
	if err != nil {
		t.Error(err)
	}
	// Remove temp file on return
	defer func(name string) {
		_ = os.Remove(name)
	}(f.Name())

	// Get temp file info
	info, err := f.Stat()
	if err != nil {
		t.Error(err)
	}

	// Close file before attempting copy
	_ = f.Close()

	srcPath := filepath.Join(tmpDir, info.Name())
	dstPath := srcPath + "_dst"

	// Do the copy
	action := CopyAction{
		Src: srcPath,
		Dst: dstPath,
	}
	if err = action.Activate(); err != nil {
		t.Error(err)
	}
	// Remove copy on return
	defer func(name string) {
		_ = os.Remove(name)
	}(dstPath)

	// Ensure temp file exists
	if _, err = os.Stat(dstPath); err != nil {
		t.Error(err)
	}
}
