package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func double(nums []int) []int {

	result := make([]int, 0, len(nums))

	for _, num := range nums {
		result = append(result, num*2)
	}

	return result
}

func main() {
	list := []int{1, 2, 3, 4, 8}
	fmt.Println(double(list))
}
