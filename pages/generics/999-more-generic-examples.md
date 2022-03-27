


## Zero value

```Go
func ZeroOf[T any]() (t T) {
	return 
}

func ZeroFor[T any](T)(t T) {
	return
}

func ZeroIt[T any](t *T) {
	*t = ZeroOf[T]
}
```

##

```Go
type Number interface{int | int64 | int32 | float64 | float32}

func Abs[T Number](x T) T {
	if x < 0 {
		return -x
	}
	return x
}
```

## Free list



## 

```Go
package main

import "fmt"

type MapSet[Key, Val comparable] map[Key]map[Val]uint32

func (ms MapSet[Key, Val]) Put(k Key, v Val) {
	vm, ok := ms[k]
	if !ok {
		vm = make(map[Val]uint32)
		ms[k] = vm
	}
	vm[v]++
}

func main() {
	var x = make(MapSet[string, int])
	x.Put("Go", 2007)
	x.Put("Go", 2009)
	x.Put("Go", 2012)
	x.Put("Go", 2007)
	x.Put("Go", 2009)
	fmt.Println(x["Go"]) // map[2007:2 2009:2 2012:1
	
	type Status = MapSet[bool, string]
	var y = make(Status, 100)
	y.Put(true, "X")
	y.Put(false, "Y")
	y.Put(true, "X")
	y.Put(false, "Z")
	y.Put(true, "W")
	y.Put(false, "Z")
	fmt.Println(y[true])  // map[W:1 X:2]
	fmt.Println(y[false]) // map[Y:1 Z:2]
}
```

## 

```Go
package main

func convert[A ~string, B ~rune|~byte, C []B](x C) <-chan A {
	var as = make(chan A, len(x))
	for _, v := range x {
		as <- A(v)
	}
	close(as)
	
	return as
}

func main() {
	var runes = []rune{65, 66, 67, 68, 69}
	var stringStream = convert[string](runes)
	for str := range stringStream {
		println(str)
	}
	
	var bytes = []byte{97, 98, 99, 100, 101}
	stringStream = convert[string](bytes)
	for str := range stringStream {
		println(str)
	}
}
```

##

```Go
type Copyable[T any] interface {

	Copy() T

}

func copySlice[T Copyable[T]](s []T) []T {

	s2 := make([]T, len(s))

	for i := range s {

		s2[i] = s[i].Copy()

	}

	return s2

}
```

## Examples:

atomic
math functions

An example which is some involuted.

* the problems to solve
  * math.Max
  * atomic.StorePointer
  * container/list
  * simplfy code generation
    * MyList, MyTree, ...
    * byte slice and string


https://github.com/golang/go/issues/51909