
# More About generic functions and methods of generic types



## The type of a value used in a generic function must be either a specified type parameter, or a specified ordinary type

Go custom generics are not implemented as simple code text templates.
This is a principal differece from code generation. 
Every expression used in a generic function must have a specified type,
which is either an ordinary type, or a type parameter.

For example, in the following code snippet, the functions `foo` and `bar` both compile okay,
but the function `dot` fails to compile. The reason is simple:

* in the function `foo`, the type of `x` is `T`, which is a type parameter.
  Certainly, in uses if the function, `x` might be instantiated as `int` or `string`,
  but which doesn't change the fact that, from the view of compilers,
  its type is a type parameter.
* in the function `bar`, the type of `x[1]` is a specified ordinary type, `int`.
* in the function `dot`, the type of `x[1]` might be `int` or `string`.
  In the function, there is only one type parameter, `T`, which is surely not the type of `x[1]`.

```Go
func foo[T int|string](x T) {
	_ = x // okay
}

func bar[T [2]int|[8]int](x T) {
	_ = x[1] // okay
}

func dot[T [2]int|[2]string](x T) {
	_ = x[1] // error: invalid operation: cannot index x
}
```

For the same reason, in the following code snippet,
the functions `nop` and `jam` both compile okay,
but the function `mud` doesn't.

```Go
func nop[T *Base, Base int32|int64](x T) {
	_ = *x // okay
}

func jam[T int32|int64](x *T) {
	_ = *x // okay
}

func mud[T *int32|*int64](x T) {
	_ = *x // error: invalid operation
}
```

The same, in the following code snippet, only the function `box` fails to compile,
all the others compile okay.

```Go
func box[T chan int | chan byte](c T) {
	_ = <-c // error: no core type
}


func sed[T chan E, E int | byte](c T) {
	_ = <-c // okay
}

type Ch <-chan int
func win[T chan int | Ch](c T) {
	_ = <-c // okay
}

func cat[T []E, E any](c T, i int) () {
	_ = c[i] // okay
}
```

The rule talked about in this section is just a general rule.
For some operations, some extra restrictions might be applied.
Often, the extra restrictions are core types and specific types related.



## Type declarations inside generic functions are not allowed now (Go 1.18)

```Go
func foo[T any]() {
	type MyByte byte // error
}

type T[_ any] struct{}

func (T[_]) bar() {
	type MyInt int // error
}
```

https://golang.org/issue/47631

## Type parameters of constraints with empty type sets

could not be instantiated, but it compiles okay.

## const types can't be type parameter

func f[P int]() {
	const y P = 5 // 
}

func g[P int]() {
	const _ = P(5) // 
}

https://github.com/golang/go/issues/50202#issuecomment-1033369747

Converting a constant to a type parameter yields a non-constant value of
the argument pased to the type parameter.

```Go
package main

const S = "Go"

func f() byte {
	return 64 << len(string(S)) >> len(string(S))
}

func g[T string]() byte {
	return 64 << len(T(S)) >> len(T(S))
}

func main() {
	println(f(), g())
}
```

## local type declarations in generic code don't work

https://github.com/golang/go/issues/47631

## no generic method, use alternative

https://github.com/golang/go/issues/49085

https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md#methods-may-not-take-additional-type-arguments


This will make code even more cubersome.
And make `reflect.Method` ...

## Type parameters act

https://github.com/golang/go/issues/50421

```Go
func toString[T byte|rune](slice []T) string {
   return string(slice) // cannot convert slice (variable of type []T) to type string
}

func toString[T []byte|[]rune](slice T) string {
        return string(slice) // okay
}

func toString[T []E, E byte|rune](slice T) string {
        return string(slice) // cannot convert []E (in T) to string
}
```

## There are no predeclared `assignableTo` and `assignableFrom` constraints

## There are no predeclared `convertibleTo` and `convertibleFrom` constraints

Whether or not a conversion involving type parameters is legal is determined
by the specific type sets of the constraints (at most two) of involving type parameters.

unspecific conversions are not supported, there are not the `convertibleFrom` and `convertibleTo` constraints.

```Go
func Convert[From, To any](in []From, f func(From) To) []To {
	var out = make([]To, len(in))
	for i := range in {
		out[i] = f(in[i])
	}
	return out
}
```

```Go
func foo[T int](x *T) *int {
	// cannot convert x (variable of type *T) to type *int
	return (*int)(x)
}

```

```Go
func foo[T string](x T) string  {
	return x // error
}

func bar[T []byte](x T) []byte {
	return x // okay
}
```

```Go
func f[A, B int](x A, y B){
	x = y // error
	y = x // error
}
```

It is worth making this valid. Maybe later

## 

The current constraint design lacks of two abilities:
1. The ability of specifying a type argument must be an interface type.
2. The ability of specifying a type argument must not be an interface type.

https://groups.google.com/g/golang-nuts/c/EL6A2jFa92k

