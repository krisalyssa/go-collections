# go-collections

Functions for working with collections, inspired by Elixir's Enum module.

## Types

### `Pair`

```golang
type Pair[T, U any] struct {
	First  T
	Second U
}
```

## Functions

### `func All`

```golang
func All[E ~[]I, I any](enumerable E, fun func(I) bool) bool
```

Returns `true` if `fun` returns `true` for all items in `enumerable`. If `fun` ever returns `false`, iteration stops immediately and `false` is returned.

### `func Any`

```golang
func Any[E ~[]I, I any](enumerable E, fun func(I) bool) bool
```

Returns `true` if `fun` returns `true` for at least one item in `enumerable`. If `fun` ever returns `true`, iteration stops immediately and `true` is returned. In all other cases `false` is returned.

### `func At`

```golang
func At[E ~[]I, I any](enumerable E, index int) I
```

Finds the item at the given `index` (zero-based). A negative `index` can be passed, in which case `index` is counted from the end. For example, `At(enumerable, -1)` finds the last item in `enumerable`.

If `index` is greater than or equal to the number of items in `enumerable`, the last item is returned. If `index` is negative and its absolute value is greater than the number of items in `enumerable`, the first item is returned.

### `func Count`

```golang
func Count[E ~[]any](enumerable E) int
```

Returns the count of items in `enumerable`.

### `func CountWhere`

```golang
func CountWhere[E ~[]I, I any](enumerable E, fun func(I) bool) int
```

Returns the count of items in `enumerable` for which `fun` returns `true`.

### `func Filter`

```golang
func Filter[E ~[]I, I any](enumerable E, fun func(I) bool) E
```

Returns a list of only those elements in `enumerable` for which `fun` returns a truthy value.

See also `Reject` which discards all elements where the function returns a truthy value.

### `func Map`

```golang
func Map[E1 ~[]I1, I1, I2 any](enumerable E1, fun func(I1) I2) []I2
```

Returns a list where each element is the result of invoking `fun` on each corresponding element of `enumerable`.

For maps, `fun` will be passed the `value` of each item in the map.

_TODO: Define a different implementation for maps, which passes the key-value pair._

### `func Reduce`

```golang
func Reduce[E ~[]I, I, A any](enumerable E, initial A, fun func(I, A) A) A
```

Invokes `fun` for each element in `enumerable` with an accumulator.

The initial value of the accumulator is `initial`. The result returned by `fun` is used as the accumulator for the next iteration. The function returns the last accumulator.

### `func ReduceWhile`

```golang
func ReduceWhile[E ~[]I, I, A any](enumerable E, initial A, fun func(I, A) (A, bool)) A
```

Reduces `enumerable` until `fun` returns `(_, false)` or every item has been visited.

### `func Reject`

```golang
func Reject[E ~[]I, I any](enumerable E, fun func(I) bool) E
```

Returns a list of only those elements in `enumerable` for which `fun` does not return a truthy value.

See also `Filter` which includes all elements where the function returns a truthy value.

### `func WithIndex`

```golang
func WithIndex[E ~[]I, I any, P [T, U any]struct](enumerable E) P[I, int]
```

Returns `enumerable` with each item wrapped in a `Pair` with its index.
