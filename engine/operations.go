package engine

import (
	"database/sql"
	"fmt"
	"kubernetes-postgres/datastore"
	"kubernetes-postgres/servicelog"
	"kubernetes-postgres/utility"
	"net/http"
	"strconv"
	"time"
)

// @Title GetCollection
// @Description get all collection items
// @Produce  json
// @Success 200 {string} OK
// @Failure 404 {string} not found
// @Failure 500 {string} internal server error
// @Router /todos [GET]
func GetCollection(w http.ResponseWriter, r *http.Request) {
	fmt.Println("sono qui")
	// Send the all collection elements
	todos, err := datastore.GetCollection()
	if err != nil {
		logger := servicelog.GetInstance()
		logger.Println(time.Now().UTC(), "Get Collection failed")
		// send error
		utility.EncodeToJsonError(w)
		return
	}
	utility.EncodeToJsonWithBody(w, todos)
}

// @Title GetElement
// @Description get an Element of the collection
// @Produce  json
// @Param id path string false "itemtoget"
// @Success 200 {string} OK
// @Failure 500 {string} not found
// @Failure 500 {string} internal server error
// @Router /todos/{id} [get]
func GetElement(w http.ResponseWriter, r *http.Request) {
	logger := servicelog.GetInstance()
	var todo *datastore.TodoElement
	// taking the id
	indexString := r.URL.Path[len("/todos/"):]
	if todoId, err := strconv.Atoi(indexString); err == nil {
		todo, err = datastore.Get(todoId)
		if err != nil {
			// Not found
			if err == sql.ErrNoRows {
				utility.EncodeToJsonNotFound(w)
				return
			}
			logger.Println(time.Now().UTC(), "GetElement failed")
			// send error
			utility.EncodeToJsonError(w)
			return
		}
		// send the element requested
		utility.EncodeToJsonWithBody(w, todo)
	} else {
		logger.Println(time.Now().UTC(), "GetElement failed")
		// send error
		utility.EncodeToJsonError(w)
	}
}

// @Title CreateElement
// @Description create an element for the collection
// @Accept  json
// @Param newtodoitem body string false "{Topic:New TodoElem, Completed:0}"
// @Produce  json
// @Success 200 {string} OK
// @Failure 500 {string} internal server error
// @Router /todos [PUT]
func CreateElement(w http.ResponseWriter, r *http.Request) {
	var newTodo datastore.TodoElement
	var err error
	var id int64
	logger := servicelog.GetInstance()
	if newTodo, err = utility.MarshallJsonAndResponse(w, r); err != nil {
		logger.Println(time.Now().UTC(), "CreateElement failed")
		// returns error
		utility.EncodeToJsonError(w)
		return
	}
	if id, err = datastore.Put(newTodo); err != nil {
		logger.Println(time.Now().UTC(), "CreateElement failed")
		// send error
		utility.EncodeToJsonError(w)
		return
	}
	// Send created with new resource id
	utility.EncodeToJson(w, r, id)
}

// @Title UpdateElement
// @Description Update an element in the collection
// @Accept  json
// @Param newtodoitem body string false "{Topic:New TodoElem, Completed:0}"
// @Produce  json
// @Success 200 {string} OK
// @Failure 500 {string} not found
// @Failure 500 {string} internal server error
// @Router /todos [POST]
func UpdateElement(w http.ResponseWriter, r *http.Request) {
	var updatedTodo datastore.TodoElement
	var err error
	var updatedid int64
	logger := servicelog.GetInstance()
	if updatedTodo, err = utility.MarshallJsonAndResponse(w, r); err != nil {
		// returns error
		utility.EncodeToJsonError(w)
		return
	}

	if updatedid, err = datastore.Update(updatedTodo); err != nil {
		// Not found
		if err == sql.ErrNoRows {
			utility.EncodeToJsonNotFound(w)
			return
		}
		logger.Println(time.Now().UTC(), "UpdateElement failed")
		// send error
		utility.EncodeToJsonError(w)
		return
	}
	// Send created with new resource id
	utility.EncodeToJson(w, r, updatedid)
}

// @Title GetCollection
// @Description Delete all items in the collection
// @Produce  json
// @Success 200 {string} OK
// @Failure 500 {string} not found
// @Failure 500 {string} internal server error
// @Router /todos [DELETE]
func DeleteCollection(w http.ResponseWriter, r *http.Request) {
	datastore.DeleteCollection()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

// @Title DeleteElement
// @Description Delete an element in the collection
// @Accept  json
// @Param id path string false "itemtodelete"
// @Produce  json
// @Param itemtodelete query string false "{id}"
// @Success 200 {string} OK
// @Failure 500 {string} not found
// @Failure 500 {string} internal server error
// @Router /todos/{id} [DELETE]
func DeleteElement(w http.ResponseWriter, r *http.Request) {
	// taking the id
	indexString := r.URL.Path[len("/todos/"):]
	logger := servicelog.GetInstance()
	if todoId, err := strconv.Atoi(indexString); err == nil {
		if err := datastore.DeleteElement(todoId); err != nil {
			// Not found
			if err == sql.ErrNoRows {
				utility.EncodeToJsonNotFound(w)
				return
			}
			logger.Println(time.Now().UTC(), "DeleteElement failed")
			// send error
			utility.EncodeToJsonError(w)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	} else {
		utility.EncodeToJsonError(w)
	}
}
