
# No Safe Efficient Ways to Do Three-way String Comparisons in Go

Three-way string comparison is [widely used] in programming.
But up to now (Go 1.19), the [strings.Compare] function in the standard library
is (intentionally) implemented with a non-opitimized way:

[widely used]: https://sourcegraph.com/search?q=context:global+switch+strings.Compare+lang:Go+&patternType=literal

```Go
func Compare(a, b string) int {
	// NOTE(rsc): This function does NOT call the runtime cmpstring function,
	// because we do not want to provide any performance justification for
	// using strings.Compare. Basically no one should use strings.Compare.
	// As the comment above says, it is here only for symmetry with package bytes.
	// If performance is important, the compiler should be changed to recognize
	// the pattern so that all code doing three-way comparisons, not just code
	// using strings.Compare, can benefit.
	if a == b {
		return 0
	}
	if a < b {
		return -1
	}
	return +1
}
```

When comparing two unequal strings, the implementation will perform two comparison operations.
Whereas a perfect implementation only needs one,
just like the implementation of [bytes.Compare] shown below,
in which [bytealg.Compare] is impplemented using assembly code
on most architectures.

```Go
func Compare(a, b []byte) int {
	return bytealg.Compare(a, b)
}
```

The `strings.Compare` implementation is comparatively inefficent.
Specifically, it is less efficent when the two string operands are not equal but their lengths are equal.

[strings.Compare]: https://github.com/golang/go/blob/go1.19/src/strings/compare.go#L7-L28
[bytes.Compare]: https://github.com/golang/go/blob/go1.19/src/bytes/bytes.go#L23-L28
[bytealg.Compare]: https://github.com/golang/go/blob/go1.19/src/internal/bytealg/compare_native.go#L12

Benchmark code constantly shows `strings.Compare` uses 2x CPU time of `bytes.Compare`
when comparing unequal same-length byte sequences (we view both strings and byte slices as byte sequences here).

The internal comment for the current `strings.Compare` implementation
is some interesting. The comment suggests that we should not use
`strings.Compare` in Go at all, but no alternative efficient ways are available now yet.
It mentions that the compiler should make special optimizations to automatically
convert the code using comparision operators into internal optimized three-way comparisons if possible.
However, such compiler optimizations have never been made,
and there are no plans to make such optimizations yet as far as I know.
Personally, I doubt such optimizations are feasible to make for any use case.
So I think the `strings.Compare` should be implemented efficiently,
to avoid breaking user expectations.


_(This is one of the dozens of facts mentioned in the [Go Optimizations 101] book.)_

[Go Optimizations 101]: https://go101.org/optimizations/101.html





