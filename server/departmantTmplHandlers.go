package main

import (
	"feklistova/models"
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"strconv"
)

// Ручки для отрисовки простейшего фронтенда с использованием шаблонов

// DeptsHandlerTmpl - отображает все департаменты компании
func DeptsHandlerTmpl(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/allDepts.html")
	type ViewData struct {
		Title string
		Depts []models.Department
	}
	data := ViewData{
		Title: "Departments",
		Depts: repo.GetAllDepartments(),
	}
	if err := tmpl.Execute(w, data); err != nil {
		return
	}
}

// NewDeptHandlerTmpl - ручка создания нового департамента в шаблоне
func NewDeptHandlerTmpl(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/newDept.html")
	type ViewData struct {
		Title string
	}
	data := ViewData{
		Title: "New department",
	}
	if err := tmpl.Execute(w, data); err != nil {
		return
	}
}

// DeptHandlerTmpl - ручка отображения департамента и его сотрудников по ID в шаблоне
func DeptHandlerTmpl(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/dept.html")
	if err != nil {
		fmt.Print(err)
	}
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	type ViewData struct {
		Dept      models.Department
		Employees []models.Employee
	}
	data := ViewData{
		Dept:      repo.GetDepartmentByID(id),
		Employees: repo.GetAllEmployeesInDept(id),
	}
	if err := tmpl.Execute(w, data); err != nil {
		return
	}
}

// EditDeptHandlerTmpl - ручка шаблона изменения департамента
func EditDeptHandlerTmpl(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/editDept.html")
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	dept := repo.GetDepartmentByID(id)
	type ViewData struct {
		Title      string
		Department models.Department
	}
	data := ViewData{
		Title:      "Edit Department",
		Department: dept,
	}
	if err := tmpl.Execute(w, data); err != nil {
		return
	}
}
