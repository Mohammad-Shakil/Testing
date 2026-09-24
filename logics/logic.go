package logics

func Reversal(P string) string {

	runes := []rune(P)
	low := 0
	high := len(runes) - 1

	for low < high {
		runes[low], runes[high] = runes[high], runes[low]
		low++
		high--
	}
	return string(runes)
}
