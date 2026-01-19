/*
	✅ Number
	Di Golang terdapat 2 tipe data number, yaitu:
	▪️Integer
	▪️Floating Point

	int8   			// -128 s/d 127
	int16  			// -32,768 s/d 32,767
	int32  			// -2,147,483,648 s/d 2,147,483,647
	int64  			// -9,223,372,036,854,775,808 s/d 9,223,372,036,854,775,807
	uint8  (byte) 	// 0 s/d 255
	uint16        	// 0 s/d 65,535
	uint32        	// 0 s/d 4,294,967,295
	uint64        	// 0 s/d 18,446,744,073,709,551,615
	float32 		// ~1.4e-45 s/d ~3.4e+38
	float64 		// ~5e-324 s/d ~1.8e+308
*/

package main

import "fmt";

func main(){
	var num1 int = 10 // Size/ukuran tergantung OS
	fmt.Println(num1)

	var num2 int8  = 127        
	var num3 int16 = 32000
	var num4 int32 = 200000
	var num5 int64 = 9000000000
	fmt.Println(num2, num3, num4, num5)

	var num6 uint   = 10
	var num7 uint8  = 255       
	var num8 uint16 = 65000
	var num9 uint32 = 400000
	var num10 uint64 = 18000000000
	fmt.Println(num6, num7, num8, num9, num10)

	var pi32 float32 = 3.14
	fmt.Println(pi32)
	var pi64 float64 = 3.141592653589793
	fmt.Println(pi64)
};
