





## Try not to use unions of types which have identical underlying type `T`, use `~T` instead

## Making uyse of type argument inference

Put type parameters without depending others at the beginning, for convenience.

put `_ ResultType` in type parameter list

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

