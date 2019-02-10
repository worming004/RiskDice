package logic

type byReverseVal []int8

func (s byReverseVal) Len() int {
	return len(s)
}
func (s byReverseVal) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byReverseVal) Less(i, j int) bool {
	return s[i] > s[j]
}
