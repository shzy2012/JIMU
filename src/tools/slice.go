package tools

func SliceContains(slice []string, str string) bool {
	for _, ele := range slice {
		if ele == str {
			return true
		}
	}
	return false
}
