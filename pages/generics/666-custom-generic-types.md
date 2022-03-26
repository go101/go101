
# More about generic types

##

type List[T any] struct {
	next *List[T]
	val  T
}


## []~B is not supported

## Currently (Go 1.18), aliases to non-fully instantiated types are not supported now.

For example, in the following code, the lines of declaring `A` and `B` don't compile.

```Go
type T[X, Y any] func(X) Y

type A[X] = T[X, string] // syntax error

type B[X, Y, Z any] = T[X, Y] // error: generic type cannot be alias

type D = T[int, string] // okay, for T[int, string] is instantiated.
```

Aliases to custom generic types might be supported in Go 1.19: https://github.com/golang/go/issues/46477



type C[X, Y any] struct{}
type C3 = C[int, int] // okay

```
package main

type G[T any] struct{
	m T
}

func (g G[T]) M() T {
	return g.m
}

func main() {
	type Alias = G[bool]
	_ = Alias.M // okay
	Alias.M(Alias{true})
	
	type Defined G[bool]
	_ = Defined.M // error: Defined.M undefined
}
```

## Generic interface types

```Go
type C[X any] interface {
	*X
	Double() X
}

func foo[D C[T], T any](v D) T {
	return v.Double()
}
```

## Anonymous type parameters

In the following generic type declaration,
the only type parameter is useless at all.

```Go
type Int[_ any] int
```

The pattern is not totally useless.

```Go

```

## A generic type may be not defined as a type parameter

```Go
type P[T any] T // error
type P[T any] [2]T // okay
```



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

## No ways to define a type parameter which only accepts certain interface type arguments



## Embedding type parameters

Embedding a type parameter, or a pointer to a type parameter, as
an unnamed field in a struct/interface type is not permitted. Similarly
embedding a type parameter in an interface type is not permitted.
Whether these will ever be permitted is unclear at present.

https://github.com/golang/go/issues/49030

https://github.com/golang/go/issues/24062

## Custom generic type instantiations are not inconsistent with built-in generic types

My personal opinion

===========================================================

## Instantiate

inconsistent with built-in generics

each instantiated type is named type.

Two instantiated types are identical if their defined types and all type arguments are identical. 

```Go
func f[A, B any](x A) B {
	type C []B // type declarations inside generic functions are not currently supported
	return any(x).(B)
}
```

```Go
// could not be instantiated.
func F[X chan Y, Y [2]X]() {}

type S[T []T] struct{}

type P[T *T,] struct{}
```

```Go
// could be instantiated.
func F[X ~chan Y, Y ~[2]X]() {}

type S[T ~[]T] struct{}

type P[T ~*T,] struct{}
```



## No ways to declare a constraint which is only satisfied by interface types.






