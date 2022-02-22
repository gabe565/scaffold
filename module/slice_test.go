package module

import (
	"sort"
	"testing"
)

func TestSlice_Sort(t *testing.T) {
	t.Parallel()

	slice := Slice{
		&Module{Name: "third"},
		&Module{Name: "second", Priority: 5},
		&Module{Name: "first", Priority: 10},
	}
	sort.Sort(slice)
	if slice[0].Name != "first" {
		t.Errorf(`slice did not sort correctly. Expected first name to be "first", got "%s"`, slice[0].Name)
	}
	if slice[1].Name != "second" {
		t.Errorf(`slice did not sort correctly. Expected second name to be "second", got "%s"`, slice[1].Name)
	}
}
