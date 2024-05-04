package main

import (
	"fmt"
)

type User struct {
	name   *string
	age    int
	values map[string]string
	list   []string
}

func structMain() {
	// user1 := User{
	// 	name: "PH",
	// 	age:  33,
	// }
	user1 := &User{}
	name := "PH"
	user1.name = &name
	*user1.name = "PH22"
	user1.age = 33
	user1.values = nil

	fmt.Println(*user1)
}
