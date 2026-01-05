# Go Fundamentals

[Course](https://github.com/firtman/go-fundamentals)

## Introduction

Go -> By Google 2011, from scratch, open source, multi platform

Created by -> Ken Thompson (B and C Language), Rob Pike (Unix), Robert Griesemer (JVM)

Efficient Compilation + Easy to Code + Efficient Execution

## History & Timeline

- strong, static type system
- C-like syntax
- compiled
- multi-paradigm
- garbage-collected
- fast
- single binary compilation

start -> 2007, public released 2009, 1st 2012 -> Backward compatible

## Philosophy

- Simplicity
- Built-in Network and Concurrency
- Vanilla
- CLI

## Setup

Every file must be under a package -> package means folder

default -> main package

entry point -> func main

run -> go run fileName.go

## Use Cases

- can generate executable binary files for different platforms (arm, 86x) and operating systems (Mac, Window, WebAssembly, JS)

## Basics

.go -> file extension

code block -> {}

no style freedom

optional semicolon

case-sensitive

no try, catch

folder is a package

package name can be url or names

entry main function

modules is a group of packages, project with go.mod file with configuration and metadata

go mod init, go build, go run ., go test, go test

## Modules

go mod init moduleName

moduleName -> url pattern -> go.mod

workspaces -> top layer of module -> go work init -> go.work

workspace > module > package

## Variables

constant -> compile time

immutable -> runtime

```go
var name type // default value

const name = value // fixed value -> bool, string or number (NOT IMMUTABLE)

var text string = "Hello" // double quotes

otherText := "Hello" // shortcut of variable initialization (ONLY within function)

```

## Types

- string
- byte (int8), int (int32), int8, int16, int32, int64, unit, uint8, uint16, unit32, uint64
- int8 (-127 - 127), uint8 (0 - 255)
- float32, float64 (json)
- bool, true, false
- ==, !=, <, >, <=, >=, &&, ||, !
- complex64 complex128
- pointer (*, &)
