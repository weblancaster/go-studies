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
