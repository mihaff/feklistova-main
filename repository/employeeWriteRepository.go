package repository

import (
	"feklistova/models"
)

// NewEmployee является методом создания нового сотрудника.
func (r *Repository) NewEmployee(employee models.Employee) (int, error) {
	var id int
	query := "INSERT INTO employees(surname, name, department_id) VALUES ($1, $2, $3) RETURNING id"
	row := r.Db.QueryRow(query, employee.Surname, employee.Name, employee.Department.ID)
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// UpdateEmployee является методом изменения профиля сотрудника
func (r *Repository) UpdateEmployee(employee models.Employee) {
	query := "UPDATE employees SET surname = $1, name = $2, department_id = $3 WHERE id = $4"
	_, err := r.Db.Exec(query, employee.Surname, employee.Name, employee.Department.ID, employee.ID)
	if err != nil {
		return
	}
}

// DeleteEmployee удаляет профиль сотрудника из базы данных
func (r *Repository) DeleteEmployee(id int) {
	query := "DELETE FROM employees WHERE id = $1"
	_, err := r.Db.Exec(query, id)
	if err != nil {
		return
	}
}
