package arraysandslices

// Sums up all the items contained in the array then returns the result
func Sum(numbers []int) (result int) {
	for _, number := range numbers {
		result += number
	}

	return result
}

func SumAllTails(x ...[]int) (result []int) {
	for _, collection := range x {
		if len(collection) == 0 {
			result = append(result, 0)
		} else {
			tail := collection[1:]
			result = append(result, Sum(tail))
		}
	}
	return result
}
