
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
* modify some descriptions for
  * [program resource initialization order](https://go101.org/article/packages-and-imports.html#initialization-order)

### 1.11.d (2018/Oct/18)

* remove "Comparing Interface Values" from [Syntax/Semantics Exceptions In Go](https://go101.org/article/exceptions.html).
* add "Comparisons 2" to [Syntax/Semantics Exceptions In Go](https://go101.org/article/exceptions.html).
* modify some descriptions for
  * [comparison rules](https://go101.org/article/value-conversions-assignments-and-comparisons.html#comparison-rules)
  * [package hierarchy](https://go101.org/article/packages-and-imports.html#package)

### 1.11.c (2018/Sep/22)

* add a new tip [How to make a struct type uncomparable?](https://go101.org/article/tips.html#make-struct-type-uncomparable).
* add a new tip [Try to reset pointers in freed-up slice elements](https://go101.org/article/tips.html#reset-pointers-for-dead-elements).
* add a new tip [Make optimizations by using BCE](https://go101.org/article/tips.html#make-using-of-bce).
* remove "Precedences Of Unary Operators" from [Syntax/Semantics Exceptions In Go](https://go101.org/article/exceptions.html).

### 1.11.b (2018/Sep/09)

* published [Go Tips 101](https://go101.org/article/tips.html).

### 1.11.a (2018/Sep/01)

* mention 1.11 new `wasm` GOARCH in [More Go Related Knowledges](https://go101.org/article/more.html#cross-platform-compiling).
* mention 1.11 new `go mod` command in [The Official Go SDK](https://go101.org/article/go-sdk.html).

### 1.10.g (2018/Jun/02)

* published [About Go 101](https://go101.org/article/101-about.html).
* published [Acknowledgements](https://go101.org/article/acknowledgements.html).
* updated [license](LICENSE).

### 1.10.f (2018/May/15)

* published [Relections in Go](https://go101.org/article/reflection.html).
* added a channel use case: [rate limiting](https://go101.org/article/channel-use-cases.html#rate-limiting).


### 1.10.e (2018/Apr/28)

* added a new detail: [Exit a program with a <code>os.Exit</code> function call and exit a goroutine with a <code>runtime.Goexit</code> function call.](https://go101.org/article/details.html#os-exit-runtime-goexit).

### 1.10.d (2018/Apr/18)

* added a new detail: [Non-exported method names and struct field names from different packages are viewed as diffferent names.](https://go101.org/article/details.html#non-exported-names-from-different-packages).
* added a FAQ question: [What does the compiler error message <code>declared and not used</code> mean?](https://go101.org/article/unofficial-faq.html#error-declared-not-used")
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

