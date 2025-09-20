package gocollections

// Thank you to Vincent Taneri for https://vitaneri.com/posts/implementing-map-filter-and-reduce-using-generic-in-go.
// The original code is licensed [CC BY-NC-SA 4.0](https://creativecommons.org/licenses/by-nc-sa/4.0/).

// declare type for mapping function?

func Map[T1, T2 any](enumerable []T1, fun func(T1) T2) []T2 {
	retval := make([]T2, len(enumerable))
	for index, item := range enumerable {
		retval[index] = fun(item)
	}
	return retval
}
