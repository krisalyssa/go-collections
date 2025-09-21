package gocollections

// declare type for filtering function?

func All[T any](enumerable []T, fun func(T) bool) bool {
	var retval bool = true
	for _, item := range enumerable {
		if !fun(item) {
			retval = false
			break
		}
	}
	return retval
}
