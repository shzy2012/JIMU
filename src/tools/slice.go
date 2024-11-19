package tools

func SliceContains(slice []string, str string) bool {
	for _, ele := range slice {
		if ele == str {
			return true
		}
	}
	return false
}

func Reverse[T any](arr []T) []T {
	// 定义两个索引，一个从开头，一个从末尾开始
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		// 交换元素
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}
