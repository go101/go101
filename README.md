<b>[Go 101](https://go101.org)</b> is a book focusing on Go syntax/semantics and all kinds of details.
The book also tries to help gophers gain a deep and thorough understanding of Go.
The book is expected to be helpful for both beginner and experienced Go programmers.

### Install And Update

Run

```
$ go get -u gitlab.com/go101/go101
```

or

```
$ go get -u github.com/go101/go101
```

to install and update ***Go 101***.

*(NOTE: if your last `go get -u` command run was before __July 28th, 2018__,
please run this command again to install the latest `go101` program.)*

### Run Locally

Add the path of the `bin` folder under `GOPATH`
into `PATH` environment variable to run `go101`.
The default value of the `GOPATH` environment variable
is the path of the `go` folder under the home directory.

```
$ go101
Server started:
   http://localhost:55555 (non-cached version)
   http://127.0.0.1:55555 (cached version)
```

The start page should be opened in a browser automatically.
If it is not opened, please visit http://localhost:55555.

### Contributing
Welcome to improve Go 101 by:
* Submitting corrections for all kinds of mistakes, such as typos, grammar errors, wording inaccuracies, description flaws, code bugs and broken links. 
* Suggesting interesting Go related contents.

Current contributors are listed on [this page](https://go101.org/article/acknowledgements.html).

Translations are also welcome. Here is a list of the ongoing translation projects:
*  [Chinese translation version](https://github.com/Golang101/golang101).

### License
Copyright (c) Tapir Liu. All rights reserved. Please read the [LICENSE](LICENSE) for more details.
