package main

import "fmt"




func countryName (s string){
	fmt.Printf("%s is my country \n", s)
}

func stateName (s string) {
	fmt.Printf("%s is where i live \n", s) 
}

func weExpact (x, y int) int {
	if x > y {
 
		return  x
	} else {

		 return  y
	 }
		
	
	    
}


//function to swap two string


func swap(x, y string)(string, string) {
	return y, x
}



//function arguments
//call by value

func increment (i int)  {

	i++
fmt.Println("inside function: ", i)

}

//call by reference

func modify (slice []int){
	slice[0] =999
	fmt.Println("inside modify: ", slice)

}


func main(){
	countryName("india")

	stateName("odisha")

	a, b := 100, 200
	result := weExpact(a, b)
	fmt.Printf("Max value is  %d\n", result)

	FirstName , LastName := swap("sourav", "kumar")

	fmt.Printf("swaped name: %s %s\n", FirstName,LastName)

	//calling increment function value

	x :=10
	increment(x)

	fmt.Println("after incrementing value: ",x)


	//calling modify function by reference

	mySlice := []int{1,2,3,4}

	modify(mySlice)

	fmt.Println("in modify after:", mySlice)

	
}