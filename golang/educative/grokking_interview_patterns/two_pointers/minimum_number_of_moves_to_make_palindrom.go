package main

func minMovesToMakePalindrome(s string) int {

	chars := []rune(s)
	moves := 0

	i, j := 0, len(chars)-1

	for i < j {
		if chars[i] == chars[j] {
			i++
			j++
		} else {
			k := j
			for k > i && chars[k] != chars[i] {
				k--
			}

			if k == i {
				chars[i], chars[i+1] = chars[i+1], chars[i]
				moves++
			} else {
				for k < j {
					chars[k], chars[k+1] = chars[k+1], chars[k]
					k++
					moves++
				}

				i++
				j--
			}
		}
	}

	return moves
}
