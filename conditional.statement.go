package main

import "fmt"

func main() {
	score := 75;

	if score > 85{
		fmt.Println("Student Grade is A")
	}else if score >= 75 {
		fmt.Println("Student Grade is B")
	}else{
		fmt.Println("Student Grade is C")
	}
}
