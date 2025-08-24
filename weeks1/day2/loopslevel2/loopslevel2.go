package main

import "fmt"

func PrintMultiples( n int) []string {
	var result []string
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			line := fmt.Sprintf("%d x %d = %d", i, j, i*j)
			result = append(result, line)
		}
	}
	return result
}

func main() {
	for _, d := range PrintMultiples(6) {
		fmt.Printf(d)
	}
}
