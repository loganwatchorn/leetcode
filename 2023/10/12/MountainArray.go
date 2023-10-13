package main

type MountainArray struct {
	arr []int
	callsToGet int
}

func (m *MountainArray) get(i int) int {
	m.callsToGet++
	return m.arr[i]
}
func (m *MountainArray) length() int {
	return len(m.arr)
}
