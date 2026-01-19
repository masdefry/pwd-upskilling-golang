package main;

import "fmt";

func main(){

	/*
		___NUMBER CONVERSION___
	*/
	var nilai32 int32 = 130;
	var nilai64 int64 = int64(nilai32);
	var nilai8 int8 = int8(nilai32); // Konversi ini menyebabkan integer overflow karena nilai berada di luar rentang int8 (-128 hingga 127)

	fmt.Println(nilai32);
	fmt.Println(nilai64);
	fmt.Println(nilai8);

	


	/*
		___STRING CONVERSION___
	*/
	name := "M Defryan";
	firstCharacterByte := name[0];
	firstCharacterString := string(firstCharacterByte);

	fmt.Println(name);
	fmt.Println(firstCharacterByte);
	fmt.Println(firstCharacterString);
}