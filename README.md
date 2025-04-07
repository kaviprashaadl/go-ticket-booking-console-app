# 🎟️ Go Ticket Booking Console App

A terminal-based movie ticket booking system developed in **Golang**. This project demonstrates real-world concepts like user authentication, password validation, ticket management, and more — all through the command line.

---

## 🚀 Features

- ✅ **User Login & Signup**
- 🔐 **Strong Password Validation** (uppercase, lowercase, digits, special characters)
- 🎬 **Multiple Movie Options**
- 🎟️ **Real-Time Ticket Availability**
- 🛑 **Prevents Overbooking**
- 🔁 **Interactive Navigation with Clean Flow**

---

## 🧠 What I Learned

- Core Go syntax and logic flow  
- Input handling using `fmt.Scan()`  
- String and character validation using `unicode`  
- Managing arrays for user data and inventory  

---

## 🛠️ Tech Stack

- **Language:** Go (Golang)
- **Packages:** `fmt`, `unicode`

---

## 📷 Sample Output

```bash
PS E:\GO\basics> go run main.go
Enter 1 if you want to login (or) 0 if you want to SignUp
0
Enter your username
kaviprashaad
Enter your Password
Kavi@123
You are successfully signed up!!
To login press 1 (or) To quit press 0
1
Enter your username
kaviprashaad
Enter your Password
Kavi@123
Successfully loged in!
To check Tickets Availability press 1 (or) To book the Tickets press 0 (or) To quit press 2
1

Movie:JaiBheem || Tickets Available:10

Movie:God Bad Ugly || Tickets Available:10

Movie:Master || Tickets Available:10

Movie:Kodi || Tickets Available:10
To check Tickets Availability press 1 (or) To book the Tickets press 0 (or) To quit press 2
0

 For Movie :JaiBheem press 0

 For Movie :God Bad Ugly press 1

 For Movie :Master press 2

 For Movie :Kodi press 3
0
Enter How many tickets You want
10
Thanks for booking 10 tickets for JaiBheem
Returning Back to Home Page .......
To check Tickets Availability press 1 (or) To book the Tickets press 0 (or) To quit press 2
0

 For Movie :JaiBheem press 0

 For Movie :God Bad Ugly press 1

 For Movie :Master press 2

 For Movie :Kodi press 3
0
Enter How many tickets You want
2
Sry not Enough Ticket
Returning to home page ........
To check Tickets Availability press 1 (or) To book the Tickets press 0 (or) To quit press 2
2
App quit .........
