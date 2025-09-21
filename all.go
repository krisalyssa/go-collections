package gocollections

// declare type for filtering function?

func All[E ~[]I, I any](enumerable E, fun func(I) bool) bool {
	var retval bool = true
	for _, item := range enumerable {
		if !fun(item) {
			retval = false
			break
		}
	}
	return retval
}
