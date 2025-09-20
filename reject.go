package gocollections

// Thank you to Vincent Taneri for https://vitaneri.com/posts/implementing-map-filter-and-reduce-using-generic-in-go.
// The original code is licensed [CC BY-NC-SA 4.0](https://creativecommons.org/licenses/by-nc-sa/4.0/).

// declare type for filtering function?

func Reject[T any](enumerable []T, fun func(T) bool) []T {
	return Filter(enumerable, func(item T) bool { return !fun(item) })
}
