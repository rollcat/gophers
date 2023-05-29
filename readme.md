# Small gophers that help you write Go programs

Functional-style utilities. You can think of this as the missing
toplevel builtins, things you can find e.g. in Python.

Requires Go 1.18 (generics).

All functions have short, familiar names, like `Assert` or `Filter`,
and you are encouraged to dot-import the module for easy access.

## Handling errors

```go
package main
import (
    . "github.com/rollcat/gophers"
)
```

Now instead of writing this:

```go
x, err := GetX()
if err != nil {
    panic(err)
}
```

You can write this:

```go
x := Must(GetX())
```

It will crash just the same.

## Functional style

For those, who do not believe that Go has a Lisp nature,
I present the familiar repertoire of `Map`, `Filter`, and `Reduce`.

There is also a handful of other small utilities such as `GroupBy`
(which classifies elements from a list into a map, according to a
function), or `SortBy` (duh). 

These are all made possible because of the introduction of generics.
