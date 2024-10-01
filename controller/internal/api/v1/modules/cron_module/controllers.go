package cron_module

import (
	"fmt"
	res "go-http/internal/api/response"
	"go-http/pkg/validate"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

const (
	GetCronsUrl   = "/cron"
	GetCronUrl    = "/cron/{id}"
	CreateCronUrl = "/cron"
	DeleteCronUrl = "/cron/{id}"
	UpdateCronUrl = "/cron/{id}"
)

func Router(api *chi.Mux) {
	api.Get(GetCronsUrl, GetCrons)
	api.Get(GetCronUrl, GetCron)
	api.Post(CreateCronUrl, CreateCron)
	api.Delete(DeleteCronUrl, DeleteCron)
	api.Put(UpdateCronUrl, UpdateCron)
}

// GET URL/cron
func GetCrons(w http.ResponseWriter, r *http.Request) {

	// replace the function to return the data from the db
	data, err := placeholderFunc()
	if err != nil {
		fmt.Printf("Error occured while calling db: %v\n", err) // replace with logger
		res.Response(w, 400, nil, res.FailedDbConnMessage)
		return
	}

	res.Response(w, 200, data, "", res.PaginationLinks{})
}

// GET URL/cron/[id:int]
func GetCron(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		res.Response(w, 400, nil, res.ParamIsNotIntMessage)
		return
	}
	_ = id

	// replace the function to return the data from the db
	data, err := placeholderFunc()
	if err != nil {
		fmt.Printf("Error occured while calling db: %v\n", err) // replace with logger
		res.Response(w, 400, nil, res.FailedDbConnMessage)
		return
	}

	res.Response(w, 200, data, "")
}

// POST URL/cron
func CreateCron(w http.ResponseWriter, r *http.Request) {

	// validate the sent body
	payload := new(CreateCronBody)
	if err := validate.RequestBody(r, payload); err != nil {
		res.Response(w, 400, err, res.FailedPayloadValidationMessage)
		return
	}

	// replace the function to return the data from the db
	data, err := placeholderFunc()
	if err != nil {
		fmt.Printf("Error occured while calling db: %v\n", err) // replace with logger
		res.Response(w, 400, nil, res.FailedDbConnMessage)
		return
	}

	res.Response(w, 200, data, "")
}

// DELETE URL/cron/[id:string]
func DeleteCron(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")
	_ = id

	// replace the function to return the data from the db
	data, err := placeholderFunc()
	if err != nil {
		fmt.Printf("Error occured while calling db: %v\n", err) // replace with logger
		res.Response(w, 400, nil, res.FailedDbConnMessage)
		return
	}

	res.Response(w, 200, data, "")
}

// PUT URL/cron/[id:int]
func UpdateCron(w http.ResponseWriter, r *http.Request) {

	// validate the sent body
	payload := new(UpdateCronBody)
	if err := validate.RequestBody(r, payload); err != nil {
		res.Response(w, 400, err, res.FailedPayloadValidationMessage)
		return
	}

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		res.Response(w, 400, nil, res.ParamIsNotIntMessage)
		return
	}
	_ = id

	// replace the function to return the data from the db
	data, err := placeholderFunc()
	if err != nil {
		fmt.Printf("Error occured while calling db: %v\n", err) // replace with logger
		res.Response(w, 400, nil, res.FailedDbConnMessage)
		return
	}

	res.Response(w, 200, data, "")
}

// Placeholder function to be replaced in the controllers after codegen.
// Usually replaced with the function which returns the data from the db.
func placeholderFunc() (struct{}, error) {
	var placeholder struct{}
	return placeholder, nil
}
