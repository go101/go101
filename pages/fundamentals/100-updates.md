
# Go (Fundamentals) 101 Update History

### v1.23.a (2024/Oct/16)

* add a [Range Over Functions](https://go101.org/article/function.html#range) section.

### v1.22.a (2024/Mar/18)

* Go 1.22 introduced `for range Integer` loops and changed the semantics of for-loop code blocks. Related articles are modified accordingly:
  * [Basic Control Flows](https://go101.org/article/control-flows.html#for-semantic-change)
  * [Goroutines, Deferred Function Calls and Panic/Recover](https://go101.org/article/control-flows-more.html#argument-evaluation-moment)
  * [Arrays, Slices and Maps in Go](https://go101.org/article/container.html#iteration)
  
### v1.21.a (2023/Oct/11)

* Go 1.21 [added a `clear` builtin function](https://go101.org/article/container.html#clear).
* Since Go 1.21, [the call `panic(nil)` will produce a new non-nil runtime panic](https://go101.org/article/panic-and-recover-use-cases.html#avoid-verbose).

### v1.20.a (2023/Feb/01)

* Go 1.20 [started to support slice to array conversions](https://go101.org/article/container.html#slice-to-array).
* Go 1.20 [added three new functions in the "unsafe" package](https://go101.org/article/unsafe.html): `SliceData`, `String`, and `StringData`.
* Since Go 1.20, the global random generator in the `math/rand` package will be auto seeded.
* mention that `-gcflags="-d=ssa/check_bce"` doesn't work for some generic functions in the [Bounds Check Elimination](https://go101.org/article/bounds-check-elimination.html) chapter.

### v1.19.a (2022/Aug/29)

* Go 1.19 [added some atomic types](https://go101.org/article/concurrent-atomic-operation.html).
* Go 1.19 [updated memory order guarantees](https://go101.org/article/memory-model.html#atomic).

### v1.18.a (2022/Apr/06)

* Go 1.18 added custom generics support. Several articles in Go 101 are adjusted accordingly, mainly including
  * [Interfaces in Go](https://go101.org/article/interface.html)
  * [Go Type System Overview](https://go101.org/article/type-system-overview.html) (note that the "named type" terminology was added back).

### v1.17.b (2021/Sep/10)

* improve implicit method value evaluation explanations. Read [this](https://go101.org/article/method.html#method-value-Normalization) and [this](https://go101.org/article/type-embedding.html#method-value-evaluation) for details.

### 1.16.b (2021/May/18)

* add Go 1.17 contents

### 1.16.a (2021/Feb/18)

* support Go 1.16 embedding feature.

### 1.15.b (2020/Oct/09)

* describe more about [reflect.DeepEqual](https://go101.org/article/details.html#reflect-deep-equal) related details.
* add [a new syntax exception](https://go101.org/article/exceptions.html#code-block-following-else).

### 1.15.a (2020/Aug/07)

* point out that, since Go Toolchain 1.15, using make+copy to clone slices is always more efficient than using append to clone.

### 1.14.g (2020/Jun/12)

* remove the new detail added in 1.14.e: The behavior of comparing struct values with both comparable and incomparable fields or array values with both comparable and incomparable elements is unspecified. The reason is [the behavior will be specified](https://github.com/golang/go/issues/8606).
* add [a new detail](https://go101.org/article/details.html#impossible-to-interface-assertion): About the impossible to-interface assertions which can be detected at compile time.


### 1.14.f (2020/Jun/02)

* All "Go SDK" uses are changed to "Go Toolchain".

### 1.14.e (2020/May/06)

* add [a new detail](https://go101.org/article/details.html#compare-values-with-both-comparable-and-incomparable-parts): The behavior of comparing struct values with both comparable and incomparable fields or array values with both comparable and incomparable elements is unspecified.
* add [a new detail](https://go101.org/article/details.html#blank-fields-are-ignored-in-comparisons): In struct value comparisons, blank fields will be ignored.

### 1.14.d (2020/Apr/25)

* `runtime.KeepAlive` related concents are removed from [Unsafe Pointers](https://go101.org/article/unsafe.html) article.
  I'm sorry for spreading some wrong information in this article before.

### 1.13.i (2019/Oct/31)

* fix a bug in the example code in the [Delete a segment of slice elements](https://go101.org/article/container.html#delete-slice-elements) section
  of the "Arrays, Slices and Maps in Go" article.
* correct explainations in [The Evaluation Moment of Deferred Function Values](https://go101.org/article/function.html#function-evaluation-time).
* the article "The Right Places to Call the Built-in <code>recover</code> Function" is [renamed to](https://go101.org/article/panic-and-recover-more.html) "Explain Panic/Recover Mechanism in Detail". It was almost wholly re-written.

### 1.13.h (2019/Oct/18)

* correct the explanations for the [Evaluation and Assignment Orders in Assignment Statements](https://go101.org/article/evaluation-orders.html#value-assignment) section
  in the "Expression Evaluation Orders" article.
* add [two new summaries](https://go101.org/article/101.html#compiler-optimizations).

### 1.13.e (2019/Oct/07)

* I decided to withdraw the last erratum in 1.13.d. (Re-added in 1.14.d)

### 1.13.d (2019/Sep/30)

* add <a href="https://go101.org/article/unsafe.html#fact-value-address-might-change">a new fact</a> to the "type-unsafe pointers" article
  and pointed out <a href="https://go101.org/article/unsafe.html#pattern-convert-to-uintptr-and-back">a serious mistake</a> was made in this article.

### 1.13.c (2019/Sep/25)

* remove the section containing a stupid code mistake from the "The Right Places to Call the recover Function" article.

### 1.13.b (2019/Sep/19)

* remove the inaccurate description "the address of a variable will never change"

### 1.13.a (2019/Sep/05)

* Go 1.13 ready.
* add two situations in the article [How to Gracefully Close Channels](https://go101.org/article/channel-closing.html).
* [the "named type" and "unnamed type" terminologies](https://go101.org/article/type-system-overview.html#unnamed-type) are added back,
  but they are eqivalent to "defined type" and "non-defined type" now.

### 1.12.d (2019/May/18)

* enrich the [Package-level Variables Initialization Order](https://go101.org/article/evaluation-orders.html#package-level-variables) section.

### 1.12.c (2019/April/09)

* remove the "named type" and "unnamed type" terminology.
* adjust some discriptions in [Type Embdding](https://go101.org/article/type-embedding.html).

### 1.12.b (2019/April/06)

* add a [Package-level Variables Initialization Order](https://go101.org/article/evaluation-orders.html#package-level-variables) section.

### 1.12.a (2019/March/02)

* Go 1.12 ready.

### 1.11.f (2019/Jan/02)

* remove "Unused Variables" from [Syntax/Semantics Exceptions In Go](https://go101.org/article/exceptions.html).

### 1.11.g (2018/Dec/27)

* a serious mistake was just fixed in this book.
  Before, the book said the starting index in a subslice syntax
  can't be larger than the length of the base slice. This is wrong.
  Please read <a href="container.html#subslice">the corrected section</a> again for details.

### 1.11.f (2018/Nov/09)

* rearrange [Go Details 101](https://go101.org/article/details.html), more details are added.

### 1.11.e (2018/Oct/26)

* published [Evaluation Orders](https://go101.org/article/evaluation-orders.html).
* modify some descriptions for [program resource initialization order](https://go101.org/article/packages-and-imports.html#initialization-order)

### 1.11.d (2018/Oct/18)

* remove "Comparing Interface Values" from [Syntax/Semantics Exceptions In Go](https://go101.org/article/exceptions.html).
* add "Comparisons 2" to [Syntax/Semantics Exceptions In Go](https://go101.org/article/exceptions.html).
* modify some descriptions for
  * [comparison rules](https://go101.org/article/value-conversions-assignments-and-comparisons.html#comparison-rules)
  * [package hierarchy](https://go101.org/article/packages-and-imports.html#package)

### 1.11.c (2018/Sep/22)

* add a new tip [How to make a struct type incomparable?](https://go101.org/article/tips.html#make-struct-type-uncomparable).
* add a new tip [Try to reset pointers in freed-up slice elements](https://go101.org/article/tips.html#reset-pointers-for-dead-elements).
* add a new tip [Make optimizations by using BCE](https://go101.org/article/tips.html#make-using-of-bce).
* remove "Precedences Of Unary Operators" from [Syntax/Semantics Exceptions In Go](https://go101.org/article/exceptions.html).

### 1.11.b (2018/Sep/09)

* published [Go Tips 101](https://go101.org/article/tips.html).

### 1.11.a (2018/Sep/01)

* mention 1.11 new `wasm` GOARCH in [More Go Related Knowledges](https://go101.org/article/more.html#cross-platform-compiling).
* mention 1.11 new `go mod` command in [The Go Toolchain](https://go101.org/article/go-toolchain.html).

### 1.10.g (2018/Jun/02)

* published [About Go 101](https://go101.org/article/101-about.html).
* published [Acknowledgments](https://go101.org/article/acknowledgements.html).

### 1.10.f (2018/May/15)

* published [Relections in Go](https://go101.org/article/reflection.html).
* added a channel use case: [rate limiting](https://go101.org/article/channel-use-cases.html#rate-limiting).


### 1.10.e (2018/Apr/28)

* added a new detail: [Exit a program with a <code>os.Exit</code> function call and exit a goroutine with a <code>runtime.Goexit</code> function call.](https://go101.org/article/details.html#os-exit-runtime-goexit).

### 1.10.d (2018/Apr/18)

* added a new detail: [Non-exported method names and struct field names from different packages are viewed as diffferent names.](https://go101.org/article/details.html#non-exported-names-from-different-packages).
* added a FAQ question: [What does the compiler error message `declared and not used` mean?](https://go101.org/article/unofficial-faq.html#error-declared-not-used)
* added a FAQ question: [What is the difference between the function call <code>time.Sleep(d)</code> and the channel receive operation <code>&lt;-time.After(d)</code>?](https://go101.org/article/unofficial-faq.html#time-sleep-after)
* added a FAQ question: [What is the difference between the random numbers produced by the <code>math/rand</code> standard package and the <code>crypto/rand</code> standard package?](https://go101.org/article/unofficial-faq.html#math-crypto-rand)
* added a FAQ question: [What are the differences between the <code>fmt.Print</code> and <code>fmt.Println</code> functions?](https://go101.org/article/unofficial-faq.html#fmt-print-println)
* added a FAQ question: [What are the differences between the built-in <code>print</code>/<code>println</code> functions and the corresponding print functions in the <code>fmt</code> and <code>log</code> standard packages?](https://go101.org/article/unofficial-faq.html#print-builtin-fmt-log)
* added a FAQ question: [Why isn't there a <code>math.Round</code> function?](https://go101.org/article/unofficial-faq.html#math-round)
* added a FAQ question: [What does the word <b><i>gopher</i></b> mean in Go community?](https://go101.org/article/unofficial-faq.html#gopher)

### 1.10.c (2018/Apr/14)

* finished the article [some common concurrent programming mistakes](https://go101.org/article/concurrent-common-mistakes.html).
* published [Go details 101](https://go101.org/article/details.html).
* unhid [Go FAQ 101](https://go101.org/article/unofficial-faq.html).

### 1.10.b (2018/Apr/06)

* added [an interesting type embedding example](https://go101.org/article/type-embedding.html#dead-loop-example)
* mentioned [receive-only channels can't be closed](https://go101.org/article/channel.html#assign-and-compare)
* mentioned [indexes in array and slice composite literals must be constants](https://go101.org/article/container.html#value-literals)

### 1.10.a (2018/Mar/31)

First release, though some articles are still not finished.


