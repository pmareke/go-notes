# Go: The Programm Langauge

## Program Structure

### Names

- The names in Go always start with a letter or a underscore.
- `Keywords`:
    - break
    - case
    - chan
    - const
    - continue 
    - default
    - defer
    - else
    - fallthrough
    - for
    - func
    - go
    - goto
    - if
    - import
    - interface
    - map
    - package
    - range
    - return
    - select
    - struct
    - switch
    - type
    - var
- `Constants`: 
    - true
    - fal
    - nil
    - iota
- `Types`:
    - int
    - int8
    - int16
    - int32
    - int64
    - uint
    - uint8
    - uint16
    - uint32
    - uint64
    - uintptr
    - float32
    - float64
    - complex128
    - complex64
    - bool
    - byte
    - rune
    - string
    - error
- `Functions`:
    - make
    - len
    - cap
    - new
    - append
    - copy
    - close
    - delete
    - complex
    - real
    - image
    - panic 
    - recover
- If an entity is declared within a function, it is local to that function.
- The case of the first letter of a name determines its visibility across package boundaries.
    - If the name begins with an upper-case letter, it is exported.
- Go use `camelCase`.

### Declarations

- There are four kind of declarations:
    - var
    - const
    - type
    - func

### Variables

- A `var` declaration creates a variable of a particular type, attaches a name to it, and sets its ini- tial value.
- Either the `type` or the `= expression` part may be omitted, but not both.
- The initial value is the zero value for the type, which is:
    - `0` for `numbers`.
    - `false` for `booleans`.
    - `""` for `strings`.
    - `nil` for `interfaces` and `reference types` (slice, pointer, map, channel, function).
- An alternate form called a short variable declaration may be used to declare and initialize local variables.
    - It takes the form `name := expression`.
    - They're used to declare and ini- tialize the majority of local variables.
- `:=` is a declaration, whereas `=` is an assignment.

### Pointers

- A variable is a piece of storage containing a value.
- A pointer value is the address of a variable.
- Not every value has an address, but every variable does.
- With a pointer, we can read or update the value of a variable indirectly.
- `&x`  means 'address of x'.
- `*p` means 'The variable to which p points'.
- Each component of a variable of aggregate type a field of a struct or an element of an array is also a variable and thus has an address too.
- The zero value for a pointer of any type is `nil`.

### Type Declarations

- A type declaration defines a new named type that has the same underlying type as an existing type.
- For every type `T`, there is a corresponding conversion operation `T(x)` that converts the value `x` to type `T`.
- Conversions are also allowed between numeric types, and between string and some slice types.
- Comparison operators like == and < can also be used to compare a value of a named type to another of the same type, 
or to a value of the underlying type.
    - Two values of different named types cannot be compared directly.
- A named type may provide notational convenience if it helps avoid writing out complex types over and over again.
- Named types also make it possible to define new behaviors for values of the type.

### Scope

- Don’t confuse scope with lifetime.
    - The scope of a declaration is a region of the program text.
    - The lifetime of a variable is the range of time during execution.
- A name declared inside a syntactic block is not visible outside that block.
- There is a lexical block for the entire source code, called the universe block.
- Declarations outside any function, that is, at package level, can be referred to from any file in the same package.
- The scope of a control-flow label, as used by break, continue, and goto statements, is the entire enclosing function.
- A program may contain multiple declarations of the same name so long as each declaration is in a different lexical block.
- When the compiler encounters a reference to a name, it looks for a declaration, starting with the innermost 
enclosing lexical block and working up to the universe block.

## Basic Data Types

- Go’s types fall into four categories: 
    - Basic types.
    - Aggregate types (Arrays and Structs).
    - Reference types (Pointers , Slices, Maps, Functions and Channels).
    - Interface types.

### Integers

### Floating-Point Numbers

### Complex Numbers

### Booleans

### Strings

### Constants