https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md#no-way-to-express-convertibility

https://gophers.slack.com/archives/C88U9BFDZ/p1647950715616299

// ConvertSlice converts each element of the slice
// as to the type To by doing a dynamic type conversion.
// Any elements of as that don't implement To will be
// omitted from the returned slice.
func ConvertSlice[To, From any](as []From) []To {
	bs := make([]To, 0, len(as))
	for _, a := range as {
		if b, ok := any(a).(To); ok {
			bs = append(bs, b)
		}
	}
	return bs
}

##

```Go
func f[A, B ~int](x A, y B){
	x = A(y)
	y = B(x)
}

func g[A, B int](x *A, y *B){
	x = (*A)(y) // error: cannot convert y
	y = (*B)(x) // error: cannot convert x
}
```

https://github.com/golang/go/issues/50815 pointer convert
https://github.com/golang/go/issues/51501 single type ...

## 


## Calls to predeclared functions

## len, cap

## Calls to the built-in `len` and `cap` functions with arguments of array type parameters always return non-constant results

For example, currently (Go 1.18), the first `cap` and `len` calls in the following code fail to compile.

```Go
func f[T [2]int]() {
	var x T
	const _ = cap(x) // error: cap(x) is not constant
	const _ = len(x) // error: len(x) is not constant

	var _ = cap(x) // okay
	var _ = len(x) // okay
	
	var y [2]int
	const _ = cap(y) // okay
	const _ = len(y) // okay
}
```

Similarly, calls to the built-in `len` and `cap` functions with arguments of array pointer type parameter also always return non-constant results.

```Go
func g[T *[2]int]() {
	var x T
	const _ = cap(x) // error: cap(x) is not constant
	const _ = len(x) // error: len(x) is not constant
}
```

This might be never changed in future Go versions: https://github.com/golang/go/issues/50226

## The built-in `len` and `cap` functions don't accept arguments of array pointer types whose base types are array type parameters

For example, currently (Go 1.18), the following function doesn't compile.

```Go
func h[T [2]int]() {
	var x T
	var _ = len(&x) // invalid argument: &x (value of type *T) for len
	var _ = cap(&x) // invalid argument: &x (value of type *T) for cap
}
```

The restriction might be removed in future Go versions, or not, I'm nor sure.

## make

## delete

## close

## A call to the built-in `close` function requires its argument has specific types and all the specific types are channels

## A call to the built-in `delete` function requires its argument has specific types and all the specific types are maps with identical key types

### Calles to predeclared `complex`, `real` and `imag` functions don't accept arguments of type parameter types

https://github.com/golang/go/issues/50937

func Real[T complex64](s T) float32 {
	return real(s)
}


## An element index operation require the container operand's specific types may not include maps and non-maps at the same time

And if all specific types are maps, them their underlying types must be identical;
Otherwise, their element types must be identical.
The elements of strings are viewed as `byte` values.

For example, currently (Go 1.18), in the following code snippet, only the functions `foo` and `tup` compile okay.

```Go
func foo[T []int | [2]int](c T) {
	_ = c[0] // okay
}

func bar[T []int | map[int]int](c T) {
	_ = c[0] // invalid operation: cannot index c
}

func vet[T map[string]int | map[int]int](c T) {
	_ = c[0] // invalid operation: cannot index c
}

func six[T map[string]int | map[int]int](c T) {
	_ = c[0] // invalid operation: cannot index c
}

type Map map[int]string
func tup[T map[int]string | Map](c T) {
	_ = c[0] // okay
}
```




## A (sub)slice operation requires the container operand has a core type

For example, currently (Go 1.18), the following function fails to compile.

```Go
func foo[T []int | [2]int](c T) {
	_ = c[:] // invalid operation: cannot slice c: T has no core type
}
```

There is an exception for this rule. If the container operand's specific types
only include string and byte slice types, then it is not required to have a core type.
For example, the following function compiles okay.

```Go
func bar[T string | []byte](c T) {
	_ = c[:] // okay
}
```

## A call to the built-in `make` function requires its first argument (container type) has a core type 

Currently (Go 1.18), in the following code snippet, the functions `foo` and `bar1` both
fail to compiler, the other two compile okay.
The reason is the first argument of a call to the built-in `make` function
is required to have a core type.
Neither of the `foo` and `bar1` functions satisfies this requirement,
whereas both of the other two functions satisfy this requirement.

```Go
func foo[T chan bool | chan int]() {
	_ = make(T) // error: invalid argument: no core type
}

func bar1[T chan<- int | <-chan int]() {
	_ = make(T) // error: invalid argument: no core type
}

type Stream chan int
type Queue Stream

func bar2[T Stream | chan int | Queue | chan<- int]() {
	_ = make(T) // okay
}

func bar3[T Stream | chan int | Queue | <-chan int]() {
	_ = make(T) // okay
}
```

