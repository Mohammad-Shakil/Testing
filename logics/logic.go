package logics

func Elcheck(p []string) []string {
	var newlist []string
	for i := 0; i < len(p); i++ {
		if len(p[i]) > 5 {
			newlist = append(newlist, p[i])
		}
	}
	return newlist
}
