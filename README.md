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

## Dependency Injection
- you can do pretty cool things with this, such as rerouting where your print statements go!! this is very neat and useful
- i wonder what other applications of DI exist
- the idea is, you have set methods that you use that are calling specific methods that can realistically put information anywhere a specific interface is implemented
    - using this method, you can create your own methods that will do this, eg printing somewhere else
- I want to learn more about how to use DI. 

### Mocking
- mocking lets you test out a function that would take a lot of resources with fewer resources
- eg
    - test out a database with a test version of said database
    - test out output with a delay with a DId version of output
- it is a way of writing code alongside Dependency Injection that makes it easier and simpler to write tests for certain bits of software
- i dont fully understand so i need more experience and usage of it to find out more

- note: when testing, test functionality, rather than implementation, unless implementation is important

- ask:
    - am i testing behaviour or implementation details?


## Concurrency

- concurrency is basically the idea of using multiple cores to perform an instruction, since programming languages unless prompted otherwise, will simply execute on a single core, since there is often a lot of stuff that happens step by step, and cannot be done simultaneously
- Go does concurrency through the use of the keyword `go`, which is followed by a function that gets done concurrently
    - eg: `go MakeAPie()` where `MakeAPie()` can be executed concurrently

- one thing to make note of, often using multiple cores results in an overhead in managing instructions and which core does what when, as a result it can increase the amount of time take to do an instruction over just one core doing it, same as when you have a lot of humans that are each doing separate tasks, it takes time to make sure they are doing what they need to
    - see [Amdahls Law](https://en.wikipedia.org/wiki/Amdahl's_law)


- concurrency in Go often makes use of anonymous functions, eg:
```go
func(...params) (...returnTypes) {
    return
}
```
- anonymous in the same way that anonymous datatypes are, benefits are:
    - single use* (can be called multiple times in a loop but it makes programming easier)
    - limited scope of variables

### Issues with Concurrency
- you have to be careful with concurrency, go 1.22+ fixed a big issue where if you used a go routine inside a loop, it would use the pointer of the values in the go routine, instead of the values, and throughout the loop the values in question are presumably gonna change, so making it so that the program would execute multiple times using the last value of the looped variables due to it using the reference, instead of using the values during the loop

eg: initially the following: 

```go
for _, url := range urls {
	go func() {
		result[url] = wc(url)
	}()
}
```

would wait to execute the function until after the loop ends, storing the pointer to url, and thus giving the last value of url to all instances of the function. 

Why it executes the function after finishing to run the for loop probably has a good reason, in that go probably assigns each loop to each core, and then they all start executing after they are passed through or something, and since it passes the pointer to url, which changes as the for loop is ran, they will be pointing to the location storing the final value. 

However the following would fix this problem because the method would make a copy of the values passed through. 

```go
for _, url := range urls {
	go func(url string) {
		result[url] = wc(url)
	}(url)
}
```

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



&*a = 300 (i think)
<!-- \*&\*a =  -->

# Test Driven Development

## Testing
- testing is used to figure out if a function is doing what it was intended to do


### Test Tables (Table Driven Tests)
- test tables is a semi-automated way to perform a test when you have lots of cases that use the same general structure that you wanna test
- in Go you do this by creating an anonymous struct (a struct that is only used 1 time), and then iterating through it, and performing the test function on it every single time

There exists something called table driven tests, where you create test cases that the program will automatically run. 

By default, VSCode does it in GO by generating a slice of anonymous structs {basically a datastructure made with no specific name to it, only such that it can be used locally}. Technically you have more freedom in how you do it, but this is a very neat way of doing it. 


## Benchmarks
- benchmarks are an implementation of a test that is used to figure out how long a piece of code runs for

- structure of a benchmark:
    - generate necessary information (either dummy values or nothing)
    - b.ResetTimer() - b of type *testing.B, presumably done mainly if there are multiple benchmarks and to reset the benchmark timer
    - make a for loop in range 0 b.N 
    - run test 


## Errors
- you should be creating errors when testing
- better to have precise errors than reuse errors
- errors are not always necessary, when testing, think whether an error actually matters or not

