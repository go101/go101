# Constraints and Type Parameters

A constraint means a type constraint, it is used to constrained some type parameters.
We could view constraints as types of types.

The relation between a constraint and a type parameter is like
the relation between a type and a value.
If we say types are value templates (and values are type instances),
then constraints are type templates (and types are constraint instances).

A type parameter is a type which is declared in a type parameter list
and could be used in a generic type specificaiton or a generic function/method declaration.
Each type parameter is a distinct named type.

Since Go 1.18, value types in Go could be categorized in two categories:

* type parameter types: the types declared in type parameter lists.
* ordinary types: the types not declared in type parameter lists.
  Before Go 1.18, there are only ordinary types.

Type parameter lists will be explained in detail in a later section.

As mentioned in the previous chapter, type constraints are actually
[interface types](https://go101.org/article/interface.html).
In order to let interface types be competent to act as the constraint role,
Go 1.18 enhances the expressiveness of interface types by supporting several new notations.

## Enhanced interface syntax 

Some new notations are introduced into Go to make it is possible to use interface types as constraints.

* The `~T` form, where `T` is a type literal or type name.
  `T` must denote a non-interface type whose underlying type is itself
  (so `T` may not be a type parameter, which is explained below).
  The form denotes a type set, which include all types whose
  [underlying type](https://go101.org/article/type-system-overview.html#underlying-type) is `T`.
  The `~T` form is called a tilde form or type tilde in this book
  (or underlying term and approximation type elsewhere).
* The `T1 | T2 | ... | Tn` form, which is called a union of terms (or type/term union in this book).
  Each `Tx` term is a tilde form, type literal, or type name,
  and it may not denote a type parameter.
  There are some restrictions of using union terms.
  These restrictions will be described in a section below.

Note that, a type literal always denotes an unnamed type,
whereas a type name may denote a named type or unnamed type.

Some legal examples of the new notations:

```Go
// tilde forms
~int
~[]byte
~map[int]string
~chan struct{}
~struct{x int}

// unions of terms
uint8 | uint16 | uint32 | uint64
~[]byte | ~string
map[int]int | []int | [16]int | any
chan struct{} | ~struct{x int}
```

We know that, before Go 1.18, an interface type may embed

* arbitrary number of method specifications (method elements, one kind of interface elements);
* arbitrary number of type names (type elements, the other kind of interface elements),
  but the type names must denote interface types.

Go 1.18 relaxed the limitaitons of type elements, so that now an interface type
may embed the following type elements:

* any type literals or type names, whether or not they denote interface types, but they must not denote type parameters.
* tilde forms.
* term unions.

The orders of interface elements embedded in an interface type are not important.

The following code snippet shows some interface type declarations,
in which the interface type literals in the declarations of `N` and `O`
are only legal since Go 1.18.

```Go
type L interface {
	Run() error
	Stop()
}

type M interface {
	L
	Step() error
}

type N interface {
	M
	interface{ Resume() }
	~map[int]bool
	~[]byte | string
}

type O interface {
	Pause()
	N
	string
	int64 | ~chan int | any
}
```

Embedding an interface type in another one is equivalent to (recursively) expanding the elements in the former into the latter. In the above example, the declarations of `M`, `N` and `O` are equivalent to the following ones:

```Go
type M interface {
	Run() error
	Stop()
	Step() error
}

type N interface {
	Run() error
	Stop()
	Step() error
	Resume()
	~map[int]bool
	~[]byte | string
}

type O interface {
	Run() error
	Stop()
	Step() error
	Pause()
	Resume()
	~map[int]bool
	~[]byte | string
	string
	int64 | ~chan int | any
}
```

We could view a single type literal, type name or tilde form as a term union with only one term.
So simply speaking, since Go 118, an interface type may specify some methods and embed some term unions.

An interface type without any embedding elements is called an empty interface.
For example, the predeclared `any` type alias denotes an empty interface type.

## Type sets and method sets

Before Go 1.18, an interface type is defined as a method set.
Since Go 1.18, an interface type is defined as a type set.
A type set only consists of non-interface types.

* The type set of a non-interface type literal or type name only contains the type denoted by the type literal or type name.
* As just mentioned above, the type set of a tilde form `~T` is the set of types whose underlying types are `T`. In theory, this is an infinite set.
* The type set of a method specification is the set of non-interface types whose method sets include the method specification.
  In theory, this is an infinite set.
* The type set of the empty interface is the set of all non-interface types.
  In theory, this is an infinite set.
* The type set of a union of terms `T1 | T2 | ... | Tn` is the union of the type sets of the terms.
* The type set of a non-empty interface is the intersection of the type sets of its interface elements.

As the type set of an empty interface type contains all non-interface types.
It is a super set of any type set.

By the current specification,
two unnamed constraints are equivalent to each other if their type sets are equal.

Given the types declared in the following code snippet,
for each interface type, its type set is shown in its preceding comment.

```Go
type Bytes []byte  // underlying type is []byte
type Letters Bytes // underlying type is []byte
type Blank struct{}
type MyString string // underlying type is string

func (MyString) M() {}
func (Bytes) M() {}
func (Blank) M() {}

// The type set of P only contians one type:
// []byte.
type P interface {[]byte}

// The type set of Q contains
// []bytes, Bytes, and Letters.
type Q interface {~[]byte}

// The type set of R contains only two types:
// []byte and string.
type R interface {[]byte | string}

// The type set of S is empty.
type S interface {R; M()}

// The type set of T contains:
// []byte, Bytes, Letters, string, and MyString.
type T interface {~[]byte | ~string}

// The type set of U contains:
// MyString, Bytes, and Blank.
type U interface {M()}

// V <=> P
type V interface {[]byte; any}

// The type set of W contains:
// Bytes and MyString.
type W interface {T; U}

// Z <=> any. Z is a blank interface. Its
// type set contains all interface types.
type Z interface {~[]byte | ~string | any}
```

Please note that interface elements are seperated with semicolon (`;`),
either explicitly or implicitly (Go compilers will
[insert some missing semicolons as needed in compilations](https://go101.org/article/line-break-rules.html)).
The following interface type literals are equivalent to each other.
The type set of the interface type denoted by them is empty.
The interface type and the underlying type of the type `S`
shown in the above code snippet are actually identical.

```Go
interface {~string; string; M();}
interface {~string; string; M()}
interface {
	~string
	string
	M()
}
```

If the type set of a type `X` is a subset of an interface type `Y`,
we say `X` implements (or satisfies) `Y`.
Here, `X` may be an interface type or a non-interface type.

All types implement the predeclared `any` type.

In the above example,

* the interface type `S`, whose type set is empty, implements any interface types.
* all types implement the interface type `Z`, which is identical to the predeclared `any` type.

When a type argument is passed to a type parameter,
the type argument must implements (or satisfies) the constraint of the type parameter.
In other words, type sets are used to check the validity of type arguments.

The list of methods specified by an interface is called the method set of the interface.
If an interface type `X` implements another interface type `Y`, then 

* the method set of `X` must be a superset of `Y`, and
* the type set of `X` must be a subset of `Y`.

Interfaces whose type sets can be defined entirely by a method set (may be empty)
are called basic interfaces.
Go versions before 1.18 only support basic interfaces.
Basic interfaces may be used as either value types or type constraints,
but non-basic interfaces may only be used as type constraints (as of Go 1.18).

In the above examples, `L`, `M`, `U`, `Z` and `any` are basic types.

In the following code, the declaration lines for `x` and `y` both compile okay, but the line for `z` fails to compile.

```Go
var x any
var y interface {M()}

// error: interface contains type constraints
var z interface {~[]byte}
```

Using no-baisc interface types as value types might be supported in later Go versions.

BTW, currently (Go 1.18), there is an unintended restriction: non-basic interface types may not be aliased.
For example, the following type alais declarations are illegal:

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

The unintended restriction will [be removed in Go 1.19](https://github.com/golang/go/issues/51616).

## More about the predeclared `comparable` constraint

As aforementioned, besides `any`, Go 1.18 introduces another new predeclared identifier `comparable`,
which denotes an interface type that is implemented by all comparable types.

The `comparable` interface could be embedded in other interface types to filter out incomprable types from their type sets.
For example, the type set of the following declared constraint `C` contains only one type: `string`.

```Go
type C interface {
	comparable
	[]byte | string | func() | map[int]bool
}
```

Currently (Go 1.18), the `comparable` interface is treated as a non-basic interface type.
So, now, it may only be used as type parameter constraints, not as the type of values.
The following code is illegal:

```Go
var x comparable = 123
```

The type set of the `comparable` interface is the set of all comparable types.
The set is a subset of the type set of the `any` interface,
so `comparable` undoubtedly implements `any`, and not vice versa.

On the other hand, starting from Go 1.0, all basic interface types are treated as comparable types.
The blank interface type `any` is not an exception.
So it looks that `any` (as a value type) should satisfy (implement) the `comparable` constraint.
This is odd.

After deliberation, Go core team believe that
[it is a design flaw](https://github.com/golang/go/issues/50646#issuecomment-1023706545)
to treat all interface types as comparable types and it is a pity that
the `comparable` type has not been supported since Go 1.0 to avoid this flaw.

Go core team try to make up for this flaw in Go custom generics age.
So they decided that `any` (as a value type) doesn't satisfy (implement) the `comparable` constraint.
A consequence of this decision is that there is not an interface type
(as a value type) satisfying the `comparable` constraint,
[which causes diffculties to some code designs](https://github.com/golang/go/issues/51257).

To avoid the consequence, a proposal has been made to
[permit using `comparable` as a value type](https://github.com/golang/go/issues/51338).
Whether or not it should be accepted is still under discuss.
It could be accepted in as earlier as Go 1.19.
That means the consequence exists in Go 1.18 for sure.

Another benefit brought by the proposal is that it provides a way to
ensure some interface comparisons will never panic.
For example, calls to the following function might panic at run time:

```Go
func foo(x, y any) bool {
	return x == y
}

var _ = foo([]int{}, []int{}) // panics
```

If the `comparable` type could be used as a value type,
then we could change the parameter types of the `foo` function
to `comparable` to ensure the calls to the `foo` function will never panic.

```Go
func foo(x, y comparable) bool {
	return x == y
}

var _ = foo([]int{}, []int{}) // fails to compile
```

## Some restrictions of using union terms

The above has mentioned that a union term may not be a type parameter. There are two other restrictions.

The first is an implementation specific restriction: a term union with more than one term cannot contain the predeclared identifier `comparable` or interfaces that have methods. 
For example, the following term unions are all illegal (as of Go toolchain 1.18):

```Go
[]byte | comprable
string | error
```

Another restriction is that the type sets of all non-interface type terms in a term union must have intersections.
Interface type terms have no this restriction, but the current implementation (Go toolchain 1.18) disallows identical interface type terms. For example, in the following code snippet, the term unions in the first two type declarations fail to compile, but the last two compile okay.

```Go
type _ interface {
	int | ~int // error
}

type _ interface {
	interface{int} | interface{int} | interface{~int} // error
}

type _ interface {
	interface{int} | interface{~int} // okay
}

type _ interface {
	int | interface{~int} // okay
}
```

The four term unions in the above code snippet are equivalent to each other in logic,
which means this restriction is not very reasonable.
So it might be removed in later Go versions, of become stricter to defeat the walkaround.

<!--
https://github.com/golang/go/issues/51607
https://github.com/golang/go/issues/45346#issuecomment-862505803
-->

## The specific type set (possibly empty) and the core type (possibly non-existing) of a constraint

The two properties of a constraint control the behaviors/operations supported by values of type parameters of the constraint.
In other words, the two properties are used to check the validity of code in function/method bodies.

In calcualting the specific type set of a constriant,
the following interface elements are called unspecified elements and are ignored:

* method elements.
* type literals or type names which denote the empty interface.
* type unions which contain an empty interface term.

The specific type sets of other specified elements are described below.

* The specific type set of a non-interface type literal or type name only contains the type denoted by the type literal or type name.
* The specific type set of a tilde form `~T` only contains `T`.
* The type set of a union of terms `T1 | T2 | ... | Tn`,
  where every `Tx` is not an empty interface type term,
  is the union of the specific type sets of the terms.

The specific type set of a constraint is the intersection of the specific type sets of its specified elements.

The specific type set of a basic interface is always empty.

A constraint has a core type if its specific type set is not empty and

1. all specific types of the constraint have the same underlying type.
1. or all specific types are channel types with an identical element type `E`
   and all directional channels have the same direction.

For the first case, the core type of the constraint is the same underlying type.

For the second case, the core type of the constraint is the type `chan<- E` or `<-chan E`,
depending on the direction of the directional channels.

For other cases, the constraint has not a core type.

In the following code snippet, the type set, the specific type set
and the core type of each constraint are listed in the comment
of that constraint. Only the types presenting in the code are listed.

```Go
type Age int
type Num int

func (Age) M() {}

// Type set:       Age
// Specific types: Age
// Core type:      int
type _ interface{Age}

// Type set:       Age, Num
// Specific types: Age, Num
// Core type:      int
type _ interface{Age | Num}

// Type set:       int, Age, Num
// Specific types: int
// Core type:      int
type _ interface{~int}

// Type set:       (empty)
// Specific types: Num
// Core type:      int
type _ interface{Num; M()}

// Type set:       int, bool
// Specific types: int, bool
// Core type:      (no core type)
type _ interface{int | bool}

// Type set:       Age
// Specific types: Age
// Core type:      int
type _ interface{Age | bool; ~int | string; M()}

// Type set:       <-chan int, chan<- int
// Specific types: <-chan int, chan<- int
// Core type:      (no core type)
type _ interface{<-chan int | chan<- int}

// Type set:       chan int, chan<- int
// Specific types: chan int, chan<- int
// Core type:      chan<- int
type _ interface{chan int | chan<- int}
```

## Type parameter lists

From the examples shown in the last chapter, we know type parameter lists
are used in generic type specifications, method declarations for generic base types
and generic function declarations.

A type parameter list contains at least one type parameter declaration
and is enclosed in square brackets.
Each parameter declaration is composed of a name portion and a constraint portion
(we can think the constraints are implicit in method declarations for generic base types).
Parameter declarations are comma-seperated in a type parameter list.

In a type parameter list, all type parameter names must be present.
They may be the blank identifier `_` (called blank name).
All non-blank names in a type parameter list must be unique.

Similar to value parameter lists, if the constraint portions of
some successive type parameter declarations in a type parameter list are identical,
then these type parameter declarations could share a common
constraint portion in the type parameter list.
For example, the following two type parameter lists are equivalent.

```Go
[A any, B any, _ comparable, _ comparable]
[A, B any, _, _ comparable]
```

If the end `]` token of a type parameter list is put on a new line,
then a comma is required (by [Go line break rules](https://go101.org/article/line-break-rules.html))
to follow the last constraint in the type parameter list;
otherwise, the comma is optional.

In the following code, the beginning lines are legal, the ending lines are not.

```Go
// Legal ones:
[T interface{~map[int]string}]
[T interface{~map[int]string},]
[T interface{~map[int]string},
]
[A, B any, _, _ comparable]
[A, B any, _, _ comparable,]
[A, B any, _, _ comparable,
]
[A, B any,
_, _ comparable]

// Illegal ones:
[A, B any, _, _ comparable
]
[T interface{~map[int]string}
]
```

Variadic type parameters are not supported.

## Simplified constraint form

In a type parametr list, if a constraint only contains one element
and that element is a type element,
then the enclosing `interface{}` may be omitted for convenience.
For example, the following two type parameter lists are equivalent.

```Go
[X interface{string|[]byte}, Y interface{~int}]
[X string|[]byte, Y ~int]
```

The simplified constraint forms make code look much cleaner.
For most cases, they don't cause any problems.
However, it might cause parsing ambiguities for some special cases.
In particular, parsing ambiguities might arise when the type parameter list
of a generic type specification declares a single type parameter
which constraint presents in simplified form and starts with `*` or `(`.

<!--
Either the spec is not accurate or the implementaiton is still not perfect yet.

type bar[T **string] struct{} // *string (type) is not an expression
-->

For example, does the following code declare a generic type?

```Go
type G[T *int] struct{}
```

It depends on what the `int` identifier denotes.
If it denotes a type (very possible, not absolutely),
then compilers will think the code declares a generic type.
If it denotes a constant (it is possible), then compilers
will treat `T *int` as a multiplication expression and
think the code declares an ordinary array type.

It is possible for compilers to distinguish what the `int` identifier denotes,
but there are some costs to achieve this. To avoid the costs,
compilers always treat the `int` identifier as a value expression
and think the above declaration is an ordinary array type declaration.
So the above declaration line will fail to compile if
either `T` or `int` is undefined or couldn't denote an integer constant.

Then how to declare a generic type with a single type parameter with `*int` as the constraint?
There are two ways to accomplish thiis:

1. use the full constraint form, or
1. let a comma follow the simplified constraint form.

The two ways are shown in the following code:

```Go
// Assume int is a predeclared type.
type G[T interface{*int}] struct{}
type G[T *int,] struct{}
```

The two ways shown above are also helpful for
some other special cases which might cause parsing ambiguities.
For example,

```Go
// PA might be array pointer variable, or a type name.
// Compilers don't treat it as a type name.
type K[cap (*PA)] struct{}

// S might be a string constant, or a type name.
// Compilers don't treat it as a type name.
type L[len (S)] struct{}
```

The following is another case which might cause parsing ambiguity.

```Go
// T, int and bool might be three constant integers,
// or int and bool are both predeclared types.
type C5[T *int|bool] struct{}
```

As of Go toolchain 1.18, inserting a comma after the presumed constraint `*int|bool` doesn't work
(It is [a bug](https://github.com/golang/go/issues/51488)
in Go toolcain 1,18 and will be fixed in Go toolchain 1.19).

Now, we could use full constraint form or exchange the places of `*int` and `bool` to make it compile okay.

```Go
// Assume int and bool are predeclared types.
type C5[T interface{*int|bool}] struct{}
type C5[T bool|*int] struct{}
```

On the other hand, the following two weird generic type declarations are both legal.

```Go
// "make" is a declard type parameter.
// Its constraint is interface{chan int}.
type PtrToChan[make (chan int)] *make

// "new" is a declard type parameter.
// Its constraint is interface{[3]float64}.
type Matrix33[new ([3]float64)] [3]new
```

The two declarations are really bad practices. Don't use them in serious code.

<!--
https://github.com/golang/go/issues/49482
https://github.com/golang/go/issues/49485
https://github.com/golang/go/issues/51488
-->

## Each type parameter is a distinct named type and its underlying type is an interface type

Since Go 1.18, named types include

* predeclared types, such as `int`, `bool` and `string`.
* defined non-generic types.
* instantiated types of generic types.
* type parameter types (the types declared in type parameter lists).

Two type parameters are never identical.

The type of a type parameter is a constraint, a.k.a an interface type.
This means the underlying type of a type parameter type is an interface type.
However, this doesn't mean a type parameter behaves like an interface type.
Its values may not be type asserted and box non-interface values.

In fact, a type parameter is just a placeholder for the types in its type set.
So it generally behaves as (the common traits of) the types in its type set in many situations.

As the underlying type of a type parameter type is not the type parameter type itself,
the tilde form `~T` is illegal if `T` is type parameter.
So the following type parameter list is illegal.
Because, as mentioned above, the type in a tilde form mustn't be an interface type
and its underlying type must be itself. Here the both conditions are not satisfied.

```Go
[A int, B ~A] // error
```

For the same reason, the following generic type declarations is also illegal.

```Go
type C[T int] interface {
	~T // error: cannot embed a type parameter
}
```

## Composite type literals containing type parameters are ordinary types

For example, `*T` is always an ordinary (pointer) type.
It is a type literal, so its underlying type is itself, whether or not `T` is a type parameter.
The following type parameter list is legal.

```Go
[A int, B ~*A] // okay
```

For the same reason, the following type parameter lists are also legal.

```Go
[T ~string|~int, A ~[2]T, B ~chan T]            // okay
[T comparable, M ~map[T]int32, F ~func(T) bool] // okay
```

{#type-parameter-scope}
## The scopes of a type parameters

The following type parameter list is valid, even if the use of `E` is ahead of the declaration of `E`.
The type parameter `E` is used in the constraint of the type parameter `S`,

```Go
[S ~[]E, E int]
```

Please note,

* as mentioned in the last section, although `E` is a type parameter type, `[]E` is an ordinary (slice) type.
* the undrelying type of `S` is `interface{~[]E}`, not `[]E`.
* the undrelying type of `E` is `interface{int}`, not `int`.

For ordinary function and method declarations, a (value) parameter/result name
is allowed to be the same as a parameter/result type name.
For example, the following funciton and method declarations are valid.

```Go
type C int

func foo(C C) {}

func (C C) Bar() {}
```

The scope of a type parameter of a generic function or a method of a generic type
also include the function/method body and value parameter/result lists.
Simply speaking, type parameters and value parameters/results are all declared in 
the top block of the function/method body.

This means the generic function declarations and method declarations for generic types
in the following code snippet all fail to compile (as of Go 1.18).

```Go
type C any
func foo1[C C]() {}    // error: C redeclared
func foo2[T C](T T) {} // error: T redeclared

type G[E any] struct{x E}
func (E G[E]) Bar1() {} // error: E redeclared
func (v G[G]) Bar2() {} // error: G is not a generic type
```

The `Bar2` method declaraiton might become legal
[since a later Go version](https://github.com/golang/go/issues/51503).

<!--
https://github.com/golang/go/issues/51503
-->

## Generic type/function instantiations and type argument inference

(need interface argument examples ...)

Generic types must be instantiated to be used as types of values.
Generic functions must be instantiated to be called or used as function values.

A generic function (type) is instantiated by substituting a type argument list
for the type parameter list of its declaration (specification).
The lengths of the type argument is the same the type parameter list.
We say each type argument is passed to the corresponding type parameter.
A type argument must be a non-generic typem and it is valid only if
it satisfies the constraint of its corresponding type parameter.

Instantiated functions are non-generic functions. 
Instantiated types are non-generic named types.

Same as type parameter lists, a type argument list is also enclosed in square brackets
and type arguments are also comma-seperated in the type argument list.

Two type argument lists are identical if their lengths are equal and all of their corresponding types are identical.
Two instantiated types are identical if they are instantiated from the same generic type and with the same type argument lists.

In the following program, the generic type `Data` is instantiated four times.
Three of the four instantiations have the same type argument list
(please note that the predeclared `byte` is an alais of the predeclared `uint8` type).
The type `W` is a defined type, which underlying type is `Z`, which is the type of the variable `x`.


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

// type T = Data[int64, uint8, []uint8] // fails to compile

func main() {
	println(reflect.TypeOf(x) == reflect.TypeOf(Z{})) // true
	println(reflect.TypeOf(x) == reflect.TypeOf(Y{})) // false
	println(reflect.TypeOf(x) == reflect.TypeOf(W{})) // false
	fmt.Printf("%T\n", x)   // main.Data[int64,uint8,[8]uint8]
	fmt.Printf("%T\n", Z{}) // main.Data[int64,uint8,[8]uint8]
	fmt.Printf("%T\n", W{}) // main.W
}
```

The following is an example using instantiated functions.

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

func main() {
	{
		type MyInt int
		var values = []MyInt {99, 12, 55, 67, 32, 3}
		var maxInt = Max[[]MyInt, MyInt]
		var x = maxInt(values)
		var y = Max[[]MyInt, MyInt](values)
		println(x, y) // 99 99
	}
	{
		var values = []string {"C", "Go", "C++"}
		var maxString = Max[[]string, string]
		var x = maxString(values)
		var y = Max[[]string, string](values)
		println(x, y) // Go Go
	}
	
	// var _ = Max[[]bool, bool] // doesn't compile
}
```

In the above example, the generic function `Max` is instantiated four times.
The instantiations result in two different non-generic functions.

Benefiting from great type inference support,

* a type argument list may be partial (but may not be empty),
as long as all type arguments could be inferred from the partial type argument list.
In a partial type argument list, left argument are kept and right ones are omitted.
* For a call to an instantiated function, the type argument list could be even totally omitted,
as long as all type arguments could be inferred from the value parameters of the call.

For example, the instantiations in the last example may be simplified as the following code shows.

```Go
func main() {
	{
		type MyInt int
		var values = []MyInt {99, 12, 55, 67, 32, 3}
		var maxInt = Max[[]MyInt] // the 2nd arg is omitted
		var x = maxInt(values)
		var y = Max(values) // the argument list is omitted
		println(x, y) // 99 99
	}
	{
		var values = []string {"C", "Go", "C++"}
		var maxString = Max[[]string] // the 2nd arg is omitted
		var x = maxString(values)
		var y = Max(values) // the argument list is omitted
		println(x, y) // Go Go
	}
}
```

In the above code snippet,

* the second type arguments for the type parameter `E` could be inferred from the partical
  type argument lists `[[]MyInt]` and `[[]string]`. They are inferred as the element types of
  the first type arguments, `MyInt` and `string`.
* the first type arguments for the type parameter `S` could be inferred from the value argument types
  `[]MyInt` and `[]string`. They are inferred as the value argument types.
  The consequently, the second type arguments could be inferred from the first type arguments.

If the type set of the constraint of a type parameter contains only one type
and the type parameter is used as a value parameter type in a generic function,
then compilers will attempt to infer the type of an untyped value argument
passed to the value parameter as that only one type. If the attempt fails,
then that untyped value argument is viewed as invalid.
For example, in the following program, the first two function calls
in the `main` function compile okay, but the last two fail to compile.

```Go
package main

func foo[T int](x T) {}
func bar[T ~float64](x T) {}

func main() {
	foo(1)          // okay
	bar[float64](1) // okay
	
	foo(1.23) // error
	bar(1)    // error
}
```

## Restrictions in the current type argument inference design and implementation

Currrenty (Go 1.18), inferring type parameters of generic types from value literals is supported limitedly.
For example, in the following code snippet, the declaration line for variable `x` is valid,
but the declaration line for variable `y` is invalid.

```Go
type Set[E comparable] map[E]struct{}

// compiles okay
var x = Set[int16]{123: struct{}{}, 789: struct{}{}}

// cannot use generic type Set[E comparable] without instantiation.
var y = Set{int16(123): struct{}{}, int16(789): struct{}{}}
```

<!--
https://github.com/golang/go/issues/50484
https://github.com/golang/go/issues/41176
-->

```Go
func foo[A, B, C any](v B) {}

var _ = foo[int, _, bool]("Go") // error: cannot use _ as value or type
```

Currently (Go 1.18), too-complicated type argument inference is not supported.

<!--
https://github.com/golang/go/issues/50482
-->

It looks value arguments are scanned firstly, 

```Go
func Append[S ~[]T, T any](s S, x ...T) S { /* implementation of append */ return s }

func _() {
        type MyPtr *int
        var x []MyPtr
        _ = append(x, new(int))  // built-in append: ok
        _ = Append(x, new(int))  // okay
        _ = Append[[]MyPtr](x, new(int))         // []MyPtr does not implement ~[]*int
        _ = Append[[]MyPtr, MyPtr](x, new(int))  // []MyPtr does not implement ~[]*int
}
```

<!--
https://github.com/golang/go/issues/51139
-->