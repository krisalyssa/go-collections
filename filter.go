package gocollections

// Thank you to Vincent Taneri for https://vitaneri.com/posts/implementing-map-filter-and-reduce-using-generic-in-go.
// The original code is licensed [CC BY-NC-SA 4.0](https://creativecommons.org/licenses/by-nc-sa/4.0/).

// declare type for filtering function?

func Filter[E ~[]I, I any](enumerable E, fun func(I) bool) E {
	var retval E
	for _, item := range enumerable {
		if fun(item) {
			retval = append(retval, item)
		}
	}
	return retval
}
