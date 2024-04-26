# GO Notes

## Resources

- [Learn Go with Tests](https://quii.gitbook.io/learn-go-with-tests)
- [Go Exercism track](https://exercism.org/tracks/go/concepts)
- [Go Koans](https://github.com/cdarwin/go-koans)
- [Go Documentation](https://golang.org/doc/)
- [Effective Go](https://golang.org/doc/effective_go.html)
- [Go by Example](https://gobyexample.com/)

## Basics

- Go is a statically typed, compiled programming language.
- All source files in a directory must share the same package name.
- When a package is imported, only entities whose names start with a capital letter can be used / accessed.
- Variables can be defined by explicitly specifying a type.
- You can also use an initializer, and the compiler will assign the variable type to match the type of the initializer.
- Constants are defined using the `const`.
- Go functions accept zero or more parameters.
    - Parameters must be explicitly typed, there is no type inference.
- Values are returned from functions using the `return` keyword.

## Comments

- Go supports two types of comments:
    - Single line comments are preceded by `//`.
    - Multiline comments are inserted between `/*` and `*/`.
- They are used by the `godoc` command.
- A documentation comment should be a complete sentence that starts with the name of the thing being described and ends with a period.
- Comments should precede packages as well as exported identifiers.

## Numbers

- Go contains basic numeric types that can represent sets of either integer or floating-point values.
    - uint: the set of all unsigned integers.
    - int: the set of all integers.
    - float: the set of all floating-point numbers
    - complex: the set of all complex numbers.
    - byte: alias for uint8.
    - rune: alias for int32.

## Arithmetic Operators

- Arithmetic operations on different types in Go gives an error. 
- You can convert the type of a variable using a type conversion.
- It's possible to use `x++` or `x--`.

## Booleans

- Booleans in Go are represented by the `bool` type.
- A `bool` is either `true` or `false`.
- Go supports three boolean operators: `!` (NOT), `&&` (AND), and `||` (OR).
- The preference order is `!`, `&&`, and `||`.

## Strings Package

## Strings
