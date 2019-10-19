# GRPC GoLang Example

Project created from the [Udemy](https://www.udemy.com/) course, [GRPC [Golang] Master Class: Build Modern API & Microservices](https://www.udemy.com/course/grpc-golang/).

___

## About Golang
![](https://raw.githubusercontent.com/felipesulzbach/grpc-go-example/master/things/go.png)


### What is?

&nbsp;&nbsp;&nbsp;&nbsp;Golang, or simply Go, is an open source language created in 2009 by [Google](https://about.google/intl/en_US/) (by engineers [Rob Pike](https://en.wikipedia.org/wiki/Rob_Pike) and [Ken Thompson](https://en.wikipedia.org/wiki/Ken_Thompson)).
The Go language was created with the goal of having **C** language performance but also focusing more readable and easier to program from more robust languages like **Java**.


### Some advantages of language
- Incredibly light in terms of memory usage;
- Suppose several concurrent processing because it uses goroutines instead of threads that are found in most programming languages. Competition is one of the language's strengths;
- Compiles very fast;
- Has garbage collector (has been incorporated into its core in order to prioritize performance);
- It is strongly typed.

&nbsp;&nbsp;&nbsp;&nbsp;GoLang intentionally leaves out many features of modern *OOP* languages. Everything is divided into packages. [Google](https://about.google/intl/en_US/) technology has only *structs* instead of *classes*.


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


## About GRPC
![](https://raw.githubusercontent.com/felipesulzbach/grpc-go-example/master/things/grpc.png)


### What is?

&nbsp;&nbsp;&nbsp;&nbsp;GRPC is a Remote Call Procedures (RPC) framework, a service that handles Remote Call Procedures (RPC) calls. It was developed by the company [Google](https://about.google/intl/en_US/) and was made available for free and open source.

&nbsp;&nbsp;&nbsp;&nbsp;GRPC supports the [Protobuf protocol](https://developers.google.com/protocol-buffers/) by default, making inter-service communication even more efficient. It also supports [HTTP2](https://en.wikipedia.org/wiki/HTTP/2) and [QUIC](https://en.wikipedia.org/wiki/QUIC) communication. But you can also use other messaging protocols such as [JSON](http://www.json.org/) and [XML](https://en.wikipedia.org/wiki/XML).


### What is it for?

&nbsp;&nbsp;&nbsp;&nbsp;GRPC was created by [Google](https://about.google/intl/en_US/) for the purpose of connecting microservices to their data centers. In addition to being able to apply it to communication between microservices, it can also be applied to connect mobile applications and browsers to backend services.


### GRPC vs REST (by [Stephane Maarek](https://www.udemy.com/user/stephane-maarek/))
| GRPC                                                                           | REST                                                                          |
|--------------------------------------------------------------------------------|-------------------------------------------------------------------------------|
| Protocol Buffers - smaller, faster                                             | JSON - text based, slower, bigger                                             |
| HTTP/2 (lower latency) - from 2015                                             | HTTP1.1 (higher latency) - from 1997                                          |
| Bidirectional and Async                                                        | Client => Server request only                                                 |
| Stream Support                                                                 | Request / Response support only                                               |
| API Oriented - no constraints, free design                                     | CRUD Oriented (Create - Retrieve - Update - Delete / POST GET PUT DELETE)     |
| Code Generation through Protocol Buffers in any language - first class citizen | Code generation through OpenAPI / Swagger (add-on) - second class citizen     |
| RPC Based - GRPC does the plumbing for us                                      | HTTP verbs based - we have to write the plumbing or use a third party library |


## About this project

### Prerequisites

- [Visual Studio Code](https://code.visualstudio.com/) or other IDE (Integrated Development Environment);
- [Golang](https://golang.org/);
- Technologies defined in the next topics of this documentation.


### Setting environment variables for Golang

&nbsp;&nbsp;&nbsp;&nbsp;In the Golang installation, the required environment variables are added. The variables and paths are as follows:

**GOROOT** - Golang installation path.
- By default, the path is `C:\Go\bin`.

**GOPATH** - The path where the workspace path will be.
- By default, the path is `%USERPROFILE%\go`.

  **NOTE:** *In Golang installation, the `go` directory is not created by default, only referenced in environment variables. Therefore, create the `go` directory in `%USERPROFILE%` if it does not exist. (Example: The directory should look like: `C:\Users\Felipe\go`.)*


### Clone project

&nbsp;&nbsp;&nbsp;&nbsp;To maintain project package imports, the project must be cloned into the `go` subdirectory. It will be `%USERPROFILE%\go\src\github.com\_dev\`.

**IMPORTANT:** *If you chose to clone the project in another directory structure, you will need to adjust `imports` to resolve the errors.*


### Preparing the environment

&nbsp;&nbsp;&nbsp;&nbsp;The following technologies are critical for running/compiling application sources:

- Download the Protoc installation [link here](https://github.com/protocolbuffers/protobuf/releases), for `*.proto` files compilations and follow these steps (Example: https://github.com/google/protobuf/releases/download/v3.10.0/protoc-3.10.0-win64.zip):
  - Extract all to `C:\proto3`;
  - Add the **PROTOROOT** environment variable with the value `C:\proto3\bin`.
- Access the application directory from the terminal and execute:
  - Install the [GRPC](https://github.com/grpc/grpc-go):
    > go get -u google.golang.org/grpc
  - Install the [Protobuf](https://github.com/golang/protobuf):
    > go get -u github.com/golang/protobuf/proto

    > go get -u github.com/golang/protobuf/protoc-gen-go
