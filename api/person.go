package api

import (
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/OPENCBS/server/repo"
	"github.com/OPENCBS/server/model"
)

/*
func GetPerson(ctx *app.AppContext, w http.ResponseWriter, r *http.Request) {
	idString := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idString)
	repo := repo.NewPersonRepo(ctx.DbProvider)
	person, err := repo.GetById(id)
	if err != nil {
		fail(w, err)
		return
	}
	if person == nil {
		http.Error(w, "Person not found", http.StatusNotFound)
		return
	}

	sendJson(w, person)
}
*/

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
	fmt.Printf("%+v\n", person)

	repo := repo.NewPersonRepo(r.DbProvider)
	result, err := repo.Add(&person)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%+v\n", result)
	}
}

func validatePerson(person model.Person) error {
	return nil
}
