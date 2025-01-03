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

go mod init path => create go.mod module file

go build => create OS-level executable file

./file => run executable file

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

Print => same line

if variable is not assigned, there is a default value.

float64 => 0

if there is no initialization, need to add var keyword and set a type => var age int

user console input => fmt.Scan(&variable)

[fmt](https://pkg.go.dev/fmt)

backtick ` for multiline string

Sprintf for returning formatted string

function => can return multiple return values, separated by comma and can return explicit value

```go
fun name(arg type, ...) (returnValue type...){
    return ...
}
```

no need to add return value if there is explicit return values
