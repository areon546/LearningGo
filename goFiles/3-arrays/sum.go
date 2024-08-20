package arraysAndSlices

func Sum(xs []int) int {
	sum := 0

	for _, number := range xs {
		sum += number
	}

	return sum
}

func SumAll(intArrays ...[]int) []int {
	length := len(intArrays)
	sums := make([]int, length)

	for i, array := range intArrays {
		sums[i] = Sum(array)
	}

	return sums
}

// Returns a slice with each value being the sum of all elements (after the first one) for the respectively entered int slice
func SumAllTails(intArrays ...[]int) []int {
	length := len(intArrays)
	sums := make([]int, length)

	for i, array := range intArrays {

		if len(array) == 0 {
			sums[i] = 0
			continue
		}

		tail := array[1:]
		sums[i] = Sum(tail)
	}

	return sums
}
