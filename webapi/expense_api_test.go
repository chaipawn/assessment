//go:build integration

package webapi_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
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
