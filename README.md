# go-collections

Functions for working with collections, inspired by Elixir's Enum module.

## `func All`

```golang
func All[T any](enumerable []T, fun func(T) bool) bool
```

Returns `true` if `fun` returns `true` for all items in `enumerable`. If `fun` ever returns `false`, iteration stops immediately and `false` is returned.

## `func Any`

```golang
func Any[T any](enumerable []T, fun func(T) bool) bool
```

Returns `true` if `fun` returns `true` for at least one item in `enumerable`. If `fun` ever returns `true`, iteration stops immediately and `true` is returned. In all other cases `false` is returned.

## `func At`

```golang
func At[T ~[]E, E any](enumerable T, index int) E
```

Finds the item at the given `index` (zero-based). A negative `index` can be passed, in which case `index` is counted from the end. For example, `At(enumerable, -1)` finds the last item in `enumerable`.

If `index` is greater than or equal to the number of items in `enumerable`, the last item is returned. If `index` is negative and its absolute value is greater than the number of items in `enumerable`, the first item is returned.

## `func Filter`

```golang
func Filter[T any](enumerable []T, fun func(T) bool) []T
```

Returns a list of only those elements in `enumerable` for which `fun` returns a truthy value.

See also `Reject` which discards all elements where the function returns a truthy value.

## `func Map`

```golang
func Map[T any](enumerable []T, fun func(T) bool) []T
```

Returns a list where each element is the result of invoking `fun` on each corresponding element of `enumerable`.

For maps, `fun` will be passed the `value` of each item in the map.

_TODO: Define a different implementation for maps, which passes the key-value pair._

## `func Reduce`

```golang
func Reduce[T1, T2 any](enumerable []T1, initial T2, fun func(T1, T2) T2) T2
```

Invokes `fun` for each element in `enumerable` with an accumulator.

The initial value of the accumulator is `initial`. The result returned by `fun` is used as the accumulator for the next iteration. The function returns the last accumulator.

## `func Reject`

```golang
func Reject[T any](enumerable []T, fun func(T) bool) []T
```

Returns a list of only those elements in `enumerable` for which `fun` does not return a truthy value.

See also `Filter` which includes all elements where the function returns a truthy value.
