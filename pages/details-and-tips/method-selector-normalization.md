<h1>Normalization of method selectors</h1>

<p>Go allows simplified forms of some selectors.</p>

<p>For example, in the following program, <code>t1.M1</code> is a simplified form of <code>(*t1).M1</code>,
and <code>t2.M2</code> is a simplified form of <code>(&amp;t2).M2</code>. At compile time, the compiler will normalize the simplified forms to their original full form.</p>

<p>The following program prints <code>0</code> and <code>9</code>, because the modification to <code>t1.X</code> has no effects on the evaluation result of <code>*t1</code> during evaluating <code>(*t1).M1</code>.</p>

<pre><code class="language-Go">package main

type T struct {
	X int
}

func (t T) M1() int {
	return t.X
}

func (t *T) M2() int {
	return t.X
}

func main() {
	var t1 = new(T)
	var f1 = t1.M1 // &lt;=&gt; (*t1).M1
	t1.X = 9
	println(f1()) // 0 
	
	var t2 T
	var f2 = t2.M2 // &lt;=&gt; (&amp;t2).M2
	t2.X = 9
	println(f2()) // 9
}
</code></pre>

<p>In the following code, the function <code>foo</code> runs okay, but the function <code>bar</code> will produce a panic. The reason is <code>s.M</code> is a simplified form of <code>(*s.T).M</code>. At compile time, the compiler will normalize the simplified form to it original full form. At runtime, if <code>s.T</code> is nil, then the evaluation of <code>*s.T</code> will cause a panic. The two modifications to <code>s.T</code> have no effects on the evaluation of <code>*s.T</code>.</p>

<pre><code class="language-Go">package main

type T struct {
	X int
}

func (t T) M() int {
	return t.X
}

type S struct {
	*T
}

func foo() {
	var s = S{T: new(T)}
	var f = s.M // &lt;=&gt; (*s.T).M
	s.T = nil
	f()
}

func bar() {
	var s S
	var f = s.M // panic
	s.T = new(T)
	f()
}

func main() {
	foo()
	bar()
}
</code></pre>

<p>Please note that, selector normalizatons will not made at run time.
For example, in the following program, the modification to <code>s.T.X</code>
has effects on the method values got through reflection and interface ways.</p>

<pre><code class="language-Go">package main

import &quot;reflect&quot;

type T struct {
	X int
}

func (t T) M() int {
	return t.X
}

type S struct {
	*T
}

func main() {
	var s = S{T: new(T)}
	var f = s.M // &lt;=&gt; (*s.T).M
	var g = reflect.ValueOf(&amp;s).Elem().
		MethodByName(&quot;M&quot;).
		Interface().(func() int)
	var h = interface{M() int}(s).M
	s.T.X = 3
	println( f() ) // 0
	println( g() ) // 3
	println( h() ) // 3
}
</code></pre>
