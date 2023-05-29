package gophers

import (
	"sort"

	"golang.org/x/exp/constraints"
)

// Sort xs in place.
func Sort[T constraints.Ordered](xs []T) {
	sort.Slice(xs, func(i, j int) bool {
		return xs[i] < xs[j]
	})
}

// Sort xs in place, comparing elements using keyfunc f.
func SortBy[T any, K constraints.Ordered](xs []T, f func(T) K) {
	sort.Slice(xs, func(i, j int) bool {
		return f(xs[i]) < f(xs[j])
	})
}

// Create a sorted copy of xs.
func Sorted[T constraints.Ordered](xs []T) []T {
	out := make([]T, len(xs))
	for i := range xs {
		out[i] = xs[i]
	}
	Sort(out)
	return out
}

// Search for the position of x in xs, using binary search.
// The input slice xs must be sorted. Return len(xs) if no match is found.
func BSearch[T comparable](xs []T, x T) int {
	return sort.Search(len(xs), func(i int) bool { return xs[i] == x })
}

// Check if xs (which must be sorted) contains x, using binary search.
func BContains[T comparable](xs []T, x T) bool {
	i := BSearch(xs, x)
	return i < len(xs) && xs[i] == x
}

// Check if x exists in xs (linear time).
func Contains[T comparable](xs []T, x T) bool {
	for _, e := range xs {
		if x == e {
			return true
		}
	}
	return false
}

// Apply f to each element in xs and return an array with the results.
func Map[A, B any](f func(A) B, xs []A) []B {
	out := make([]B, 0, len(xs))
	for _, x := range xs {
		out = append(out, f(x))
	}
	return out
}

// Filter elements in xs based on boolean predicate f.
func Filter[T any](f func(T) bool, xs []T) []T {
	out := make([]T, 0, len(xs))
	for _, x := range xs {
		if f(x) {
			out = append(out, x)
		}
	}
	return out
}

// Reduce the array of xs into an accumulator value using reducer f.
func Reduce[A, V any](f func(A, V) A, acc A, xs []V) A {
	for _, x := range xs {
		acc = f(acc, x)
	}
	return acc
}

// Create a function that composes f and g into a single function n,
// such that n(x) = g(f(x)).
func Compose[A, B, C any](f func(A) B, g func(B) C) func(A) C {
	return func(a A) C { return g(f(a)) }
}

// Create a function that composes f, g and h into a single function n,
// such that n(x) = h(g(f(x))).
func Compose3[A, B, C, D any](f func(A) B, g func(B) C, h func(C) D) func(A) D {
	return func(a A) D { return h(g(f(a))) }
}

// Group xs by key function f into a map of arrays. Relative order of
// xs in map values is preserved.
func GroupBy[T any, K comparable](f func(T) K, xs []T) map[K][]T {
	out := make(map[K][]T)
	for _, x := range xs {
		k := f(x)
		if a, ok := out[k]; ok {
			out[k] = append(a, x)
		} else {
			out[k] = []T{x}
		}
	}
	return out
}

// Return the keys in the map.
func MapKeys[K comparable, V any](m map[K]V) (keys []K) {
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// Return the values in the map.
func MapValues[K comparable, V any](m map[K]V) (values []V) {
	for _, v := range m {
		values = append(values, v)
	}
	return values
}

// Return only unique values in xs, disregarding order.
func Uniq[T comparable](xs []T) []T {
	seen := map[T]bool{}
	for _, x := range xs {
		if _, ok := seen[x]; !ok {
			seen[x] = true
		}
	}
	return MapKeys(seen)
}

// Return (newxs, value), with value at idx removed from newxs.
func SlicePop[T any](xs []T, idx int) ([]T, T) {
	return append(xs[:idx], xs[idx+1:]...), xs[idx]
}
