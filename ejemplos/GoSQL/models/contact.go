package models

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Contact struct {
	Id    int
	Name  string
	Email string
	Phone string
}

func CreateContact() Contact {
	reader := bufio.NewReader(os.Stdin)
	var newContact Contact
			
	fmt.Print("Nombre: ")
	name, _ := reader.ReadString('\n')
	newContact.Name = strings.TrimSpace(name)
	fmt.Print("Email: ")
	email, _ := reader.ReadString('\n')
	newContact.Email = strings.TrimSpace(email)
	fmt.Print("Telefono: ")
	phone, _ := reader.ReadString('\n')
	newContact.Phone = strings.TrimSpace(phone)

	return newContact
	
}

func (c *Contact)EditContact() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Nombre: (%s) ",c.Name)
	name, _ := reader.ReadString('\n')
	c.Name = strings.TrimSpace(name)
	fmt.Printf("Email: (%s) ",c.Email)
	email, _ := reader.ReadString('\n')
	c.Email = strings.TrimSpace(email)
	fmt.Printf("Telefono: (%s) ",c.Phone)
	phone, _ := reader.ReadString('\n')
	c.Phone = strings.TrimSpace(phone)
	
}