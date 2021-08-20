package handlers

import (
	"encoding/json"
	"net/http"
	"school/models"
	"school/storage"
	"strconv"
)

func createCareer(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response := newResponse(Error, "not allowed method", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}
	data := models.Career{}
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
	err = models.CreateCareer(conn, data.Name, data.StudyPlanID)
	if err != nil {
		response := newResponse(Error, "a problem occurred while creating the model", nil)
		responseJSON(w, http.StatusInternalServerError, response)
	}
	response := newResponse(Message, "career created succesfuly", nil)
	responseJSON(w, http.StatusCreated, response)
}

func getAllCareers(w http.ResponseWriter, r *http.Request) {
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
	res := models.GetAllCareers(conn)
	response := newResponse(Message, "ok", res)
	responseJSON(w, http.StatusInternalServerError, response)
}

func getCareerById(w http.ResponseWriter, r *http.Request) {
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
	res := models.GetCareerById(conn, ID)
	response := newResponse(Message, "ok", res)
	responseJSON(w, http.StatusInternalServerError, response)

}

func updateCareer(w http.ResponseWriter, r *http.Request) {
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

	data := models.Career{}
	err = json.NewDecoder(r.Body).Decode(&data)
	data.ID = uint(ID)
	if err != nil {
		response := newResponse(Error, "career's struct isn't correct", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	conn := storage.NewConnection(driver)
	models.UpdateCareer(conn, data)
	response := newResponse(Message, "ok", nil)
	responseJSON(w, http.StatusInternalServerError, response)

}

func deleteCareer(w http.ResponseWriter, r *http.Request) {
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
	models.DeleteCareer(conn, ID)
	response := newResponse(Message, "model deleted succesfuly", nil)
	responseJSON(w, http.StatusInternalServerError, response)
}
