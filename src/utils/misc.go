package utils

func ForEach[T any](array []T, consumer func(T)) {
	for _, v := range array {
		consumer(v)
	}
}

func Map[T any, O any](array []T, mapper func(T) O) []O {
	result := make([]O, len(array))
	for i, v := range array {
		result[i] = mapper(v)
	}
	return result
}

func Filter[T any](array []T, filter func(T) bool) []T {
	result := make([]T, 0)
	for _, v := range array {
		if filter(v) {
			result = append(result, v)
		}
	}
	return result
}
