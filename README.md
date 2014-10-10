Go Pre-Hackathon Web Example
============================

A simple example of a web server in Go.

Go's [`net/http`](http://golang.org/pkg/net/http/) library is kind of a micro-web-framework with almost everything you need for [building a web application](https://golang.org/doc/articles/wiki/).

I chose [Martini](https://github.com/go-martini/martini), a very lightweight framework,
on top on `net/http` to allow for faster newbie start. Martini is compatible with [`http.HandlerFunc`](http://godoc.org/net/http#HandlerFunc) interface which makes it a great complimentary to `net/http` rather than replacing it so you should learn a bit about how `net/http` and `http.HandlerFunc` are working. 

## What does it do?

In this example you will see several examples of processing requests:

* Serving static files. Everything under `/public` is served as a static page. In this example the root path `/` will serve the `/public/index.html` file.
* The path for `GET /hello` will return a simple string that will be rendered as HTML and return a 200 status code
  * The path for `GET /hello/:name` will be a named parameter path
* The path `GET /hello_template` shows the usage of an HTML template file and reading parameters from the query string. (i.e. `/hello_template?name=myname`)
* The path `GET /hello.json` Uses a struct as a template to return a JSON response.
* The path `POST /hello.json` reads a JSON body and returns a JSON response with `201` status code.

## Getting Started

After installing Go and setting up your [GOPATH](http://golang.org/doc/code.html#GOPATH), clone this repository and install the dependencies:

```
$ go get github.com/go-martini/martini
$ go get github.com/martini-contrib/render
```

Now you are ready to run the server with `go run server.go`

####Tip
Use [Gin](github.com/codegangsta/gin) to allow for live code reload
