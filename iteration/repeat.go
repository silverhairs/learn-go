package iteration

// Returns a new string with the passed value repeated four times
func Repeat(letter string, count int) (result string) {
	for i := 0; i < count; i++ {
		result += letter
	}
	return result
}
