package arrayslices

func Sum(numbers []int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}

func SumAllTails(manyInputs ...[]int) []int {

	ansList := []int{}

	for _, input := range manyInputs {
		tail := input[1:]
		ansList = append(ansList, Sum(tail))
	}

	return ansList
}
