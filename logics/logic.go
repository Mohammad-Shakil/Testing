package logics

func ListCompress(p []int) []int {
	var newlist []int
	for i := 0; i < len(p); i++ {
		if p[i] >= 50 {
			newlist = append(newlist, p[i])
		}
	}
	p = newlist
	return p
}
