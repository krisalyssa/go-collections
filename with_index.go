package gocollections

func WithIndex[E ~[]I, I any, P Pair[I, int]](enumerable E) []P {
	retval := make([]P, len(enumerable))
	for index, item := range enumerable {
		retval[index] = P{item, index}
	}
	return retval
}
