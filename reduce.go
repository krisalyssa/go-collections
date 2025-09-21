package gocollections

// Thank you to Vincent Taneri for https://vitaneri.com/posts/implementing-map-filter-and-reduce-using-generic-in-go.
// The original code is licensed [CC BY-NC-SA 4.0](https://creativecommons.org/licenses/by-nc-sa/4.0/).

// declare type for reducing function?

func Reduce[E ~[]I, I, A any](enumerable E, initial A, fun func(I, A) A) A {
	accumulator := initial
	for _, item := range enumerable {
		accumulator = fun(item, accumulator)
	}
	return accumulator
}
