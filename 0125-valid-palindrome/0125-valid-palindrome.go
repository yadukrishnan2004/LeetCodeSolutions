func isPalindrome(s string) bool {
    if s == " " || len(s)==1 {
    return true
}
    var r strings.Builder
        result := strings.ReplaceAll(s, " ", "")
        result = strings.ToLower(result) 
    	for _,v:=range result {
            if unicode.IsLetter(v) {
                r.WriteRune(v)
            }
        }

str := r.String()


i := len(str) - 1
j := 0


if i < 1{
    return false
}
for j < i {
    if str[j] != str[i] {
        return false
    }
    i--
    j++
}

	return true
}