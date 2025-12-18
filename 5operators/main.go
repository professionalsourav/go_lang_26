package main

import"fmt"

func main(){
	A := 10
	B := 20
	C := 11
	D := 3
	fmt.Println("A + B =", A + B) //addition
	fmt.Println("A - B =", A - B )  //subtraction
	fmt.Println("A * B =", A * B)  //multiplication
	fmt.Println("A / B =", A/B)   //Division
	fmt.Println("C % D =", C%D) //modulous
	A++
	fmt.Println("A++ =", A)  //increment

	A--
	fmt.Println("A-- =", A)

	//relational operators

	fmt.Println("A == B", A == B) //equals- false
	//NOT equal to
	fmt.Println("A != B", A != B) //not quals - true

	// greater than
	fmt.Println("A > B", A > B)

	//less than
	fmt.Println("A < B", A < B)

	//greater than or equal to 

		fmt.Println("A >= B", A >= B)

//logical operator

x := true
y := false

//AND
fmt.Println("x && y:", x && y)

//OR
fmt.Println("x || y:", x || y)

//NOT
fmt.Println("!x:", !x)
fmt.Println("!y:", !y)

//assignment operator
  p := 34
  q := 67

  p += q

  fmt.Println("p += q:",p )

  p -= q
   fmt.Println("q -= p:",q )
	

   //bitwise operators

   e := 87

   //bitwise AND assignment
   e &= 2
   fmt.Println("e &= 2 =", e)

   //bitwise XOR assignment

   e ^= 2 
    fmt.Println("e ^= 2 =", e)

	//bitwise OR assignment

	e |= 2
	fmt.Println("e |= 2 =", e)


	//miscellaneous operator

	g := 90

	//address-of operator
	ptr := &g

	fmt.Println("Address of g:", ptr)

	//pointer dereference operator
     fmt.Println("value of *ptr:", *ptr)


}