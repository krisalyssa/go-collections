package gocollections

func At[T ~[]E, E any](enumerable T, index int) E {
	count := len(enumerable)

	if index >= 0 {
		if index >= count {
			return enumerable[count-1]
		}

		return enumerable[index]
	} else {
		if (count + index) < 0 {
			return enumerable[0]
		}

		return enumerable[count+index]
	}
}
