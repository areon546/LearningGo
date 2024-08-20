package iteration

// const repeatCount = 5

func Repeat(input string, repeatCount int) (repeated string) {
	for i := 0; i < repeatCount; i++ {
		repeated += input
	}

	return
}
