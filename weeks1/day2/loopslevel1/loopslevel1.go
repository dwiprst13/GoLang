package main

import "fmt"

func PrintNumbers(n int) []int {
	var result []int
	for i := 1; i <= n; i++ {
		result = append(result, i)
	}
	return result
}

func main() {
	for _, v := range PrintNumbers(10) {
		fmt.Println(v)
	}
}
