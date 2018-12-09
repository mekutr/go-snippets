package main

import "fmt"

/*
Go allows to use existing types and extend/change their behaviour by type embedding.
-This feature improves code re-usability and changing the behaviour of an exiting type.
*/

type user struct {
	firstName string
	lastName  string
	email     string
}

// notify implements a method that can be called by type user
func (u *user) notify() {
	fmt.Printf("first_name: %v\nlast_name: %v\nemail: %v", u.firstName, u.lastName, u.email)
}

// admin represents an admin user
type admin struct {
	user // embedded type
	// -- to embed a type all that needs to happen is for the type name to be declared
	privileges []string
}

func main() {
	a := admin{
		user: user{
			firstName: "Alec",
			lastName:  "Baldwin",
			email:     "alec@baldwin.com",
		},
		privileges: []string{"site", "database"},
	}

	// inner type's method can be reached directly
	//a.user.notify()

	// the inner type method is promoted.
	a.notify()
}
