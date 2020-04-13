<!-- @format -->

## package main

`package main` declaration is the first thing in a go file and its meant to tell go what is the workspace/project/package the file is part of. \
The reason we have called `main` is because go has two types of packages, `executable` and `resusable`.

**Executable:** \
Is when the code that is compiled and bundled as executable file. Must have a `func` called `main`.

**Reusable:** \
Are packages that are used as dependencies and not generated as executable. One way to differ a resusable from executable is the by the `package main` declaration.

## import

import is used to get code/access defined in another package

## file pattern

We will always find this file pattern in go files.

```
// package declaration
package main

// import required/needed packages
import ("fmt")

// functions and whatnots to do things
func main() {}
```

## pointers

Let's say we have an instance of a struct called `person` which we have the field `name` with the value as `"jon doe"` and we try to change this value like so:

```go
p := person{
  name: "jon doe"
}
p.name = "new name"
```

The above code won't work because at the time of the instantiation we are saving/allocating in memory the person object with name and value, and then when trying to change the name although go won't through any errors we won't be changing the values/data in the original allocation of the person object.

For that we need to use pointer operators:

`&` Get the memory address where originally saved, example `pPointer := &p` \
`*` Get the value of the memory address is pointing to, example `*pPointer`

That being said, Go will allow to omit the `&` when the receiver is a pointer "type".

When to think about the pointers being passed arround? when using value types such as `int`, `float`, `string`, `bool` and `structs`
