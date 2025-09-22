package gocollections

func ReduceWhile[E ~[]I, I, A any](enumerable E, initial A, fun func(I, A) (A, bool)) A {
	accumulator := initial
	for _, item := range enumerable {
		var cont bool

		accumulator, cont = fun(item, accumulator)
		if !cont {
			break
		}
	}
	return accumulator
}
