package main

func reverseWords(s string) string {
	str := []rune(s)
	left, right := 0, 0
	for i, ch := range s {
		if ch != ' ' {
			right = i
		}

		if ch == ' ' || i+1 == len(str) {
			for left < right {
				str[left], str[right] = str[right], str[left]
				left++
				right--
			}

			left = i + 1
		}
	}

	return string(str)
}

func main() {
	println(reverseWords("God is a girl"))
}
