
# Generic Instantiations and Type Argument Inferences

In the last two chapters, we have used several instantiations of generic types and functions.
Here, this chapter makes a formal introduction for instantiated types and functions.

## Generic type/function instantiations

Generic types must be instantiated to be used as types of values, and
generic functions must be instantiated to be called or used as function values.

A generic function (type) is instantiated by substituting a type argument list
for the type parameter list of its declaration (specification).
The lengths of the (full) type argument list is the same as the type parameter list.
Each type argument corresponds a type parameter and is passed to that type parameter.
Assume the constraint of the corresponding type parameter of a type argument is `C`,
then the type argument must be

* either an ordinary type (a non-interface type or a basic interface type)
  which [satisfies][satisfies] (but is not required to implement) `C`;
* or a type parameter which [implements] `C` if the instantiation
  is within a containing generic function declaration or type specification.
  The type parameter is declared by the containing generic function
  declaration or type specification.

[satisfies]: 555-type-constraints-and-parameters.md#implementation-vs-satisfaction
[implements]: 555-type-constraints-and-parameters.md#type-sets-and-implementations

Instantiated functions are non-generic functions.
Instantiated types are named ordinary value types.

Same as type parameter lists, a type argument list is also enclosed in square brackets
and type arguments are also comma-separated in the type argument list.
The comma insertion rule for type argument lists is also the same as type parameter lists.

Two type argument lists are identical if the lengths of their full forms are equal and all of their corresponding argument types are identical.
Two instantiated types are identical if they are instantiated from the same generic type and with the same full-from type argument list.

As indicated several times in the above descriptions, a type argument list might be partial.
Please read the following "type argument inferences" section for details.

In the following program, the generic type `Data` is instantiated four times.
Three of the four instantiations have the same type argument list
(please note that the predeclared `byte` is an alias of the predeclared `uint8` type).
So the type of variable `x`, the type denoted by alias `Z`, and the underlying type of
the defined type `W` are the same type.

```Go
package main

import (
	"fmt"	
	"reflect"
)

type Data[A int64 | int32, B byte | bool, C comparable] struct {
	a A
	b B
	c C
}

var x = Data[int64, byte, [8]byte]{1<<62, 255, [8]byte{}}
type Y = Data[int32, bool, string]
type Z = Data[int64, uint8, [8]uint8]
type W Data[int64, byte, [8]byte]

// The following line fails to compile because
// []uint8 doesn't satisfy the comparable constraint.
// type T = Data[int64, uint8, []uint8] // error

func main() {
	println(reflect.TypeOf(x) == reflect.TypeOf(Z{})) // true
	println(reflect.TypeOf(x) == reflect.TypeOf(Y{})) // false
	fmt.Printf("%T\n", x)   // main.Data[int64,uint8,[8]uint8]
	fmt.Printf("%T\n", Z{}) // main.Data[int64,uint8,[8]uint8]
}
```

The following is an example using some instantiated functions
of a generic function.

```Go
package main

type Ordered interface {
	~int | ~uint | ~int8 | ~uint8 | ~int16 | ~uint16 |
	~int32 | ~uint32 | ~int64 | ~uint64 | ~uintptr |
	~float32 | ~float64 | ~string
}

func Max[S ~[]E, E Ordered](vs S) E {
	if len(vs) == 0 {
		panic("no elements")
	}
	
	var r = vs[0]
	for i := range vs[1:] {
		if vs[i] > r {
			r = vs[i]
		}
	}
	return r
}

type Age int
var ages = []Age{99, 12, 55, 67, 32, 3}

var langs = []string {"C", "Go", "C++"}

func main() {
	var maxAge = Max[[]Age, Age]
	println(maxAge(ages)) // 99
	
	var maxStr = Max[[]string, string]
	println(maxStr(langs)) // Go
}
```

In the above example, the generic function `Max` is instantiated twice.

