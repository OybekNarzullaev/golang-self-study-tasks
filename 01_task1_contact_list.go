package main

import "fmt"

type Contact struct {
	ID        int
	FirstName string
	LastName  string
	Phone     string
	Email     string
	Position  string
}

type ContactListType []Contact

// create add
func (cl *ContactListType) add(
	ID int,
	FirstName string,
	LastName string,
	Phone string,
	Email string,
	Position string,
) Contact {
	c := Contact{
		ID:        ID,
		FirstName: FirstName,
		LastName:  LastName,
		Phone:     Phone,
		Email:     Email,
		Position:  Position,
	}
	*cl = append(*cl, c)
	return c
}

// get one
func (cl ContactListType) getOne(ID int) Contact {
	var c Contact
	for _, v := range cl {
		if v.ID == ID {
			c = v
			fmt.Printf("%+v\n", v)
		}
	}
	return c
}

// get all
func (cl ContactListType) getAll() ContactListType {
	for i, v := range cl {
		fmt.Printf("%v -> %+v\n", i, v)
	}

	return cl
}

// Update update
func (cl *ContactListType) Update(
	ID int,
	FirstName string,
	LastName string,
	Phone string,
	Email string,
	Position string,
) ContactListType {
	c := Contact{
		ID:        ID,
		FirstName: FirstName,
		LastName:  LastName,
		Phone:     Phone,
		Email:     Email,
		Position:  Position,
	}
	list := ContactListType(*cl)
	for i, v := range list {
		if v.ID == ID {
			list[i] = c
			*cl = list
			return *cl
		}
	}
	return *cl
}

// Delete delete
func (cl *ContactListType) Delete(ID int) ContactListType {
	list := ContactListType(*cl)
	for i, v := range list {
		if v.ID == ID {
			list[i] = list[len(list)-1]
			list = list[:len(list)-1]
			*cl = list
			return *cl
		}
	}
	return *cl
}

func main() {
	var contactList ContactListType
	fmt.Println("add contact to contactlist")
	contactList.add(1, "Oybek", "Narzullayev", "+998946462499", "oybek@gmail.com", "active")
	contactList.add(2, "Shohzod", "Narzullayev", "+998904964647", "Shohzod@gmail.com", "active")
	contactList.add(3, "Madinabonu", "Narzullayeva", "+998931269976", "madinabonu@gmail.com", "active")

	fmt.Println("Show a contact from contact list")
	contactList.getOne(1)

	fmt.Println("Show all contacts")
	contactList.getAll()

	fmt.Println("Delete some contact")
	fmt.Println(contactList.Delete(3))

	fmt.Println("Show all contacts")
	contactList.getAll()

	fmt.Println("Update some contact")
	contactList.Update(1, "Oybek update", "Narzullayev", "+998946462499", "oybek@gmail.com", "active")

	fmt.Println("Show all contacts")
	contactList.getAll()
}
