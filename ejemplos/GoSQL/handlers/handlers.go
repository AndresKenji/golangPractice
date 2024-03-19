package handlers

import (
	"database/sql"
	"fmt"
	"log"

	"kenji.gomysql/gomysql/models"
)

// Listar contactos desde la base de datos
func ListContacts(db *sql.DB) {
	// Realizar consulta mediante querys
	query := "SELECT * FROM contact"
	
	// Ejecutar consulta
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Iterar sobre las rows 
	fmt.Println("Lista de contactos")
	fmt.Println("--------------------------------------------------")
	
	for rows.Next(){
		// Instancia el modelo contact
		contact := models.Contact{}
		var valueEmail sql.NullString
		err := rows.Scan(&contact.Id,&contact.Name,&valueEmail,&contact.Phone)
		if err != nil {
			log.Fatal(err)
		}
		if valueEmail.Valid{
			contact.Email = valueEmail.String
		}else{
			contact.Email = "Sin correo"
		}

		fmt.Printf("ID: %d, Nombre: %s, Email: %s, Telefono: %s\n",
			contact.Id, contact.Name, contact.Email, contact.Phone)
	}

	fmt.Println("--------------------------------------------------")
}



func GetContactByID(db *sql.DB, id int) models.Contact{
	// Realizar consulta mediante querys
	query := "SELECT * FROM contact WHERE id = ?"
	
	// Ejecutar consulta
	row := db.QueryRow(query, id)
	
	contact := models.Contact{}
	var valueEmail sql.NullString
	err := row.Scan(&contact.Id,&contact.Name,&valueEmail,&contact.Phone)
	if err != nil {
		if err == sql.ErrNoRows{
			log.Fatalf("No se encontro ningun contacto con el ID %d", id)
		}
		log.Fatal(err)
	}
	if valueEmail.Valid{
		contact.Email = valueEmail.String
	}else{
		contact.Email = "Sin correo"
	}
	fmt.Println("Contacto")
	fmt.Println("--------------------------------------------------")
	fmt.Printf("ID: %d, Nombre: %s, Email: %s, Telefono: %s\n",
		contact.Id, contact.Name, contact.Email, contact.Phone)
	fmt.Println("--------------------------------------------------")

	return contact
}


func CreateContact(db *sql.DB, newContact models.Contact){
	query := "INSERT INTO contact (name,email,phone) VALUES (?, ?, ?)"

	_, err := db.Exec(query, newContact.Name, newContact.Email, newContact.Phone)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Se ha registrado un contacto nuevo!")
}

func UpdateContact(db *sql.DB, newContact models.Contact){
	query := "UPDATE contact SET name = ?, email = ?, phone = ? WHERE id = ?"

	_, err := db.Exec(query, newContact.Name, newContact.Email, newContact.Phone, newContact.Id)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Se ha registrado un contacto nuevo!")
}


func DeleteContact(db *sql.DB, id int){
	query := "DELETE FROM contact WHERE Id = ?"

	_, err := db.Exec(query, id)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Se ha eliminado el contacto!")
}