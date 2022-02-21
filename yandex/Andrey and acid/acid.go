package main

import (
	"fmt"
	"os"
)

func isSorted(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			return false
		}
	}
	return true
}
func groupTanks(tanks []int) []int {
	l := len(tanks)
	arr := make([]int, 0)
	for i := 0; i < l-1; i++ {
		if tanks[i] < tanks[i+1] {
			arr = append(arr, tanks[i])
		}
	}
	arr = append(arr, tanks[l-1])
	return arr
}

func retNum(tanks []int) int {
	if !isSorted(tanks) {
		return -1
	}
	num := 0
	arr := groupTanks(tanks)
	for i := 1; i < len(arr); i++ {
		num += arr[i] - arr[i-1]
	}
	return num
}

func main() {
	var n int
	fmt.Fscan(os.Stdin, &n)
	tanks := make([]int, n)
	for i, _ := range tanks {
		fmt.Fscan(os.Stdin, &tanks[i])
	}
	fmt.Fprintln(os.Stdout, retNum(tanks))
}
