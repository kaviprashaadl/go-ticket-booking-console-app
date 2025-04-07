package main

import (
	"fmt"
	"unicode"
)

var pUserName [100]string
var pPassword [100]string
var ticketCount [4]int = [4]int{10, 10, 10, 10}
var ticketName [4]string = [4]string{"JaiBheem", "God Bad Ugly", "Master", "Kodi"}
var pSize int = 0

func loginOrregister() int {
	var a int
	fmt.Println("Enter 1 if you want to login (or) 0 if you want to SignUp")
	for {
		fmt.Scan(&a)
		if a <= 1 {
			return a
		}
		fmt.Println("Retype the status")
	}

}
func passwordCheck(password string) int {
	var size int = len(password)
	var lower int = 0
	var upper int = 0
	var special int = 0
	var digit int = 0
	for _, ch := range password {
		if unicode.IsLower(ch) {
			lower = lower + 1

		}
		if unicode.IsUpper(ch) {
			upper = upper + 1

		}
		if unicode.IsPunct(ch) {
			special = special + 1

		}
		if unicode.IsDigit(ch) {
			digit = digit + 1

		}
	}
	if special >= 1 && upper >= 1 && lower >= 1 && size >= 8 && digit > 1 {
		return 1
	} else {
		return 0
	}
}
func checkRegisterStatus(userName string) int {
	for _, ch := range pUserName {
		if ch == userName {
			return 0

		}
	}
	return 1

}

func validateLoginStatus(userName string, password string) int {
	for i := 0; i < pSize; i++ {
		if pUserName[i] == userName && pPassword[i] == password {
			return 1
		}
	}
	
	return 0

}

func register() {
	fmt.Println("Enter your username")
	var userName string
	fmt.Scan(&userName)
	fmt.Println("Enter your Password")
	var password string
	fmt.Scan(&password)
	for {
		if passwordCheck(password) == 1 {
			break
		} else {
			fmt.Println("Retpe you Password!!")
			fmt.Scan(&password)
		}
	}
	if checkRegisterStatus(userName) == 1 {
		pUserName[pSize] = userName
		pPassword[pSize] = password
		pSize = pSize + 1
		fmt.Println("You are successfully signed up!!")
	} else {
		fmt.Println("You are already signed up!!")
	}
	var b int
	start :
		fmt.Println("To login press 1 (or) To quit press 0")
		fmt.Scan(&b)
		if b == 1 {
			login()
		}
		if b>1{
			fmt.Println("Invalid Input.Retype.......")
			goto start
		}

}
func ticketBook() {
	var movie int
	var tCount int
	for i, ch := range ticketName {
		fmt.Printf("\n For Movie :%v press %v\n", ch, i)
	}
	fmt.Scan(&movie)
start:
	fmt.Println("Enter How many tickets You want")
	fmt.Scan(&tCount)
	if tCount >= 0 {
		if ticketCount[movie]-tCount >= 0 {
			ticketCount[movie] = ticketCount[movie] - tCount
			fmt.Printf("Thanks for booking %v tickets for %v\n", tCount, ticketName[movie])
			fmt.Println("Returning Back to Home Page .......")
			bookingStatus()

		} else {
			fmt.Println("Sry not Enough Ticket")
			fmt.Println("Returning to home page ........")
			bookingStatus()

		}
	} else {
		fmt.Println("Invalid Input")
		goto start
	}

}
func bookingStatus() {
	var a int
start:
	fmt.Println("To check Tickets Availability press 1 (or) To book the Tickets press 0 (or) To quit press 2")
	fmt.Scan(&a)
	if a == 1 {
		ticketAvailable()
	} else if a == 0 {
		ticketBook()
	} else if a == 2 {
		fmt.Println("App quit .........")
	} else {
		fmt.Println("Enter a valid input")
		goto start
	}

}
func ticketAvailable() {
	for i := 0; i < 4; i++ {
		fmt.Printf("\nMovie:%v || Tickets Available:%v\n", ticketName[i], ticketCount[i])
	}
	bookingStatus()

}
func login() {
	fmt.Println("Enter your username")
	var userName string
	fmt.Scan(&userName)
	fmt.Println("Enter your Password")
	var password string
	fmt.Scan(&password)
	var status int = validateLoginStatus(userName, password)
	if status == 1 {
		fmt.Println("Successfully loged in!")
		bookingStatus()
	} else {
		fmt.Println("Incorrect Input")
	}

}
func main() {
	var a int = loginOrregister()
	if a == 1 {
		login()
	}
	if a == 0 {
		register()
	}

}
