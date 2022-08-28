
# About Go Custom Generics

The main purpose of custom generics is to avoid code repetitions,
or in other words, to increase code reusability.

For some situations, generics could also lead to cleaner code and simpler APIs
(not always).

For some situations, generics could also improve code execution performance
(again not always).

Before version 1.18, for many Go programmers, the lack of custom generics caused pains in Go programming under some situations.

Indeed, the pains caused by the lack of custom generics were alleviated to a certain extend by the following facts.

* Since version 1.0, Go has been supported built-in generics, which include some built-in generic type kinds (such as map and channel) and generic functions (`new`, `make`, `len`, `close`, etc).
* Go supports reflection well (through interfaces and the `reflect` standard package).
* Repetitive code could be generated automatically by using some tools (such as the `//go:generate` comment directive supported by the official Go toolchain).

However, the pains are still there for many use cases.
The demand for custom generics became stronger and stronger.
In the end, the Go core team decided to support custom generics in Go.

For all sorts of reasons, including considerations of syntax/semantics backward compatibility and implementation difficulties, the Go core team settled down on [the type parameters proposal](https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md) to implement custom generics.

<!--
https://dl.acm.org/doi/10.1145/3428217
https://dl.acm.org/doi/pdf/10.1145/3428217
-->

The first Go version supporting custom generics is 1.18.

The type parameters proposal tries to solve many code reuse problems, but not all.
And please note that, not all the features mentioned in the parameters proposal have been implemented yet currently (Go 1.19). The custom generics design and implementation will continue to evolve and get improved in future Go versions. And please note that the proposal is not the ceiling of Go custom generics.

Despite the restrictions (temporary or permanent ones) in the current Go custom generics design and implementation,
I also have found there are some details which are handled gracefully and beautifully in the implementation.

Although Go custom generics couldn't solve all code reuse problems,
personally, I believe Go custom generics will be used widely in future Go programming.

