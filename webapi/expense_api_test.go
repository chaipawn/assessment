//go:build integration

package webapi_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/chaipawn/assessment/webapi"
)

type response struct {
	*http.Response
	err error
}

func (r *response) Decode(v interface{}) error {
	if r.err != nil {
		return r.err
	}

	return json.NewDecoder(r.Body).Decode(v)
}

func request(method, url string, body io.Reader) *response {
	req, _ := http.NewRequest(method, url, body)
	req.Header.Add("Content-Type", "application/json")
	client := http.Client{}
	res, err := client.Do(req)
	return &response{res, err}
}

func uri(paths ...string) string {
	host := os.Getenv("APP_URL")
	if host == "" {
		host = "http://localhost:2565"
	}

	if paths == nil {
		return host
	}

	url := append([]string{host}, paths...)
	return strings.Join(url, "/")
}

func TestAPICreateExpense(t *testing.T) {
	body := bytes.NewBufferString(`{
		"title": "strawberry smoothie",
		"amount": 79,
		"note": "night market promotion discount 10 bath", 
		"tags": ["food", "beverage"]
	}`)

	var response webapi.CreateExpenseResponse
	err := request(http.MethodPost, uri("expenses"), body).Decode(&response)
	if err != nil {
		t.Fatal(err)
	}

	if response.Title != "strawberry smoothie" {
		t.Errorf("expense title expect %s, but got %s", "strawberry smoothie", response.Title)
	}

	if response.Amount != float64(79) {
		t.Errorf("expense amount expect %f, but got %f", float64(79), response.Amount)
	}

	if response.Note != "night market promotion discount 10 bath" {
		t.Errorf("expense note expect %s, but got %s", "night market promotion discount 10 bath", response.Note)
	}

	if len(response.Tags) != 2 {
		t.Errorf("expense tags count expect %d, but got %d", 2, len(response.Tags))
	}

	for _, tag := range response.Tags {
		if tag != "food" && tag != "beverage" {
			t.Errorf("expense tags expect %v but got %s", []string{"food", "beverage"}, tag)
		}
	}
}

func TestAPIGetExpense(t *testing.T) {
	body := bytes.NewBufferString(`{
		"title": "strawberry smoothie",
		"amount": 79,
		"note": "night market promotion discount 10 bath", 
		"tags": ["food", "beverage"]
	}`)

	var responseCreated webapi.CreateExpenseResponse
	_ = request(http.MethodPost, uri("expenses"), body).Decode(&responseCreated)
	expectExpenseId := responseCreated.Id

	var response webapi.GetExpenseResponse
	rawResponse := request(http.MethodGet, uri("expenses", strconv.Itoa(expectExpenseId)), bytes.NewBufferString(""))
	err := rawResponse.Decode(&response)
	if err != nil {
		t.Fatal(err)
	}

	if rawResponse.StatusCode != http.StatusOK {
		t.Errorf("response status code expect %d, but got %d", http.StatusOK, rawResponse.StatusCode)
	}

	if response.Title != "strawberry smoothie" {
		t.Errorf("expense title expect %s, but got %s", "strawberry smoothie", response.Title)
	}

	if response.Amount != float64(79) {
		t.Errorf("expense amount expect %f, but got %f", float64(79), response.Amount)
	}

	if response.Note != "night market promotion discount 10 bath" {
		t.Errorf("expense note expect %s, but got %s", "night market promotion discount 10 bath", response.Note)
	}

	if len(response.Tags) != 2 {
		t.Errorf("expense tags count expect %d, but got %d", 2, len(response.Tags))
	}

	for _, tag := range response.Tags {
		if tag != "food" && tag != "beverage" {
			t.Errorf("expense tags expect %v but got %s", []string{"food", "beverage"}, tag)
		}
	}
}

