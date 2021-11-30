package main

func main() {

}

func compare(a, b int, reverse bool) bool {
	if reverse {
		return a > b
	} else {
		return a < b
	}
}

func Sort(array []int, reverse bool) []int {

	if len(array) < 2 {
		return array
	}

	splitter := int(len(array) / 2)
	leftPart := Sort(array[:splitter], reverse)
	rightPart := Sort(array[splitter:], reverse)

	l, r := 0, 0
	res := make([]int, len(array), len(array))
	for i := 0; i < len(array); i++ {
		if r >= len(rightPart) || l < len(leftPart) && compare(leftPart[l], rightPart[r], reverse) {
			res[i] = leftPart[l]
			l++
		} else {
			res[i] = rightPart[r]
			r++
		}
	}

	return res
}
