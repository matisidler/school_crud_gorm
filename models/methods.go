package models

import (
	"gorm.io/gorm"
)

func (s *Student) AfterCreate(tx *gorm.DB) (err error) {
	id := s.CareerID
	tx.Table("careers").Where("id", id).Update("students_quantity", gorm.Expr("students_quantity + ?", 1))
	tx.Model(&s).Update("career_name", tx.Model(&Career{}).Select("name").Where("students.career_id = careers.id"))
	return
}

func (s *Teacher) AfterCreate(tx *gorm.DB) (err error) {
	id := s.CareerID
	tx.Table("careers").Where("id", id).Update("teachers_quantity", gorm.Expr("teachers_quantity + ?", 1))
	tx.Model(&s).Update("career_name", tx.Model(&Career{}).Select("name").Where("teachers.career_id = careers.id"))
	tx.Model(&s).Update("subject_name", tx.Model(&Subject{}).Select("name").Where("teachers.career_id = subjects.id"))
	return
}
