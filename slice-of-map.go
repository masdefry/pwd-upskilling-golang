package main

import "fmt";

func main(){
	students := []map[string]string{
		{
			"name": "Defryan", 
			"finalScore": "90",
			"campus": "Purwadhika BSD",
		},
		{
			"name": "Tito", 
			"finalScore": "80",
			"campus": "Purwadhika JKT",
		},
	}

	for _, item := range students {
		fmt.Println(item["name"]);
		fmt.Println(item["campus"]);
	}
}