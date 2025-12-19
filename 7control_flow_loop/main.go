package main

import "fmt"

func main(){
 //for loop
	for i := 0; i < 11; i++ {
		fmt.Println(i)
	}

	//nested loop control
    //Multiplication table

	for j := 1; j <= 10; j++ {
		for k := 1; k <= 10; k++ {
			fmt.Printf("%d * %d = %d\t", j, k, j*k)
		}
		fmt.Println()
	}

	//loop control statements

	//break statement

	for n := 0; n< 10; n++ {
		if n%5 == 0{
			break
		}

		fmt.Println(n)
	}
     
	//continue statement
  
	for m := 0; m < 10; m++ {
		if m%2 == 0{
			continue
		}

		fmt.Println(m)
	}

     fmt.Println("go to statement start")
	// GOTO

	for p := 0; p < 10; p++ {
		fmt.Println(p)
		if p == 5 {
            goto end
		}
	}
	end:
	fmt.Println("loop ended")

// second example of goto
	fmt.Println("second example of goto start")

	z := 0
start:
    if z < 10 {
		fmt.Println(z)
		z++
		goto start
	}



}

