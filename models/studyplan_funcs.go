package models

import (
	"fmt"
	"log"

	"gorm.io/gorm"
)

type StudyPlanGA struct {
	gorm.Model
	FirstYear  string
	SecondYear string
	ThirdYear  string
}

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

func GetAll(conn *gorm.DB) []StudyPlanGA {
	studyplan := make([]StudyPlan, 0)
	conn.Find(&studyplan)
	studyslice := make([]StudyPlanGA, 0)
	for _, sp := range studyplan {
		s := StudyPlanGA{
			Model:      sp.Model,
			FirstYear:  sp.FirstYear,
			SecondYear: sp.SecondYear,
			ThirdYear:  sp.ThirdYear,
		}
		fmt.Println(s)
		studyslice = append(studyslice, s)
	}

	return studyslice
}

func GetById(conn *gorm.DB, id int) StudyPlanGA {
	studyplan := StudyPlan{}
	conn.First(&studyplan, id)
	s := StudyPlanGA{
		Model:      studyplan.Model,
		FirstYear:  studyplan.FirstYear,
		SecondYear: studyplan.SecondYear,
		ThirdYear:  studyplan.ThirdYear,
	}

	return s
}

func UpdateStudyPlan(conn *gorm.DB, sp StudyPlan) StudyPlan {
	studyplan := StudyPlan{}
	studyplan.ID = sp.ID
	conn.Model(&studyplan).Updates(sp)
	return studyplan
}

func DeleteStudyPlan(conn *gorm.DB, id int) {
	sp := StudyPlan{}
	sp.ID = uint(id)
	conn.Delete(&sp)
}
