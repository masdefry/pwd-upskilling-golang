package main

import "fmt"

func main() {
	First();

	result := Multiply(3, 1000);
	fmt.Println(result);

	resultArea, resultCircumferance := Calculate(10, 3);
	fmt.Println(resultArea);
	fmt.Println(resultCircumferance);
}

func First() {
	fmt.Println("My first function!");
};

func Multiply(num1 int, num2 int) int{
	return num1*num2;
}

func Calculate(length int, height int)(int, int){
	area := length * height;
	circumferance := 2 * (length + height);

	return area, circumferance
}