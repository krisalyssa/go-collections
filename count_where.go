package gocollections

func CountWhere[E ~[]I, I any](enumerable E, fun func(I) bool) int {
	return Count(Filter(enumerable, fun))
}
