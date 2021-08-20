package models

import (
	"fmt"

	"gorm.io/gorm"
)

type StudyPlanGA struct {
	gorm.Model
	FirstYear  string
	SecondYear string
	ThirdYear  string
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
