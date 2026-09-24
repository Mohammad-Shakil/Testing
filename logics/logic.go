package logics

func Elcheck(s []string, minlength int) []string {
	var newlist []string
	for _, p := range s {
		if len(p) > minlength {
			newlist = append(newlist, p)
		}
	}
	return newlist
}
