[link on formatting](https://pkg.go.dev/fmt)
%q - strings
%d - numbers
%v - default formatting of the datatype
%#v - default formatting, showing field value pairs



- types of function are categorised by their number of parameters
    - variadic functions - variable number of parameters, see [variadic functions](https://gobyexample.com/variadic-functions)


- GO doesnt like comparing slices, reasonable, so to do that you can call `reflect.DeepEqual(x, y)` on 2 variables to see if they are equal
    - this is like doing .equals in Java


- you can assign a function to a variable in GO and use it locally as that func, neat

i feel like i should spend more time on arrays here