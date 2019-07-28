package testing

import (
	"ricardo/engine"
	"testing"
	"net/http"
	"net/http/httptest"
	"reflect"
)

// Test which mock an http request. Create with POST an element and check with GET
func TestCreateAndGetItem(t *testing.T) {
	// Post a to-do element to the database
	elemCreated, reqCreate := MockPostRequest(t)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(engine.CreateElement)
	handler.ServeHTTP(rr, reqCreate)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}
	//Get the ID of the element just inserted contained in the location of the header
	location := rr.Header().Get("Location")
	index := location[len("/"):]

	// Get the element
	reqGet := MockGetRequest(t, index)
	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(engine.GetElement)
	handler.ServeHTTP(rr, reqGet)
	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	// Check body
	if reflect.DeepEqual(rr.Body.Bytes(), elemCreated)  {
		t.Errorf("Body returned wrong body %v want %v", rr.Body, reqCreate.Body)
	}
}

// Test which mock an http request. Create with POST an element and check with GET
func TestCreateUpdateAndGetItem(t *testing.T) {
	// Post a to-do element to the database
	elemCreated, reqCreate := MockPostRequest(t)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(engine.CreateElement)
	handler.ServeHTTP(rr, reqCreate)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}
	//Get the ID of the element just inserted contained in the location of the header
	location := rr.Header().Get("Location")
	index := location[len("/"):]

	// Update the element just created
	updatedElem, reqUpdate := MockUpdateRequest(t, index)
	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(engine.UpdateElement)
	handler.ServeHTTP(rr, reqUpdate)
	if reflect.DeepEqual(rr.Body.Bytes(), elemCreated)  {
		t.Errorf("Body returned wrong body %v want %v", rr.Body, reqCreate.Body)
	}
	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}

	// Get the element just updated
	reqGet := MockGetRequest(t, index)
	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(engine.GetElement)
	handler.ServeHTTP(rr, reqGet)
	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	// Check body
	if reflect.DeepEqual(rr.Body.Bytes(), updatedElem)  {
		t.Errorf("Body returned wrong body %v want %v", rr.Body, reqCreate.Body)
	}
}

// Test which mock an http request. Create with POST an element and check with GET
func TestCreateDeleteAndGetItem(t *testing.T) {
	// Post a to-do element to the database
	_, reqCreate := MockPostRequest(t)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(engine.CreateElement)
	handler.ServeHTTP(rr, reqCreate)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}
	//Get the ID of the element just inserted contained in the location of the header
	location := rr.Header().Get("Location")
	index := location[len("/"):]

	// Update the element just created
	reqDelete := MockDeleteItemRequest(t, index)
	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(engine.DeleteElement)
	handler.ServeHTTP(rr, reqDelete)
	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Get the element just updated
	reqGet := MockGetRequest(t, index)
	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(engine.GetElement)
	handler.ServeHTTP(rr, reqGet)
	// Check the status code is what we expect. (no found we deleted it previously)
	if status := rr.Code; status != http.StatusNoContent {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNoContent)
	}
}


func TestDeleteCollection(t *testing.T) {
	// First of all delete the database to ensure I POST and PUT the first element
	reqDelete := MockDeleteRequest(t)
	handler := http.HandlerFunc(engine.DeleteCollection)
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, reqDelete)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}