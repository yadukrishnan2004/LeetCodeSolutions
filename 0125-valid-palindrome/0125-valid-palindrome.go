func isPalindrome(s string) bool {

	var r strings.Builder
	result := strings.ToLower(s)

	for _, v := range result {
		if unicode.IsLetter(v) || unicode.IsDigit(v) {
			r.WriteRune(v)
		}
	}

	str := r.String()

	i := len(str) - 1
	j := 0

	for j < i {
		if str[j] != str[i] {
			return false
		}
		i--
		j++
	}

	return true
}