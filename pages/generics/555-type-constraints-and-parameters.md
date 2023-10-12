# Constraints and Type Parameters

A constraint means a type constraint, it is used to constrain some type parameters.
We could view constraints as types of types.

The relation between a constraint and a type parameter is like
the relation between a type and a value.
If we say types are value templates (and values are type instances),
then constraints are type templates (and types are constraint instances).

A type parameter is a type which is declared in a type parameter list
and could be used in a generic type specification or a generic function/method declaration.
Each type parameter is a distinct named type.

Type parameter lists will be explained in detail in a later section.

As mentioned in the previous chapter, type constraints are actually
[interface types](https://go101.org/article/interface.html).
In order to let interface types be competent to act as the constraint role,
Go 1.18 enhances the expressiveness of interface types by supporting several new notations.

## Enhanced interface syntax

Some new notations are introduced into Go to make it possible to use interface types as constraints.

* The `~T` form, where `T` is a type literal or type name.
  `T` must denote a non-interface type whose underlying type is itself
  (so `T` may not be a type parameter, which is explained below).
  The form denotes a type set, which include all types whose
  [underlying type](https://go101.org/article/type-system-overview.html#underlying-type) is `T`.
  The `~T` form is called a tilde form or type tilde in this book
  (or underlying term and approximation type elsewhere).
* The `T1 | T2 | ... | Tn` form, which is called a union of terms (or type/term union in this book).
  Each `Tx` term must be a tilde form, a type literal, or a type name,
  and it may not denote a type parameter.
  There are some requirements for union terms.
  These requirements will be described in a section below.

<!--
https://github.com/golang/go/issues/52391
-->

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

Go 1.18 relaxed the limitations of type elements, so that now an interface type
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

Embedding an interface type in another one is equivalent to (recursively) expanding the elements of the former into the latter. In the above example, the declarations of `M`, `N` and `O` are equivalent to the following ones:

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
So simply speaking, since Go 1.18, an interface type may specify some methods and embed some term unions.

An interface type without any embedding elements is called an empty interface.
For example, the predeclared `any` type alias denotes an empty interface type.

{#type-sets-and-implementations}
## Type sets and type implementations

Before Go 1.18, an interface type is defined as a method set.
Since Go 1.18, an interface type is defined as a type set.
A type set consists of only non-interface types.

In fact, every type term defines a method set.
Calculations of type sets follow the following rules:

* The type set of a non-interface type literal or type name only contains the type denoted by the type literal or type name.
* As just mentioned above, the type set of a tilde form `~T` is the set of types whose underlying types are `T`. In theory, this is an infinite set.
* The type set of a method specification is the set of non-interface types whose method sets include the method specification.
  In theory, this is an infinite set.
* The type set of an empty interface (such as the predeclared `any`) is the set of all non-interface types.
  In theory, this is an infinite set.
* The type set of a union of terms `T1 | T2 | ... | Tn` is the union of the type sets of the terms.
* The type set of a non-empty interface is the intersection of the type sets of its interface elements.

By the current specification, two unnamed constraints are equivalent to each other if their type sets are equal.

Given the types declared in the following code snippet,
the type set of each interface type is described in the preceding comment of that interface type.

```Go
type Bytes []byte  // underlying type is []byte
type Letters Bytes // underlying type is []byte
type Blank struct{}
type MyString string // underlying type is string

func (MyString) M() {}
func (Bytes) M() {}
func (Blank) M() {}

// The type set of P only contains one type:
// []byte.
type P interface {[]byte}

// The type set of Q contains
// []byte, Bytes, and Letters.
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
// type set contains all non-interface types.
type Z interface {~[]byte | ~string | any}
```

Please note, interface elements are separated with semicolon (`;`),
either explicitly or implicitly (Go compilers will
[insert some missing semicolons as needed during compilations](https://go101.org/article/line-break-rules.html)).
The following interface type literals are equivalent to each other,
they all denote an unnamed interface type which type set is empty.
The denoted unnamed interface type and the underlying type of the type `S`
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

If the type set of a type `X` is a subset of an interface type `Y`, we say `X` implements `Y`.
Here, `X` may be an interface type or a non-interface type.

Because the type set of an empty interface type is a super set of the type sets of any types,
all types implement an empty interface type.
On the other hand, because an empty type set is a subset of any type sets,
an empty-type-set interface type implements any interface types.

Take the types declared above as an example,

* the interface type `S`, whose type set is empty, implements all interface types.
* all types implement the interface type `Z`, which is actually a blank interface type.

The list of the methods specified by an interface type is called the method set of the interface type.
If an interface type `X` implements another interface type `Y`, then the method set of `X` must be a super set of `Y`.

Interface types whose type sets can be defined entirely by a method set (may be empty)
are called basic interface types.
Before 1.18, Go only supports basic interface types.
Basic interfaces may be used as either value types or type constraints,
but non-basic interfaces may only be used as type constraints (as of Go 1.21).

Take the types declared above as an example,, `L`, `M`, `U`, `Z` and `any` are basic types.

In the following code, the declaration lines for `x` and `y` both compile okay,
but the line declaring `z` fails to compile.

```Go
var x any
var y interface {M()}

// error: interface contains type constraints
var z interface {~[]byte}
```

Whether or not to support non-basic interface types as value types in future Go versions in unclear now.

Note, before Go toolchain 1.19, aliases to non-basic interface types were not supported.
The following type alias declarations [are only legal since Go toolchain 1.19](https://github.com/golang/go/issues/51616).

```Go
type C[T any] interface{~int; M() T}
type C1 = C[bool]
type C2 = comparable
type C3 = interface {~[]byte | ~string}
```

## More about the predeclared `comparable` constraint

As aforementioned, besides `any`, Go 1.18 also introduced another new predeclared identifier `comparable`,
which denotes an interface type whose method set is composed of all strictly comparable (non-interface) types
("strictly comparable" is explained in the next section).

The `comparable` interface type could be embedded in other interface types
to filter out types which are not strictly comparable from their type sets.
For example, the type set of the following declared constraint `C` contains only one type: `string`,
because the other types in the union are either
[incomprarable](https://go101.org/article/type-system-overview.html#types-not-support-comparison)
(the first three) or not [strictly comparable](#strictly-comparable) (the last one).

```Go
type C interface {
	comparable
	[]byte | func() | map[int]bool | string | [2]any
}
```

_(Note, some earlier Go toolchain 1.18 and 1.19 versions failed to exclude `[2]any` from the type set of `C`. The bug has been fixed in newer Go toolchain 1.18 and 1.19 versions.)_

Currently (Go 1.21), the `comparable` interface is treated as a non-basic interface type.
So, now, it may only be used as type parameter constraints, not as value types.
The following code is illegal:

```Go
var x comparable = 123
```

The type set of the `comparable` interface is obviously a subset of the type set of the `any` interface,
so `comparable` undoubtedly implements `any`, and not vice versa.

{#strictly-comparable}
## Comparable vs. strictly comparable

We know that, although interface types are comparable, comparing two values of an interface type
will produce a panic at run time if the dynamic types
of the two interface values are identical and the identical dynamic type is incomparable.
Comparing values of struct or array types which contain interface components might also produce a panic.

For example, any of the comparisons shown in the `main` function will produce a panic at run time.

```Go
package main

var m any = map[int]string{}
var a = [2]any{1, func(){}}
var s = struct{x any}{x: m}

func main() {
	_ = m == m
	_ = a == a
	_ = s == s
}
```

The concept of "strictly comparable" is introduced in Go 1.20.
Comparing values of a strictly comparable type is guaranteed to be run-time panic free.
An ordinary type is strictly comparable if it is comparable and neither an interface type nor composed of interface types.

The following value types are not strictly comparable:

* (basic) interface types.
* struct types which fields contain interfaces.
* array types which elements contain interfaces.
* type parameters which type set contains at least one type of the above two cases (structs and arrays).

In the above example, the types `any`, `[2]any` and `struct{x any}` are all comparable but not strictly comparable.

As mentioned in the last section, all types in the type set of the predeclared `comparable` interface are strictly comparable.

{#implementation-vs-satisfaction}
## Type implementation vs. type satisfaction

Before Go 1.20, the two terminologies were used interchangeably.
In other words, the following two descriptions were equivalent to each other before Go 1.20:

* a type `X` implements an interface type `Y`.
* a type `X` satisfies an interface type `Y`.

Before Go 1.20, they both meant the type set of the type `X` is a sub-set of the interface type `Y`.
So, before Go 1.20, if a type `X` implements an interface type `Y`,
then it must also satisfy `Y`; and vice versa.

Since Go 1.20, the meaning of type implementation remains the same.
However, sometimes, an ordinary value type `X` might satisfy an interface type `Y`, even if it doesn't implement `Y`.
In other words, since Go 1.20, if an ordinary value type `X` implements an interface type `Y`,
then it must also satisfy `Y`; but not vice versa.

Specifically speaking, since Go 1.20, if a comparable ordinary value type `X` is not strictly comparable
and it doesn't implement a type constraint `Y`,
but the underlying type of `Y` can be written as `interface{ comparable; E }`,
where `E` is a basic interface type and the ordinary value type `X` implements `E`,
then `X` also satisfies `Y`.

For example, the type `any` surely doesn't implements the type constraint `comparable`.
But when it is used as an ordinary value type, it satisfies `comparable`.
Because the underlying type of `comparable` can be written as `interface{ comparable; any }`
and `any` implements `any`.

In the following code, the types `*A` and `*B` both satisfy (but don't implement) the interface type `C`.
Because `C` can be written as `interface{ comparable; interface{ M() } }` and
both `*A` and `*B` implement `interface{ M() }`.

```Go
type A [2]any

func (a *A) M() {}

type B struct{
	A
	x any
}

type C interface {
	comparable
	M()
}
```

As of Go 1.21, type satisfactions are used to verify whether or not an ordinary value type
can be used as a type argument of an instantiation of a generic type/function.
Please read the next chapter for details.

## More requirements for union terms

The above has mentioned that a union term may not be a type parameter. There are two other requirements for union terms.

The first is an implementation specific requirement: a term union with more than one term cannot contain the predeclared identifier `comparable` or interfaces that have methods. 
For example, the following term unions are both illegal (as of Go toolchain 1.21):

```Go
[]byte | comparable
string | error
```

To make descriptions simple, this book will view the predeclared `comparable` interface type
as an interface type having a method (but not view it as a basic interface type).

Another requirement (restriction) is that the type sets of all non-interface type terms in a term union must have no intersections. For example, in the following code snippet, the term unions in the first declaration fails to compile, but the last two compile okay.

```Go
type _ interface {
	int | ~int // error
}

type _ interface {
	interface{int} | interface{~int} // okay
}

type _ interface {
	int | interface{~int} // okay
}
```

The three term unions in the above code snippet are equivalent to each other in logic,
which means this restriction is not very reasonable.
So it might be removed in later Go versions, or become stricter to defeat the workaround.

<!--
https://github.com/golang/go/issues/51607
https://github.com/golang/go/issues/45346#issuecomment-862505803
-->

## Type parameter lists

From the examples shown in the last chapter, we know type parameter lists
are used in generic type specifications, method declarations for generic base types
and generic function declarations.

A type parameter list contains at least one type parameter declaration
and is enclosed in square brackets.
Each parameter declaration is composed of a name part and a constraint part
(we can think the constraints are implicit in method declarations for generic base types).
The name represents a type parameter constrained by the constraint. 
Parameter declarations are comma-separated in a type parameter list.

In a type parameter list, all type parameter names must be present.
They may be the blank identifier `_` (called blank name).
All non-blank names in a type parameter list must be unique.

Similar to value parameter lists, if the constraints of
some successive type parameter declarations in a type parameter list are identical,
then these type parameter declarations could share a common
constraint part in the type parameter list.
For example, the following two type parameter lists are equivalent.

```Go
[A any, B any, X comparable, _ comparable]
[A, B any, X, _ comparable]
```

Similar to value parameter lists, if the right `]` token in a type parameter list
and the last constraint in the list are at the same line, an optional comma is allowed
to be inserted between them.
[The comma is required](https://go101.org/article/line-break-rules.html#commas) if the two are not at the same line.

For example, in the following code, the beginning lines are legal, the ending lines are not.

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

To make descriptions simple, the type set of the constraint of a type parameter
is also called the type set of the type parameter and type set of a value of
the type parameter in this book.

## Simplified constraint form

In a type parameter list, if a constraint only contains one element
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
then compilers should think the code declares a generic type.
If it denotes a constant (it is possible), then compilers
will treat `T *int` as a multiplication expression and
think the code declares an ordinary array type.

It is possible for compilers to distinguish what the `int` identifier denotes,
but there are some costs to achieve this. To avoid the costs,
compilers always treat the `int` identifier as a value expression
and think the above declaration is an ordinary array type declaration.
So the above declaration line will fail to compile
if `T` or `int` don't denote integer constants.

Then how to declare a generic type with a single type parameter with `*int` as the constraint?
There are two ways to accomplish this:

1. use the full constraint form, or
1. let a comma follow the simplified constraint form.

The two ways are shown in the following code snippet:

```Go
// Assume int is a predeclared type.
type G[T interface{*int}] struct{}
type G[T *int,] struct{}
```

The two ways shown above are also helpful for
some other special cases which might also cause parsing ambiguities.
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

We should insert a comma after the presumed constraint `*int|bool` to remove the ambiguity.

```Go
type C5[T *int|bool, ] struct{} // compiles okay
```

(Note: this way doesn't work with Go toolchain 1.18. It was [a bug](https://github.com/golang/go/issues/51488) and has been fixed since Go toolchain 1.19.)

We could also use full constraint form or swap the positions of `*int` and `bool` to make it compile okay.

```Go
// Assume int and bool are predeclared types.
type C5[T interface{*int|bool}] struct{}
type C5[T bool|*int] struct{}
```

On the other hand, the following two weird generic type declarations are both legal.

```Go
// "make" is a declared type parameter.
// Its constraint is interface{chan int}.
type PtrToChan[make (chan int)] *make

// "new" is a declared type parameter.
// Its constraint is interface{[3]float64}.
type Matrix33[new ([3]float64)] [3]new
```

The two declarations are really bad practices. Don't use them in serious code.

<!--
https://github.com/golang/go/issues/49482
https://github.com/golang/go/issues/49485
https://github.com/golang/go/issues/51488
-->

{#type-parameters-are-interfaces}
## Each type parameter is a distinct named type

Since Go 1.18, named types include

* predeclared types, such as `int`, `bool` and `string`.
* defined non-generic types.
* instantiated types of generic types.
* type parameter types (the types declared in type parameter lists).

Two different type parameters are never identical.

The type of a type parameter is a constraint, a.k.a an interface type.
This means the underlying type of a type parameter type should be an interface type.
However, this doesn't mean a type parameter behaves like an interface type.
Its values may neither box non-interface values nor be type asserted (as of Go 1.21).
In fact, it is almost totally meaningless to talk about underlying types of type parameters.
We just need to know that the underlying type of a type parameter is not itself.
And we ought to think that two type parameters never share an identical underlying type,
even if the constraints of the two type parameters are identical.

In fact, a type parameter is just a placeholder for the types in its type set.
Generally speaking, it represents a type which owns the common traits of the types in its type set.

As the underlying type of a type parameter type is not the type parameter type itself,
the tilde form `~T` is illegal if `T` is type parameter.
So the following (equivalent) type parameter lists are illegal.

```Go
[A int, B ~A]                       // error
[A interface{int}, B interface{~A}] // error
```

As mentioned above, type parameters are also disallowed to be embedded
as type names and type terms in an interface type.
The following declarations are also illegal. 

```Go
type Cx[T int] interface {
	T
}

type Cy[T int] interface {
	T | []string
}
```

In fact, currently (Go 1.21), [type parameters may not be embedded in struct types](888-the-status-quo-of-Go-custom-generics.md#embed-type-parameter), too.

{#type-parameter-scope}
## The scopes of a type parameters

Go specification says:

* The scope of an identifier denoting a type parameter of a function or declared by a method receiver begins after the name of the function and ends at the end of the function body.
* The scope of an identifier denoting a type parameter of a type begins after the name of the type and ends at the end of the specification of the type.

So the following type declaration is valid, even if the use of type parameter `E` is ahead of its declaration.
The type parameter `E` is used in the constraint of the type parameter `S`,

```Go
type G[S ~[]E, E int] struct{}
```

Please note,

* as mentioned in the last section, although `E` is a type parameter type, `[]E` is an ordinary (slice) type.
* the underlying type of `S` is `interface{~[]E}`, not `[]E`.
* the underlying type of `E` is `interface{int}`, not `int`.

By Go specification, the function and method declarations in the following code all fail to compile.

```Go
type C any
func foo1[C C]() {}    // error: C redeclared
func foo2[T C](T T) {} // error: T redeclared

type G[G any] struct{x G} // okay
func (E G[E]) Bar1() {}   // error: E redeclared
```

The following `Bar2` method declaration should compile okay, but it doesn't now (Go toolchain 1.21). This is [a bug which will be fixed in Go toolchain 1.21](https://github.com/golang/go/issues/51503).

```Go
type G[G any] struct{x G} // okay
func (v G[G]) Bar2() {}   // error: G is not a generic type
```

## Composite type literals (unnamed types) containing type parameters denote ordinary types

For example, `*T` is always an ordinary (pointer) type.
It is a type literal, so its underlying type is itself, whether or not `T` is a type parameter.
The following type parameter list is legal.

```Go
[A int, B *A] // okay
```

For the same reason, the following type parameter lists are also legal.

```Go
[T ~string|~int, A ~[2]T, B ~chan T]            // okay
[T comparable, M ~map[T]int32, F ~func(T) bool] // okay
```

## More about generic type and function declarations

We have seen a generic type declaration and some generic function declarations in the last chapter.
Different from ordinary type and function declarations, each of the generic ones has a
type parameter list part.

This book doesn't plan to further talk about generic type and function declaration syntax.

The source type part of a generic type declaration must be an ordinary type.
So it might be

* a composite type literal. As mentioned above, a composite type literal always represents an ordinary type.
* a type name which denotes an ordinary type.
* an instantiated type. Type instantiations will be explained in detail in the next chapter.
  For now, we only need to know each instantiated type is a named ordinary type.

The following code shows some generic type declarations with all sorts of source types.
All of these declarations are valid.

```Go
// The source types are ordinary type names.
type (
	Fake1[T any] int
	Fake2[_ any] []bool
)

// The source type is an unnamed type (composite type).
type MyData [A any, B ~bool, C comparable] struct {
	x A
	y B
	z C
}

// The source type is an instantiated type.
type YourData[C comparable] MyData[string, bool, C]
```

Type parameters may not be used as the source types in generic type declarations.
For example, the following code doesn't compile.

```Go
type G[T any] T // error
```

