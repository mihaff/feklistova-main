package repository

import (
	"feklistova/models"
	"fmt"
	"log"
)

// GetAllDepartments - метод для получения всех отделов в компании
func (r *Repository) GetAllDepartments() []models.Department {
	var depts []models.Department
	var dept models.Department
	query := "SELECT d.id, department_name, count(e.id) FROM departments d LEFT JOIN employees e on d.id = e.department_id GROUP BY d.id ORDER BY d.id"
	rows, err := r.Db.Query(query)
	if err != nil {
		log.Println(err)
		return nil
	}
	// После выполнения запроса сканируем каждую строку и записываем значения в структуру.
	for rows.Next() {
		err = rows.Scan(&dept.ID, &dept.DepartmentName, &dept.DepartmentEmployeeCount)
		if err != nil {
			log.Println(err)
			return nil
		}
		depts = append(depts, dept)
	}
	return depts
}

// GetDepartmentByID - метод получения отдела по его ID
func (r *Repository) GetDepartmentByID(id int) models.Department {
	var dept models.Department
	query := "SELECT d.id, department_name, count(e.id) FROM departments d LEFT JOIN employees e on d.id = e.department_id where d.id = $1 GROUP BY d.id ORDER BY d.id"
	err := r.Db.QueryRow(query, id).Scan(&dept.ID, &dept.DepartmentName, &dept.DepartmentEmployeeCount)
	if err != nil {
		fmt.Print(err)
	}
	return dept
}
