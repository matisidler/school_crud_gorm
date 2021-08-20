package models

import (
	"fmt"

	"gorm.io/gorm"
)

type CareerGA struct {
	gorm.Model
	Name             string
	StudentsQuantity int
	TeachersQuantity int
	StudyPlanID      int
}

func CreateCareer(conn *gorm.DB, name string, spid int) error {
	car := Career{
		Name:        name,
		StudyPlanID: spid,
	}
	conn.Create(&car)
	return nil
}

func GetAllCareers(conn *gorm.DB) []CareerGA {
	careers := make([]Career, 0)
	conn.Find(&careers)
	sslice := make([]CareerGA, 0)
	for _, sp := range careers {
		s := CareerGA{
			Model:            sp.Model,
			Name:             sp.Name,
			StudentsQuantity: sp.StudentsQuantity,
			TeachersQuantity: sp.TeachersQuantity,
			StudyPlanID:      sp.StudyPlanID,
		}
		fmt.Println(s)
		sslice = append(sslice, s)
	}
	return sslice
}

func GetCareerById(conn *gorm.DB, id int) CareerGA {
	career := Career{}
	conn.First(&career, id)
	s := CareerGA{
		Model:            career.Model,
		Name:             career.Name,
		StudentsQuantity: career.StudentsQuantity,
		TeachersQuantity: career.TeachersQuantity,
		StudyPlanID:      career.StudyPlanID,
	}

	return s
}

func UpdateCareer(conn *gorm.DB, sp Career) Career {
	career := Career{}
	career.ID = sp.ID
	conn.Model(&career).Updates(sp)
	return career
}

func DeleteCareer(conn *gorm.DB, id int) {
	sp := Career{}
	sp.ID = uint(id)
	conn.Delete(&sp)
}
