package filterOdd

/*
Написать функцию filterOdd(), которая принимает на вход слайс целых чисел.
Функция должна отфильтровать слайс так, чтобы в нем остались только нечетные числа.
Функция возвращает обработанный слайс.
При этом, необходимо минимизировать использование памяти, максимум, что можно выделить - это переменная для копии слайса.
*/

func filterOdd(nums []int) []int {
	// Your code here
	var result = make([]int, 0, len(nums))
	for _, num := range nums {
		if num%2 != 0 {
			result = append(result, num)
		}
	}
	return result
}
