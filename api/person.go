package api

import (
	"net/http"
	"encoding/json"
	"github.com/OPENCBS/server/repo"
	"github.com/OPENCBS/server/model"
)

func GetPeople(w http.ResponseWriter, r *APIRequest) {
	repo := repo.NewPersonRepo(r.DbProvider)
	offset, limit := r.GetRange()
	people, err := repo.GetPeople(offset, limit)
	if err != nil {
		sendInternalServerError(w, err)
		return
	}
	if r.CanAcceptGzip() {
		sendCompressedJson(w, people)
	} else {
		sendJson(w, people)
	}
}

func AddPerson(w http.ResponseWriter, r *APIRequest) {
	decoder := json.NewDecoder(r.Body);
	var person model.Person
	err := decoder.Decode(&person)
	if err != nil {
		sendInternalServerError(w, err)
		return
	}
	err = validatePerson(person)
	if err != nil {
		apiError := &ApiError{err.Error(), err.Error(), ""}
		sendJsonWithStatus(w, apiError, 422)
		return
	}

	repo := repo.NewPersonRepo(r.DbProvider)
	result, err := repo.Add(&person)
	if err != nil {
		sendInternalServerError(w, err)
		return
	}
	sendJsonWithStatus(w, result, http.StatusCreated)
}

func validatePerson(person model.Person) error {
	return nil
}
