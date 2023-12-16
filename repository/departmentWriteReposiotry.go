package repository

import (
	"feklistova/models"
	"log"
)

// CreateDepartment - метод записи нового отдела в базу данных. Для работы использует открытое подключение из Repository
func (r *Repository) CreateDepartment(department models.Department) (int, error) {
	var id int
	query := "INSERT INTO departments (department_name) VALUES ($1) RETURNING id"

	// db.Query выполняет SQL запрос, который указан в строке выше, при этом может возвращать значение.
	row, err := r.Db.Query(query, department.DepartmentName)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	// Так как запрос возвращает ID, требуется его записать в переменную
	for row.Next() {
		err = row.Scan(&id)
		if err != nil {
			return 0, err
		}
	}
	return id, nil
}

// UpdateDepartment - метод изменения отдела. На вход принимает структуру отдела, редактируя его в БД
func (r *Repository) UpdateDepartment(department models.Department) {
	query := "UPDATE departments SET department_name = $1 WHERE id = $2"
	// db.Exec выполняет SQL запрос без возврата значений
	_, err := r.Db.Exec(query, department.DepartmentName, department.ID)
	if err != nil {
		return
	}
}

// DeleteDepartment - метод удаления отдела и увольнения всех сотрудников по ID отдела
func (r *Repository) DeleteDepartment(id int) {
	query := "DELETE FROM employees WHERE department_id = $1"
	_, err := r.Db.Exec(query, id)
	if err != nil {
		return
	}
	query = "DELETE FROM departments WHERE id = $1"
	_, err = r.Db.Exec(query, id)
	if err != nil {
		return
	}
}
