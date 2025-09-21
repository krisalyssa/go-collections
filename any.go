package gocollections

// declare type for filtering function?

func Any[T any](enumerable []T, fun func(T) bool) bool {
	var retval bool = false
	for _, item := range enumerable {
		if fun(item) {
			retval = true
			break
		}
	}
	return retval
}
