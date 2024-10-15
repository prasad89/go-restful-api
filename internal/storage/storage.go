package storage

import "github.com/prasad89/go-restful-api/internal/types"

type Storage interface {
	CreateStudent(name string, email string, age int) (int64, error)
	GetStudents() ([]types.Student, error)
	GetStudent(id int64) (types.Student, error)
	UpdateStudent(name string, email string, age int) (int64, error)
	DeleteStudent(id int64) (int64, error)
}
