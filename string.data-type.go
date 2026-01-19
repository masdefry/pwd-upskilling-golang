/*
	✅ String
		String di Golang adalah kumpulan byte (bytes) yang merepresentasikan teks.

		📄Declaration
			➡️ Double Quote ("")
				▪️ Mendukung escape character (\n, \t, dll)

			➡️ Backtick (``)
				▪️ Raw string (string mentah)
				▪️ Tidak memproses escape character
				▪️ Cocok digunakan untuk raw query untuk menghindari SQL injection

		📄Format Verb
			Format verb di Go adalah kode khusus (placeholder) yang digunakan bersama fungsi formatting seperti:
			▪️fmt.Printf
			▪️fmt.Sprintf
			▪️fmt.Fprintf
			Tujuannya untuk mengontrol cara sebuah nilai ditampilkan ke output (terminal, string, file, dll).

			➡️ Jenis-jenis format verb:
				▪️%s   // string
				▪️%q   // string dengan tanda kutip
				▪️%d   // desimal (base 10)
				▪️%b   // biner
				▪️%o   // oktal
				▪️%x   // hex
				▪️%f   // float default
				▪️%.2f // 2 angka di belakang koma
				▪️%t   // boolean
				▪️%T   // menampilkan tipe data
				▪️%v   // default value
				▪️%+v  // detail struct
				▪️%#v  // Go syntax representation

*/

package main

import "fmt";

func main(){
	name := "Go Language"
	fmt.Println(name);
	fmt.Println(len(name)); // 📝 `len` mirip seperti `.length` di Javascript
	fmt.Println(name[0]); // 📝 Saat kita ambil index ke-0 seperti ini, maka hasil yang didapat masih berupa byte. Byte dari `D` yaitu 68
	fmt.Println(name[1]); // 📝 Saat kita ambil index ke-1 seperti ini, maka hasil yang didapat masih berupa byte. Byte dari `e` yaitu 101
	fmt.Println(string(name[0])); // G
	fmt.Println(string(name[1])); // o

	/* ___CONCATENATION___ */
	first := "Purwadhika";
	var second string = "School";
	result := first + " " + second;
	fmt.Println(result);

	/* ___FORMAT VERB___ */
	fullName := "M. Defryan";
	output := fmt.Sprintf("Hello, my name is %s", fullName); // Sprintf: Me-return hasil format string
	fmt.Println(output);

	var score int = 90;
	var isPassed bool = true; 
	fmt.Printf("Exam Score is %d. Passed: %t", score, isPassed) // Printf: Menampilkan langsung hasil format string
};
