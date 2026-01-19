package main

import "fmt";

func main(){
	/*___ARRAY___*/
	framework := [3]string{"Next.js", "React.js", "Nest.js"};
	
	for i := 0; i < len(framework); i++{
		fmt.Println(framework[i]);
	}

	for index, value := range framework{
		fmt.Println(value)
		fmt.Println(index)
	}

	/*___SLICE___*/
	var carBrands []string;
	carBrands = []string{"Honda", "Toyota", "Suzuki"};
	
	for i := 0; i < len(carBrands); i++{
		fmt.Println(carBrands[i]);
	}

	for index, value := range carBrands{
		fmt.Println(value)
		fmt.Println(index)
	}
}