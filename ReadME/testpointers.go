// this is just me trying to learn more about pointers
// it has nothing to do with the application
package main

import "fmt"

func main() {

	var x int = 50
	// var s string = "string"
	// &x is a pointer :D
	var y *int = &x // will not accept var s
	var k *int = y
	fmt.Println(y)  // 0xc00000a0b8
	fmt.Println(&y) // 0xc000044020
	fmt.Println(*y) // 50
	fmt.Println(k)  // 0xc00000a0b8
	fmt.Println(&k) // 0xc000044028
	fmt.Println(*k) // 50
	*y = 100
	fmt.Println(k)  // 0xc00000a0b8
	fmt.Println(&k) // 0xc000044028
	fmt.Println(*k) // 100

	fmt.Println(y)  // 0xc00000a0b8
	fmt.Println(&y) // 0xc000044020
	fmt.Println(*y) // 100

	fmt.Println(x) // 100

	var i int = 40
	j := &i
	fmt.Println(j)  // 0xc00000a108
	fmt.Println(&j) // 0xc000044030
	fmt.Println(*j) // 40
	*j = 80
	fmt.Println(j)  // 0xc00000a108
	fmt.Println(&j) // 0xc000044030
	fmt.Println(*j) // 80

	fmt.Println(i) // 80

}

/*
a pointer is a variable that stores the memory address of another variable. Pointers are useful for sharing large amounts of data efficiently and for modifying the data that is passed to a function.

Here’s a breakdown of the key concepts related to pointers in Go:

Declaring Pointers: You declare a pointer to a specific type using an asterisk (*) before the type. For example, var ptr *int declares a pointer to an integer.
Go

var x int = 1
var p *int = &x
In this example, p is a pointer to an integer. The & operator is used to get the memory address of x, so p now points to x.

Dereferencing Pointers: You can access the value that a pointer points to using the * operator. This is known as “dereferencing” the pointer.
Go

fmt.Println(*p) // prints 1
In this example, *p gives you the value that p points to, which is 1.

Pointer to a Struct: Pointers can also point to complex types like structs. This is often used in functions that need to modify the struct or when the struct is large and would be expensive to copy.
Go

type Person struct {
    Name string
    Age  int
}

func birthday(p *Person) {
    p.Age++
}

func main() {
    var bob Person = Person{"Bob", 30}
    birthday(&bob)
    fmt.Println(bob.Age)  // prints 31
}
In this example, the birthday function takes a pointer to a Person struct. It increments the Age field of the struct that the pointer points to.

Nil Pointers: A nil pointer is a pointer that doesn’t point to any value. The zero value for pointers of any type is nil.
Go

var p *int = nil
In this example, p is a pointer to an integer that doesn’t point to any value.

Pointers and Arrays: When you pass an array to a function in Go, it receives a copy of the array, not a pointer to it. If you want to modify the original array, you need to pass a pointer to it.
Go

func addOne(p *[3]int) {
    for i := range *p {
        (*p)[i]++
    }
}

func main() {
    arr := [3]int{1, 2, 3}
    addOne(&arr)
    fmt.Println(arr)  // prints [2 3 4]
}
In this example, the addOne function takes a pointer to an array of integers. It increments each element in the array that the pointer points to.
*/
