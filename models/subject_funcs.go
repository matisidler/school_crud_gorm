package models

import (
	"fmt"

	"gorm.io/gorm"
)

type SubjectsGA struct {
	gorm.Model
	Name string
}

func CreateSubject(conn *gorm.DB, name string) error {
	subj := Subject{
		Name: name,
	}
	conn.Create(&subj)
	return nil
}
func GetAllSubjects(conn *gorm.DB) []SubjectsGA {
	subjects := make([]Subject, 0)
	conn.Find(&subjects)
	sslice := make([]SubjectsGA, 0)
	for _, sp := range subjects {
		s := SubjectsGA{
			Model: sp.Model,
			Name:  sp.Name,
		}
		fmt.Println(s)
		sslice = append(sslice, s)
	}
	return sslice
}

func GetSubjectById(conn *gorm.DB, id int) SubjectsGA {
	subject := Subject{}
	conn.First(&subject, id)
	s := SubjectsGA{
		Model: subject.Model,
		Name:  subject.Name,
	}

	return s
}

func UpdateSubject(conn *gorm.DB, sp Subject) Subject {
	subject := Subject{}
	subject.ID = sp.ID
	conn.Model(&subject).Updates(sp)
	return subject
}

func DeleteSubject(conn *gorm.DB, id int) {
	sp := Subject{}
	sp.ID = uint(id)
	conn.Delete(&sp)
}
