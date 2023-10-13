
# No Safe Efficient Ways to Do Three-way String Comparisons in Go

Three-way string comparison is widely used in programming ([proof 1] and [proof 2]).
But up to now (Go 1.19), the [strings.Compare] function in the standard library
is ([intentionally]) implemented with [a non-opitimized way]:

[proof 1]: https://sourcegraph.com/search?q=context:global+switch+strings.Compare+lang:Go+&patternType=literal
[proof 2]: https://sourcegraph.com/search?q=context:global+strings.Compare+lang:Go&patternType=standard
[intentionally]: https://news.ycombinator.com/item?id=33353106
[a non-opitimized way]: https://go-review.googlesource.com/c/go/+/3012

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
in which [bytealg.Compare] is implemented using assembly code
on most architectures.

```Go
func Compare(a, b []byte) int {
	return bytealg.Compare(a, b)
}
```

The `strings.Compare` implementation is comparatively inefficient.
Specifically, it is less efficient when the two string operands are not equal but their lengths are equal.

[strings.Compare]: https://github.com/golang/go/blob/go1.19/src/strings/compare.go#L7-L28
[bytes.Compare]: https://github.com/golang/go/blob/go1.19/src/bytes/bytes.go#L23-L28
[bytealg.Compare]: https://github.com/golang/go/blob/go1.19/src/internal/bytealg/compare_native.go#L12

Benchmark code constantly shows `strings.Compare` uses 2x CPU time of `bytes.Compare`
when comparing unequal same-length byte sequences (we view both strings and byte slices as byte sequences here).

The internal comment for the current `strings.Compare` implementation
is some interesting. The comment suggests that we should not use
`strings.Compare` in Go at all, but no alternative efficient ways are available now yet
(ironically, this function is used in [Go toolchain code] and recommended by [a standard library function]).
It mentions that the compiler should make special optimizations to automatically
convert the code using comparison operators into internal optimized three-way comparisons if possible.
However, such compiler optimizations have never been made,
and there are no plans to make such optimizations yet as far as I know.
Personally, I doubt such optimizations are feasible to be made for any use case.
So I think [the `strings.Compare` should be implemented efficiently][issue 50167],
to avoid breaking user expectations.

[Go toolchain code]: https://github.com/golang/go/blob/go1.19/src/cmd/go/internal/modindex/read.go#L822
[a standard library function]: https://github.com/golang/go/blob/go1.19/src/sort/search.go#L88-L99
[issue 50167]: https://github.com/golang/go/issues/50167

_(This is one of the dozens of facts mentioned in the [Go Optimizations 101] book.)_

[Go Optimizations 101]: https://go101.org/optimizations/101.html

Update: view discussions on [reddit] and [HN].

[reddit]: https://old.reddit.com/r/programming/comments/ycz0va/no_safe_efficient_ways_to_do_threeway_string/
[HN]: https://news.ycombinator.com/item?id=33316402





