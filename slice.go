/*
	📝
	Slice merupakan struktur data yang mirip array, namun ukurannya dinamis
	dan bisa di manipulasi.
*/

package main

import "fmt"

func main(){
	var pcGame []string;
	pcGame = append(pcGame, "Fifa 25");
	pcGame = append(pcGame, "eFootball 25");
	pcGame = append(pcGame, "F1");

	numbers := make([]int, 0); // Membuat `slice` tanpa membatasi jumlah data (bisa di append banyak data)
	alphabets := make([]string, 0, 10); // Membuat `slice` dengan length = 0 dan batas maksimum jumlah data = 10
	campus := make([]string, 10) // Membuat `slice` dengan maksimum jumlah data = 10

	purwadhikaCampus := []string{"BDG"};
	purwadhikaCampus = append(purwadhikaCampus, "BSD");
	purwadhikaCampus = append(purwadhikaCampus, "JKT");

	fmt.Println(pcGame);
	fmt.Println(purwadhikaCampus);
}