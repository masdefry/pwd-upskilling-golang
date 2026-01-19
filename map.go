/*
	📝
	Map di Golang (Go) adalah struktur data key–value (pasangan kunci dan nilai), mirip object di JavaScript atau dictionary di Python.

	📌 Map dipakai untuk pencarian data cepat berdasarkan key.
*/

package main

import "fmt"

func main(){
	data := map[string]string{
		"id": "1", 
		"email": "ryan@gmail.com",
	}
	fmt.Println(data);
	fmt.Println(data["id"])
	fmt.Println(data["email"])

	data["email"] = "defryan@gmail.com"
	fmt.Println(data);

	delete(data, "id");
	fmt.Println(data)
}