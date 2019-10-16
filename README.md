# GRPC GoLang Example

Project created in course [gRPC [Golang] Master Class: Build Modern API & Microservices](https://www.udemy.com/course/grpc-golang/).

___

## About Golang
![](https://raw.githubusercontent.com/felipesulzbach/grpc-go-example/master/things/go.png)


### What is?

Golang, or simply Go, is an open source language created in 2009 by [Google](https://about.google/intl/en_US/) (by engineers [Rob Pike](https://en.wikipedia.org/wiki/Rob_Pike) and [Ken Thompson](https://en.wikipedia.org/wiki/Ken_Thompson)).
The Go language was created with the goal of having **C** language performance but also focusing more readable and easier to program like **Java** language.


### Some advantages of language
- Incredibly light in terms of memory usage;
- Suppose several concurrent processing because it uses goroutines instead of threads that are found in most programming languages. Competition is one of the language's strengths;
- Compiles very fast;
- Has garbage collector (has been incorporated into its core in order to prioritize performance);
- It is strongly typed.

GoLang intentionally leaves out many features of modern *OOP* languages. Everything is divided into packages. [Google](https://about.google/intl/en_US/) technology has only *structs* instead of *classes*.


### Some companies that have adopted Golang:
- Netflix
- The Economist
- IBM
- GitHub
- Uber
- Docker
- Dropbox
- OpenShift
- Twitter
- [Complete list by country (link here)](https://github.com/golang/go/wiki/GoUsers)


## About gRPC
![](https://raw.githubusercontent.com/felipesulzbach/grpc-go-example/master/things/grpc.png)


### What is?
Bla bla bla...


## About this project

### Prerequisites

- [Visual Studio Code](https://code.visualstudio.com/) or other IDE (Integrated Development Environment);
- [Golang](https://golang.org/);
- Technologies defined in the next topics of this documentation.


### Setting environment variables for Golang

**GOROOT** - Add Golang installation path.
- By default, the path is *C:\Go\bin*.

**GOPATH** - Add the path where the work path will be.
- By default, the path is *C:\Users\user\go\bin*.


### GOPATH structure
![](https://raw.githubusercontent.com/felipesulzbach/grpc-go-example/master/things/default-estructure-go.png)

- BIN: Contains executable commands;
- PKG: Contains compiled files from some libraries.
- SRC:  Contains Go source files, and the created projects;


### Preparing the environment

The following technologies are critical for running/compiling application sources.

Install the [gRPC](https://github.com/grpc/grpc-go):
> go get -u google.golang.org/grpc

Install the [Protobuf](https://github.com/golang/protobuf):
> go get -d -u github.com/golang/protobuf/protoc-gen-go
