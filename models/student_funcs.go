package models

import (
	"fmt"
	"school/storage"

	"gorm.io/gorm"
)

type StudentJSON struct {
	gorm.Model
	FullName string `json:"name"`
	CareerID int    `json:"carid"`
	Mail     string `json:"mail"`
	Phone    string `json:"phone"`
}

func CreateStudent(conn *gorm.DB, name string, careerid int, mail string, phone string) error {
	stud := Student{
		FullName: name,
		CareerID: careerid,
		Mail:     storage.StringToNull(mail),
		Phone:    storage.StringToNull(phone),
	}
	conn.Create(&stud)
	return nil
}

func GetAllStudents(conn *gorm.DB) []StudentJSON {
	student := make([]Student, 0)
	conn.Find(&student)
	sslice := make([]StudentJSON, 0)
	for _, sp := range student {
		s := StudentJSON{
			Model:    sp.Model,
			FullName: sp.FullName,
			CareerID: sp.CareerID,
			Mail:     sp.Mail.String,
			Phone:    sp.Phone.String,
		}
		fmt.Println(s)
		sslice = append(sslice, s)
	}
	return sslice
}

func GetStudentById(conn *gorm.DB, id int) StudentJSON {
	student := Student{}
	conn.First(&student, id)
	s := StudentJSON{
		Model:    student.Model,
		FullName: student.FullName,
		CareerID: student.CareerID,
		Mail:     student.Mail.String,
		Phone:    student.Phone.String,
	}

	return s
}

func UpdateStudent(conn *gorm.DB, sp StudentJSON) {
	student := Student{
		Model:    sp.Model,
		FullName: sp.FullName,

		CareerID: sp.CareerID,
		Mail:     storage.StringToNull(sp.Mail),
		Phone:    storage.StringToNull(sp.Phone),
	}
	student.ID = sp.ID
	conn.Model(&student).Updates(student)
}

func DeleteStudent(conn *gorm.DB, id int) {
	sp := Student{}
	sp.ID = uint(id)
	conn.Delete(&sp)
}
