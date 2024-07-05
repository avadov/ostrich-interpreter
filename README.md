# Ostrich

Ostrich is an experimental interpreted programming language. The interpreter has been written in Golang. Ostrich is dynamically typed.

## Getting Started

### Install dependencies

Install Go on Ubuntu

```sh
sudo apt update
sudo apt upgrade
sudo apt install golang
```

Install Go on Arch Linux

```sh
sudo pacman -Syu
sudo pacman -Sy go
```

### Run

```sh
git clone https://github.com/avadov/ostrich-interpreter.git
cd ostrich-interpreter
go run main.go
```

## Basic Syntax

### String literals

```
"This is a string"
```

### Booleans

```
true
false
```

### Variable declarations

```
let num = 123;
let name = "Some text";
let ok = true;
let result = 10 * (20 / 2);
```

### Assignment statements

The assignment statement assigns a new value to a variable

```
name = "Another string"
```

### Arithmetic operators

Unary:

`-` negation

Binary:

`+` sum   
`-` subtraction   
`*` multiplication   
`/` division   

### Comparison operators

`==` equal  
`!=` not equal  
`<` less  
`>` greater  

### String operators

`+` concatenation

### Arrays

An array is an ordered list of elements of possibly different types.

```
let myArray = [true, "Some text", 28, func(x) { x * x }];
```

Each element in the array can be accessed individually.

```
myArray[0];
myArray[3](2);
```

There are some built-in functions that work with arrays:

`len(myArray)` - the number of elements in `myArray`;  
`first(myArray)` - returns the first element of `myArray`;  
`rest(myArray)` - returns a *new* array containing all elements of `myArray`, except the first one (the `cdr` function). *It returns a newly allocated array!*;  
`last(myArray)` - returns the last element of `myArray`;  
`push(myArray, "four")` - allocates a new array with the same elements as the old one plus the new, pushed element. Doesn’t modify the given array.

### Dictionaries

A dictionary maps keys to values.

```
let myHash = {"age": 33, true: "a boolean", 99: "an integer"};
```

The index operator gets values out of the dictionary:

```
myHash[true]
myHash["age"]
myHash[99]
```

### Conditionals

```
if (x > 10) {
    puts("everything okay!");
} else {
    puts("x is too low!");
}
```

### Functions

```
let add = func(a, b) { return a + b; };
add(5, 9);
```

Functions are ordinary variables. They can be passed as arguments to other functions.

```
let twice = func(f, x) {
    return f(f(x));
};

let addTwo = func(x) {
    return x + 2;
};

twice(addTwo, 2);
```

### Output

Print the given arguments on new lines:

```
puts("hello", "world", 55, false)
```

## License

Ostrich is available under the permissive MIT license.
