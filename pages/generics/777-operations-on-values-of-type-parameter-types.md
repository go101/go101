

# Operations on Values of Type Parameter Types

This chapter will talk about which operations on values of type parameters
are valid and which are invalid in generic function bodies.

Within a generic function body,
an operation on a value of a type parameter is valid only if it is
valid for values of every type in the type set of the constraint of the type parameter.
In the current custom generic design and implementation (Go 1.19),
it is not always vice versa.
Some extra requirements must be met to make the operation valid.

Currently, there are many such restrictions. Some of them are temporary
and might be removed from future Go versions, some are permanent.
The temporary ones are mainly caused by implementation workload,
so they need some time and efforts to be removed eventually.
The permanent ones are caused by the custom generics design principles.

The following contents of this chapter will list these restrictions.
Some facts and related concepts will also be listed.

## The type of a value used in a generic function must be a specified type

As mentioned in a previous chapter, since Go 1.18,
value types in Go could be categorized in two categories:

* type parameter types: the types declared in type parameter lists.
* ordinary types: the value types not declared in type parameter lists.
  Before Go 1.18, there are only ordinary types.

Go custom generics are not implemented as simple code text templates.
This is a fundamental difference from code generation.
There is a principle rule in Go programming:
every typed expression must have a specified type,
which may be either an ordinary type, or a type parameter.

For example, in the following code snippet, only the function `dot` doesn't compile.
the other ones compile okay.
The reasons are simple:

* in the function `foo`, the type of `x` is `T`, which is a type parameter.
  Certainly, in uses of the function, `x` might be instantiated as `int` or `string`,
  but which doesn't change the fact that, from the view of compilers,
  its type is a type parameter.
* in the function `bar`, the types of `x[i]` and `x[y]` are both a type parameter, `E`.
* in the function `win`, the types of `x[1]` and `x[y]` are both a specified ordinary type, `int`.
* in the function `dot`, the types of `x[1]` and `x[y]` are might be `int` or `string` (two different ordinary types), though they are always identical.

```Go
func foo[T int | string](x T) {
	var _ interface{} = x // okay
}

func bar[T []E, E any](x T, i, j int) () {
	x[i] = x[j] // okay
}

func win[T ~[2]int | ~[8]int](x T, i, j int) {
	x[i] = x[j] // okay
}

func dot[T [2]int | [2]string](x T, i, j int) {
	x[i] = x[j]      // error: invalid operation
	var _ any = x[i] // error: invalid operation
}
```

The element types of strings are viewed as `byte`, so the following code compiles,

```Go
func ele[ByteSeq ~string|~[]byte](x ByteSeq, n int) {
	_ = x[n] // okay
}
```

For the same reason (the principle rule), in the following code snippet,
the functions `nop` and `jam` both compile okay,
but the function `mud` doesn't.

```Go
func nop[T *Base, Base int32|int64](x T) {
	*x = *x + 1 // okay
}

func jam[T int32|int64](x *T) {
	*x = *x + 1 // okay
}

func mud[T *int32|*int64](x T) {
	*x = *x + 1 // error: invalid operation
}
```

The same, in the following code snippet, only the function `box` fails to compile,
the other two both compile okay.

```Go
func box[T chan int | chan byte](c T) {
	_ = <-c // error: no core type
}


func sed[T chan E, E int | byte](c T) {
	_ = <-c // okay
}

type Ch <-chan int
func cat[T chan int | Ch](c T) {
	_ = <-c // okay
}
```

