package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("golang", "go")) // true
	fmt.Println(strings.Split("a,b,c", ",")) // [a b c]
	fmt.Println(strings.Join([]string{"a", "b"}, "-")) // a-b
	fmt.Println(strings.ToUpper("go")) // GO
	fmt.Println(strings.ToLower("GO")) // go
	fmt.Println(strings.TrimSpace(" go ")) // go
	fmt.Println(strings.Replace("go go", "go", "js", -1)) // js js
	fmt.Println(strings.HasPrefix("golang", "go")) // true
	fmt.Println(strings.HasSuffix("main.go", ".go")) // true

	// Atoi: string -> int
	num, err := strconv.Atoi("10")
	fmt.Println(num) // 10
	fmt.Println(err) // <nil>

	// Itoa: int -> string
	str := strconv.Itoa(10)
	fmt.Println(str) // "10"

	// ParseFloat: string -> float64
	f, err := strconv.ParseFloat("3.14", 64)
	fmt.Println(f) // 3.14
	fmt.Println(err) // <nil>
}
