# dont-interface

```
> go run main.go --file test.go
Among your codebase, there are...
    8 parameters declared in functions, 5 of them are evil! (40.00% typed)
    5 values returned from functions, 3 of them are evil! (40.00% typed)
    3 fields declared in structs, 2 of them are evil! (33.33% typed)
    2 values declared, 1 of them are evil! (50.00% typed)
    3 type aliases introduced, 1 of them are evil! (66.67% typed)
Overall, 42.86% of your types are strictly typed (not interface{}).
```

`dont-interface` calculates how many `interface{}` are declared or used in:

- function parameters, e.g. `i1 interface{}, i2 []interface{}, i3 ...interface{}`
- function return values, e.g. `(interface{}, string)`
- struct fields, e.g. `i interface{}`
- value declarations, e.g. `var i interface{}`
- type aliases, e.g. `type NotInterface interface{}`. Note that `type I1 interface{ test1() }` is _not_ a type alias.

This tiny snippet is partially inspired by Section 4.1.1
of [LLVM: A Compilation Framework for Lifelong Program Analysis & Transformation](https://doi.org/10.1109/CGO.2004.1281665)
, where the authors investigated to what extent is the LLVM's type speculation correct (and I know this paper have
nothing to do with structural typing).

### Usage

```
> go install github.com/Ray-Eldath/dont-interface
> dont-interface --file test.go --file main.go
```

### The Last Thing...

> [Don't `interface{}`!](https://youtu.be/yWeuUwpEQfs)
>
> :wink:
