package models

import (
	"fmt"
	"school/storage"

	"gorm.io/gorm"
)

type TeacherJSON struct {
	gorm.Model
	Name      string  `json:"name"`
	Salary    float32 `json:"salary"`
	CareerID  int     `json:"carid"`
	SubjectID int     `json:"subid"`
	Mail      string  `json:"mail"`
	Phone     string  `json:"phone"`
}

func CreateTeacher(conn *gorm.DB, name string, salary float32, subjectid int, careerid int, mail string, phone string) error {
	teac := Teacher{
		Name:      name,
		Salary:    salary,
		SubjectID: subjectid,
		CareerID:  careerid,
		Mail:      storage.StringToNull(mail),
		Phone:     storage.StringToNull(phone),
	}
	conn.Create(&teac)
	return nil
}

func GetAllTeachers(conn *gorm.DB) []TeacherJSON {
	teachers := make([]Teacher, 0)
	conn.Find(&teachers)
	sslice := make([]TeacherJSON, 0)
	for _, sp := range teachers {
		s := TeacherJSON{
			Model:     sp.Model,
			Name:      sp.Name,
			Salary:    sp.Salary,
			SubjectID: sp.SubjectID,
			CareerID:  sp.CareerID,
			Mail:      sp.Mail.String,
			Phone:     sp.Phone.String,
		}
		fmt.Println(s)
		sslice = append(sslice, s)
	}
	return sslice
}

func GetTeacherById(conn *gorm.DB, id int) TeacherJSON {
	teacher := Teacher{}
	conn.First(&teacher, id)
	s := TeacherJSON{
		Model:     teacher.Model,
		Name:      teacher.Name,
		Salary:    teacher.Salary,
		SubjectID: teacher.SubjectID,
		CareerID:  teacher.CareerID,
		Mail:      teacher.Mail.String,
		Phone:     teacher.Phone.String,
	}

	return s
}

func UpdateTeacher(conn *gorm.DB, sp TeacherJSON) {

	teacher := Teacher{
		Model:     sp.Model,
		Name:      sp.Name,
		Salary:    sp.Salary,
		SubjectID: sp.SubjectID,
		CareerID:  sp.CareerID,
		Mail:      storage.StringToNull(sp.Mail),
		Phone:     storage.StringToNull(sp.Phone),
	}
	teacher.ID = sp.ID
	conn.Model(&teacher).Updates(teacher)
}

func DeleteTeacher(conn *gorm.DB, id int) {
	sp := Teacher{}
	sp.ID = uint(id)
	conn.Delete(&sp)
}
