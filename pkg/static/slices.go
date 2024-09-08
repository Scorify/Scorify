package static

import (
	"slices"
)

func FilterSlice[T any](slice []T, filter func(int, T) bool) []T {
	filteredSlice := make([]T, 0, len(slice))
	for i, item := range slice {
		if filter(i, item) {
			filteredSlice = append(filteredSlice, item)
		}
	}
	return slices.Clip(filteredSlice)
}

func MapSlice[T, U any](slice []T, mapper func(int, T) U) []U {
	mappedSlice := make([]U, len(slice))
	for i, item := range slice {
		mappedSlice[i] = mapper(i, item)
	}
	return mappedSlice
}
