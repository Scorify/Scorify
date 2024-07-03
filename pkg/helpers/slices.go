package helpers

func FilterSlice[T any](slice []T, filter func(T) bool) []T {
	var filteredSlice []T
	for _, item := range slice {
		if filter(item) {
			filteredSlice = append(filteredSlice, item)
		}
	}
	return filteredSlice
}

func MapSlice[T, U any](slice []T, mapper func(T) U) []U {
	var mappedSlice []U
	for _, item := range slice {
		mappedSlice = append(mappedSlice, mapper(item))
	}
	return mappedSlice
}