By my understanding, this requirement is in order to make subsequent operations
on the made containers (they are channels in the above example) always legal.
For example, to prevent make sure a value received from the made
channel has a specified type (either a type parameter, or an ordinary type).

Personally, I think the requirement is over strict.
After all, the assumed subsequent operations might not happen for many use cases
(such as he functions `foo_a` and `bar1_a` below), and the containers may present as (value) parameters
even if it has not a core type, as the following example shows:

```Go
func g(any) {}

func foo_a[T chan bool | chan int](x T) {
	g(x)
}

func bar1_a[T chan<- int | <-chan int](x T) {
	g(x)
}
```

Because of the same requirement, neither of the following three functions compile.

```Go
func zig[T ~[]int | map[int]int](c T) {
	_ = make(T) // error: invalid argument: no core type
}

func fat[T ~[]int | ~[]bool](c T) {
	_ = make(T) // error: invalid argument: no core type
}
```

Calls to the built-in `new` function have not this requirement.

## The type literal in a composite literal must have a core type

This restriciton is smimilar to the last one.
For example, currently (Go 1.18), in the following code snippet,
the functions `foo` and `bar` compile okay,
but the other ones don't.

```Go
func foo[T ~[]int] () {
	_ = T{}
}

type Ints []int

func bar[T []int | Ints] () {
	_ = T{}
}

func ken[T []int | []string] () {
	_ = T{} // error: invalid composite literal type T
}

func jup[T [2]int | map[int]int] () {
	_ = T{} // error: invalid composite literal type T
}
```

## In a `for-range` loop, the ranged container is required to have a core type

For example, currently (Go 1.18), in the following code, 
only the last two functions, `dot1` and `dot2` compile okay.

```Go
func values[T []E | map[int]E, E any](kvs T) []E {
	r := make([]E, 0, len(kvs))
	for _, v := range kvs { // cannot range over kvs (T has no core type)
		r = append(r, v)
	}
	return r
}

func keys[T map[int]string | map[int]int](kvs T) []int {
	r := make([]int, 0, len(kvs))
	for k := range kvs { // cannot range over kvs (T has no core type)
		r = append(r, k)
	}
	return r
}

func sum[M map[int]int | map[string]int](m M) (sum int) {
	for _, v := range m {
		sum += v
	}
	return
}

func foo[T []int | []string] (v T) {
	for range v {} // error: cannot range over v (T has no core type)
}

func bar[T [3]int | [6]int] (v T) {
	for range v {} // error: cannot range over v (T has no core type)
}

type MyInt int

func cat[T []int | []MyInt] (v T) {
	for range v {} // error: cannot range over v (T has no core type)
}

type Slice []int

func dot1[T []int | Slice] (v T) {
	for range v {} // okay
}

func dot2[T ~[]int] (v T) {
	for range v {} // okay
}
```

Need a cire type.
```Go
func f[T []int | map[int]int] (t T, g func(int)) {
	for _, v := range t { // error
	g(v)
	}
}
```

The restriction is intended. I think its intention is to ensure both of the two iteration variables
always have a specified type (either an ordinary type or a type parameter type).
Howwver, this restriction is over strict for this intention.
Becease, in practice, the key types or element types of some containers are identical,
even if the underlying type of the containers are different.
And in many use cases, either of the two iteration variables is ignored.

I'm not sure whether or not the restriction will be removed in future Go versions: https://github.com/golang/go/issues/49551.
In my opinion, the restriciton reduces the usefulness of Go custom generics in some extent.

If all possible types are slice and arrays, and their element types are identical,
we could use plan `for` loops to walk around this restriction.

```Go
func cat[T [3]int | [6]int | []int] (v T) {
	for i := 0; i < len(v); i++ { // okay
		_ = v[i] // okay
	}
}
```

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

The conversion `[]byte(v)` (if it follows the `range` keyword) is specifically
optimization by the official standard Go compiler so that it doesn't duplicate
underlying bytes.

The following code doesn't compile now (Go 1.18).
Whether or not it will compile later is unknown.

```Go
func PrintEach[T string | []rune](runes T) {
	for _, r := range runes { // cannot range over runes
		_ = r
	}
}
```

<!--
https://github.com/golang/go/issues/51053
-->

## A function is required to have a core type to be callable

For example, currently (Go 1.18), in the following code, the functions `foo` and `bar` don't compile,
bit the `dot` function does.

```Go
func foo[F func(int) | func(any)] (f F) {
	f(1) // error: invalid operation: cannot call non-function f
}

func bar[F func(int) | func(int)int] (f F) {
	f(1) // error: invalid operation: cannot call non-function f
}

type Fun func(int)

func dot[F func(int) | Fun] (f F) {
	f(1) // okay
}
```

Not sure whether or not the restriction will be lifted in future Go versions.

<!--

https://github.com/golang/go/issues/50285 inference from results (not a good idea, and may not be get supported for ever)

-->

## Compile-time type switch

https://github.com/golang/go/issues/45380#issuecomment-1074153465



