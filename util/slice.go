package util

func NotNullList[T any](list []T) []T {
	if len(list) == 0 {
		return []T{}
	}
	return list
}
