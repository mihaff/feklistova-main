package main

import (
	"feklistova/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// REST - ручки доступны по адресу http://<host>/api/<route>
// Ответом REST ручек являются JSON файлы
// Также для каждой ручки прописана Swagger документация

// NewEmployeeHandler - ручка регистрации нового сотрудника
//
//	@Summary		REST ручка для регистрации нового сотрудника
//	@ID				new-employee
//	@Accept mpfd
//	@Param		surname	formData	string	true	"surname"
//	@Param		name	formData	string	true	"name"
//	@Param		deptID	formData	string	true	"deptID"
//	@Produce json
//	@Tags employees
//	@Router			/empl/newemployee [post]
func NewEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	// Для обеспечения работы в рамках localhost нужно явно указать допуск CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	// Парсим данные из формы регистраци
	err := r.ParseMultipartForm(10)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error parsing form data", http.StatusBadRequest)
		return
	}

	// Получаем данные из полей
	surname := r.Form.Get("surname")
	name := r.Form.Get("name")
	dept, _ := strconv.Atoi(r.Form.Get("deptID"))
	fmt.Println(dept)
	var employee models.Employee
	// Явно передаем их в структуру пользователя
	employee.Surname = surname
	employee.Name = name
	employee.Department.ID = dept
	// Вызываем метод регистрации пользователя
	id, err := repo.NewEmployee(employee)
	if err != nil {
		return
	}
	// Записываем ответ об успешном создании пользователя
	resp := []byte(fmt.Sprintf("Employee %d in %d dept hired successfully", id, dept))
	_, err = w.Write(resp)
	if err != nil {
		return
	}
}

// EditEmployeeHandler - ручка изменения профиля сотрудника
//
//	@Summary		REST - ручка для изменения профиля сотрудника
//	@ID				edit-empl
//	@Accept mpfd
//	@Param		id			path		int		true	"userId"
//	@Param		surname	formData	string	true	"surname"
//	@Param		name	formData	string	true	"name"
//	@Param		deptID	formData	string	true	"deptID"
//	@Produce json
//	@Tags employees
//	@Router			/empl/{id}/edit [put]
func EditEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")                                // Change '*' to the specific origin you want to allow or use multiple origins
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS") // Include the HTTP methods you want to allow
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	err := r.ParseMultipartForm(10)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error parsing form data", http.StatusBadRequest)
		return
	}
	surname := r.Form.Get("surname")
	name := r.Form.Get("name")
	dept, _ := strconv.Atoi(r.Form.Get("deptID"))
	var employee models.Employee
	// Явно передаем их в структуру пользователя
	employee.ID = id
	employee.Surname = surname
	employee.Name = name
	employee.Department.ID = dept
	repo.UpdateEmployee(employee)
	if err != nil {
		return
	}
}

// DeleteEmployeeHandler - ручка удаления профиля сотрудника
//
//	@Summary		REST ручка для удаления профиля сотрудника
//	@ID				delete-empl
//	@Accept json
//	@Param		id	path		int		true	"User ID"
//
// @Produce json
// @Tags employees
// @Router			/empl/{id}/delete [delete]
func DeleteEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")                                // Change '*' to the specific origin you want to allow or use multiple origins
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS") // Include the HTTP methods you want to allow
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	// Получаем ID пользователя из адресной строки
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	repo.DeleteEmployee(id)
	_, err := w.Write([]byte("Employee fired successfully"))
	if err != nil {
		return
	}

}

// EmployeesHandler - ручка получения всех сотрудников
//
//	@Summary		REST ручка вывода всех пользователей
//	@ID				all-empl
//	@Accept json
//	@Produce json
//	@Tags employees
//	@Router			/empl [get]
func EmployeesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")                                // Change '*' to the specific origin you want to allow or use multiple origins
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS") // Include the HTTP methods you want to allow
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	// Сериализуем полученные данные из БД, чтобы вывести их в Swagger
	res, _ := json.Marshal(repo.GetAllEmployees())
	// Явно указываем тип возвращаемого значения
	w.Header().Set("Content-Type", "application/json")
	_, err := w.Write(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

}

