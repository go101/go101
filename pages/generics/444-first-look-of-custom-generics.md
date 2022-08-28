
# First Look of Custom Generics

In the custom generics world, a type may be declared as a generic type,
and a function may be declared as generic function.
In addition, generic types are defined types, so they may have methods.

The declaration of a generic type, generic function, or method of a generic type
contains a type parameter list part, which is the main difference from
the declaration of an ordinary type, function, or method.

## A generic type example

Firstly, let's view an example to show how generic types look like.
It might be not a perfect example, but it does show the usefulness of custom generic types. 

```Go
package main

import "sync"

type Lockable[T any] struct {
	sync.Mutex
	Data T
}

func main() {
	var n Lockable[uint32]
	n.Lock()
	n.Data++
	n.Unlock()
	
	var f Lockable[float64]
	f.Lock()
	f.Data += 1.23
	f.Unlock()
	
	var b Lockable[bool]
	b.Lock()
	b.Data = !b.Data
	b.Unlock()
	
	var bs Lockable[[]byte]
	bs.Lock()
	bs.Data = append(bs.Data, "Go"...)
	bs.Unlock()
}
```

In the above example, `Lockable` is a generic type.
Comparing to non-generic types, there is an extra part, a type parameter list,
in the declaration (specification, more precisely speaking) of a generic type.
Here, the type parameter list of the `Lockable` generic type is `[T any]`.

A type parameter list may contain one and more type parameter declarations
which are enclosed in square brackets and separated by commas.
Each parameter declaration is composed of a type parameter name and a (type) constraint.
For the above example, `T` is the type parameter name and `any` is the constraint of `T`.

Please note that `any` is a new predeclared identifier introduced in Go 1.18.
It is an alias of the blank interface type `interface{}`.
We should have known that all types implements the blank interface type.

_(Note, generally, Go 101 books don't say a type alias is a type.
They just say a type alias denotes a type.
But for convenience, the Go Generics 101 book often says `any` is a type.)_

We could view constraints as types of (type parameter) types.
All type constraints are actually interface types.
Constraints are the core of custom generics and
will be explained in detail in the next chapter.

`T` denotes a type parameter type.
Its scope begins after the name of the declared generic type
and ends at the end of the specification of the generic type.
Here it is used as the type of the `Data` field.

Since Go 1.18, value types fall into two categories:

1. type parameter type;
1. ordinary types.

Before Go 1.18, all values types are ordinary types.

