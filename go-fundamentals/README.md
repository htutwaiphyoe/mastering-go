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
