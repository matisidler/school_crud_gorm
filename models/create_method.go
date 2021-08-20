package models

import (
	"log"
	"school/storage"

	"gorm.io/gorm"
)

func CreateStudyPlan(conn *gorm.DB, fy string, sy string, ty string) error {
	plan := StudyPlan{
		FirstYear:  fy,
		SecondYear: sy,
		ThirdYear:  ty,
	}
	conn.Create(&plan)
	log.Println("Created succesfuly")
	return nil
}

func CreateSubject(conn *gorm.DB, name string) error {
	subj := Subject{
		Name: name,
	}
	conn.Create(&subj)
	return nil
}

func CreateCareer(conn *gorm.DB, name string, spid int) error {
	car := Career{
		Name:        name,
		StudyPlanID: spid,
	}
	conn.Create(&car)
	return nil
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

func CreateStudent(conn *gorm.DB, name string, careerid int, mail string, phone string) error {
	stud := Student{
		FullName: "Matias Sidler",
		CareerID: 1,
		Mail:     storage.StringToNull(mail),
		Phone:    storage.StringToNull(phone),
	}
	conn.Create(&stud)
	return nil
}
