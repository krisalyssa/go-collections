package gocollections

// Thank you to Vincent Taneri for https://vitaneri.com/posts/implementing-map-filter-and-reduce-using-generic-in-go.
// The original code is licensed [CC BY-NC-SA 4.0](https://creativecommons.org/licenses/by-nc-sa/4.0/).

// declare type for filtering function?

func Reject[T any](enumberable []T, fun func(T) bool) []T {
	var retval []T
	for _, item := range enumberable {
		if !fun(item) {
			retval = append(retval, item)
		}
	}
	return retval
}
