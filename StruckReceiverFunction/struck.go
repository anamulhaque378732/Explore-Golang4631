package main

import "fmt"

// struck  (read only) code segment

type User struct { //compile time , code segment
	Name string //member variable or property
	Age  int    //member variable or property
}

func printUserDetails(usr User) {

	fmt.Println("Name : ", usr.Name)

	fmt.Println("Age : ", usr.Age)

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

	user1 = User{ // this is "instance" like to object and when you instance this is instantiate
		Name: "Anamul",
		Age:  26,
	}

	// fmt.Println("Name -", user1.Name)
	// fmt.Println("Age - ", user1.Age)

	// printUserDetails(user1)

	user2 := User{
		Name: "Sumona",
		Age:  24,
	}

	// fmt.Println("Name -", user2.Name)
	// fmt.Println("Age -", user2.Age)

	// printUserDetails(user2)
	// user1.printDetails()
	user2.printDetails()

	user1.call(45)
}
