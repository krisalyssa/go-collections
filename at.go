package gocollections

func At[E ~[]I, I any](enumerable E, index int) I {
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
