package gophers

import (
	"fmt"
	"testing"
)

func checkSlice[T comparable](t *testing.T, expected, got []T) {
	if len(got) != len(expected) {
		t.Fatalf("expected %#v; got %#v", expected, got)
		return
	}
	for i := range expected {
		if expected[i] != got[i] {
			t.Fatalf("expected %#v; got %#v", expected, got)
			return
		}
	}
}

func TestSortInts(t *testing.T) {
	t.Parallel()
	data := []int{4, 222, 6, 1, 17, 0}
	expected := []int{0, 1, 4, 6, 17, 222}
	Sort(data)
	checkSlice(t, expected, data)
}

func TestSortStrings(t *testing.T) {
	t.Parallel()
	data := []string{"4", "222", "6", "1", "17", "0"}
	expected := []string{"0", "1", "17", "222", "4", "6"}
	Sort(data)
	checkSlice(t, expected, data)
}

func TestSortedInts(t *testing.T) {
	t.Parallel()
	data := []int{4, 222, 6, 1, 17, 0}
	expected := []int{0, 1, 4, 6, 17, 222}
	got := Sorted(data)
	checkSlice(t, expected, got)
}

func TestSortBy(t *testing.T) {
	t.Parallel()
	data := []int{4, 222, 6, 1, 17, 0}
	expected := []int{222, 17, 6, 4, 1, 0}
	SortBy(data, func(i int) int { return -i })
	checkSlice(t, expected, data)
}

func TestBSearchSome(t *testing.T) {
	t.Parallel()
	data := []int{0, 1, 4, 6, 17, 222}
	if idx := BSearch(data, 222); idx != 5 {
		t.Fatalf("did not find 222 at position 5: %#v", idx)
	}
}

func TestBSearchNone(t *testing.T) {
	t.Parallel()
	data := []int{0, 1, 4, 6, 17, 222}
	if idx := BSearch(data, 224); idx != 6 {
		t.Fatalf("expected 224 to not be found: %#v", idx)
	}
}

func TestBContainsOK(t *testing.T) {
	t.Parallel()
	data := []int{0, 1, 4, 6, 17, 222}
	if !BContains(data, 222) {
		t.Fatalf("expected to find 222 in: %#v", data)
	}
}

func TestBContainsFail(t *testing.T) {
	t.Parallel()
	data := []int{0, 1, 4, 6, 17, 222}
	if BContains(data, 224) {
		t.Fatalf("expected to not find 224 in: %#v", data)
	}
}

func TestContainsOK(t *testing.T) {
	t.Parallel()
	data := []int{4, 222, 6, 1, 17, 0}
	if !Contains(data, 222) {
		t.Fatalf("expected to find 222 in: %#v", data)
	}
}

func TestContainsFail(t *testing.T) {
	t.Parallel()
	data := []int{4, 222, 6, 1, 17, 0}
	if Contains(data, 224) {
		t.Fatalf("expected to not find 224 in: %#v", data)
	}
}

func TestMap(t *testing.T) {
	t.Parallel()
	data := []int{4, 222, 6, 1, 17, 0}
	expected := []int{6, 224, 8, 3, 19, 2}
	got := Map(func(a int) int { return a + 2 }, data)
	checkSlice(t, expected, got)
}

func TestFilterEven(t *testing.T) {
	t.Parallel()
	data := []int{4, 222, 6, 1, 17, 0}
	expected := []int{4, 222, 6, 0}
	got := Filter(func(a int) bool { return a%2 == 0 }, data)
	checkSlice(t, expected, got)
}

func TestFilterOdd(t *testing.T) {
	t.Parallel()
	data := []int{4, 222, 6, 1, 17, 0}
	expected := []int{1, 17}
	got := Filter(func(a int) bool { return a%2 == 1 }, data)
	checkSlice(t, expected, got)
}

func TestReduceInt(t *testing.T) {
	t.Parallel()
	data := []int{4, 222, 6, 1, 17, 0}
	expected := 250
	got := Reduce(func(a, b int) int { return a + b }, 0, data)
	if got != expected {
		t.Fatalf("expected sum of %#v is %#v; got %#v", data, expected, got)
	}
}

func TestReduceAsFilter(t *testing.T) {
	t.Parallel()
	data := []int{4, 222, 6, 1, 17, 0}
	expected := []int{4, 222, 6, 0}
	got := Reduce(func(acc []int, v int) []int {
		if v%2 == 0 {
			return append(acc, v)
		}
		return acc
	}, []int{}, data)
	checkSlice(t, expected, got)
}

func TestReduceAsMap(t *testing.T) {
	t.Parallel()
	data := []int{4, 222, 6, 1, 17, 0}
	expected := []int{6, 224, 8, 3, 19, 2}
	got := Reduce(func(acc []int, v int) []int {
		return append(acc, v+2)
	}, []int{}, data)
	checkSlice(t, expected, got)
}

func TestComposeInt(t *testing.T) {
	t.Parallel()
	plusOne := func(a int) int { return a + 1 }
	plusTwo := Compose(plusOne, plusOne)
	if plusTwo(0) != 2 {
		t.Fail()
	}
}

func TestCompose3Int(t *testing.T) {
	t.Parallel()
	plusOne := func(a int) int { return a + 1 }
	plusThree := Compose3(plusOne, plusOne, plusOne)
	if plusThree(0) != 3 {
		t.Fail()
	}
}

func TestComposeMixed(t *testing.T) {
	t.Parallel()
	plusOne := func(a int) int { return a + 1 }
	toString := func(a int) string { return fmt.Sprintf("%d", a) }
	f := Compose(plusOne, toString)
	if f(0) != "1" {
		t.Fail()
	}
}

func TestCompose3Mixed(t *testing.T) {
	t.Parallel()
	length := func(xs []int) int { return len(xs) }
	plusOne := func(a int) int { return a + 1 }
	toString := func(a int) string { return fmt.Sprintf("%d", a) }
	f := Compose3(length, plusOne, toString)
	if f([]int{}) != "1" {
		t.Fail()
	}
}

func TestUniq(t *testing.T) {
	t.Parallel()
	data := []int{4, 222, 4, 6, 1, 0, 17, 0, 222}
	expected := []int{0, 1, 4, 6, 17, 222}
	got := Uniq(data)
	Sort(got)
	checkSlice(t, expected, got)
}

func TestGroupBy(t *testing.T) {
	t.Parallel()
	data := []int{4, 222, 6, 1, 17, 0}
	expected := map[int][]int{
		1: []int{1, 17},
		2: []int{4, 222, 6, 0},
	}
	got := GroupBy(func(a int) int {
		if a%2 == 0 {
			return 2
		} else {
			return 1
		}
	}, data)
	if odds, ok := got[1]; ok {
		checkSlice(t, expected[1], odds)
	} else {
		t.Fail()
	}
	if evens, ok := got[2]; ok {
		checkSlice(t, expected[2], evens)
	} else {
		t.Fail()
	}
}
