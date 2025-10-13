package handlers

import (
	"bufio"
	"fmt"
	"mysql/internal/connect"
	"os"
	"time"
)

func List() {
	connect.Connect()
	defer connect.CloseConnection()
	sql := "SELECT id, nombre, correo, telefono FROM clientes;"
	rows, err := connect.Db.Query(sql)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var nombre, correo, telefono string
		err = rows.Scan(&id, &nombre, &correo, &telefono)
		if err != nil {
			panic(err)
		}
		fmt.Printf("ID: %d, Nombre: %s, Correo: %s, Telefono: %s\n", id, nombre, correo, telefono)
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}

}

func ListByID(id int) {
	connect.Connect()
	defer connect.CloseConnection()
	sql := "SELECT id, nombre, correo, telefono FROM clientes WHERE id = ?;"
	row := connect.Db.QueryRow(sql, id)
	var nombre, correo, telefono string
	err := row.Scan(&id, &nombre, &correo, &telefono)
	if err != nil {
		panic(err)
	}
	fmt.Printf("ID: %d, Nombre: %s, Correo: %s, Telefono: %s\n", id, nombre, correo, telefono)
}

func Create(nombre, correo, telefono string) {
	connect.Connect()
	defer connect.CloseConnection()
	dateNow := time.Now()
	sql := "INSERT INTO clientes (nombre, correo, telefono, fecha) VALUES (?, ?, ?, ?);"
	result, err := connect.Db.Exec(sql, nombre, correo, telefono, dateNow)
	if err != nil {
		panic(err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Cliente creado con ID: %d\n", id)
}


func Update(id int, nombre, correo, telefono string) {
	connect.Connect()
	defer connect.CloseConnection()
	sql := "UPDATE clientes SET nombre = ?, correo = ?, telefono = ? WHERE id = ?;"
	result, err := connect.Db.Exec(sql, nombre, correo, telefono, id)
	if err != nil {
		panic(err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Cliente con ID %d actualizado, filas afectadas: %d\n", id, rowsAffected)
}

func Delete(id int) {
	connect.Connect()
	defer connect.CloseConnection()
	sql := "DELETE FROM clientes WHERE id = ?;"
	result, err := connect.Db.Exec(sql, id)
	if err != nil {
		panic(err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Cliente con ID %d eliminado, filas afectadas: %d\n", id, rowsAffected)
}

var ID int
var nombres, correos, telefonos string
func ExecuteTransaction() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Seleccione una opción:")
		fmt.Println("-------------------------------------")
		fmt.Println("1. Listar todos los clientes")
		fmt.Println("2. Listar cliente por ID")
		fmt.Println("3. Crear nuevo cliente")
		fmt.Println("4. Actualizar cliente existente")
		fmt.Println("5. Eliminar cliente")
		fmt.Println("6. Salir")
		fmt.Print("\nIngrese el número de la opción deseada: ")
		if scanner.Scan() {
		if scanner.Text() == "1" {
				List()
				break
			}
			if scanner.Text() == "2" {
				fmt.Print("Ingrese el ID del cliente: ")
				if scanner.Scan() {
					_, err := fmt.Sscan(scanner.Text(), &ID)
					if err != nil {
						fmt.Println("ID inválido. Por favor, ingrese un número entero.")
						return
					}
					ListByID(ID)
				}
			}
			if scanner.Text() == "3" {
				fmt.Print("Ingrese el nombre del cliente: ")
				if scanner.Scan() {
					nombres = scanner.Text()
				}
				fmt.Print("Ingrese el correo del cliente: ")
				if scanner.Scan() {
					correos = scanner.Text()
				}
				fmt.Print("Ingrese el teléfono del cliente: ")
				if scanner.Scan() {
					telefonos = scanner.Text()
				}
				Create(nombres, correos, telefonos)
			}
			if scanner.Text() == "4" {
				fmt.Print("Ingrese el ID del cliente a actualizar: ")
				if scanner.Scan() {
					_, err := fmt.Sscan(scanner.Text(), &ID)
					if err != nil {
						fmt.Println("ID inválido. Por favor, ingrese un número entero.")
						return
					}
					fmt.Print("Ingrese el nuevo nombre del cliente: ")
					if scanner.Scan() {
						nombres = scanner.Text()
					}
					fmt.Print("Ingrese el nuevo correo del cliente: ")
					if scanner.Scan() {
						correos = scanner.Text()
					}
					fmt.Print("Ingrese el nuevo teléfono del cliente: ")
					if scanner.Scan() {
						telefonos = scanner.Text()
					}
					Update(ID, nombres, correos, telefonos)
				}
			}
			if scanner.Text() == "5" {
				fmt.Print("Ingrese el ID del cliente a eliminar: ")
				if scanner.Scan() {
					_, err := fmt.Sscan(scanner.Text(), &ID)
					if err != nil {
						fmt.Println("ID inválido. Por favor, ingrese un número entero.")
						return
					}
					Delete(ID)
				}
			}
		}
	}
}