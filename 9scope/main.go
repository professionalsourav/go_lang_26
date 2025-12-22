package main


import "fmt"

var g int //global variable 

var x int = 25 //variable shadowing

func main (){
	//local variable

	var a , b, c int

	a = 10
	b = 20
	c = a+b
	fmt.Printf("local variable \n")

	fmt.Printf("a = %d, b = %d, c = %d\n", a, b, c)


	//global variable
  fmt.Printf("global variable \n")
	g = b-a 

	fmt.Printf("a = %d, b = %d, global variable g = %d\n", a, b , g)

	 x := 100

	 fmt.Printf("x = %d\n", x)

}