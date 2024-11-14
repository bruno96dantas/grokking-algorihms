package main

import "fmt"

func binarySearch(numberList []int, item int) int {

	lowNumber := 0
	highNumber := len(numberList) - 1

	for lowNumber <= highNumber {
		mid := (lowNumber + highNumber) / 2

		if numberList[mid] == item {
			return mid
		} else if numberList[mid] < item {
			lowNumber = mid + 1
		} else {
			highNumber = mid - 1
		}
	}

	return 0
}

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func main() {
	x := binarySearch(makeRange(1, 20), 3)
	fmt.Println(x)
}
