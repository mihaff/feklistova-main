package models

// Employee представляет собой сущность сотрудника компании
type Employee struct {
	ID         int
	Surname    string
	Name       string
	Department Department
}
