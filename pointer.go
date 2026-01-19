/*
	📝
	Pointer di Golang adalah variabel yang menyimpan alamat memori dari variabel lain (bukan menyimpan value aslinya).
	Dengan pointer, kita bisa mengakses dan mengubah nilai asli dari suatu variabel.
	Di JavaScript, object dan array itu dikirim sebagai reference. Di Go, semua data dikirim sebagai value, kecuali kalau kita secara eksplisit pakai pointer.

	▪️Core Concept Pointer
		Misalnya kita mempunyai variabel biasa:

			x := 10

		Di memori, variable `x` tersebut disimpan di suatu alamat (misalnya 0x1234).
		Jika kita membuat pointer:

			p := &x

		Artinya:
			➡️ &x 	→ alamat memori dari x
			➡️ p 	→ pointer yang menyimpan alamat tersebut

	▪️Pointer Operator
		➡️ & (address-of)	→ Mengambil alamat memori suatu variabel
		➡️ * (dereference)	→ Mengambil atau mengubah nilai di alamat tersebut

	✅ Kapan sebaiknya menggunakan pointer?
		➡️ Ingin mengubah data asli
		➡️ Data berukuran besar (struct besar)
		➡️ Menghindari copy data berulang

	✅ Dengan Pointer vs Tanpa Pointer (Case Function)
		❌ Tanpa Pointer (tidak berubah)

			func changeValue(x int) {
				x = 100
			}

			func main() {
				a := 10
				changeValue(a)
				fmt.Println(a) // 10
			}

		✅ Dengan Pointer (berubah)

				func changeValue(x *int) {
					*x = 100
				}

				func main() {
					a := 10
					changeValue(&a)
					fmt.Println(a) // 100
				}
*/

package main

import "fmt"

type User struct {
	ID       int
	Username string
	Email    string
	IsActive bool
}

func ActivateUser(user *User) {
	user.IsActive = true
}

func main() {
	//___SHORT DECLARATION___
	firstNum := 1000
	secondNum := &firstNum // & (Address Operator) Menyimpan alamat memori dari `firstNum`
	fmt.Println(firstNum);
	fmt.Println(secondNum);
	fmt.Println(*secondNum); // * (Dereference) Digunakan untuk mengambil dan mengubah nilai di alamat tersebut

	*secondNum = 1;
	fmt.Println(secondNum);
	fmt.Println(*secondNum); 
	fmt.Println(firstNum);

	
	
	// ___VARIABLE DECLARATION___
	var num1 int = 3;
	var num2 *int = &num1;
	fmt.Println(num1);
	fmt.Println(num2);



	// ___STRUCT___
	user := User{
		ID: 1,
		Username: "defryan",
		Email: "defryan@gmail.com",
		IsActive: false,
	};
	ActivateUser(&user)
	fmt.Println(user.IsActive) 
}