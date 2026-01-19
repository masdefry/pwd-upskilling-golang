package main

import "fmt"

type User struct {
	Id       int
	Username, Email string 
	IsActive bool
}

func (user User) display() string {
	return fmt.Sprintf("Name: %s___ Email: %s___ isActive: %t", user.Username, user.Email, user.IsActive)
}

func main() {
	user := User{
		1, "immanueljanis", "imm@gmail.com", false,
	};
	resultDisplay := user.display();
	fmt.Println(resultDisplay)
}