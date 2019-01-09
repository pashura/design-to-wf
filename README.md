# design-to-wf
Design in, webform out

# document-parsing-service

The document parsing service is a service that pulls parcel payloads and parses information from them.

## Table of Contents

- [1 Go Setup](#1-go-setup)
  - [1.1 Installing Go](#11-installing-go)
  - [1.2 Add to path](#12-add-to-path)
  - [1.3 The Go Workspace](#13-setting-up-your-go-workspace)
  - [1.4 Go Examples](#14-go-examples)
- [2 Install](#2-installing-this-project)
- [3 Build](#3-building-the-project)
  - [3.1 Local](#31-local)
  - [3.2 Docker](#32-docker)
- [4 Run](#4-running-the-project)
  - [4.1 Local](#41-local)
  - [4.2 Docker](#42-docker)
- [5 TODOs](#6-todos)

## 1 Go Setup

### 1.1 Installing Go

This project is written in Go, so you will need to install it if you have not already. Depending on your OS this can be
done in many different ways. I recommend that you just install it using that instructions found [HERE](https://golang.org/doc/install).
Alternatively you can also do the following:

- On Mac OS you can use brew : `brew install go`
- On Ubuntu you can use apt-get or snap : `sudo apt-get install golang-1.9-go` or `snap install --classic go`

*Note* it is recommended to use a Go version 1.10.0 or up.

### 1.2 Add to PATH

Add the go bin folder to path for easy running of programs later. `export PATH=$PATH:$(go env GOPATH)/bin`

### 1.3 Setting up your Go Workspace

Unlike other programming languages Go programmers typically keep all their Go code in a single workspace. This space is called your
GOPATH. This is by default set to $HOME/go on mac and linux. I would recommend just keeping this setting. If you would like to change
this though just set an environment variable: `$GOPATH = /some/other/place/go-workspace`

In your Go path you will find find three folders(after you install something, you can make these too if you want).

```text
go/
    bin/
    pkg/
    src/
```

- `src` contains all of your Go source files
- `pkg` contains package objects used to build up binaries
- `bin` contains your executable programs

For more information on the Go workspace reference [THIS](https://golang.org/doc/code.html#Workspaces) document.

### 1.4 Go Examples

Any of this code confusing? Or would you like a little more clarity on the syntax? Although Stack Overflow is great I recommend [Go by Example](https://gobyexample.com/)!
Check it out!

## 2 Installing this project

To pull this package into your GOPATH run the following command: `go get github.com/mpashura/design-to-wf`
If this does not work try adding the following to your `~/.gitconfig`

```text
[url "git@github.com:"]
    insteadOf = https://github.com/
```

After you have pulled the project run navigate to the project home:
`cd $GOPATH/src/github.com/mpashura/design-to-wf`
Once there get everything setup by running: ``

## 3 Running unit tests

```bash
go test ./...
```

..to be continued....


## 3 Building the Project

### 3.1 Local

To locally build the project run the following command from the root dir of the project: 

### 3.2 Docker

The docker build/deploy assume you have Docker installed on your machine. If you don't you can download it from [HERE](https://www.docker.com/community-edition)

To build the Dockerfile run: 

## 4 Running the Project

### 4.1 Local

Run: 

### 4.2 Docker

Run: 


## TODOS

 - 