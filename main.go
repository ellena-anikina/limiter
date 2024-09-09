package main

import "fmt"

func Append(list []int, el ...int) []int {
	var res []int
	resLen := len(list) + len(el)

	if resLen <= cap(list) {
		res = list[:resLen]
	} else {
		resCap := 3 * len(list)
		res = make([]int, resLen, resCap)
		copy(res, list)
	}

	for i := 0; i < len(el); i++ {
		res[len(list)+i] = el[i]
	}

	return res
}

func main() {

	list := make([]int, 4, 5)

	list = Append(list, 1, 2, 3)

	fmt.Println(list, len(list), cap(list))
}
