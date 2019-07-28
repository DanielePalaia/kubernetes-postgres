/* Mock requests for tests */
package testing

import (
	"testing"
	"net/http"
	"bytes"
)

func MockPostRequest(t *testing.T) ([]byte, *http.Request) {
var newElement = []byte(`{
  		"Topic": "New Todo",
  		"Completed": 0,
  		"Due": "2012-11-01T22:08:41+00:00"
	}`)

	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	req, err := http.NewRequest("POST", "/todos", bytes.NewBuffer(newElement))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	return newElement, req
}

// Review this part too much complexity to concatenate string to []byte
func MockUpdateRequest(t *testing.T, index string) ([]byte, *http.Request) {
	// Defining a new element
	newElement := make([]byte, 0)
	var firstPiece = []byte(`{
		"ID":`)

	var thirdPiece = []byte (`,
  		"Topic": "New Todo Updated",
  		"Completed": 1,
  		"Due": "2012-11-01T22:08:41+00:00"
	}`)

	for _, v := range firstPiece{
		newElement = append(newElement, v)
	}
	for _, v := range index{
		newElement = append(newElement, byte(v))
	}
	for _, v := range thirdPiece{
		newElement = append(newElement, v)
	}

	req, err := http.NewRequest("PUT", "/todos", bytes.NewBuffer(newElement))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	return newElement, req
}

func MockGetRequest(t *testing.T, index string) *http.Request  {
	req, err := http.NewRequest("GET", "/todos/" + index, nil)
	if err != nil {
		t.Fatal(err)
	}
	return req
}

func MockDeleteItemRequest(t *testing.T, index string) *http.Request  {
	req, err := http.NewRequest("DELETE", "/todos/" + index, nil)
	if err != nil {
		t.Fatal(err)
	}
	return req
}


func MockDeleteRequest(t *testing.T) *http.Request  {
	req, err := http.NewRequest("DELETE", "/todos", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	return req
}