package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"todolist/application"
	"todolist/domain"
	"todolist/infrastructure"
)

type appContext struct {
	listUsersUC        application.ListUsersUseCase
	selectOrCreateUser application.SelectOrCreateUserUseCase
	listUserTasksUC    application.ListUserTasksUseCase
	createTaskUC       application.CreateTaskUseCase
	updateTaskUC       application.UpdateTaskUseCase
	deleteTaskUC       application.DeleteTaskUseCase
	selectedUser       *domain.User
	reader             *bufio.Reader
}

func newAppContext() *appContext {
	taskRepo := &infrastructure.In_memory_task_repository{Path: "tasks.json"}
	userRepo := &infrastructure.In_memory_user_repository{Path: "users.json"}

	return &appContext{
		listUsersUC:        application.ListUsersUseCase{UserRepo: userRepo},
		selectOrCreateUser: application.SelectOrCreateUserUseCase{UserRepo: userRepo},
		listUserTasksUC:    application.ListUserTasksUseCase{TaskRepo: taskRepo},
		createTaskUC:       application.CreateTaskUseCase{TaskRepo: taskRepo},
		updateTaskUC:       application.UpdateTaskUseCase{TaskRepo: taskRepo},
		deleteTaskUC:       application.DeleteTaskUseCase{TaskRepo: taskRepo},
		reader:             bufio.NewReader(os.Stdin),
	}
}

func (a *appContext) readLine(prompt string) string {
	fmt.Print(prompt)
	value, _ := a.reader.ReadString('\n')
	return strings.TrimSpace(value)
}

func (a *appContext) readInt(prompt string) int {
	for {
		text := a.readLine(prompt)
		value, err := strconv.Atoi(text)
		if err == nil {
			return value
		}
		fmt.Println("Ingresa un numero valido")
	}
}

func (a *appContext) chooseUser() error {
	for {
		fmt.Println("\n===== Usuarios =====")
		users, err := a.listUsersUC.Execute()
		if err != nil {
			return err
		}

		if len(users) > 0 {
			for _, u := range users {
				fmt.Printf("%d) %s\n", u.ID, u.Name)
			}
		} else {
			fmt.Println("No hay usuarios registrados")
		}

		fmt.Println("n) Crear o seleccionar por nombre")
		choice := strings.ToLower(a.readLine("> "))

		if choice == "n" {
			name := a.readLine("Nombre de usuario: ")
			user, err := a.selectOrCreateUser.Execute(name)
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}
			a.selectedUser = user
			return nil
		}

		userID, err := strconv.Atoi(choice)
		if err != nil {
			fmt.Println("Opcion invalida")
			continue
		}

		for _, u := range users {
			if u.ID == userID {
				a.selectedUser = u
				return nil
			}
		}

		fmt.Println("Usuario no encontrado")
	}
}

func statusLabel(status domain.TaskStatus) string {
	switch status {
	case domain.Pending:
		return "pendiente"
	case domain.InProgress:
		return "en progreso"
	case domain.Completed:
		return "completada"
	default:
		return string(status)
	}
}

func (a *appContext) listTasks() ([]*domain.Task, error) {
	tasks, err := a.listUserTasksUC.Execute(a.selectedUser.ID)
	if err != nil {
		return nil, err
	}

	fmt.Printf("\n===== Tareas de %s =====\n", a.selectedUser.Name)
	if len(tasks) == 0 {
		fmt.Println("No hay tareas")
		return tasks, nil
	}

	for _, t := range tasks {
		fmt.Printf("ID: %d | %s | Estado: %s\n", t.ID, t.Name, statusLabel(t.Status))
		fmt.Printf("   Detalle: %s\n", t.Detail)
	}

	return tasks, nil
}

func (a *appContext) parseStatusInput(input string) (domain.TaskStatus, bool) {
	value := strings.ToLower(strings.TrimSpace(input))
	switch value {
	case "1", "pending", "pendiente":
		return domain.Pending, true
	case "2", "in_progress", "en progreso", "progreso":
		return domain.InProgress, true
	case "3", "completed", "completada":
		return domain.Completed, true
	default:
		return "", false
	}
}

func (a *appContext) createTask() {
	name := a.readLine("Nombre de la tarea: ")
	detail := a.readLine("Detalle: ")

	task, err := a.createTaskUC.Execute(a.selectedUser.ID, name, detail)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Tarea creada con ID %d\n", task.ID)
}

func (a *appContext) editTask() {
	tasks, err := a.listTasks()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	if len(tasks) == 0 {
		return
	}

	taskID := a.readInt("ID de la tarea a editar: ")

	var current *domain.Task
	for _, t := range tasks {
		if t.ID == taskID {
			current = t
			break
		}
	}

	if current == nil {
		fmt.Println("Tarea no encontrada")
		return
	}

	name := a.readLine(fmt.Sprintf("Nombre [%s]: ", current.Name))
	if strings.TrimSpace(name) == "" {
		name = current.Name
	}

	detail := a.readLine(fmt.Sprintf("Detalle [%s]: ", current.Detail))
	if strings.TrimSpace(detail) == "" {
		detail = current.Detail
	}

	fmt.Printf("Estado actual: %s\n", statusLabel(current.Status))
	fmt.Println("Nuevo estado: 1) pendiente 2) en progreso 3) completada")
	statusInput := a.readLine("Estado [Enter para mantener]: ")
	newStatus := current.Status
	if strings.TrimSpace(statusInput) != "" {
		parsed, ok := a.parseStatusInput(statusInput)
		if !ok {
			fmt.Println("Estado invalido")
			return
		}
		newStatus = parsed
	}

	updated, err := a.updateTaskUC.Execute(a.selectedUser.ID, taskID, name, detail, newStatus)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Tarea %d actualizada (%s)\n", updated.ID, statusLabel(updated.Status))
}

func (a *appContext) deleteTask() {
	tasks, err := a.listTasks()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	if len(tasks) == 0 {
		return
	}

	taskID := a.readInt("ID de la tarea a eliminar: ")

	confirm := strings.ToLower(a.readLine("Confirmar eliminacion (s/n): "))
	if confirm != "s" && confirm != "si" {
		fmt.Println("Operacion cancelada")
		return
	}

	if err = a.deleteTaskUC.Execute(a.selectedUser.ID, taskID); err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Tarea eliminada")
}

func (a *appContext) runTaskMenu() bool {
	fmt.Printf("\n===== Usuario actual: %s =====\n", a.selectedUser.Name)
	fmt.Println("1) Ver tareas")
	fmt.Println("2) Agregar tarea")
	fmt.Println("3) Editar tarea")
	fmt.Println("4) Eliminar tarea")
	fmt.Println("5) Cambiar usuario")
	fmt.Println("6) Salir")

	option := a.readInt("> ")

	switch option {
	case 1:
		if _, err := a.listTasks(); err != nil {
			fmt.Println("Error:", err)
		}
	case 2:
		a.createTask()
	case 3:
		a.editTask()
	case 4:
		a.deleteTask()
	case 5:
		if err := a.chooseUser(); err != nil {
			fmt.Println("Error:", err)
		}
	case 6:
		return false
	default:
		fmt.Println("Opcion invalida")
	}

	return true
}

func main() {
	app := newAppContext()

	fmt.Println("Gestor de tareas CLI")
	if err := app.chooseUser(); err != nil {
		fmt.Println("Error al inicializar usuario:", err)
		os.Exit(1)
	}

	for app.runTaskMenu() {
	}

	fmt.Println("Hasta luego")
}
