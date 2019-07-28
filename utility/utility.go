package utility

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"kubernetes-postgres/datastore"
	"kubernetes-postgres/servicelog"
	"net/http"
	"strconv"
	"time"
)

var Credentials string = ""
var Host string = ""

func MarshallJson(r *http.Request) (datastore.TodoElement, error) {
	var todo datastore.TodoElement
	logger := servicelog.GetInstance()
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		logger.Println(time.Now().UTC(), "Error in MarshallJson")
		return todo, err
	}
	if err := r.Body.Close(); err != nil {
		logger.Println(time.Now().UTC(), "Error in MarshallJson")
		return todo, err
	}
	if err := json.Unmarshal(body, &todo); err != nil {
		return todo, err
	}
	return todo, nil
}

// Marshall the json data in input inside the todoElement structure (from network)
func MarshallJsonAndResponse(w http.ResponseWriter, r *http.Request) (datastore.TodoElement, error) {
	var todo datastore.TodoElement
	logger := servicelog.GetInstance()
	var err error

	if todo, err = MarshallJson(r); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			logger.Println(time.Now().UTC(), "Error in MarshallJson")
			return todo, err
		}
	}

	return todo, nil
}

// Encode the json response for GET and PUT (with body)
func EncodeToJsonWithBody(w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	js, err := json.Marshal(response)
	if err != nil {
		logger := servicelog.GetInstance()
		logger.Println(time.Now().UTC(), "Decoding error")
		// returns error
		w.WriteHeader(http.StatusInternalServerError)
		return err
	}
	w.Write(js)
	return nil
}

// Encode the json response without body (PUT and POST)
func EncodeToJson(w http.ResponseWriter, r *http.Request, index int64) {
	// send the index of the element just created in the header location field
	resource := r.Host + "/" + strconv.Itoa(int(index))
	w.Header().Set("Location", resource)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func EncodeToJsonError(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
}

func EncodeToJsonNotFound(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)

}
