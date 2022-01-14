package arraysandslices

// Multiplies all the items contained in the slice
func Product(numbers []int) (result int) {
	result = 1
	for _, number := range numbers {
		result *= number
	}
	return result
}
