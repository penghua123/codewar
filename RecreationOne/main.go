package main

import (
	"fmt"
	"math"
)

func main() {
	result := ListSquared(1, 100)
	fmt.Println(result)
}

func ListSquared(m, n int) [][]int {
	var squared [][]int
	for i := m; i <= n; i++ {
		div := divisor(i)
		res := detectRecreation(div)
		if res {
			itemSuccess := []int{i, int(math.Pow(float64(i), 2))}
			squared = append(squared, itemSuccess)
		}
	}
	return squared
}

func divisor(m int) []int {
	var div []int
	for i := 1; i <= m; i++ {
		if m%i == 0 {
			div = append(div, i)
		}
	}
	return div
}

func detectRecreation(div []int) bool {
	sum := float64(0)
	for _, num := range div {
		sum += math.Pow(float64(num), 2)
	}
	sumSqr := math.Sqrt(sum)
	sumSqrInt := int(sumSqr)
	if sumSqr == float64(sumSqrInt) {
		return true
	} else {
		return false
	}
}
