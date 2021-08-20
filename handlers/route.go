package handlers

import (
	"net/http"
)

//esta funcion routePerson me permite crear todos los handlers de person (create, update, etc)
func Route(mux *http.ServeMux) {
	sp := StudyPlan{}
	subj := Subject{}
	car := Career{}
	teac := Teacher{}
	stud := Student{}

	mux.HandleFunc("/studyplan/create", sp.create)
	mux.HandleFunc("/subject/create", subj.create)
	mux.HandleFunc("/career/create", car.create)
	mux.HandleFunc("/teacher/create", teac.create)
	mux.HandleFunc("/student/create", stud.create)

	mux.HandleFunc("/studyplan/getall", sp.getAllStudyPlan)

	/* mux.HandleFunc("/v1/persons/getall", middleware.Log(h.getAll))
	mux.HandleFunc("/v1/persons/update", middleware.Log(h.update))
	mux.HandleFunc("/v1/persons/delete", middleware.Log(h.delete))
	mux.HandleFunc("/v1/persons/getbyid", middleware.Log(h.getById)) */
}
