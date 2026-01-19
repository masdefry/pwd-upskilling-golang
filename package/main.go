package main

import (
	"day01/calculation"
	"fmt"
);

/*
	Penamaan `func main` bersifat WAJIB ketika 
	kita menggunakan package yg executable___ dalam hal ini 
	package `main`
*/
func main(){
	fmt.Println("Hello, Golang!");

	// message:= TestOnly();

	result:=calculation.Add(3, 10)

	// fmt.Println(message);
	fmt.Println(result);
}

/*
	📝
	Pemilihan package terdapat 2 jenis: 
	1. Executable
	2. Library

	Untuk file yang executable, gunakan package main. 
	Untuk file yang hanya sekedar menyimpan function seperti layaknya library, gunakan nama package SELAIN main.

	📝 
	Func TestOnly yang ada di file `entity.go` tidak harus di import. 
	Karena `main.go` dan `entity.go` berada di 1 package yang sama, yaitu `main`.

	📝
	Untuk menjalankannya, maka kita harus `go run` untuk kedua file tsb (main.go & entity.go)
*/