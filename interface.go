/*
	📝
	Dalam bahasa pemrograman Go (Golang), interface adalah sebuah tipe data yang berisi kumpulan definisi method (tanda tangan metode)
	tanpa isi/implementasi.
*/

package main

import "fmt"

type PaymentProcessor interface {
    Pay(amount float64) string
}

// ___IMPLEMENTASI-01___
type Midtrans struct{}
func (m Midtrans) Pay(amount float64) string {
    return fmt.Sprintf("Membayar %.2f menggunakan Midtrans", amount)
}

// ___IMPLEMENTASI-02___
type Xendit struct{}
func (x Xendit) Pay(amount float64) string {
    return fmt.Sprintf("Membayar %.2f menggunakan Xendit", amount)
}


func main(){
	var payment PaymentProcessor;

	// ___If Payment Using Midtrans___
	payment = Midtrans{}
	fmt.Println(payment.Pay(150000))

	// ___If Payment Using Xendit___
	payment = Xendit{}
	fmt.Println(payment.Pay(200000))
}