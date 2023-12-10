# Maintaining all notes here from Nana's teachings and others

# Self-task : Try rewriting this code using a different data structure.

- All our code must belong to go package. Go programs are organized into packages.
package main
	```
	package main
	```

- A package for formatted I/O
import "fmt"

- For one go program, you can have only one main function. main is a keyword and the only entry-point to a go project
	```
	func main() {
		fmt.Print("Hello, World!\n")
		fmt.Println("This function adds a new line at the end of this statement.")
	}
	```

> A common practice to define variables at the beginning of a function.

- fmt.Printf is used to print formatted data, i.e. it tells go how to format the passed variable before displaying. Does NOT add a newline character at the EOL.

- Go is a **statically typed** language. Thus you need to tell the Go compiler the data type during a variable declaration. However, if the assignment is done on the same line of declaration, Go compiler can detect and assign the data type on its own.

	Why that?
	-> Explicit type declaration ensure more robust code and reduces the likelihood of errors. It also helps developers to catch mismatches sooner(at compile time). 

	```
	const conferenceName = "Go Conference"

	var userName string
	userName = "Tom"
	```

> A big advantage of Go is that **errors** are detected **during compile time** rather than runtime as against many other programming languages.

- Arrays in Golang cannot contain multipe data types. It can store only a single data type.

	*An array of 10 string elements of which first 3 elements are declared right away.*	
	```
	var names = [10]string{"john", "alice", "bob"}
	```

- **Slices are dynamic arrays!**
```
// var bookings []string
// var bookings = []string{}
bookings := []string{}
bookings = append(bookings, thing_to_append)

fmt.Printf(bookings[0]) // accessing the values from slice is the same as static arrays
```

### In Golang, you can return as many number of variables from a function as you want.
