package module

type ModuleSlice []*Module

func (slice ModuleSlice) Len() int {
	return len(slice)
}

func (slice ModuleSlice) Less(i, j int) bool {
	return slice[j].Priority < slice[i].Priority
}

func (slice ModuleSlice) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}
