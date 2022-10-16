package main       //for standalone executable

import "fmt"     // formatted I/O

// main function is the first function to be called when program is executed
func main() {
	fmt.Println("Hello, World!") 

	// variables()
	// types()
	// loop()
	// arrays()
	// maps()
	// pointers()
	structs()
}



func variables(){
	// declare a single variable a
	var a int       
	const constant_value = 32

	// declare multiple variables
	var (
		b bool
		c float32
		d string
	)

	a = 54
	b,c = true, 3.142
	d = "learning go programming language"
	fmt.Println(a,b,c,d,constant_value)


	x := 42            // Initialize and assign to a single variable
	y, z := true, 32.0 // Initialize and assign to multiple variables
	w := "string"
	fmt.Println(x, y, z, w) // 42 true 32 string
}




// Go offers a rich collection of types, including numericas, booleans, strings, error, and the ability to create custom types.
func types(){
	/* User specified types */
	const a int32 = 12         // 32-bit integer
	const b float32 = 20.5      // 32-bit float
	var c complex128 = 1 + 4i  // 128-bit complex number
	var d uint16 = 14          // 16-bit unsigned integer

	 /* Default types */
	 n := 42              // int
	 pi := 3.14           // float64
	 x, y := true, false  // bool
	 z := "Go is awesome" // string

	// %T is a conversion character called a "verb" in Go. verb denotes the type of the passed variable. It prints the type of the variables;
	 fmt.Printf("user-specified types:\n\t%T %T %T %T\n", a, b, c, d)
	 fmt.Printf("default types:\n\t%T %T %T %T %T\n", n, pi, x, y, z)
}



func loop(){
	var loop_type int
	
	fmt.Println("Loop type:\n1. For loop\n2.While loop\n")
	fmt.Println("Make choice (1 or 2) :")
	fmt.Scan(&loop_type)

	switch loop_type{
	case 1:
		str := "For loop complete"
		for i := 0; i<len(str); i++ {
			fmt.Println("Inside for loop")
		}
	case 2:
		i := 0
		for i < 10 {
			fmt.Println("While uses for,.... :)")
			i++
		}
	default:
		fmt.Println("It's the default case :(")
}
}

func arrays(){
	// an array of size 4 containing string data type
	// var arr = [4]string{"ayanokoji", "horikita", "sudo-ken", "ryuen"}

	// for i := 0; i < len(arr); i++ {
	// 	fmt.Println(arr[i])
	// }

	// for index, val := range arr {
	// 	fmt.Println(index, val)
	// }

	x := []float64{2.15, 3.14, 42.0, 29.5}
	fmt.Printf("%v", x)

	// array slicing
	anime_list := [8]string{"classroom of the elites", "death note", "Attack on Titans", "The Promised Neverland", "Spy-X Family", "Fullmetal Alchemist", "Tokyo Ghoul", "Code Geass"}

	seen := anime_list[0:5]

	fmt.Printf("\nseen anime : %v\n", seen)
}


// maps in go are similar to dictionaries in python
func maps(){
	/* Define a map containing the release year of several languages */
	firstReleases := map[string]int{
		"C": 1972, "C++": 1985, "Java": 1996,
		"Python": 1991, "JavaScript": 1996, "Go": 2012,
	}

	/* Loop through each entry and output the name and release year */
	for k, v := range firstReleases {
		fmt.Printf("%s was first released in %d\n", k, v)
	}
}


func pointers(){
	var ptr *int     // declare pointer to an int
	num := 19
	ptr = &num
	value := *ptr

	fmt.Printf("Address of num : %v\n", ptr)
	fmt.Printf("Value of num : %v\n", value)
}
