# Constraints and Type Parameters

A constraint means a type constraint, it is used to constrained some type parameters.
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
  There are some requirements for union terms.
  These requirements will be described in a section below.

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
* The type set of an empty interface is the set of all non-interface types.
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

// The type set of P only contains one type:
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

Please note that interface elements are separated with semicolon (`;`),
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

Because the type set of an empty interface type is a super set of the type sets of any types,
all types implement an empty interface type.

In the above example,

* the interface type `S`, whose type set is empty, implements all interface types.
* all types implement the interface type `Z`, which is actually a blank interface type.

The list of methods specified by an interface type is called the method set of the interface type.
If an interface type `X` implements another interface type `Y`, then the method set of `X` must be a super set of `Y`.

Interface types whose type sets can be defined entirely by a method set (may be empty)
are called basic interface types.
Before 1.18, Go only supports basic interface types.
Basic interfaces may be used as either value types or type constraints,
but non-basic interfaces may only be used as type constraints (as of Go 1.18).

In the above examples, `L`, `M`, `U`, `Z` and `any` are basic types.

In the following code, the declaration lines for `x` and `y` both compile okay,
but the line declaring `z` fails to compile.

```Go
var x any
var y interface {M()}

// error: interface contains type constraints
var z interface {~[]byte}
```

Using non-basic interface types as value types might be supported in future Go versions.

## More about the predeclared `comparable` constraint

As aforementioned, besides `any`, Go 1.18 introduces another new predeclared identifier `comparable`,
which denotes an interface type that is implemented by all comparable types.

The `comparable` interface type could be embedded in other interface types
to filter out incomparable types from their type sets.
For example, the type set of the following declared constraint `C` contains only one type: `string`.

```Go
type C interface {
	comparable
	[]byte | string | func() | map[int]bool
}
```

Currently (Go 1.18), the `comparable` interface is treated as a non-basic interface type.
So, now, it may only be used as type parameter constraints, not as value types.
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
This is quite odd.

After deliberation, Go core team believe that
[it is a design flaw](https://github.com/golang/go/issues/50646#issuecomment-1023706545)
to treat all interface types as comparable types and it is a pity that
the `comparable` type has not been supported since Go 1.0 to avoid this flaw.

Go core team try to make up for this flaw in Go custom generics age.
So they decided that all basic interface types don't satisfy (implement) the `comparable` constraint.
A consequence of this decision is [it causes diffculties to some code designs](https://github.com/golang/go/issues/51257).

To avoid the consequence, a proposal has been made to
[permit using `comparable` as value types](https://github.com/golang/go/issues/51338).
Whether or not it should be accepted is still under discuss.
It could be accepted in as earlier as Go 1.19.

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

## More requirements for union terms

The above has mentioned that a union term may not be a type parameter. There are two other requirements for union terms.

The first is an implementation specific requirement: a term union with more than one term cannot contain the predeclared identifier `comparable` or interfaces that have methods. 
For example, the following term unions are both illegal (as of Go toolchain 1.18):

```Go
[]byte | comparable
string | error
```

To make descriptions simple, this book will view the predeclared `comparable` interface type
as an interface type having a method (but not view it as a basic interface type).

Another requirement (restriction) is that the type sets of all non-interface type terms in a term union must have no intersections.
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
So it might be removed in later Go versions
(or as earlier as 1.18.x), or become stricter to defeat the workaround.

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

When the type set of a type parameter is mentioned, it means
the type set of the constraint of the type parameter.

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

As of Go toolchain 1.18, inserting a comma after the presumed constraint `*int|bool` doesn't work
(It is [a bug](https://github.com/golang/go/issues/51488)
in Go toolchain 1.18 and will be fixed in Go toolchain 1.19).

Now, we could use full constraint form or exchange the places of `*int` and `bool` to make it compile okay.

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

## Each type parameter is a distinct named type and its underlying type is an interface type

Since Go 1.18, named types include

* predeclared types, such as `int`, `bool` and `string`.
* defined non-generic types.
* instantiated types of generic types.
* type parameter types (the types declared in type parameter lists).

Two different type parameters are never identical.

The type of a type parameter is a constraint, a.k.a an interface type.
This means the underlying type of a type parameter type is an interface type.
However, this doesn't mean a type parameter behaves like an interface type.
Its values may not box non-interface values and be type asserted (as of Go 1.18).

In fact, a type parameter is just a placeholder for the types in its type set.
So it generally behaves as (the common traits of) the types in its type set in many situations.

As the underlying type of a type parameter type is not the type parameter type itself,
the tilde form `~T` is illegal if `T` is type parameter.
So the following type parameter list is illegal.
Because, as mentioned above, the type in a tilde form mustn't be an interface type
and its underlying type must be itself. Here the both of the conditions are not satisfied.

```Go
[A int, B ~A] // error
```

For the same reason, the following generic type declaration is also illegal.

```Go
type C[T int] interface {
	~T // error: cannot embed a type parameter
}
```

## Composite type literals (unnamed types) containing type parameters are ordinary types

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

{#type-parameter-scope}
## The scopes of a type parameters

The following type parameter list is valid, even if the use of `E` is ahead of the declaration of `E`.
The type parameter `E` is used in the constraint of the type parameter `S`,

```Go
[S ~[]E, E int]
```

Please note,

* as mentioned in the last section, although `E` is a type parameter type, `[]E` is an ordinary (slice) type.
* the underlying type of `S` is `interface{~[]E}`, not `[]E`.
* the underlying type of `E` is `interface{int}`, not `int`.

For ordinary function and method declarations, a (value) parameter/result name
is allowed to be the same as a parameter/result type name.
For example, the following function and method declarations are all valid.

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

type G[G any] struct{x G} // okay
func (E G[E]) Bar1() {}   // error: E redeclared
func (v G[G]) Bar2() {}   // error: G is not a generic type
```

The `Bar2` method declaration might become legal
[since a future Go version](https://github.com/golang/go/issues/51503).

<!--
https://github.com/golang/go/issues/51503
-->

## More about generic type and function declarations

We have seen a generic type declaraton and some generic function declarations in the last chapter.
Different from ordinary type and function declarations, each of the generic ones has a
type parameter list part.

This book doesn't plan to further talk about generic type and function declartion syntax.

The source type part of a generic type declaration must be an ordinary type.
So it might be

* a composite type literal. As mentioned above, a composite type literal always represents an ordinary type.
* a type name which denotes an ordinary type.
* an instantiated type. Type instantiations will be explained in detail in the next chapter.
  For now, we only need to know each instantiated type is a named ordinary type.

The following code showns some generic type delcarations with all sorts of source types.
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

Type parameters may not be used as the soruce types in generic type declarations.
The following code doesn't compile.

```Go
type G[T any] T // error
```
