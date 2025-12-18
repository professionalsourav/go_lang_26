package main
import "fmt"

func main(){

	//way1: declare and assign in two different line
	var mango string

	mango = "this is a fruit"

	var weight int

	weight = 45

	fmt.Println(mango)
	fmt.Println( weight)

	//way2: declare and assign in same line

	var height int = 56

	fmt.Println("mango: ", mango)
	fmt.Println("weight: ", weight)
	fmt.Println("height: ", height)

	//way3: shorthand
	//shorthand notation :=
	//type inference

	// var age = 89

	age:=89
	city := "nyc"
	fmt.Println(age)
	fmt.Println("the name of my city is: ", city)

	var apples, oranges int = 34,67
	fmt.Println("i have", apples, "apples and", oranges, "oranges")

	var fruits =  apples + oranges

	fmt.Println("fruits:",fruits)
 
	var windows, mac, linux string = "windows is ok\n", "mac is meh\n", "linux is goat!\n"

	print(windows,mac,linux)

	//static type declaration

	var x float64 = 20.9
	fmt.Println(x)
	fmt.Printf("x is of type: %T\n", x)

	//dynamic type declaration

	 y := 45

	 fmt.Println(y)
	fmt.Printf("y is of type: %T\n", y)


	//mixed variable declaration

	var a, b, c = 234.7777, 45, "hello world"

	fmt.Println(a)
		fmt.Println(b)
			fmt.Println(c)


			fmt.Printf("a is of type: %T\n", a)
			fmt.Printf("b is of type: %T\n", b)
			fmt.Printf("c is of type: %T\n", c)



}
