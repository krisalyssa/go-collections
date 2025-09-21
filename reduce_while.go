package gocollections

func ReduceWhile[E ~[]I, I, A any](enumerable E, initial A, fun func(I, A) (A, bool)) A {
	accumulator := initial
	for _, item := range enumerable {
		var new_accumulator A
		var cont bool

		new_accumulator, cont = fun(item, accumulator)
		if cont {
			accumulator = new_accumulator
		} else {
			break
		}
	}
	return accumulator
}
