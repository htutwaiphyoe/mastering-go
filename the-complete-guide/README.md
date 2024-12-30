# Go - The Complete Guide

GO => open-source programming language by Google

Simplicity, Clarity, Scalability, High Performance, Concurrency, built-in standard library, static typing

Networking, APIs, Microservices, CLI tools

Go => compiled language

<https://go.dev/>

file extension => .go

go run file.go

each file needs to have under a package

main => main entry point

module => multiple packages

go mod init path

go build

main function will be called when program is started

don't have two main functions in different files under same main packages

variable => var variableName

operators => +, *, /

types => int, float64

can't mixed different types since static types

int to float => float64(int)

Println => new line

go run . => run module

go can infer type of values

explicit type assignment => int, float64, string, bool

:= => declare + assign + infer => convention

multiple variable declaration and assignment in one line but cannot add multiple type assignments

constant => const

go formatter settings

```json
"[go]": {
    "editor.insertSpaces": true,
    "editor.formatOnSave": true,
    "editor.defaultFormatter": "golang.go"
}
```

user console input => fmt.Scan(&variable)
