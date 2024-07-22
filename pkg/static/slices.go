package static

func FilterSlice[T any](slice []T, filter func(int, T) bool) []T {
	var filteredSlice []T
	for i, item := range slice {
		if filter(i, item) {
			filteredSlice = append(filteredSlice, item)
		}
	}
	return filteredSlice
}

func MapSlice[T, U any](slice []T, mapper func(int, T) U) []U {
	var mappedSlice []U
	for i, item := range slice {
		mappedSlice = append(mappedSlice, mapper(i, item))
	}
	return mappedSlice
}
