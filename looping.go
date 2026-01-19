/*
	📝
	Golang hanya punya 1 jenis looping, yaitu looping `for`
*/

package main

import "fmt"

func main(){

	/*
		___LOOPING FOR___
	*/
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	/*
		___LOOPING WHILE___
	*/
	start := 1;

	for start <= 5{
		fmt.Println(start);
		start++
	}

	text := "Belajar Pemrograman dengan Golang"

	for index, value := range text{
		fmt.Println(index)
		fmt.Println(value) // Masih dalam bentuk byte
		fmt.Println(string(value)) // Dikonversi menjadi string
	}
}