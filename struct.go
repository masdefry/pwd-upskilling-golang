package main

import "fmt"

type User struct {
	Id       int
	Username string
	Email    string
	IsActive bool
}

type Group struct {
	Name        string
	Admin       User
	Members     []User
	IsAvailable bool
}

func main() {
	// Cara-01: Assign satu per satu
	user := User{}
	user.Id = 1
	user.Username = "defryan"
	user.Email = "defryan@gmail.com"
	user.IsActive = true
	fmt.Println(user)

	// Cara-02: Struct Literal
	userNew := User{
		Id:       2,
		Username: "aboy01",
		Email:    "aboy01@gmail.com",
		IsActive: false,
	}
	fmt.Println(userNew)

	user.Display();

	/*
		___EMBED STRUCT___
	*/
	users := []User{user, userNew}

	group := Group{
		Name: "JCWDBSDAM37",
		Admin: User{
			Id:       99,
			Username: "superadmin",
			Email:    "admin@gmail.com",
			IsActive: true,
		},
		IsAvailable: true,
		Members:     users,
	}

	group.DisplayGroup();
}



// ___STRUCT METHOD___
func (u User) Display() string {
	return fmt.Sprintf(
		"Username: %s, Email: %s, IsActive: %t",
		u.Username,
		u.Email,
		u.IsActive,
	)
}



// ___STRUCT METHOD___
func (group Group) DisplayGroup() {
	fmt.Printf("Group Name: %s\n", group.Name)
	fmt.Printf("Total Members: %d\n", len(group.Members))
	fmt.Println("Member List:")

	for _, user := range group.Members {
		fmt.Println("-", user.Username)
	}
}
