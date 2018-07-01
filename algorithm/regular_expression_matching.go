package algorithm

func IsRegluarExpressionMatch(str, patten string) bool {
	if len(str) == 0 && len(patten) == 0 {
		return true
	} else if len(str) == 0 {
		return false
	} else if len(patten) == 0 {
		return false
	} else if patten == ".*" {
		return true
	}

	i, j := 0, 0
	for i < len(str) && j < len(patten) {
		currentChar := patten[0]
		if currentChar == str[i] {
			i++
			j++
			continue
		}

		if currentChar == byte('.') {
			i++
			j++
			continue
		}

		if currentChar == byte('*') {

		}
		if str[i] == currentChar {
			i++
			j++
		} else {
			if patten[j] == byte('*') {

			}
			i++
		}
	}

	return false
}
