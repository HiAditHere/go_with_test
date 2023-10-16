package iteration

const repeatCount = 5

func Repeat(char string) string {
	ans := char

	for i := 1; i < repeatCount; i++ {
		ans += char
	}

	return ans
}
