package mymath

func IsBigNum(n int) bool {
	if n >= 1000 {
		return true
	}
	return false
}

func IsSmallNum(n int) bool {
	if n < 1000 {
		return true
	}
	return false
}
