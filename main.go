package main

import "fmt"

func main() {
	var n int
	var curr int
	var total int = 0
	fmt.Scanln(&n)
	slice := make([]int, 0, n)
	for i := 0; i < n; i++ {
		fmt.Scanln(&curr)
		slice = append(slice, curr)
	}
	for i := 0; i < len(slice); i++ {
		total += slice[i]
	}
	fmt.Println(total)
}