func TestAPIGetExpenseNotFound(t *testing.T) {
	rawResponse := request(http.MethodGet, uri("expenses", "999999999"), bytes.NewBufferString(""))
	if rawResponse.err != nil {
		t.Fatal(rawResponse.err)
	}

	if rawResponse.StatusCode != http.StatusNotFound {
		t.Errorf("response status code expect %d, but got %d", http.StatusNotFound, rawResponse.StatusCode)
	}
}

func TestAPIUpdateExpense(t *testing.T) {
	createExpenseBody := bytes.NewBufferString(`{
		"title": "strawberry smoothie",
		"amount": 79,
		"note": "night market promotion discount 10 bath", 
		"tags": ["food", "beverage"]
	}`)

	var responseCreated webapi.CreateExpenseResponse
	_ = request(http.MethodPost, uri("expenses"), createExpenseBody).Decode(&responseCreated)
	expectExpenseId := responseCreated.Id

	updateExpenseBody := bytes.NewBufferString(`{
		"title": "strawberry smoothie updated",
		"amount": 88,
		"note": "night market promotion discount 10 bath updated", 
		"tags": ["fooded", "beverage", "drink"]
	}`)
	var response webapi.GetExpenseResponse
	rawResponse := request(http.MethodPut, uri("expenses", strconv.Itoa(expectExpenseId)), updateExpenseBody)
	err := rawResponse.Decode(&response)
	if err != nil {
		t.Fatal(err)
	}

	if rawResponse.StatusCode != http.StatusOK {
		t.Errorf("response status code expect %d, but got %d", http.StatusOK, rawResponse.StatusCode)
	}

	if response.Title != "strawberry smoothie updated" {
		t.Errorf("expense title expect %s, but got %s", "strawberry smoothie updated", response.Title)
	}

	if response.Amount != float64(88) {
		t.Errorf("expense amount expect %f, but got %f", float64(88), response.Amount)
	}

	if response.Note != "night market promotion discount 10 bath updated" {
		t.Errorf("expense note expect %s, but got %s", "night market promotion discount 10 bath updated", response.Note)
	}

	if len(response.Tags) != 3 {
		t.Errorf("expense tags count expect %d, but got %d", 3, len(response.Tags))
	}

	for _, tag := range response.Tags {
		if tag != "fooded" && tag != "beverage" && tag != "drink" {
			t.Errorf("expense tags expect %v but got %v", []string{"fooded", "beverage", "drink"}, response.Tags)
		}
	}
}

func TestAPIUpdateExpenseNotFound(t *testing.T) {
	updateExpenseBody := bytes.NewBufferString(`{
		"title": "strawberry smoothie updated",
		"amount": 88,
		"note": "night market promotion discount 10 bath updated", 
		"tags": ["fooded", "beverage", "drink"]
	}`)
	rawResponse := request(http.MethodPut, uri("expenses", "999999999"), updateExpenseBody)
	if rawResponse.err != nil {
		t.Fatal(rawResponse.err)
	}

	if rawResponse.StatusCode != http.StatusNotFound {
		t.Errorf("response status code expect %d, but got %d", http.StatusNotFound, rawResponse.StatusCode)
	}
}

func TestAPIGetAllExpense(t *testing.T) {
	body := bytes.NewBufferString(`{
		"title": "strawberry smoothie",
		"amount": 79,
		"note": "night market promotion discount 10 bath", 
		"tags": ["food", "beverage"]
	}`)

	_ = request(http.MethodPost, uri("expenses"), body)

	var responses []webapi.GetExpenseResponse
	rawResponse := request(http.MethodGet, uri("expenses"), bytes.NewBufferString(""))
	err := rawResponse.Decode(&responses)
	if err != nil {
		t.Fatal(err)
	}

	if rawResponse.StatusCode != http.StatusOK {
		t.Errorf("response status code expect %d, but got %d", http.StatusOK, rawResponse.StatusCode)
	}

	if len(responses) < 1 {
		t.Errorf("expenses count expect greater than 1, but got %d", len(responses))
	}
}