A generic type is a [defined type](https://go101.org/article/type-system-overview.html#type-definition).
It must be instantiated to be used as value types.
The notation `Lockable[uint32]` is called an instantiated type (of the generic type `Lockable`).
In the notation, `[uint32]` is called a type argument list and `uint32` is called a type argument,
which is passed to the corresponding `T` type parameter.
That means the type of the `Data` field of the instantiated type `Lockable[uint32]` is `uint32`.

A type argument must implement the constraint of its corresponding type parameter.
The constraint `any` is the loosest constraint, any value type could be passed to the `T` type parameter.
The other type arguments used in the above example are: `float64`, `bool` and `[]byte`.

Every instantiated type is a [named type](https://go101.org/article/type-system-overview.html#named-type) and an ordinary type.
For example, `Lockable[uint32]` and `Lockable[[]byte]` are both named types.

The above example shows how custom generics avoid code repetitions for type declarations.
Without custom generics, several struct types are needed to be declared,
like the following code shows.

```Go
package main

import "sync"

type LockableUint32 struct {
	sync.Mutex
	Data uint32
}

type LockableFloat64 struct {
	sync.Mutex
	Data float64
}

type LockableBool struct {
	sync.Mutex
	Data bool
}

type LockableBytes struct {
	sync.Mutex
	Data []byte
}

func main() {
	var n LockableUint32
	n.Lock()
	n.Data++
	n.Unlock()
	
	var f LockableFloat64
	f.Lock()
	f.Data += 1.23
	f.Unlock()
	
	var b LockableBool
	b.Lock()
	b.Data = !b.Data
	b.Unlock()
	
	var bs LockableBytes
	bs.Lock()
	bs.Data = append(bs.Data, "Go"...)
	bs.Unlock()
}
```

The non-generic code contains many code repetitions,
which could be avoided by using the generic type demonstrated above.

## An example of a method of a generic type

Some people might not appreciate the implementation of the above generic type.
Instead, they prefer to use a different implementation as the following code shows.
Comparing with the `Lockable` implementation in the last section, the new one
hides the struct fields from outside package users of the generic type.

```Go
package main

import "sync"

type Lockable[T any] struct {
	mu sync.Mutex
	data T
}

func (l *Lockable[T]) Do(f func(*T)) {
	l.mu.Lock()
	defer l.mu.Unlock()
	f(&l.data)
}

func main() {
	var n Lockable[uint32]
	n.Do(func(v *uint32) {
		*v++
	})
	
	var f Lockable[float64]
	f.Do(func(v *float64) {
		*v += 1.23
	})
	
	var b Lockable[bool]
	b.Do(func(v *bool) {
		*v = !*v
	})
	
	var bs Lockable[[]byte]
	bs.Do(func(v *[]byte) {
		*v = append(*v, "Go"...)
	})
}
```

In the above code, a method `Do` is declared for the generic base type `Lockable`.
Here, the receiver type is a pointer type, which base type is the generic type `Lockable`.
Different from method declarations for ordinary base types,
there is a type parameter list part
following the receiver generic type name `Lockable` in the receiver part.
Here, the type parameter list is `[T]`.

The type parameter list in a method declaration for a generic base type
is actually a duplication of the type parameter list specified
in the generic receiver base type specification. To make code clean,
type parameter constraints are (and must be) omitted from the list.
That is why here the type parameter list is `[T]`, instead of `[T any]`.

Here, `T` is used in a value parameter type, `func(*T)`.

* The type of its method `Do` of the instantiated type `Lockable[uint32]` is `func(f func(*uint32))`.
* The type of its method `Do` of the instantiated type `Lockable[float64]` is `func(f func(*float64))`.
* The type of its method `Do` of the instantiated type `Lockable[bool]` is `func(f func(*bool))`.
* The type of its method `Do` of the instantiated type `Lockable[[]byte]` is `func(f func(*[]byte))`.

Please note that, the type parameter names used in a method declaration for a generic base type
are not required to be the same as the corresponding ones used in the generic type specification.
For example, the above method declaration is equivalent to the following rewritten one.

```Go
func (l *Lockable[Foo]) Do(f func(*Foo)) {
	...
}
```

Though, it is a bad practice not to keep the corresponding type parameter names the same.

BTW, the name of a type parameter may even be the blank identifier `_`
if it is not used (this is also true for the type parameters in generic type
and function declarations). For example,

```Go
func (l *Lockable[_]) DoNothing() {
}
```

## A generic function example

Now, let's view an example of how to declare and use generic (non-method) functions.

```Go
package main

// NoDiff checks whether or not a collection
// of values are all identical.
func NoDiff[V comparable](vs ...V) bool {
	if len(vs) == 0 {
		return true
	}
	
	v := vs[0]
	for _, x := range vs[1:] {
		if v != x {
			return false
		}
	}
	return true
}

func main() {
	var NoDiffString = NoDiff[string]
	println(NoDiff("Go", "Go", "Go")) // true
	println(NoDiffString("Go", "go")) // false
	
	println(NoDiff(123, 123, 123, 123)) // true
	println(NoDiff[int](123, 123, 789)) // false
	
	type A = [2]int
	println(NoDiff(A{}, A{}, A{}))     // true
	println(NoDiff(A{}, A{}, A{1, 2})) // false
	
	println(NoDiff(new(int)))           // true
	println(NoDiff(new(int), new(int))) // false

	println(NoDiff[bool]())   // true

	// _ = NoDiff() // error: cannot infer V
	
	// error: *** does not implement comparable
	// _ = NoDiff([]int{}, []int{})
	// _ = NoDiff(map[string]int{})
	// _ = NoDiff(any(1), any(1))
}
```

In the above example, `NoDiff` is a generic function.
Different from non-generic functions, and similar to generic types,
there is an extra part, a type parameter list, in the declaration of a generic function.
Here, the type parameter list of the `NoDiff` generic function is `[V comparable]`, in which
`V` is the type parameter name and `comparable` is the constraint of `V`.

`comparable` is new predeclared identifier introduced in Go 1.18.
It denotes an interface that is implemented by all comparable types.
It will be explained with more details in the next chapter.

Here, the type parameter `V` is used as the variadic (value) parameter type.

The notations `NoDiff[string]`, `NoDiff[int]` and `NoDiff[bool]` are called instantiated functions (of the generic function `NoDiff`).
Similar to instantiated types, in the notations, `[string]`, `[int]` and `[bool]` are called type argument lists.
In the lists, `string`, `int` and `bool` are called type arguments, all of which are passed to the `V` type parameter.

The whole type argument list may be totally omitted from an instantiated function expression
if all the type arguments could be inferred from the value arguments.
That is why some calls to the `NoDiff` generic function have no type argument lists in the above example.

* In the call `NoDiff("Go", "Go", "Go")`, the type argument is inferred as `string`, the default type of the value arguments.
* In the call `NoDiff(123, 123, 123, 123)`, the type argument is inferred as `int`, the default type of the value arguments.
* In the two calls with `A{}` and `A{1, 2}` as value arguments, the type argument is inferred as `[2]int`, the type of the value arguments.
* In the two calls with `new(int)` as value arguments, the type argument is inferred as `*int`, the type of the value arguments.

Please note that all of these type arguments implement the `comparable` interface.
Incomparable types, such as `[]int` and `map[string]int` may not be passed as type arguments
of calls to the `NoDiff` generic function.
And please note that, although `any` is a comparable (value) type, it doesn't implement `comparable`, so it is also not an eligible type argument.
This will be talked about in detail in the next chapter.

The above example shows how generics avoid code repetitions for function declarations.
Without custom generics, we need to declare a function for each type argument used in the example.
The bodies of these function declarations would be almost the same.

## Generic functions could be viewed as simplified forms of methods of generic types

The generic function shown in the above section could be viewed as a simplified
form of a method of a generic type, as shown in the following code.

```Go
package main

type NoDiff[V comparable] struct{}

func (nd NoDiff[V]) Do(vs ...V) bool {
	... // same as the body of the above generic function
}

func main() {
	var NoDiffString = NoDiff[string]{}.Do
	println(NoDiffString("Go", "go")) // false
	
	println(NoDiff[int]{}.Do(123, 123, 789)) // false
	
	println(NoDiff[*int]{}.Do(new(int))) // true
}
```

In the above code, `NoDiff[string]{}.Do`, `NoDiff[int]{}.Do`
and `NoDiff[*int]{}.Do` are three method values of different
instantiated types.

We could view a generic type as a type parameter space,
and view all of its methods as some functions sharing the same type parameter space.

To make descriptions simple, sometimes, methods of generic types
are also called as generic functions in this book.

