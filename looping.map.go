package main

import "fmt";

func main(){
	cars := map[string]string{
		"brand": "Honda",
		"type": "Jazz Idsi",
		"color": "Gray Stone",
	}

	for key, value := range cars{
		fmt.Println(key);
		fmt.Println(value);
	}
}