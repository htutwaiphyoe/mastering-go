# Go - The Complete Guide

[github](https://github.com/mschwarzmueller/go-complete-guide-resources)

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

```go
if condition {
    if condition {
    
    }
} else if condition {

} else {

}
```

bool => boolean type

for loop only in GO Lang

```go
for i := 0; i < 2; i++ {

}
```

infinite loop

```go
for {

}
```

break => to exit the loop
continue => to skip current loop iteration

switch => no need break for each case

```go
switch variable {
 case value:
 ...
}
```

file module => os
os.WriteFile(fileName, []byte(string), permission)

0644 => file permission in linux (read/write)

_ for unused value

os.ReadFile(fileName)
change data []byte to string and change string to float64 with strconv.ParseFloat(string(data), 64)

readFile does not break app if there is no file. just return empty byte

error handling => err != nil

custom error => errors package => errors.New(string)

panic(string) => console.error => break the app

every go file must be under a package

if same package, no need to import/export

import are not shared

every package must be in an own sub folder of project

name of the folder is the same name as the package name

file can be any name

import need the whole module path and file path

to export functions or variables, must be start with Uppercase

import => package.Uppercase

third-party lib => [pkg](https://pkg.go.dev/)

go get library => npm i library

go get => npm i

add library in go.mod file

go has pointers like C

pointers => variable that stores value address instead of value

& => &variable => address of stored values

1. avoid unnecessary value copies => &variable, *variable
2. direct data mutation => *variable =*variable

pointer type => *type

get the pointer value => dereferencing => *variable

Scan uses pointers internally to direct data mutation with user input value

can access properties of struct with pointer address => no need to dereference => shortcut way

method => functions in struct

```go
func (v type) name(){

}
```

receiver

use pointer to mutate struct

custom constructor function for convention

Scanln => to end program when press enter

custom constructor for validation

struct field also must be exported with capital letter

use struct type in embedded struct for direct access of nested field and methods

custom types => alias for build-in types + methods in that alias

Scan => single input

multiline input

```go
reader := bufio.NewReader(os.Stdin)
value, err := reader.ReadString('\n')

if err != nil {
    return ""
}

value = strings.TrimSuffix(value, "\n")
value = strings.TrimSuffix(value, "\r")
```

json encoding => encoding/json => json.Marshal(struct) => extract public fields of struct only

struct tags => meta for struct fields => `json:"fieldName"` => library picks up metadata to process

embedded interface

```go
type A interface{

}

type B interface{
    A
}
```

any, interface{} => any value type

type switch

value.(type) => only for switch

```go
switch value.(type){
    case int:
    case string:
}
```

type from value => typeof

```go
intValue, ok := value.(int)
```

generic

```go
func add[T int | float64](a, b T) T{
    return a + b
}
```

arrays

```go
numbers := []type{value, value}
```

array => index starts from 0 => array[0]

slices from arrays => array[1:3] => new array start from 1 and until 3 but excluding 3

array[1:] => start from 1 until the end

array[:4] => start from 0 until the 4, excluding 4

if there is no assigned value, add default value => [4]string => 4 empty strings

slice => subset of the array => more than one value, not entire array

cannot start with negative indexes

array can be out of bound, last element => length + 1

slices can be created from another slices

slices reference same memory address of array, so if mutate of the item, it will affect original array

go saves metadata of slices, A slice has both a length and a capacity.

len => number of items in array/slice

cap => capacity of array => the number of elements in the underlying array, counting from the first element in the slice

[len vs cap](https://go.dev/tour/moretypes/11)

dynamic array => []type{} => automatically create slices

cannot assign directly in array index out of bound

use append(slice, item) for dynamic and return new slice with new array and re-assign make overwrite original array

append can add multiple values

spread operator to merge two array => slices...
