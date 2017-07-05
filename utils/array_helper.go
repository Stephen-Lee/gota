package utils

func IsContain(arr []interface{}, ele interface{}) bool {
	for _, i := range arr {
		if i == ele {
			return true
		}
	}
	return false
}