* The first instantiation `Max[[]Age, Age]` results in a `func([]Age] Age` function value.
* The second one, `Max[[]string, string]`, results in a `func([]string) string` function value.

{#type-argument-inferences}
## Type argument inferences for generic function instantiations

In the generic function example shown in the last section,
the two function instantiations are called full form instantiations,
in which all type arguments are presented in their containing type argument lists.
Go supports type inferences for generic function instantiations,
so that the type argument list of an instantiated function
may be partial or even be omitted totally,
as long as the missing type arguments could be inferred from value parameters
and presented type arguments.

For example, the `main` function of the last example in the last section could be rewritten as

```Go
func main() {
	var maxAge = Max[[]Age] // partial argument list
	println(maxAge(ages)) // 99
	
	var maxStr = Max[[]string] // partial argument list
	println(maxStr(langs)) // Go
}
```

A partial type argument list must be a prefix of the full argument list.
In the above code, the second arguments are both omitted,
because they could be inferred from the first ones.

If an instantiated function is called directly and some suffix type arguments
could be inferred from the value argument types, then the type argument list
could be also partial or even be omitted totally.
For example, the `main` function could be also rewritten as

```Go
func main() {
	println(Max(ages))  // 99
	println(Max(langs)) // Go
}
```

The rewritten `main` function shows that the calls to
generics functions could be as clean as ordinary functions (at least sometimes),
even if generics function declarations are more verbose.

Please note that, type argument lists may be omitted totally but may not be blank.
The following code is illegal.

```Go
func main() {
	println(Max[](ages))  // syntax error
	println(Max[](langs)) // syntax error
}
```

The above has mentioned that a partial type argument list must be a prefix of the full argument list.
This also means that the inferred type arguments in a type argument list belong to a suffix of the type argument list.
For example, the following code fails to compile.

```Go
package main

func foo[A, B, C any](v B) {}

func main() {
	// error: cannot use _ as value or type
	foo[int, _, bool]("Go")
}
```

Type arguments could be inferred from element types, field types,
parameter types and result types of value arguments.
For example,

```Go
package main

func luk[E any](v struct{x E}) {}
func kit[E any](v []E) {} 
func wet[E any](v func() E) {}

func main() {
	luk(struct{x int}{123})        // okay
	kit([]string{"go", "c"})       // okay
	wet(func() bool {return true}) // okay
}
```

If the type set of the constraint of a type parameter contains only one type
and the type parameter is used as a value parameter type in a generic function,
then compilers will attempt to infer the type of an untyped value argument
passed to the value parameter as that only one type. If the attempt fails,
then that untyped value argument is viewed as invalid.

For example, in the following program, only the first function call compiles.

```Go
package main

func foo[T int](x T) {}
func bar[T ~int](x T) {}

func main() {
	// The default type of 1.0 is float64, but
	// 1.0 is also representable as integer values.

	foo(1.0)  // okay
	foo(1.23) // error: cannot use 1.23 as int

	bar(1.0)  // error: float64 does not implement ~int
	bar(1.23) // error: float64 does not implement ~int
}
```

Type argument inferences rules and implementation are still in the process of continuous improvement,
from version to version. For example, the following code failed to compile before Go 1.21
but it compiles okay since Go 1.21 (`T` is inferred as `float64`).

```Go
func Less[T ~int | ~float64](x, y T) bool {
	return x < y
}

var _ = Less(1, 0.0) // only compiles since 1.21
```

Before Go 1.21, when using a generic function as a function value,
the type argument list of the instantiated function must be full.
Since Go 1.21, the type argument list is allowed to be partial if
some suffix type arguments could be inferred. For example

```Go
func Sum[T ~float32 | ~float64](x, y T) T {
	return x + y
}

// The following two lines compile okay only since Go 1.21.
var _ func(float32, float32) float32 = Sum
var _ func(float64, float64) float64 = Sum

// Before Go 1.21, the above two lines must be written as:
var _ func(float32, float32) float32 = Sum[float32]
var _ func(float64, float64) float64 = Sum[float64]
```

A complete example:

```Go
package main

import "fmt"

func Convert[S ~[]E, E, T any](s S, conv func(E) T) []T {
	r := make([]T, len(s))
	for i := range s {
		r[i] = conv(s[i])
	}
	return r
}

func isEven(n int) bool {
	return n&1 == 0
}

func intsToBools(s []int) []bool {
	// The following line only compiles since Go 1.21.
	var conv func([]int, func(int) bool) []bool = Convert
	// Before Go 1.21, to make the above line compile, its
	// right side must be written as Convert[[]int, int bool]
	
	return conv(s, isEven)
}

func main() {
	var s = []int{2, 3, 5, 8}
	fmt.Println(intsToBools(s)) // [true false false true]
}
```

Similarly, since Go 1.21, type arguments may be inferred
by analyzing method parameter types and result types.
For example, the following code failed to compile before Go 1.21
but it compiles okay since Go 1.21.


```Go
package main

type C[I, O any] interface {
	M(I) O
}

type Foo struct {
	v int
}

func (f Foo) M(n int) bool {
	return n == f.v
}

func bar[T C[I, O], I, O any](t T) {}

func main() {
	var foo Foo
	
	// 1.20:  error: cannot infer I. It must be written
	//        as bar[Foo, int, bool](foo) to compile.
	// 1.21+: compile okay.
	bar(foo)
}
```

A more complex example (it compiles okay since Go 1.21):

```Go
type Getter[T any] interface {
	Get() T
}

type Age[T uint8 | int16] struct {
	n T
}

func (a Age[T]) Get() T {
	return a.n
}

func doSomething[T any](g Getter[T]) T {
	return g.Get()
}

// The two lines only compiles since Go 1.21.
var _ = doSomething(Age[uint8]{255}) // error
var _ = doSomething(Age[int16]{256}) // error

// Before Go 1.21, to make the above two lines
// compile, they must be written as:
var _ = doSomething[uint8](Age[uint8]{255})
var _ = doSomething[int16](Age[int16]{256})
```

The book doesn't explain the detailed mechanisms involved in the inference process.
Please read [Go specification] and [this blog] for the details in type argument inferences.

[Go specification]: https://go.dev/ref/spec#Type_inference
[this blog]: https://go.dev/blog/type-inference

<!--
https://github.com/golang/go/issues/51139
-->

## Type argument inferences don't work for generic type instantiations

Currently (Go 1.22), inferring type arguments of instantiated types from value literals is not supported.
That means the type argument list of a generic type instantiation must be always in full form.

For example, in the following code snippet, the declaration line for variable `y` is invalid,
even if it is possible to infer the type argument as `int16`.

```Go
type Set[E comparable] map[E]bool

// compiles okay
var x = Set[int16]{123: false, 789: true}

// error: cannot use generic type without instantiation.
var y = Set{int16(123): false, int16(789): true}
```

Another example:

```Go
import "sync"

type Lockable[T any] struct {
	sync.Mutex
	Data T
}

// compiles okay
var a = Lockable[float64]{Data: 1.23}

// error: cannot use generic type without instantiation
var b = Lockable{Data: float64(1.23)} // error
```

It is unclear whether or not [type argument inferences
for generic type instantiations](https://github.com/golang/go/issues/50482)
will be supported in future Go versions.

## Differences between passing type parameters and ordinary types as type arguments

The above has mentioned that,

* when an ordinary type is used as a type argument, the ordinary type
  is only required to [satisfy][satisfies] the constraint of the corresponding
  type parameter. This is a change made in Go 1.20.
* when a type parameter is used as a type argument, the type parameter
  is required to [implement][implements] the constraint of the corresponding
  type parameter.

For example, since Go 1.20, all instantiations in the following program are valid
(they are invalid before Go 1.20).
None of the type arguments, which are all ordinary types,
implement the `comparable` constraints, but they all satisfy the constraints.

```Go
package main

type Base[T comparable] []T

type _ Base[any]
type _ Base[error]
type _ Base[[2]any]
type _ Base[struct{x any}]

func Gf[T comparable](x T) {}

var _ = Gf[any]
var _ = Gf[error]
var _ = Gf[[2]any]
var _ = Gf[struct{x any}]

func main() {
	Gf[any](123)
	Gf[error](nil)
	Gf[[2]any]([2]any{})
	Gf[struct{x any}](struct{x any}{})
}
```

On the other hand, all the instantiations in the following program are invalid.
All the type arguments are type parameters, and none of them implement
the `comparable` constraints, which is why they are all invalid type arguments.

```Go
package main

type Base[T comparable] []T

func Gf[T comparable](x T) {}

func vut[T any]() {
	_ = Gf[T]     // T does not implement comparable
	var _ Base[T] // T does not implement comparable
}

func law[T struct{x any}]() {
	_ = Gf[T]     // T does not implement comparable
	var _ Base[T] // T does not implement comparable
}

func pod[T [2]any]() {
	_ = Gf[T]     // T does not implement comparable
	var _ Base[T] // T does not implement comparable
}

func main() {}
```

All the type arguments used in the following program are valid.
They are all type parameters and implement the `comparable` constraints.


```Go
package main

type Base[T comparable] []T

func Gf[T comparable](x T) {}

func vut[T comparable]() {
	_ = Gf[T]
	var _ Base[T]
}

func law[T struct{x int}]() {
	_ = Gf[T]
	var _ Base[T] 
}

func pod[T [2]string]() {
	_ = Gf[T]
	var _ Base[T]
}

func main() {}
```
 
## More about instantiated types

As an instantiated type is an ordinary type,

* it may be used as union terms if it is an non-interface type.
* it may be used as union terms if it is an interface type without methods.
* it may be used as type constraints if it is an interface type.
* it may be embedded in struct types if it satisfies
  [the requirements for the embedded fields](https://go101.org/article/type-embedding.html).

The generic type declarations `C1` and `C2` in the following code are both valid.

```Go
package main

type Slice[T any] []T

type C1[E any] interface {
	Slice[E] // an ordinary type name
}

type C2[E any] interface {
	[]E | Slice[E] // okay
}

func foo[E any, T C2[E]](x T) {}

func main() {
	var x Slice[bool]
	var y Slice[string]
	foo(x)
	foo(y)
}
```

The following code shows an ordinary struct type declaration
which embeds two instantiations of the generic type `Set`.
To avoid duplicate field names, one of the embedded fields
uses an alias to an instantiation of the generic type.

```Go
package main

type Set[T comparable] map[T]struct{}

type Strings = Set[string]

type MyData struct {
	Set[int]
	Strings
}

func main() {
	var md = MyData {
		Set:     Set[int]{},
		Strings: Strings{},
	}
	md.Set[123] = struct{}{}
	md.Strings["Go"] = struct{}{}
}
```

## About the instantiation syntax inconsistency between custom generics and built-in generics

From previous contents, we could find that the instantiation syntax of Go custom generics
is inconsistent with Go built-in generics.

```Go
type TreeMap[K comparable, V any] struct {
	// ... // internal implementation
}

func MyNew[T any]() *T {
	return new(T)
}

// The manners of two instantiations differ.
var _ map[string]int
var _ TreeMap[string, int]

// The manners of two instantiations differ.
var _ = new(bool)
var _ = MyNew[bool]()
```

Personally, I think the inconsistency is pity and it increases the load of cognition burden in Go programming.
On the other hand, I admit that it is hard (or even impossible) to make the syntax consistent.
It is a pity that Go didn't support custom generics from the start.


<!--

infer function type parameters from generic interface arguments if necessary
https://github.com/golang/go/issues/52397

https://groups.google.com/g/golang-nuts/c/0yNl0Vz_9jA (there is an issue for this)..

-->
