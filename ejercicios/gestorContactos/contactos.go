package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Contact struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

func saveContactsToFile(contacts []Contact) error {
	file, err := os.Create("contacts.json")
	defer file.Close()
	
	encoder := json.NewEncoder(file)
	err = encoder.Encode(contacts)
	if err != nil {
		return err
	}

	return nil
}

func loadContactsFromFile(contacts *[]Contact) error {
	file, err := os.Open("contacts.json")
	defer file.Close()
	if err != nil {
		return err
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&contacts)
	if err != nil {
		return err
	}
	return nil
}



func main() {

	var contacts []Contact
	err := loadContactsFromFile(&contacts)
	if err != nil {
		fmt.Println("Error al cargar los contactos",err)
	}
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(`
#####################################
######## Gestor de Contactos ########
######## 1. Agregar Contacto ########
######## 2. Mostrar Contactos #######
######## 3. Salir            ########
#####################################

Elige una opción: `)
		var option int
		_, err = fmt.Scanln(&option)
		if err != nil {
			fmt.Println("Error al leer la opción")
			return
		}
		switch option{
		case 1:
			var c Contact
			fmt.Print("Nombre: ")
			c.Name, _ = reader.ReadString('\n')
			c.Name = strings.TrimSpace(c.Name)
			fmt.Print("Email: ")
			c.Email, _ = reader.ReadString('\n')
			c.Email = strings.TrimSpace(c.Email)
			fmt.Print("Telefono: ")
			c.Phone, _ = reader.ReadString('\n')
			c.Phone = strings.TrimSpace(c.Phone)
			contacts = append(contacts, c)
			if err := saveContactsToFile(contacts); err != nil {
				fmt.Println("Error al guardar el contacto")
			}
		case 2:
			fmt.Println("======================================")
			for index, contact := range contacts{
				fmt.Printf("%d, Nombre: %s Email: %s Telefono: %s \n",
			                index, contact.Name,contact.Email, contact.Phone)
			}
			fmt.Println("======================================")
		case 3:
			return
		}

	}
}