// EmployeeHandler - ручка получения профиля сотрудника по ID
//
//	@Summary		REST ручка для получения сотрудника по ID
//	@ID				empl-by-id
//	@Accept json
//	@Param		id	path		int		true	"Empl ID"
//	@Produce json
//	@Tags employees
//	@Router			/empl/{id} [get]
func EmployeeHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	w.Header().Set("Access-Control-Allow-Origin", "*")                                // Change '*' to the specific origin you want to allow or use multiple origins
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS") // Include the HTTP methods you want to allow
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	res, _ := json.Marshal(repo.GetEmployeeByID(id))
	w.Header().Set("Content-Type", "application/json")
	_, err := w.Write(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// NewDeptHandler - ручка создания отдела
//
//	@Summary		REST ручка для создания нового отдела
//	@ID				newdept
//	@Accept mpfd
//	@Param		deptName	formData	string	true	"department name"
//	@Produce json
//	@Tags departments
//	@Router			/dept/newdept [post]
func NewDeptHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")                                // Change '*' to the specific origin you want to allow or use multiple origins
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS") // Include the HTTP methods you want to allow
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	err := r.ParseMultipartForm(10)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error parsing form data", http.StatusBadRequest)
		return
	}
	data := r.Form.Get("deptName")
	var dept models.Department
	dept.DepartmentName = data
	deptID, err := repo.CreateDepartment(dept)
	if err != nil {
		// Если возникает ошибка, то выводим, что пользователя, который должен быть автором поста не существует
		fmt.Println(err)
		http.Error(w, "This user doesnt exist", http.StatusNotFound)
		return
	}
	resp := []byte("Department " + strconv.Itoa(deptID) + " created successfully")
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(resp)
	if err != nil {
		return
	}
}

// AllDeptsHandler - ручка вывода всех отделов
//
//	@Summary		REST ручка для получения всех отделов
//	@ID				all-depts
//	@Accept json
//	@Produce json
//	@Tags departments
//	@Router			/dept [get]
func AllDeptsHandler(w http.ResponseWriter, r *http.Request) {
	depts := repo.GetAllDepartments()
	res, _ := json.Marshal(depts)
	w.Header().Set("Content-Type", "application/json")
	_, err := w.Write(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// EditDeptHandler - ручка изменения отдела
//
//	@Summary		REST ручка для изменения отдела
//	@ID				edit-dept
//	@Accept mpfd
//	@Param		deptName	formData	string	true	"departmentName"
//	@Param		id			path		int		true	"departmentId"
//	@Produce json
//	@Tags departments
//	@Router			/dept/{id}/edit [put]
func EditDeptHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")                                // Change '*' to the specific origin you want to allow or use multiple origins
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS") // Include the HTTP methods you want to allow
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	// Parse form data
	err := r.ParseMultipartForm(10)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error parsing form data", http.StatusBadRequest)
		return
	}

	// Access form fields
	data := r.Form.Get("deptName")
	var dept models.Department
	dept.ID = id
	dept.DepartmentName = data
	repo.UpdateDepartment(dept)
	if err != nil {
		return
	}
}

// DeleteDeptHandler - ручка удаления отдела
//
//	@Summary		REST ручка удаления отдела
//	@ID				delete-dept
//	@Accept json
//	@Param		id			path		int		true	"deptID"
//	@Produce json
//	@Tags departments
//	@Router			/dept/{id}/delete [delete]
func DeleteDeptHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")                                // Change '*' to the specific origin you want to allow or use multiple origins
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS") // Include the HTTP methods you want to allow
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	repo.DeleteDepartment(id)
	_, err := w.Write([]byte("Department deleted successfully"))
	if err != nil {
		return
	}
}

// DeptHandler - ручка получения отдела по ID
//
//	@Summary		REST ручка получения отдела по ID
//	@ID				get-dept-by-id
//	@Accept json
//	@Param		id			path		int		true	"postId"
//	@Produce json
//	@Tags departments
//	@Router			/dept/{id} [get]
func DeptHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	w.Header().Set("Access-Control-Allow-Origin", "*")                                // Change '*' to the specific origin you want to allow or use multiple origins
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS") // Include the HTTP methods you want to allow
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	res, _ := json.Marshal(repo.GetDepartmentByID(id))
	w.Header().Set("Content-Type", "application/json")
	_, err := w.Write(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
