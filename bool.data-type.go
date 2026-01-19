package main

import (
	"fmt"
	"reflect"
)

func main(){
	isGraduated := true;
	isActive := false; 

	fmt.Printf("%T", isGraduated);
	fmt.Printf("%T", isActive);
	fmt.Println(reflect.TypeOf(isGraduated))
	fmt.Println(reflect.TypeOf(isActive))
}