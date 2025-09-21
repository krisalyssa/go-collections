package gocollections

func Count[E ~[]I, I any](enumerable E) int {
	return len(enumerable)
}
