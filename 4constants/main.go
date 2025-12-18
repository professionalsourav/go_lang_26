package main

import "fmt"

func main(){
	//declare strings and numeric constants

	const NAME string = "sourav kumar"
	const PRICE int = 100
	const PI float64 = 3.141

	fmt.Println("name is :", NAME)
	fmt.Println("price is :", PRICE)
    fmt.Println("Pi vaule is :", PI)


	const (
		DECIMAL = 255 //decimal with no prefix
		OCTAL = 0377  //octal with leading 0
		HEXADECIMAL = 0xff  //hexadecimal with leading 0x
	)


	fmt.Println("Decimal:", DECIMAL, "octal:", OCTAL, "hexadecimal:", HEXADECIMAL)


	//floating point literal

	//a floating-point literal can have an integer part, a decimal point, a fractional part and an exponent part.

	const PII float64 = 3.141
	 fmt.Println("Pi vaule is :", PII)

	 const AVOGADRO = 6.022e23
	 fmt.Println("AVOGADRO value is : ", AVOGADRO)

	 //escape sequences in string literals

	 //this is a new line
	 const GREETING = "Hello, Earth!\n"
	 //this is a quote
	 const QUOTE = "\"GO is simple!\" - A programmer!"
	 fmt.Println(GREETING)
	 fmt.Println(QUOTE)

	 //multi-line and concatenated string literals
	 const MULTILINE = "meeaaooooooooooooooooooooooo" +"hiiihiihiiihhiiiihihihiihiihiii"
	 const CONCATENATED = "concatenated" + "string"
	 fmt.Println(MULTILINE)
	 fmt.Println(CONCATENATED)

	 //boolean constants
	 const ACTIVE = true
	 const READY = false
	 fmt.Println("ACTIVE:", ACTIVE, "READY:", READY)

	 //constants for calculations

	 const LENGTH = 50
	 const WIDTH  = 67
	 const AREA  = LENGTH * WIDTH
	 fmt.Println("area is :", AREA)
}