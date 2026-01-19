package main

import (
	"fmt"
	"slices"
)

func main(){
	alphabets := []string{"x", "j", "a"};

	fmt.Println(slices.Contains(alphabets, "a"));
	
	slices.Sort(alphabets);
	fmt.Println(alphabets);

	slices.Replace(alphabets, 0, 1, "xyz");
	fmt.Println(alphabets);

	stadium := []string{"GBK", "GBLA", "GBT"};
	stadiumCopy := slices.Clone(stadium); 
	stadiumCopy[0] = "JIS";
	fmt.Println(stadium);
	fmt.Println(stadiumCopy);

	nums := []int{1, 1, 2, 2, 2, 3};
	compactResult := slices.Compact(nums);
	fmt.Println(compactResult);
}