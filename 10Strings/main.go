package main
 
import(
	"strings"
        "fmt"
) 

func main() {

	//creating strings

	var greeting = "hello, world!"

	fmt.Printf("normal string: ")

	fmt.Printf("%s", greeting)
	fmt.Printf("\n")


	//display the hexadecimal byte values of the string

	fmt.Printf("HEX bytes: ")
	for i := 0; i < len(greeting); i++ {
		fmt.Printf("%x", greeting[i])
	}
     fmt.Printf("\n")
	//measuring string length

	fmt.Printf("the length of this string is : ")
	fmt.Println(len(greeting))


	// string concatentation
	// create a slice of strings

	fruits := []string{"apples", "oranges", "and bananas"}

	// concatenating using strings.join

	fmt.Println(strings.Join(fruits, " "))


	
	
}