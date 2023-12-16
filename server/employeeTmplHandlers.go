package main

import (
	"feklistova/models"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

// Ручки для отрисовки простейшего фронтенда с использованием шаблонов

// HomeHandlerTmpl - ручка домашнего шаблона
func HomeHandlerTmpl(w http.ResponseWriter, r *http.Request) {
	// Явно указываем локацию шаблона
	tmpl, _ := template.ParseFiles("templates/index.html")
	// Создаем структуру отображаемых данных
	type viewData struct {
		Title string
	}
	// Выполняем шаблон
	err := tmpl.Execute(w, viewData{Title: "Web Apps Task 5"})
	if err != nil {
		log.Println(err)
	}

}

// NewEmployeeHandlerTmpl - ручка шаблона с регистрацией нового сотрудника
func NewEmployeeHandlerTmpl(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/newEmployee.html")
	type viewData struct {
		Title string
	}
	err := tmpl.Execute(w, viewData{Title: "New Employee"})
	if err != nil {
		log.Println(err)
	}
}

// EmployeeHandlerTmpl - ручка шаблона с профилем сотрудника
func EmployeeHandlerTmpl(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/employee.html")
	// Получаем ID сотрудника из адресной строки
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	// Используя методы репозитория, получаем информацию о сотруднике
	employee := repo.GetEmployeeByID(id)
	type ViewData struct {
		Title    string
		Employee models.Employee
	}
	// Явно указываем отображаемую информацию
	data := ViewData{
		Title:    "Employee profile",
		Employee: employee,
	}
	if err := tmpl.Execute(w, data); err != nil {
		return
	}
}

// EmployeesHandlerTmpl - ручка для отображения шаблона со всеми сотрудниками
func EmployeesHandlerTmpl(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/employees.html")
	employees := repo.GetAllEmployees()
	type ViewData struct {
		Title     string
		Employees []models.Employee
	}
	data := ViewData{
		Title:     "Employees List",
		Employees: employees,
	}
	if err := tmpl.Execute(w, data); err != nil {
		return
	}
}

// EditEmployeeHandlerTmpl - шаблон изменения профиля сотрудника
func EditEmployeeHandlerTmpl(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/editProfile.html")
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	employee := repo.GetEmployeeByID(id)
	type ViewData struct {
		Title    string
		Employee models.Employee
	}
	data := ViewData{
		Title:    "Edit Profile",
		Employee: employee,
	}
	if err := tmpl.Execute(w, data); err != nil {
		return
	}
}
