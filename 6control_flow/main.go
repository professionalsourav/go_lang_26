package main

import "fmt"


func main(){
	age := 17

	if age >= 18{
		fmt.Println("you are eligble to vote")
	} else {
		fmt.Println("you are not eligble to vote")
	}



	score := 8

	if score >= 90 {
		fmt.Println("grade:A")
	} else if score >=80 {
		if score >= 85 {
			fmt.Println("grade: B+")
		} else {
			fmt.Println("grade: B")
		}

	} else {
		fmt.Println("grade: C")
	}
}