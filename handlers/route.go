package handlers

import (
	"net/http"
)

//esta funcion routePerson me permite crear todos los handlers de person (create, update, etc)
func Route(mux *http.ServeMux) {

	mux.HandleFunc("/studyplan/create", createStudyPlan)
	mux.HandleFunc("/subject/create", createSubject)
	mux.HandleFunc("/career/create", createCareer)
	mux.HandleFunc("/teacher/create", createTeacher)
	mux.HandleFunc("/student/create", createStudent)

	mux.HandleFunc("/studyplan/getall", getAllStudyPlan)
	mux.HandleFunc("/subject/getall", getAllSubjects)
	mux.HandleFunc("/career/getall", getAllCareers)
	mux.HandleFunc("/teacher/getall", getAllTeachers)
	mux.HandleFunc("/student/getall", getAllStudents)

	mux.HandleFunc("/studyplan", getStudyPlanById)
	mux.HandleFunc("/subject", getSubjectById)
	mux.HandleFunc("/career", getCareerById)
	mux.HandleFunc("/teacher", getTeacherById)
	mux.HandleFunc("/student", getStudentById)

	mux.HandleFunc("/studyplan/delete", deleteStudyPlan)
	mux.HandleFunc("/subject/delete", deleteSubject)
	mux.HandleFunc("/career/delete", deleteCareer)
	mux.HandleFunc("/teacher/delete", deleteTeacher)
	mux.HandleFunc("/student/delete", deleteStudent)

	mux.HandleFunc("/studyplan/update", UpdateStudyPlan)
	mux.HandleFunc("/subject/update", UpdateSubject)
	mux.HandleFunc("/career/update", updateCareer)
	mux.HandleFunc("/teacher/update", updateTeacher)
	mux.HandleFunc("/student/update", updateStudent)
	/* mux.HandleFunc("/v1/persons/getall", middleware.Log(h.getAll))
	mux.HandleFunc("/v1/persons/update", middleware.Log(h.update))
	mux.HandleFunc("/v1/persons/delete", middleware.Log(h.delete))
	mux.HandleFunc("/v1/persons/getbyid", middleware.Log(h.getById)) */
}
