package models

import (
	"database/sql"

	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	FullName   string         `gorm:"type:varchar(50);not null"`
	CareerID   int            `gorm:"type:int;not null"`
	CareerName string         `gorm:"type:varchar(50);not null"`
	Mail       sql.NullString `gorm:"type:varchar(50);"`
	Phone      sql.NullString `gorm:"type:varchar(50);"`
}

type Career struct {
	gorm.Model
	Name             string `gorm:"type:varchar(50);not null"`
	StudentsQuantity int    `gorm:"type:int;not null"`
	TeachersQuantity int    `gorm:"type:int;not null"`
	StudyPlanID      int    `gorm:"type:int;not null"`
	Student          Student
	Teachers         Teacher
}

type Teacher struct {
	gorm.Model
	Name        string         `gorm:"type:varchar(50);not null"`
	Salary      float32        `gorm:"type:real;not null"`
	CareerID    int            `gorm:"type:int;not null"`
	CareerName  string         `gorm:"type:varchar(50)"`
	SubjectID   int            `gorm:"type:int;not null"`
	SubjectName string         `gorm:"type:varchar(50)"`
	Mail        sql.NullString `gorm:"type:varchar(50)"`
	Phone       sql.NullString `gorm:"type:varchar(50)"`
}

type StudyPlan struct {
	gorm.Model
	FirstYear  string `gorm:"type:varchar(150)"`
	SecondYear string `gorm:"type:varchar(150)"`
	ThirdYear  string `gorm:"type:varchar(150)"`
	Career     Career
}

type Subject struct {
	gorm.Model
	Name    string `gorm:"type:varchar(150)"`
	Teacher Teacher
}
