# go-collections

Functions for working with collections, inspired by Elixir's Enum module.

## `func Map`

```golang
func Map[T any](enumerable []T, fun func(T) bool) []T
```

Returns a list where each element is the result of invoking `fun` on each corresponding element of `enumerable`.

For maps, `fun` will be passed the `value` of each item in the map.

_TODO: Define a different implementation for maps, which passes the key-value pair._
