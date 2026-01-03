
# About Go Generics 101

Since version 1.18, Go has supported custom generics.

This book talks about the custom generics feature of Go programming language.
The content in this book includes:

* custom generic syntax,
* type constraints and type parameters,
* type arguments and type inference,
* how to write valid custom generic code,
* current implementation/design restrictions.

A reader needs to be familiar with Go general programming to read this book.
In particular, readers of this book should be familiar with
[Go type system](https://go101.org/article/type-system-overview.html),
including Go built-in generics, which and Go custom generics are two different systems.

Currently, the book mainly focuses on the syntax of (and concepts in) custom generics.
More practical examples will be provided when I get more experiences of using custom generics.

## About GoTV

During writing this book, the tool [GoTV](https://go101.org/apps-and-libs/gotv.html)
is used to manage installations of multiple Go toolchain versions and check
the behavior differences between Go toolchain versions.


<!--

https://github.com/golang/go/issues/60377
    A defined interface type X can be converted to a defined interface type Y,
    and vice versa, if their underlying types are identical.
    That is the difference between interface types and non-interface types.

https://github.com/golang/go/issues/67025

https://github.com/golang/go/issues/66751

https://github.com/golang/go/issues/58573 local types in different instantiations of a generic function
https://github.com/golang/go/issues/58573#issuecomment-1433898205
https://github.com/golang/go/issues/65152

https://github.com/golang/go/issues/58608

https://github.com/golang/proposal/blob/master/design/generics-implementation-dictionaries-go1.18.md

* type argument inferences still have some limitations
  * https://github.com/golang/go/issues/63750

* type argument inference needs more detailed explainations.
 

* example: how to define an expected constraint?
  * some achievable, some are not.  

* An example show the difference of using ordinary interface and generic constraint.

* more

	https://github.com/golang/go/issues/62172
		https://github.com/golang/go/issues/40301#issuecomment-885119414
		https://github.com/golang/go/issues/40301#issuecomment-754156626
		
				package main

				import "unsafe"

				func f(x int64) byte {
				  return 1 << unsafe.Sizeof(x) >> unsafe.Sizeof(x)
				}

				func g[T int64](x T) byte {
				  return 1 << unsafe.Sizeof(x) >> unsafe.Sizeof(x)
				}

				func main() {
				  var n int64 = 0
				  println(f(n), g(n))
				}

	https://github.com/golang/go/issues/61741

	https://github.com/golang/go/issues/60130
	https://github.com/golang/go/issues/60117

	https://github.com/golang/go/issues/51522 miscompilation of comparison between type parameter and interface
	https://github.com/golang/go/issues/51521 wrong panic message for method call on nil of generic interface type

	https://github.com/golang/go/issues/53477
	https://github.com/golang/go/issues/50681 // compile time type switch
	https://github.com/golang/go/issues/49206 // type switch
	https://github.com/golang/go/issues/45380 type switch on type parameters not supported
	
	https://github.com/golang/go/issues/54028
	https://github.com/golang/go/issues/53762
	
	https://github.com/golang/go/issues/53087 produce duplicate type descriptor
	
	https://github.com/golang/go/issues/53137 unsafe.Offsetof bug
	
	https://github.com/golang/go/issues/53309
	
	https://github.com/golang/go/issues/53419
	
	https://github.com/golang/go/issues/52181
	
	https://github.com/golang/go/issues/53635
	
	https://github.com/golang/go/issues/53883
	
	https://github.com/golang/go/issues/54447
	
	https://github.com/golang/go/issues/54456
	
	https://github.com/golang/go/issues/54535
	
	https://github.com/golang/go/issues/54537
	
	https://github.com/golang/go/issues/55964
	
	https://github.com/golang/go/issues/56923
	
	https://github.com/golang/go/issues/62157


==================== type argument inference https://go.dev/blog/type-inference

 * https://x.com/zigo_101/status/1714885320599302598

interace:
	
		https://x.com/zigo_101/status/1714187265310957864
		
		package main

		func f       (...A) {}
		func g[T any](...T) {}

		type A any
		type B any                                            

		var a A
		var b B

		func main(){
			g(a, b)
		}

channel

		package main

		import "fmt"

		func g[T any](...T) (_ T){return}

		type A = chan int
		type B = <-chan int
		type C chan int
		type D <-chan int

		var a A
		var b B
		var c C
		var d D

		func main(){
		  // T is infered as C
		 fmt.Printf("%T ", g(a, b, c, d)) // type D of d does not match inferred type C for T
		}

more composite types

		package main

		import "fmt"

		func g[T any](...T) (_ T){return}

		type A = []int
		type B []int
		type C []int

		func main(){
		  // T is inferred as B
		 fmt.Printf("%T ", g(A{}, B{}, C{})) // type C of C{} does not match inferred type B for T
		}

-->


  
