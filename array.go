/*
	📝
	Di Golang, array adalah tipe data koleksi dengan ukuran tetap dan berisi elemen bertipe sama,
	diakses menggunakan indeks (mulai dari 0).

	Karena sifatnya yg tetap, method array di Golang sangat terbatas:
	1. Mendapatkan jumlah item: len()
	2. Perulangan dengan menggunakan: for-range
	3. Menyalin item: copy()
*/

package main

import "fmt";

func main(){
	var numbers [3]int;
	numbers[0] = 1;
	numbers[1] = 100;
	numbers[2] = 1000;
	fmt.Println(numbers);

	programmingLanguage := [3]string{"Javascript", "Typescript", "Golang"};
	fmt.Println(programmingLanguage[0])
	fmt.Println(programmingLanguage[1])
	fmt.Println(programmingLanguage[2])
	
	programmingLanguage[1] = "PHP";
	fmt.Println(programmingLanguage);

	cars := [...]int{1, 100, 100, 10000, 100000, 1000000}; // With `...` you don’t need to manually count the number of elements
	cars[10] = 5; // ❌
	fmt.Println(cars);

	/* ___ARRAY MULTIDIMENSI___ */
	matrix := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println(matrix);




	array := [5]int{10, 20, 30, 40, 50}
	fmt.Println(array[1:5]) // Slicing dari index ke-1 sampai index ke-5
	fmt.Println(array[4:]); // Slicing setelah index ke-4 sampai index terakhir
	fmt.Println(array[:2]) // Slicing dari index awal sampai sebelum index ke-2
}