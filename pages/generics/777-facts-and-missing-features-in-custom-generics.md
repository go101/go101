

## Instantiated types are ordinary type, even if some of type arguments are type parameters

For example, the declaration for the `R` and `S` generic types are valid
in the following program.

```Go
package main

type T[X, Y any] func(X) Y

func (t T[X, Y]) Do(x X) Y {
	return t(x)
}

// The source types are partial instantiations
// of the generic type T.
type (
    R[X any]       T[X, bool] // okay
    S[X, Y, Z any] T[Z, Y]    // okay
)

func main() {
	var t T[int, bool] = func(x int) bool {
		return x%2 == 0
	}
	var a, b bool = t.Do(3), t.Do(6)
	println(a, b) // false true
	
	// The underlying type of R[int] and T[int, bool]
	// are identical, so the conversion is valid.
	var r = R[int](t)
	
	// The underlying type of S[string, bool, int] and 
    // T[int, bool] are also identical, so the
    // conversion is also valid.
	var s = S[string, bool, int](t)
	
	_, _ = r, s
	
	// The following two lines both fail to compile,
	// because the type R[int] and S[string, bool, int]
	// both have no methods.
	// var _ bool = r.Do(3)
	// var _ bool = s.Do(3)
}
```

Please note that, a defined generic type doesn't obtain the method declared directly
for the source type. This is the same as [ordinary type definitions](https://go101.org/article/method.html#method-obtaining).

## A composite types are always unnamed ordinary types, even if some of its components are type parameters

The componets of composite types include container element types, struct field types,
pointer base types, function parameter and result types.

For example, the following two generic type declarations are both valid.
In the declarations, `[]T` and `*T` are both ordinary types.

```Go
type Floats[T ~float32 | ~float64] []T
type PtrConstraint[T any] interface {~*T}
```

On the other hand, the following two declartions are both invalid,
because type parameter types (`T` here) may not be used as source types
in type specifications and type terms in interface elements.

```Go
type Floats[T ~float32 | ~float64] T     // error
type PtrConstraint[T any] interface {~T} // error
```

The following generic function declaration is valid.
In the declaration, the composite types (`[]B`, `*A`, `map[B]A`) are all ordinary types.
Please note, `[]B` and `*A` are simplified forms of
`interface{[]B}` and `interface{*A}`.

```Go
func nut[A []B, B *A]() {
	var _ map[B]A
}
```

## Type declarations inside generic functions are not currently supported

Currently (Go 1.18), local type declarations are not allowed in generic functions.
For example, in the following code, the ordinary function `f` compiles okay,
but the generic function `g` doesn't.

```Go
func f() {
	type _ int // okay
}

func g[T any]() {
	type _ int // error
}

type T[_ any] struct{}

func (T[_]) m() {
	type _ int // error
}
```

This restriction might [be removed in future Go versions](https://github.com/golang/go/issues/47631).

## Type parameters may be not used as types of (local) named constants

For example, the following function fails to compile.

```Go
func f[P int]() {
	const y P = 5 // error: invalid constant type P
}
```

This fact [will never be changed](https://github.com/golang/go/issues/50202).

Because of this fact, converting a constant to a type parameter yields
a non-constant value of the argument pased to the type parameter.
For example, in the following code, the function `h` compiles,
but the function `g` doesn't.

```Go
const N = 5

func g[P int]() {
	const y = P(N) // error: P(N) is not constant
}

func h[P int]() {
	var y = P(N) // okay
	_ = y
}
```

Because of the conversion rule, the return results of the two
functions, `mud` and `tex`, are different.

```Go
package main

const S = "Go"

func mud() byte {
	return 64 << len(string(S)) >> len(string(S))
}

func tex[T string]() byte {
	return 64 << len(T(S)) >> len(T(S))
}

func main() {
	println(mud()) // 64
	println(tex()) // 0
}
```

Please read the [strings in Go](https://go101.org/article/string.html) article
for why the two functions return different results.

## Generic type aliases are not supported currently

Currently (Go 1.18), a declared type alias may not have type parameters.
For example, in the following code, only the alias declaration for `A` is legal,
the other alias declarations are all illegal.

The alias `A` is actually an alias to an ordinary type `func(int) string`.

```Go
type T[X, Y any] func(X) Y

type A = T[int, string] // okay

// generic type cannot be alias
type B[X any] = T[X, string]   // error
type C[X, Y, Z any] = T[X, Y]  // error
type D[X any] = T[int, string] // error
```

Generic type aliases [might be supported in future Go versions](https://github.com/golang/go/issues/46477).

## Aliases to non-basic interface types are not supported currently

Currently (Go 1.18), there is an unintended restriction: non-basic interface types may not be aliased.
For example, the following type alias declarations are illegal:

```Go
type C[T any] interface{~int; M() T}
type C1 = C[bool]
type C2 = comparable
type C3 = interface {~[]byte | ~string}
```

Whereas the following ones are legal:

```Go
type Ca = any
type Cb = interface{M1(); M2() int}
```

The unintended restriction [will be removed in Go 1.19](https://github.com/golang/go/issues/51616).

## Non-basic interface types may not be used as value types

This has been mentioned in the last chapter.
Whether or not this restriciton will be removed in future Go versions in unclear now.

## Type parameters may not be embedded (as of Go 1.18)

The following type declarations are invalid,
even if the type set of (the constraint of) of the type parameter `T`
doesn't contain pointer types (named pointer types are not embeddable
by the current Go specification).

```Go
type A[T ~int] struct {
	T // error
}

type B[T ~int] struct {
	*T // error
}
```

The main reason for this restriction is that,
if non-basic interface types may be used as value types in future Go versions,
then an interface type argument `I` might be passed to the `T` type parameter,
but `*I` is also not capable of being embedded by the current Go specificaiton.

[It is unclear](https://github.com/golang/go/issues/49030)
whether or not [this restriction](https://github.com/golang/go/issues/24062)
will be removed in future Go versions.



## About constraints with empty type sets

https://github.com/golang/go/issues/51470
https://github.com/golang/go/issues/51917



## The method set of a constraint is not calculated compeletely for some cases

The Go specification states:

> The method set of an interface type is the intersection of the method sets of each type in the interface's type set.

However, currently (Go toolchain 1.18), only the methods explicitly specified in interface types are calculated. For example, in the code, the method set of the constraint should contain both `Foo` and `Bar`,
and the code should compile okay, but it doesn't (as of Go toolchain 1.18).

```Go
package main

type S struct{}

func (S) Bar() {}

type C interface {
	S
	Foo()
}

func foobar[T C](v T) {
	v.Foo() // okay
	v.Bar() // v.Bar undefined
}

func main() {}
```

This restriction is planed to be removed in Go toochain 1.19.

The related issue: https://github.com/golang/go/issues/51183

## No ways to specific a field set for a constraint

We know that an interface type may specify a method set.
But up to now (Go 1.18), it could not specify a (struct) field set.

The restriction might be lifted from future Go versions.

There is a proposal for this: https://github.com/golang/go/issues/51259

```Go
type I interface{ ~[]struct{ A, b int } | ~[]struct{ A, x int } }
func Print[T I](v T) {
	fmt.Println(v[0].A) // ✅ Works
}
```

https://github.com/golang/go/issues/51977

## No ways to use common fields of the type set of a constraint if the constraint has not a core (struct) type

Currently (Go 1.18), even if all types in the type set of a constraint
are structs and they share some common fields, the common fields still
could not be used if the structs don't share the identical underlying type.

For example, the generic functions in the following example all fail to compile.

```Go
package main

type S1 struct {
	X int
}

type S2 struct {
	X int `json:X`
}

type S3 struct {
	X int
	Z bool
}

type S4 struct {
	S1
}

func F12[T S1 | S2](v T) {
	_ = v.X // error: v.x undefined
}

func F13[T S1 | S3](v T) {
	_ = v.X // error: v.x undefined
}

func F14[T S1 | S4](v T) {
	_ = v.X // error: v.x undefined
}

func main() {}
```

There is a proposal to remove this limit: https://github.com/golang/go/issues/48522

Temp walkaround: through explicitly specified methods

https://github.com/golang/go/issues/50417
field accesses through type parameters will be disabled for Go 1.18.
https://github.com/golang/go/issues/51576
https://github.com/golang/go/issues/50233





## Compile-time type switch

https://github.com/golang/go/issues/45380#issuecomment-1074153465




## no generic method, use alternative

https://github.com/golang/go/issues/49085

https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md#methods-may-not-take-additional-type-arguments


This will make code even more cubersome.
And make `reflect.Method` ...

## No ways to construct a constraint which is only satisfied by interface types.

or by non-interface types

## No ways to construct a constraint which is only satisfied by non-interface types.

or by non-interface types

## Custom generic type instantiations are not inconsistent with built-in generic types

My personal opinion

This increases the recognizaiton burdon.

## Fact: it is impossible to get unnamed instantiated types

This is determined