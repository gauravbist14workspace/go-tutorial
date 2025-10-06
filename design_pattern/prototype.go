package main

import "fmt"

type Prototype interface {
	Clone() Prototype
	GetDetails() string
}

type User struct {
	Name    string
	Age     int
	Hobbies []string
}

func (u *User) GetDetails() string {
	return fmt.Sprintf("User{name: %s, age: %d}", u.Name, u.Age)
}

func (u *User) Clone() Prototype {
	// Deep copy of Hobby slice to avoid shared reference
	hobbyCopy := make([]string, len(u.Hobbies))
	copy(hobbyCopy, u.Hobbies)

	return &User{
		Name:    u.Name,
		Age:     u.Age,
		Hobbies: hobbyCopy,
	}
}

func main() {
	original := &User{Name: "Gaurav", Age: 30}
	fmt.Println("Original:", original.GetDetails())

	// Clone the original
	clone := original.Clone()
	// Type assertion to access User fields
	clonedUser := clone.(*User)
	clonedUser.Name = "Rahul"

	fmt.Println("Cloned:  ", clonedUser.GetDetails())
	fmt.Println("Original after clone update:", original.GetDetails())
}
