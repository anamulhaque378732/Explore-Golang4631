package main

import "fmt"

// struck  (read only) code segment

type User struct { // read only. store code segment. compile time , code segment
	Name    string // member variable or property
	Age     int    // member variable or property
	Country string
	Id      int
}

func printUserDetails(usr User) {

	fmt.Println("Name : ", usr.Name)
	fmt.Println("Age : ", usr.Age)
	fmt.Println("Country : ", usr.Country)
	fmt.Println("Id : ", usr.Id)
}

// receiver function (you can use when you create custom type)

func (usr User) printDetails() {
	fmt.Println("Name : ", usr.Name)
	fmt.Println("Age : ", usr.Age)
}

func (usr User) call(a int) {
	fmt.Println("Name : ", usr.Name)
	fmt.Println(a)
}

func main() {

	var user1 User

	user1 = User{ // this is "instance" like to object and when you instance this is call instantiate
		Name:    "Anamul",
		Age:     26,
		Country: "Bangladesh",
		Id:      52,
	}

	// fmt.Println("Name - ", user1.Name, "Age - ", user1.Age, " Country - ", user1.Country)

	// printUserDetails(user1)

	user2 := User{
		Name:    "Sumona",
		Age:     24,
		Country: "Bangladesh",
		Id:      12,
	}

	// fmt.Println("Name -", user2.Name, "Age - ", user2.Age, " Country - ", user2.Country)

	// printUserDetails(user2)

	user1.printDetails()
	user2.printDetails()

	user1.call(45)
}
