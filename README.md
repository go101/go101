**[Go 101 in Leanpub store](https://leanpub.com/go101)** | **[Go 101 in Apple Books store](https://itunes.apple.com/us/book/id1459984231)** | **[Go 101 in Kindle store](https://www.amazon.com/dp/B07Q3HWZ98)** | **[eBooks](https://github.com/go101/go101/releases)** | **[update history](UPDATES.md)** | **[wiki](https://github.com/go101/go101/wiki)**

----

<b>[Go 101](https://go101.org)</b> is a book focusing on Go syntax/semantics and all kinds of runtime related things.
It tries to help gophers gain a deep and thorough understanding of Go.
This book also collects many details of Go and in Go programming.
The book is expected to be helpful for both beginner and experienced Go programmers.

To get latest changes of Go 101, please follow the official twitter account: [@go100and1](https://twitter.com/go100and1).

### Install, Update, and Read Locally

If you use Go toolchian v1.16+, then you don't need to clone the project respository:

```shell
### Install or update.

$ go install -tags=embed go101.org/go101@latest

### Read. (GOBIN path, defaulted as GOPATH/bin, should be set in PATH)

$ go101
Server started:
   http://localhost:55555 (non-cached version)
   http://127.0.0.1:55555 (cached version)
```

If you use Go toolchian v1.15-, or you would make some modifications (for contribution, etc.):
```shell
### Install.

$ git clone https://github.com/go101/go101.git

### Update. Enter the Go 101 project directory (which
# contains the current `README.md` file), then run

$ git pull

### Read. Enter the Go 101 project directory, then run

$ go run .
Server started:
   http://localhost:55555 (non-cached version)
   http://127.0.0.1:55555 (cached version)
```

The start page should be opened in a browser automatically.
If it is not opened, please visit http://localhost:55555.

Options:
```
-port=1234
-theme=light # or dark (default)
```

### Contributing
Welcome to improve Go 101 by:
* Submitting corrections for all kinds of mistakes, such as typos, grammar errors, wording inaccuracies, description flaws, code bugs and broken links.
* Suggesting interesting Go related contents.

Current contributors are listed on [this page](https://go101.org/article/acknowledgements.html).

Translations are also welcome. Here is a list of the ongoing translation projects:
* [中文版](https://github.com/golang101/golang101)

### License

Please read the [LICENSE](LICENSE) for more details.
