package repository

import (
	"feklistova/models"
	"log"
)

// GetAllEmployees является методом получения всех сотрудников.
// Может использоваться сотрудниками HR например для мониторинга пользователей.
func (r *Repository) GetAllEmployees() []models.Employee {
	var employees []models.Employee
	rows, err := r.Db.Query("SELECT e.id, e.surname, e.name, d.id, d.department_name FROM employees e JOIN departments d on e.department_id = d.id order by e.id")
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		e := models.Employee{}
		err := rows.Scan(&e.ID, &e.Surname, &e.Name, &e.Department.ID, &e.Department.DepartmentName)
		if err != nil {
			log.Println(err)
			continue
		}
		employees = append(employees, e)
	}
	err = rows.Close()
	if err != nil {
		log.Println(err)
		return nil
	}
	return employees
}

// GetEmployeeByID является методом получения пользователя по уникальному id.
// Может использоваться например для получения собственного профиля или
// для поиска конкретного сотрудника.
func (r *Repository) GetEmployeeByID(id int) models.Employee {
	var e models.Employee
	query := "SELECT e.id, e.surname, e.name, d.id, d.department_name FROM employees e JOIN departments d on e.department_id = d.id WHERE e.id = $1"
	row := r.Db.QueryRow(query, id)
	err := row.Scan(&e.ID, &e.Surname, &e.Name, &e.Department.ID, &e.Department.DepartmentName)
	if err != nil {
		return models.Employee{}
	}
	return e
}

// GetAllEmployeesInDept - метод получения всех сотрудников конкретного отдела. На вход принимает id отдела
func (r *Repository) GetAllEmployeesInDept(id int) []models.Employee {
	var employees []models.Employee
	var e models.Employee
	query := "SELECT d.department_name, d.id, e.id, e.surname, e.name from departments d join employees e on d.id = e.department_id where d.id = $1"
	rows, err := r.Db.Query(query, id)
	if err != nil {
		log.Println(err)
		return nil
	}
	for rows.Next() {
		err = rows.Scan(&e.Department.DepartmentName, &e.Department.ID, &e.ID, &e.Surname, &e.Name)
		if err != nil {
			log.Println(err)
			return nil
		}
		employees = append(employees, e)
	}
	return employees
}
