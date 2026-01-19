/*
	📝 Type declaration digunakan untuk membuat tipe data baru (custom type) dari
	tipe data yang sudah ada.
	Tujuannya:
	1. Lebih jelas dan deskriptif
	2. Lebih aman (type safety). Karena tipe data di Go cukup beragam, maka akan lebih aman jika membuat 1 type untuk digunakan berkali-kali
	3. Lebih mudah dibaca
*/

package main

import "fmt";

func main(){
	type NoKTP string;
	type IsMarried bool;

	var idCardNumber NoKTP = "352213501920001";
	var status IsMarried = true; 

	fmt.Println(idCardNumber);
	fmt.Println(status);
}