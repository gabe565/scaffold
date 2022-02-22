package module

type Slice []*Module

func (slice Slice) Len() int {
	return len(slice)
}

func (slice Slice) Less(i, j int) bool {
	return slice[j].Priority < slice[i].Priority
}

func (slice Slice) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}
