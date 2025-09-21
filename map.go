package gocollections

// Thank you to Vincent Taneri for https://vitaneri.com/posts/implementing-map-filter-and-reduce-using-generic-in-go.
// The original code is licensed [CC BY-NC-SA 4.0](https://creativecommons.org/licenses/by-nc-sa/4.0/).

// declare type for mapping function?

func Map[E1 ~[]I1, I1, I2 any](enumerable E1, fun func(I1) I2) []I2 {
	retval := make([]I2, len(enumerable))
	for index, item := range enumerable {
		retval[index] = fun(item)
	}
	return retval
}
