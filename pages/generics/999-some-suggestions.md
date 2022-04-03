





## Try to use aliases of instantiated types

Or embed them, to make the methods declared for them available.

This is actually the same as general value types, nothing special.

```Go
package main

type Foo[T any] []T

func (f Foo[T]) M() {}

type (
	P = Foo[int]
	Q Foo[int]
	R struct{
		Foo[int]
	}
)

func main() {
	var x P
	x.M() // okay
	
	var y Q
	y.M() // y.M undefined (type Q has no field or method M)
	
	var r R
	r.M() // okay
}
```


## Try not to use unions of types which have identical underlying type `T`, use `~T` instead

## Making use of type argument inference

Put type parameters without depending others at the beginning, for convenience.

put `_ ResultType` in type parameter list

## Inferring type arguments from function fesult fypes is not supported

Use extra value parameters of the result types, just use the type info for inference


## Bad practice

Pass values of type parameters to interfaces, 

Don't use generics for try to use generics.

* bad use cases
  use parameterized types but pass values of the types to interfaces.
  	func F(T int|float64](v T) {fmt.Println(v)}
	and: call methods.

## generic functions might be less efficient

For the current implementaiton algorithm, ...

https://github.com/golang/go/issues/48849 

https://github.com/golang/go/issues/51699

## There are some bugs currently

https://github.com/golang/go/issues/51700
https://github.com/golang/go/issues/52026

