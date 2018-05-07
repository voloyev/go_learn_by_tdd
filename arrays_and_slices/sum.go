package arraysAndSlices

func Sum(numbers []int) int {
	sum := 0
	for _, el := range numbers {
		sum += el
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums
}

func SumAllTrails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
			continue
		}
		tail := numbers[1:]
		sums = append(sums, Sum(tail))
	}
	return sums
}
