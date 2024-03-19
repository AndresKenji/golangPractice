package main

// https://github.com/joho/godotenv
// https://github.com/go-sql-driver/mysql
import (
	"fmt"
	"log"
	"kenji.gomysql/gomysql/database"
	"kenji.gomysql/gomysql/handlers"
	"kenji.gomysql/gomysql/models"
)

func main() {
	fmt.Println("Iniciando conexión a la base de datos")
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	continuar := true

	for continuar {
		fmt.Println("----------------------------Agenda de contactos----------------------------")
		fmt.Println("1) Ver contactos")
		fmt.Println("2) Agregar un contacto")
		fmt.Println("3) Editar un contacto")
		fmt.Println("4) Eliminar un contacto")
		fmt.Println("5) Salir")
		fmt.Print("Elije una opción: ")
		var opcion int
		fmt.Scanln(&opcion)

		switch opcion {
		case 1:
			handlers.ListContacts(db)
		case 2:
			newContact := models.CreateContact()
			handlers.CreateContact(db, newContact)
		case 3:
			handlers.ListContacts(db)
			var contactId int
			fmt.Print("id: ")
			fmt.Scanln(&contactId)
			contact := handlers.GetContactByID(db, contactId)
			contact.EditContact()
			handlers.UpdateContact(db,contact)
			
		case 4:
			handlers.ListContacts(db)
			var contactId int
			fmt.Print("id: ")
			fmt.Scanln(&contactId)
			handlers.DeleteContact(db, contactId)

		case 5:
			continuar = false
		}


	}
	

}