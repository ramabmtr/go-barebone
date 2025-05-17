package checker

func StrIn(s string, checker ...string) bool {
	for _, c := range checker {
		if s == c {
			return true
		}
	}

	return false
}
