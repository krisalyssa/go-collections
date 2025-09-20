package gocollections

// Thank you to Vincent Taneri for https://vitaneri.com/posts/implementing-map-filter-and-reduce-using-generic-in-go.
// The original code is licensed [CC BY-NC-SA 4.0](https://creativecommons.org/licenses/by-nc-sa/4.0/).

// declare type for reducing function?

func Reduce[T1, T2 any](enumerable []T1, initial T2, fun func(T1, T2) T2) T2 {
	accumulator := initial
	for _, item := range enumerable {
		accumulator = fun(item, accumulator)
	}
	return accumulator
}
