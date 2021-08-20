package handlers

import (
	"encoding/json"
	"net/http"
	"school/models"
	"school/storage"
	"strconv"
)

func createStudyPlan(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response := newResponse(Error, "not allowed method", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}
	data := models.StudyPlan{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		response := newResponse(Error, "study plan's struct isn't correct", nil)
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
	err = models.CreateStudyPlan(conn, data.FirstYear, data.SecondYear, data.ThirdYear)
	if err != nil {
		response := newResponse(Error, "a problem occurred while creating study plan", nil)
		responseJSON(w, http.StatusInternalServerError, response)
	}
	response := newResponse(Message, "studyplan created succesfuly", nil)
	responseJSON(w, http.StatusCreated, response)
}

func getAllStudyPlan(w http.ResponseWriter, r *http.Request) {
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
	res := models.GetAll(conn)
	response := newResponse(Message, "ok", res)
	responseJSON(w, http.StatusInternalServerError, response)
}

func getStudyPlanById(w http.ResponseWriter, r *http.Request) {
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
	res := models.GetById(conn, ID)
	response := newResponse(Message, "ok", res)
	responseJSON(w, http.StatusInternalServerError, response)

}

func UpdateStudyPlan(w http.ResponseWriter, r *http.Request) {
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

	data := models.StudyPlan{}
	err = json.NewDecoder(r.Body).Decode(&data)
	data.ID = uint(ID)
	if err != nil {
		response := newResponse(Error, "study plan's struct isn't correct", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	conn := storage.NewConnection(driver)
	models.UpdateStudyPlan(conn, data)
	response := newResponse(Message, "ok", nil)
	responseJSON(w, http.StatusInternalServerError, response)

}

func deleteStudyPlan(w http.ResponseWriter, r *http.Request) {
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
	models.DeleteStudyPlan(conn, ID)
	response := newResponse(Message, "model deleted succesfuly", nil)
	responseJSON(w, http.StatusInternalServerError, response)
}
