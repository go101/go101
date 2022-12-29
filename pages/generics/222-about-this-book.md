
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


<!--
https://github.com/golang/proposal/blob/master/design/generics-implementation-dictionaries-go1.18.md

* show some basic interfeace type argument examples

* ~T is called underlying term

* type argument inference needs more detailed explainations.

* Implementation restriction: A compiler need not report an error if an operand's type is a type parameter with an empty type set. Functions with such type parameters cannot be instantiated; any attempt will lead to an error at the instantiation site. 

* example: how to define an expected constraint?
  * some achievable, some are not.  

* An example show the difference of using ordinary interface and generic constraint

* What does this constraint mean?
  interface {
  	M1()
  	M2() error
  	I
  	int | bool
  }
  使用bullet一条一条列出来。

* more

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
	
-->


  
