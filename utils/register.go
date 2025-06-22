package utils

import "fmt"

type Users struct {
	Username string
	Password string
}

func Register(userr *Users) bool {
	var n int
	var pass string
	fmt.Print("Enter username: ")
	fmt.Scan(&userr.Username)

	fmt.Print("Enter password: ")
	fmt.Scan(&userr.Password)

	for n < 5 {
		fmt.Print("Renter password: ")
		fmt.Scan(&pass)
		if userr.Password == pass {
			fmt.Println()
			fmt.Println("Registration successful!")
			return true
		} else {
			n++
			fmt.Println("Passwords do not match. Please try again. Maximum attempts are 5, Current attempts are", n)
		}

	}
	fmt.Println("Registration failed. Maximum attempts reached.")
	return false
}
