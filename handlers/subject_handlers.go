package handlers

import (
	"encoding/json"
	"net/http"
	"school/models"
	"school/storage"
	"strconv"
)

func createSubject(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response := newResponse(Error, "not allowed method", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}
	data := models.Subject{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		response := newResponse(Error, "the struct isn't correct", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}
	driver := r.Header.Get("driver")
	if driver != "psql" && driver != "mysql" {
		response := newResponse(Error, "wrong driver", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}
	conn := storage.NewConnection(driver)
	err = models.CreateSubject(conn, data.Name)
	if err != nil {
		response := newResponse(Error, "a problem occurred while creating the model", nil)
		responseJSON(w, http.StatusInternalServerError, response)
	}
	response := newResponse(Message, "subject created succesfuly", nil)
	responseJSON(w, http.StatusCreated, response)
}

func getAllSubjects(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response := newResponse(Error, "not allowed method", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}
	driver := r.Header.Get("driver")
	if driver != "psql" && driver != "mysql" {
		response := newResponse(Error, "wrong driver", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}
	conn := storage.NewConnection(driver)
	res := models.GetAllSubjects(conn)
	response := newResponse(Message, "ok", res)
	responseJSON(w, http.StatusInternalServerError, response)
}

func getSubjectById(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response := newResponse(Error, "not allowed method", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}
	driver := r.Header.Get("driver")
	if driver != "psql" && driver != "mysql" {
		response := newResponse(Error, "wrong driver", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	ID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		response := newResponse(Error, "ID not valid", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	conn := storage.NewConnection(driver)
	res := models.GetSubjectById(conn, ID)
	response := newResponse(Message, "ok", res)
	responseJSON(w, http.StatusInternalServerError, response)

}

func UpdateSubject(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		response := newResponse(Error, "not allowed method", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}
	driver := r.Header.Get("driver")
	if driver != "psql" && driver != "mysql" {
		response := newResponse(Error, "wrong driver", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	ID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		response := newResponse(Error, "ID not valid", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	data := models.Subject{}
	err = json.NewDecoder(r.Body).Decode(&data)
	data.ID = uint(ID)
	if err != nil {
		response := newResponse(Error, "subject's struct isn't correct", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	conn := storage.NewConnection(driver)
	models.UpdateSubject(conn, data)
	response := newResponse(Message, "ok", nil)
	responseJSON(w, http.StatusInternalServerError, response)

}

func deleteSubject(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		response := newResponse(Error, "not allowed method", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}
	driver := r.Header.Get("driver")
	if driver != "psql" && driver != "mysql" {
		response := newResponse(Error, "wrong driver", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	ID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		response := newResponse(Error, "ID not valid", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	conn := storage.NewConnection(driver)
	models.DeleteSubject(conn, ID)
	response := newResponse(Message, "model deleted succesfuly", nil)
	responseJSON(w, http.StatusInternalServerError, response)
}