This rule [might be relaxed to some extent in future Go versions](https://github.com/golang/go/issues/52129).

## Type parameters may be type asserted to

As a type parameter is a specified type, it may be type asserted to.
The following code compiles, even if there are duplicate `case` type expressions
at run time in the `type-switch` code block within the `wua` function.

```Go
import "fmt"

func nel[T int | string](x any) {
	if v, ok := x.(T); ok {
		fmt.Printf("x is a %T\n", v)
	} else {
		fmt.Printf("x is not a %T\n", v)
	}
}

func wua[T int | string](x any) {
	switch v := x.(type) {
	case T:
		fmt.Println(v)
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	}
}
```

## Type parameters may be not used as types of (local) named constants

That means values of type parameters are all non-constants.

For example, the following function fails to compile.

```Go
func f[P int]() {
	const y P = 5 // error: invalid constant type P
}
```

This fact [will never be changed](https://github.com/golang/go/issues/50202).

Because of this fact, converting a constant to a type parameter yields
a non-constant value of the argument passed to the type parameter.
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
and [this issue](https://github.com/golang/go/issues/28591)
for why the two functions return different results.

## The core type of a type

A non-interface type always has a core type, which
is the underlying type of the non-interface type.
Generally, we don't care about such case in using custom generics.

An interface type might have a core type or not.

1. Generally speaking, if all types in the type set of the interface type (a constraint)
share an identical [underlying type](https://go101.org/article/type-system-overview.html#underlying-type),
then the identical underlying type is called the core type of the interface type.
1. If the types in the type set of then interface type don't share an identical underlying type
but they are all [channel types](https://go101.org/article/channel.html)
which share an identical element type `E`, and all directional channels in them have the same direction,
then the core type of the interface type is the directional channel type
`chan<- E` or `<-chan E`, depending on the direction of the directional channels present.
1. For cases other than the above two, the interface type has not a core type.

For example, in the following code, each of the types shown in the first group
has a core type (indicated in the tail comments), yet the types shown in the
second group all have no core types.

```Go
type (
	Age      int                   // int
	AgeC     interface {Age}       // int
	AgeOrInt interface {Age | int} // int
	Ints     interface {~int}      // int
	
	AgeSlice  []Age                        // []Age
	AgeSlices interface{~[]Age}            // []Age
	AgeSliceC interface {[]Age | AgeSlice} // []Age
	
	C1 interface {chan int | chan<- int} // chan<- int
	C2 interface {chan int | <-chan int} // <-chan int
)

type (
	AgeOrIntSlice interface {[]Age | []int}
	OneParamFuncs interface {func(int) | func(int) bool}
	Streams       interface {chan int | chan Age}
	C3            interface {chan<- int | <-chan int}
)
```

Many operations require the constraint of a type parameter has a core type.

To make descriptions simple, sometimes, we also call the core type of the constraint
of a type parameter as the core type of the type parameter.

## A function is required to have a core type to be callable

For example, currently (Go 1.19), in the following code, the functions `foo` and `bar` don't compile, bit the `tag` function does.
The reason is the `F` type parameters in the `foo` and `bar` generic functions
both have not a core type, even

but the `F` type parameter in the `tag` generic function does have.

```Go
func foo[F func(int) | func(any)] (f F, x int) {
	f(x) // error: invalid operation: cannot call non-function f
}

func bar[F func(int) | func(int)int] (f F, x int) {
	f(x) // error: invalid operation: cannot call non-function f
}

type Fun func(int)

func tag[F func(int) | Fun] (f F, x int) {
	f(x) // okay
}
```

It is unclear whether or not the rule will be relaxed in future Go versions.

## The type literal in a composite literal must have a core type

For example, currently (Go 1.19), in the following code snippet,
the functions `foo` and `bar` compile okay, but the other ones don't.

```Go
func foo[T ~[]int] () {
	_ = T{}
}

type Ints []int

func bar[T []int | Ints] () {
	_ = T{} // okay
}

func ken[T []int | []string] () {
	_ = T{} // error: invalid composite literal type T
}

func jup[T [2]int | map[int]int] () {
	_ = T{} // error: invalid composite literal type T
}
```

## An element index operation requires the container operand's type set may not include maps and non-maps at the same time

And if all types in the type set are maps, then their underlying types must be identical
(in other words, the type of the operand must have a core type).
Otherwise, their element types must be identical.
The elements of strings are viewed as `byte` values.

For example, currently (Go 1.19), in the following code snippet, only the functions `foo` and `bar` compile okay.

```Go
func foo[T []byte | [2]byte | string](c T) {
	_ = c[0] // okay
}

type Map map[int]string
func bar[T map[int]string | Map](c T) {
	_ = c[0] // okay
}

func lag[T []int | map[int]int](c T) {
	_ = c[0] // invalid operation: cannot index c
}

func vet[T map[string]int | map[int]int](c T) {
	_ = c[0] // invalid operation: cannot index c
}
```

The restriction might be removed in the future Go versions
(just my hope, in fact I'm not sure on this).

If the type of the index expression is a type parameter,
then all types in its type set must be integers.
The following function compiles okay.

```Go
func ind[K byte | int | int16](s []int, i K) {
	_ = s[i] // okay
}
```

_(It looks the current Go specification is not correct on this.
The specification requires the index expression must has a core type.)_

## A (sub)slice operation requires the container operand has a core type

For example, currently (Go 1.19), the following two functions both fail to compile,
even if the subslice operations are valid for all types in the corresponding type sets.

```Go
func foo[T []int | [2]int](c T) {
	_ = c[:] // invalid operation: cannot slice c: T has no core type
}

func bar[T [8]int | [2]int](c T) {
	_ = c[:] // invalid operation: cannot slice c: T has no core type
}
```

The restriction might be removed in the future Go versions
(again, just my hope, in fact I'm not sure on this).

There is an exception for this rule. If the container operand's type set
only include string and byte slice types, then it is not required to have a core type.
For example, the following function compiles okay.

```Go
func lol[T string | []byte](c T) {
	_ = c[:] // okay
}
```

Same as element index operations, if the type of an index expression is a type parameter,
then all types set of its type set must be all integers.

## In a `for-range` loop, the ranged container is required to have a core type

For example, currently (Go 1.19), in the following code, 
only the last two functions, `dot1` and `dot2`, compile okay.

```Go
func values[T []E | map[int]E, E any](kvs T) []E {
	r := make([]E, 0, len(kvs))
	// error: cannot range over kvs (T has no core type)
	for _, v := range kvs {
		r = append(r, v)
	}
	return r
}

func keys[T map[int]string | map[int]int](kvs T) []int {
	r := make([]int, 0, len(kvs))
	// error: cannot range over kvs (T has no core type)
	for k := range kvs {
		r = append(r, k)
	}
	return r
}

func sum[M map[int]int | map[string]int](m M) (sum int) {
	// error: cannot range over m (M has no core type)
	for _, v := range m {
		sum += v
	}
	return
}

func foo[T []int | []string] (v T) {
	// error: cannot range over v (T has no core type)
	for range v {}
}

func bar[T [3]int | [6]int] (v T) {
	// error: cannot range over v (T has no core type)
	for range v {}
}

type MyInt int

func cat[T []int | []MyInt] (v T) {
	// error: cannot range over v (T has no core type)
	for range v {}
}

type Slice []int

func dot1[T []int | Slice] (v T) {
	for range v {} // okay
}

func dot2[T ~[]int] (v T) {
	for range v {} // okay
}
```

The restriction is intended. I think its intention is to ensure both of
the two iteration variables always have a specified type
(either an ordinary type or a type parameter type).
However, this restriction is over strict for this intention.
Because, in practice, the key types or element types of some containers are identical,
even if the underlying type of the containers are different.
And in many use cases, one of the two iteration variables is ignored.

I'm not sure whether or not [the restriction will be removed in future Go versions](https://github.com/golang/go/issues/49551).
In my opinion, the restriction reduces the usefulness of Go custom generics in some extent.

If all possible types are slice and arrays, and their element types are identical,
we could use plain `for` loops to walk around this restriction.

```Go
func cat[T [3]int | [6]int | []int] (v T) {
	for i := 0; i < len(v); i++ { // okay
		_ = v[i] // okay
	}
}
```

The call to the `len` predeclared function is valid here.
A later section will talk about this.

The following code also doesn't compile, but which is reasonable.
Because the iterated elements for `string` are `rune` values,
whereas the iterated elements for `[]byte` are `byte` values.

```Go
func mud[T string | []byte] (v T) {
	for range v {} // error: cannot range over v (T has no core type)
}
```

If it is intended to iterate the bytes in either byte slices and strings,
we could use the following code to achieve the goal.

```Go
func mud[T string | []byte] (v T) {
	for range []byte(v) {} // okay
}
```

The conversion `[]byte(v)` (if it follows the `range` keyword) is [specifically
optimized by the official standard Go compiler](https://go101.org/article/string.html#conversion-optimizations) so that it doesn't duplicate
underlying bytes.

The following function doesn't compile now (Go 1.19),
even if the types of the two iteration variables are always `int` and `rune`.
Whether or not it will compile in future Go versions is unclear.

```Go
func aka[T string | []rune](runes T) {
	// cannot range over runes (T has no core type)
	for i, r := range runes {
		_ = i
		_ = r
	}
}
```

<!--
https://github.com/golang/go/issues/51053
-->


## Type parameter involved conversions

Firstly, we should know [the conversion rules for ordinary types/values](https://go101.org/article/value-conversions-assignments-and-comparisons.html).

By the current specification (Go 1.19),
given two types `From` and `To`, assume at least one of them is a type parameter,
then a value of `From` can be converted to `To` if a value of each type in
the type set of `From` can be converted to each type in the type set of `T`
(note that the type set of an ordinary type only contains the ordinary type itself).

For example, all of the following functions compile okay.

```Go
func pet[A ~int32 | ~int64, B ~float32 | ~float64](x A, y B){
	x = A(y)
	y = B(x)
}

func dig[From ~byte | ~rune, To ~string | ~int](x From) To {
	return To(x)
}

func cov[V ~[]byte | ~[]rune](x V) string {
	return string(x)
}

func voc[V ~[]byte | ~[]rune](x string) V {
	return V(x)
}
```

But the following function fails to compile,
because `string` values may not be converted to `int`.

```Go
func eve[X, Y int | string](x X) Y {
	return Y(x) // error
}
```

The following function doesn't compile, even if the conversion in it
is valid for all possible type arguments.
The reason is `[]T` is an ordinary type, not a type parameter,
and its underlying type is itself.
There is not a rule which allows converting values from `[]T` to `string`.

```Go
func jon[T byte](x Bytes) []T {
	return []T(x) // error
}
```

Future Go versions [might relax the rules](https://github.com/golang/go/issues/50421)
to make the conversion in the above example valid.

By using the official standard Go compiler, in the following program,

* the functions `tup` and `pad` don't compile.
  The reason is values of type `AgePtr` can't be directly converted to `*int`.
* all the other three generic functions compile okay, but the `dot` function
  should not compile by the above described rule.
  This might be [a bug of the standard compiler, or the rule described in
  the current Go specification needs a small adjustment](https://github.com/golang/go/issues/50815).

```Go
package main

type Age int
type AgePtr *Age

func dot[T ~*Age](x T) *int {
	return (*int)(x) // okay
}

func tup(x AgePtr) *int {
	// error: cannot convert x (variable of type AgePtr)
	//        to type *int
	return (*int)(x)
}

func tup2(x AgePtr) *int {
	return (*int)((*Age)(x))
}

func pad[T AgePtr](x T) *int {
	// error: cannot convert x to type *int
	return (*int)(x)
}

func pad2[T AgePtr](x T) *int {
	return (*int)((*Age)(x))
}

func main() {
	var x AgePtr
	var _ = dot[AgePtr](x)
	var _ = tup2(x)
	var _ = pad2[AgePtr](x)
}
```



## Type parameter involved assignments

Firstly, we should know [the assignment rules for ordinary types/values](https://go101.org/article/value-conversions-assignments-and-comparisons.html).

In the following descriptions, the type of the destination value is called as the destination type, and the type of the source value is called as the source type.

By the current specification (Go 1.19), for a type parameter involved assignment,

* if the destination type is a type parameter and the source value is
  an untyped value, then the assignment is valid only if
  the untyped value is assignable to each type in the type set of
  the destination type parameter.
* if the destination type is a type parameter but the source type is an ordinary type,
  then the assignment is valid only if the source ordinary type is
  [unnamed](https://go101.org/article/type-system-overview.html#named-type)
  and its values is assignable to each type in the type set of the destination type parameter.
* if the source type is a type parameter but the destination type is an ordinary type,
  then the assignment is valid only if the destination ordinary type is unnamed
  and values of each type in the type set of the source type parameter
  are assignable to the destination ordinary type.
* if both of the destination type and the source type are type parameters,
  then the assignment is invalid.

From the rules, we could get that type value of a named type can not be assigned to another named type.

In the following code snippet, there are four invalid assignments.

```Go
func dat[T ~int | ~float64, S *int | []bool]() {
	var _ T = 123 // okay
	var _ S = nil // okay
}

func zum[T []byte](x []byte) {
	var t T = x // okay
	type Bytes []byte
	var y Bytes = x // okay (both are ordinary types)
	x = t // okay
	x = y // okay
	
	// Both are named types.
	t = y // error
	y = t // error
	
	// To make the above two assignments valid,
	// the sources in then must be converted.
	t = T(y)     // okay
	y = Bytes(t) // okay
}


func pet[A, B []byte](x A, y B){
	// Both are type parameters.
	x = y // error: cannot use y as type A in assignment
	y = x // error: cannot use x as type B in assignment
}
```

It is unclear whether or not the assignment rules will be relaxed in future Go versions.
It looks [the posibility is small](https://github.com/golang/go/issues/51501).

## Calls to predeclared functions

The following are some rules and details for the calls to some predeclared functions
when type parameters are involved. 

## A call to the predeclared `len` or `cap` functions is valid if it is valid for all of the types in the type set of the argument

In the following code snippet, the function `capacity` fails to compile,
the other two functions both compile okay.

```Go
type Container[T any] interface {
	~chan T | ~[]T | ~[8]T | ~*[8]T | ~map[int]T | ~string
}

func size[T Container[int]](x T) int {
	return len(x) // okay
}

func capacity[T Container[int]](x T) int {
	return cap(x) // error: invalid argument x for cap
}

func capacity2[T ~chan int | ~[]int](x T) int {
	return cap(x) // okay
}
```

Please note that a call to `len` or `cap` always returns a non-constant value
if the type of the argument of the call is a type parameter,
even of the type set of the argument only contains arrays and pointers to arrays.
For example, in the following code,
the first `cap` and `len` calls within the first two functions
all fail to compile.

```Go
func f[T [2]int](x T) {
	const _ = cap(x) // error: cap(x) is not constant
	const _ = len(x) // error: len(x) is not constant

	var _ = cap(x) // okay
	var _ = len(x) // okay
}

func g[P *[2]int](x P) {
	const _ = cap(x) // error: cap(x) is not constant
	const _ = len(x) // error: len(x) is not constant

	var _ = cap(x) // okay
	var _ = len(x) // okay
}

func h(x [2]int) {
	const _ = cap(x) // okay
	const _ = len(x) // okay
	const _ = cap(&x) // okay
	const _ = len(&x) // okay
}
```

The rule [might be changed](https://github.com/golang/go/issues/50226).
But honestly speaking, the possibility is very small.
Personally, I think the current behavior is more logical.

Because of this rule, the following two functions return different results.

```Go
package main

const S = "Go"

func ord(x [8]int) byte {
	return 1 << len(x) >> len(x)
}

func gen[T [8]int](x T) byte {
	return 1 << len(x) >> len(x)
}

func main() {
	var x [8]int
	println(ord(x), gen(x)) // 1 0
}
```

Again, please read the [strings in Go](https://go101.org/article/string.html) article
and [this issue](https://github.com/golang/go/issues/28591)
for why the two functions return different results.

Please not that, the following function doesn't compile,
because the type of `&x` is `*T`, which is a pointer
to a type parameter, instead of a pointer to an array.

```Go
func e[T [2]int]() {
	var x T
	var _ = len(&x) // invalid argument: &x for len
	var _ = cap(&x) // invalid argument: &x for cap
}
```

In other words, a type parameter which type set contains only one type
is not equivalent to that only type.
A type parameter has wave-particle duality.
For some situations, it acts as the types in its type set.
For some other situations, it acts as a distinct type.
More specifically, a type parameter acts as a distinct type
(which doesn't share underlying type with any other types)
when it is used as a component of a composite type.
In the above example. `*T` and `*[2]int` are two different (ordinary) types.

## A call to the predeclared `new` function has not extra requirements for its argument

The following function compiles okay.

```Go
func MyNew[T any]() *T {
	return new(T)
}
```

It is equivalent to

```Go
func MyNew[T any]() *T {
	var t T
	return &t
}
```

## A call to the predeclared `make` function requires its first argument (the container type) has a core type 

Currently (Go 1.19), in the following code snippet, the functions `voc` and `ted` both
fail to compile, the other two compile okay.
The reason is the first argument of a call to the predeclared `make` function
is required to have a core type.
Neither of the `voc` and `ted` functions satisfies this requirement,
whereas both of the other two functions satisfy this requirement.

```Go
func voc[T chan bool | chan int]() {
	_ = make(T) // error: invalid argument: no core type
}

func ted[T chan<- int | <-chan int]() {
	_ = make(T) // error: invalid argument: no core type
}

type Stream chan int
type Queue Stream

func fat[T Stream | chan int | Queue | chan<- int]() {
	_ = make(T) // okay
}

func nub[T Stream | chan int | Queue | <-chan int]() {
	_ = make(T) // okay
}
```

By my understanding, this requirement is in order to make subsequent operations
on the made containers (they are channels in the above example) always legal.
For example, to make sure a value received from the made
channel has a specified type (either a type parameter, or an ordinary type).

Personally, I think the requirement is over strict.
After all, for some cases, the supposed subsequent operations don't happen.

To use values of a type parameter which doesn't have a core type within a generic function,
we can pass such values as value arguments into the function, as the following code shows.

```Go
func doSomething(any) {}

func voc2[T chan bool | chan int](x T) {
	doSomething(x)
}

func ted2[T chan<- int | <-chan int](x T) {
	doSomething(x)
}
```

Because of the same requirement, neither of the following two functions compile.

```Go
func zig[T ~[]int | map[int]int](c T) {
	_ = make(T) // error: invalid argument: no core type
}

func rat[T ~[]int | ~[]bool](c T) {
	_ = make(T) // error: invalid argument: no core type
}
```

Calls to the predeclared `new` function have not this requirement.

## A call to the predeclared `delete` function requires all types in the type set of its first argument have an identical key type

Note, here, the identical key type may be ordinary type or type parameter type.

The following functions both compile okay.

```Go
func zuk[M ~map[int]string | ~map[int]bool](x M, k int) {
	delete(x, k)
}

func pod[M map[K]int | map[K]bool, K ~int | ~string](x M, k K) {
	delete(x, k)
}
```

## A call to the predeclared `close` function requires all types in the type set of its argument are channel types

The following function compiles okay.

```Go
func dig[T ~chan int | ~chan bool | ~chan<- string](x T) {
	close(x)
}
```

Note that the current Go specification requires that the argument of
a call to the predeclared `close` function must have a core type.
This is inconsistent with the implementation of the official standard Go compiler.

## Calls to predeclared `complex`, `real` and `imag` functions don't accept arguments of type parameter now

Calling the three functions with arguments of type parameters might break the principle rule mentioned in the first section of the current chapter.

This is a problem the current custom generics design is unable to solve.
There is [an issue](https://github.com/golang/go/issues/50937) for this.

## About constraints with empty type sets

The type sets of some interface types might be empty.
An empty-type-set interface type implements any interface types,
including itself.

Empty-type-set interface types are totally useless in practice,
but they might affect the implementation perfection from theory view.

There are really several imperfections in the implementation
of the current official standard Go compiler (v1.19).

For [example](https://github.com/golang/go/issues/51470),
should the following function compile?
It does with the latest official standard Go compiler (v1.19).
However, one of the above sections has mentioned that a `make` call
requires its argument must have a core type.
The type set of the constraint `C` declared in the following code
is empty, so it has not a core type, then the `make` call within
the `foo` function should not compile.

```Go
// This is an empty-type-set interface type.
type C interface {
        map[int]int
        M()
}
       
func foo[T C]() {
        var _ = make(T)
}
```

This following is [another example](https://github.com/golang/go/issues/51917#issuecomment-1084188702),
in which all the function calls in the function `g` should compile okay.
However, two of them fail to compile with
the latest official standard Go compiler (v1.19).

```Go
func f1[T any](x T) {}
func f2[T comparable](x T) {}
func f3[T []int](x T) {}
func f4[T int](x T) {}

// This is an empty-type-set interface type.
type C interface {
	[]int
	m()
}

func g[V C](v V) {
	f1(v) // okay
	f2(v) // error: V does not implement comparable
	f3(v) // okay
	f4(v) // error: V does not implement int
}
```

The current Go specification specially states:

> Implementation restriction: A compiler need not report an error if an operand's type is a type parameter with an empty type set. Functions with such type parameters cannot be instantiated; any attempt will lead to an error at the instantiation site.

So the above shown imperfections are not bugs of the official standard Go compiler.





