# LearningGo

## Hello World
Test Driven Development (TDD) Cycle:
- Write a test
- Make the compiler pass
- Run the test, see that it fails and check the error message is meaningful
- Write enough code to make the test pass
- Refactor

## Iteration
[for loops in GO](https://gobyexample.com/for)

## Arrays and Slices
in GO:
- array - fixed size
- slice - dynamic size

- types of function are categorised by their number of parameters
    - variadic functions - variable number of parameters, see [variadic functions](https://gobyexample.com/variadic-functions)


- GO doesnt like comparing slices, reasonable, so to do that you can call `reflect.DeepEqual(x, y)` on 2 variables to see if they are equal
    - this is like doing .equals in Java


- you can assign a function to a variable in GO and use it locally as that func, neat

i feel like i should spend more time on arrays here

## Formatting
[link on formatting](https://pkg.go.dev/fmt)
%q - strings
%d - numbers
%v - default formatting of the datatype
%#v - default formatting, showing field value pairs


## Structs

- structs are how GO handles datatype development

There exists something called table driven tests, where you create test cases that the program will automatically run. 

By default, VSCode does it in GO by generating a slice of anonymous structs {basically a datastructure made with no specific name to it, only such that it can be used locally}. Technically you have more freedom in how you do it, but this is a very neat way of doing it. 

## Pointers

# Pointers
prefixxes to datatypes:
`*` - pass by pointer, lets you edit the original values, not a copy
nothing - copies value, no pointer schenanigans
`&` - the memory location of the variable

# Dereferencing pointers