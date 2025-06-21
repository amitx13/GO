GOOOOOOOOlang

**WHY GO?**
```
1. Go is a statically typed, compiled language. which means that the type of a variable is known at compile time.
2. Fast and lightweight
3. Easily concurrent(Multiple tasks make progress at the same time, but they may not run simultaneously. Think of it as task switching) and (parallel)multi-threaded
4. Garbage collection
5. Go code generally runs faster than interpreted languages (Js, python, Ruby) and compiles faster than other compiled languages (Rust, C, C++, etc.)
6. Go runtime is responsible for managing memory (to cleanup unused memory), garbage collection, and scheduling goroutines.
```
**Lexer** : The Job of the lexer is to simple understand that u are following the grammer of the lang. So that the syntax is correct it does it before compilation 

**Memory Management** : Memory allocation and deallocation happens automatically in GO.
*new()* - Allocate memory but no initialization, will get a memory address, zeroed storage - we cannot put any data initialy
*make()* - Allocate memory and init, will get a memory address, non-zeroed storage - can put any data.
*GC* - Garbage collector happens automatically anything that is nil or outof scope is eligible.

**GO ENV (go env)** 
*GOGC* - The GOGC variable sets the initial garbage collection target percentage. A collection is triggered when the ratio of freshly collected data to live data remaining after the previous collection reaches this percentage. The dafault is GOGC=100

-: *Notes* :-

Case insensitive.
Variable type should be known in advance.
Almost Everything is a Type in Go.
Formating string in Go:
    %v - represents the value use as a default format
    %d - represents an integer
    %f - represents a floating-point number
    %t - represents a boolean
    %s - represents a string
    %q - represents a string as it is quoted
    %T - represents the type of the value

**Conditionals**

if size > 5 {
    fmt.Println("You are large enough!")
}

we can also define variable in the initial if stmt
if size := 7; size > 5{
    fmt.Println("f u are big")
}

why pointer?

Sometimes when u pass on these variable a copy of this variable is passed on whenever there is a case when u want to avoide such this cases and u want absolute
garentee that always actual value should be passed on then we prefer that a pointer should be passed. A pointer is just a direct reference to a memory address

