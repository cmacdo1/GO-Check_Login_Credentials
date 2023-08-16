// checks the credentials of a user logging in.
// allows 3 retries for user if first attempt is incorrect

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var loginID string = "admin"
var password string = "Pa$$w0rd"

func main() {
	checkCredentials(loginID, password)
}

func checkCredentials(loginID string, password string) {
	reader := bufio.NewReader(os.Stdin)

	for guess := 0; guess <= 3; guess++ {
		fmt.Print("\nPlease Enter Your Login ID: ")

		//reads users keyboard input and stores in "a" variable or stores an error in the "err" variable
		a, err := reader.ReadString('\n')

		//checks that error is not nil
		if err != nil {
			log.Fatal(err) //report error and stop the program
		}
		//gets rid of the new line from ReadString
		a = strings.TrimSpace(a)

		fmt.Print("Please Enter Your Password: ")

		//reads users keyboard input and stores in "b" variable or stores an error in the "err" variable
		b, err := reader.ReadString('\n')

		//checks that error is not nil
		if err != nil {
			log.Fatal(err) //report the error and stop the program
		}
		//gets rid of new line from ReadString
		b = strings.TrimSpace(b)

		//if statements below check user login id & password against set variables for login id & password
		//if login and password are correct, program will break out of loop
		if a == loginID && b == password {
			fmt.Println("\nLogin Successful!")
			break
		}

		if a == loginID && b != password {
			fmt.Println("\nPassword Incorrect!")
		}

		/*This was not part of the requirements, but I felt it should be added
		if a user inputs the wrong login id and password, no error is given
		all that is listed is the number of attempts left*/
		if a != loginID && b != password {
			fmt.Println("\nLogin & Password Incorrect!")
		}

		if a != loginID && b == password {
			fmt.Println("\nLogin Incorrect!")
		}

		//lets user know how many attempts are left or when account is locked
		if guess == 0 {
			fmt.Println("Try again - 3 attempts left!")
		}

		if guess == 1 {
			fmt.Println("Try again - 2 attempts left!")
		}

		if guess == 2 {
			fmt.Println("Try again - 1 attempt left!")
		}

		if guess == 3 {
			fmt.Println("Account Locked - Contact 800-123-4567")
		}
	}
}
