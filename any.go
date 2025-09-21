package gocollections

// declare type for filtering function?

func Any[E ~[]I, I any](enumerable E, fun func(I) bool) bool {
	var retval bool = false
	for _, item := range enumerable {
		if fun(item) {
			retval = true
			break
		}
	}
	return retval
}
