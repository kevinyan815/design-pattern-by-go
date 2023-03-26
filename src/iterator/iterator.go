package main

import "fmt"

type collection interface {
	createIterator() iterator
}

type iterator interface {
	hasNext() bool
	getNext() *User
}

type User struct {
	name string
	age int
}

type userCollection struct {
	users []*User
}

func (u *userCollection) createIterator() iterator {
	return &userIterator{
		users: u.users,
	}
}

type userIterator struct {
	index int
	users []*User
}

func (ui *userIterator) hasNext() bool {
	return ui.index < len(ui.users)
}

func (ui *userIterator) getNext() *User {
	if ui.hasNext() {
		user := ui.users[ui.index]
		ui.index++
		return user
	}

	return nil
}


func main() {
	userK := &User{
		name: "Kevin",
		age: 30,
	}
	userD := &User{
		name: "Diamond",
		age:  25,
	}

	userCollection := &userCollection{
		users: []*User{userK, userD},
	}

	iterator := userCollection.createIterator()
	for iterator.hasNext() {
		user := iterator.getNext()
		fmt.Printf("User is %v\n", user)
	}
}
