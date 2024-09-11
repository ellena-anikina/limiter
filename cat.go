package cat

func cat(s string, n int) string {

	count := make(map[rune]int)
	res := ""

	for _, char := range s {
		count[char]++
		if count[char] <= n {
			res += string(char)
		}
	}

	return res
}
