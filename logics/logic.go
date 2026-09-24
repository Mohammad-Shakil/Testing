package logics

func Elcheck(p []string, minlength int) []string {
	var newlist []string
	for i := 0; i < len(p); i++ {
		if len(p[i]) > minlength {
			newlist = append(newlist, p[i])
		}
	}
	return newlist
}
