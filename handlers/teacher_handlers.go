package handlers

import (
	"encoding/json"
	"net/http"
	"school/models"
	"school/storage"
)

func (s *Teacher) create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response := newResponse(Error, "not allowed method", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}
	data := Teacher{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		response := newResponse(Error, "the struct isn't correct", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}
	driver := data.Driver
	if driver != "psql" && driver != "mysql" {
		response := newResponse(Error, "wrong driver", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}
	conn := storage.NewConnection(driver)
	err = models.CreateTeacher(conn, data.Name, data.Salary, data.SubjectId, data.CareerId, data.Mail, data.Phone)
	if err != nil {
		response := newResponse(Error, "a problem occurred while creating the model", nil)
		responseJSON(w, http.StatusInternalServerError, response)
	}
	response := newResponse(Message, "teacher created succesfuly", nil)
	responseJSON(w, http.StatusCreated, response)
}