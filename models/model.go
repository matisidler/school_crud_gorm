package models

import (
	"database/sql"

	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	FullName   string         `gorm:"type:varchar(50);not null" json:"name"`
	CareerID   int            `gorm:"type:int;not null" json:"carid"`
	CareerName string         `gorm:"type:varchar(50)"`
	Mail       sql.NullString `gorm:"type:varchar(50);" json:"mail"`
	Phone      sql.NullString `gorm:"type:varchar(50);" json:"phone"`
}

type Career struct {
	gorm.Model
	Name             string `gorm:"type:varchar(50);not null" json:"name"`
	StudentsQuantity int    `gorm:"type:int;not null"`
	TeachersQuantity int    `gorm:"type:int;not null"`
	StudyPlanID      int    `gorm:"type:int;not null" json:"spid"`
	Student          Student
	Teachers         Teacher
}

type Teacher struct {
	gorm.Model
	Name        string         `gorm:"type:varchar(50);not null" json:"name"`
	Salary      float32        `gorm:"type:real;not null" json:"salary"`
	CareerID    int            `gorm:"type:int;not null" json:"carid"`
	CareerName  string         `gorm:"type:varchar(50)"`
	SubjectID   int            `gorm:"type:int;not null" json:"subid"`
	SubjectName string         `gorm:"type:varchar(50)"`
	Mail        sql.NullString `gorm:"type:varchar(50)" json:"mail"`
	Phone       sql.NullString `gorm:"type:varchar(50)" json:"phone"`
}

type StudyPlan struct {
	gorm.Model
	FirstYear  string `gorm:"type:varchar(150)" json:"fy"`
	SecondYear string `gorm:"type:varchar(150)" json:"sy"`
	ThirdYear  string `gorm:"type:varchar(150)" json:"ty"`
	Career     Career
}

type Subject struct {
	gorm.Model
	Name    string `gorm:"type:varchar(150)" json:"name"`
	Teacher Teacher
}
