# GOUF - Utility Functions for generic types

Go team [released Go 1.18 beta](https://tip.golang.org/doc/go1.18) recently with support for [Generics(a.k.a type parameters)](https://go.dev/doc/tutorial/generics). This package provides much needed and frequently used utility functions for generic data types.

## Supported Operations

### Slices

- All - Check if all elements in a slice satisfy a condition
- Any - Check if any element in a slice satisfies a condition
- Check - Check if slice can accept a given indices
- Count - count all elements in a slice that satisfy a condition
- Delete - Delete element at given index
- DeleteUnordered - Delete element at given index, without preserving order(efficient than Delete)
- DeleteFn - Delete element that satisfies a condition
- DeletePlaces - Delete elements at given indices
- Filter - Filter elements in a slice that satisfy a condition
- FilterInPlace - Filter elements in a slice that satisfy a condition without allocating new slice(efficient than Filter)
- FilterInPlaceGC - Filter elements in a slice that satisfy a condition without allocating new slice and zeroing wanted elements(efficient than Filter)
- FindFn - Find element in a slice(for comparable types)
- FindFn - Find element in a slice that satisfies a condition
- Map - Map elements in a slice
- Reduce - Reduce elements in a slice to a single value
- Sum - Sum all elements in a slice (for numeric/string types)
- SumFn - Sum all elements in a slice that satisfy a condition(for custom types)
- Unique - Remove duplicate elements in a slice (for comparable types)
- UniqueFn - Remove duplicate elements in a slice that satisfy a condition(for custom types)
- UniqueInPlace - Remove duplicate elements in a slice without allocating new slice(efficient than Unique)
- UniqueFnInPlace - Remove duplicate elements in a slice without allocating new slice(efficient than UniqueFn)

And a lot more(in pipeline).

### Maps

TBD

## Want to contribute?

see [guide](CONTRIBUTING.md).

## Contact

[Akshay Bharambe](https://twitter.com/akshaybharambe1)

## Feedback

This is very much important to improve this package. Do let me know if you have any feedback. Leave a ⭐ if you find this helpful.
