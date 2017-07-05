package utils

func IsContainString(arr []string, ele string) bool {
	for _, i := range arr {
		if i == ele {
			return true
		}
	}
	return false
}
