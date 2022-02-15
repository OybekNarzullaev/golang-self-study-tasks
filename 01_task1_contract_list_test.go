package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var contactList ContactListType
var contacts ContactListType = ContactListType{
	Contact{1, "Oybek", "Narzullayev", "+998946462499", "oybek@gmail.com", "active"},
	Contact{2, "Shohzod", "Narzullayev", "+998904964647", "Shohzod@gmail.com", "active"},
	Contact{3, "Madinabonu", "Narzullayeva", "+998931269976", "Madina@gmail.com", "active"},
}

func TestAdd(t *testing.T) {

	for _, v := range contacts {
		assert.Equal(t, v, contactList.add(
			v.ID,
			v.FirstName,
			v.LastName,
			v.Phone,
			v.Email,
			v.Position,
		))
	}
}

func TestContactListType_Update(t *testing.T) {
	type item struct {
		c        Contact
		contacts ContactListType
	}

	items := []item{
		item{
			Contact{1, "Oybek update", "Narzullayev", "+998946462499", "oybek@gmail.com", "active"},
			ContactListType{
				Contact{1, "Oybek update", "Narzullayev", "+998946462499", "oybek@gmail.com", "active"},
				Contact{2, "Shohzod", "Narzullayev", "+998904964647", "Shohzod@gmail.com", "active"},
				Contact{3, "Madinabonu", "Narzullayeva", "+998931269976", "Madina@gmail.com", "active"},
			},
		},

		item{
			Contact{2, "Shohzod update", "Narzullayev", "+998904964647", "Shohzod@gmail.com", "active"},
			ContactListType{
				Contact{1, "Oybek", "Narzullayev", "+998946462499", "oybek@gmail.com", "active"},
				Contact{2, "Shohzod update", "Narzullayev", "+998904964647", "Shohzod@gmail.com", "active"},
				Contact{3, "Madinabonu", "Narzullayeva", "+998931269976", "Madina@gmail.com", "active"},
			},
		},

		item{
			Contact{3, "Madinabonu update", "Narzullayeva", "+998946462499", "Madina@gmail.com", "active"},
			ContactListType{
				Contact{1, "Oybek", "Narzullayev", "+998946462499", "oybek@gmail.com", "active"},
				Contact{2, "Shohzod", "Narzullayev", "+998904964647", "Shohzod@gmail.com", "active"},
				Contact{3, "Madinabonu update", "Narzullayeva", "+998946462499", "Madina@gmail.com", "active"},
			},
		},
	}

	for _, v := range items {
		v.contacts = contacts
		assert.Equal(t, v.contacts, contacts.Update(
			v.c.ID,
			v.c.FirstName,
			v.c.LastName,
			v.c.Phone,
			v.c.Email,
			v.c.Position,
		))
	}
}

func TestContactListType_Delete(t *testing.T) {
	type item struct {
		id       int
		contacts ContactListType
	}

	items := []item{
		item{
			1,
			ContactListType{
				Contact{3, "Madinabonu", "Narzullayeva", "+998931269976", "Madina@gmail.com", "active"},
				Contact{2, "Shohzod", "Narzullayev", "+998904964647", "Shohzod@gmail.com", "active"},
			},
		},

		item{
			2,
			ContactListType{
				Contact{1, "Oybek", "Narzullayev", "+998946462499", "oybek@gmail.com", "active"},
				Contact{3, "Madinabonu", "Narzullayeva", "+998931269976", "Madina@gmail.com", "active"},
			},
		},

		item{
			3,
			ContactListType{
				Contact{1, "Oybek", "Narzullayev", "+998946462499", "oybek@gmail.com", "active"},
				Contact{2, "Shohzod", "Narzullayev", "+998904964647", "Shohzod@gmail.com", "active"},
			},
		},
	}

	contacts = ContactListType{
		Contact{1, "Oybek", "Narzullayev", "+998946462499", "oybek@gmail.com", "active"},
		Contact{2, "Shohzod", "Narzullayev", "+998904964647", "Shohzod@gmail.com", "active"},
		Contact{3, "Madinabonu", "Narzullayeva", "+998931269976", "Madina@gmail.com", "active"},
	}
	for _, v := range items {
		assert.Equal(t, v.contacts, contacts.Delete(v.id))
		contacts = ContactListType{
			Contact{1, "Oybek", "Narzullayev", "+998946462499", "oybek@gmail.com", "active"},
			Contact{2, "Shohzod", "Narzullayev", "+998904964647", "Shohzod@gmail.com", "active"},
			Contact{3, "Madinabonu", "Narzullayeva", "+998931269976", "Madina@gmail.com", "active"},
		}
	}
}

func TestGetOne(t *testing.T) {
	assert.Equal(t, contacts, contacts.getAll())
}

func TestGetAll(t *testing.T) {
	for _, v := range contacts {
		assert.Equal(t, v, contacts.getOne(v.ID))
	}
}
