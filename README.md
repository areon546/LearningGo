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

## Structs

- structs are how GO handles datatype development

There exists something called table driven tests, where you create test cases that the program will automatically run. 

By default, VSCode does it in GO by generating a slice of anonymous structs {basically a datastructure made with no specific name to it, only such that it can be used locally}. Technically you have more freedom in how you do it, but this is a very neat way of doing it. 

There are 2 ways to create a datatype in GO:
1. ```type ND struct {
        aND int
        b string
} ```
2. ```type NDS datatype ```

The second is used if you dont wanna create any unique functionality but wanna call it a different thing for type safety and the like. 

The first is basically used to create classes, however, you cannot create methods in that way. You need to do the following for methods:

```go
func (n ND) fName(...params) (...returnTypes) {}
```
outside of the struct declaration. 

To assign values to a datatype, 
1. int(1) can be used
2. ND{a, b} can be used - more for structs, I presume it uses order of usage or assignment via aND:a and so on
3. var ndVar NDS = datatype() (implicitly or explicitly)


## Maps

- interesting feature of maps: 
- can be nil
    - nil maps can be read from and return that there is nothing there, but NEVER write to a nil map
    - \therefore always initialise a map
        - eg use make(map[string]string, or map[string]string{})
- can return 2 values
    - value, found := m["key"], value being the value stored at "key" (if there is one, else it returns ""), and found being a bool of whether it is found









# Formatting
[link on formatting](https://pkg.go.dev/fmt)
%q - strings
%d - numbers
%v - default formatting of the datatype
%#v - default formatting, showing field value pairs
%T - datatype



# Pointers
prefixxes to datatypes:
`*` - pass by pointer, lets you edit the original values, not a copy
nothing - copies value, no pointer schenanigans
`&` - the memory location of the variable
`*&` - the * makes the compiler use the following section as the pointer to a spot, the & makes the section a mem address, essentially a pointer, so `*&` is like doing nothing

## Dereferencing pointers
- dereferencing pointers and referencing values is a way of moving around through memory addresses through the use of  stuff like `*` and `&`, `*` being the pointer, `&` being the memory address. im still not entirely sure what they mean

    | Memory Address | Value |
a = | 100            | 200   |
    | 200            | 300   |
    | 300            | 100   |

a = 200
&a = 100
*&a = 300


# Errors
- you should be creating errors when testing
- better to have precise errors than reuse errors
- errors are not always necessary, when testing, think whether an error actually matters or not






&*a = 300 (i think)
<!-- \*&\*a =  -->
