package test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"unicode/utf8"

	"github.com/dixxe/personal-website/iternal/controllers"
)

func TestPostMessage(t *testing.T) {

	test_value := `{"content":"banan","msg_type":10}`

	// Creating a body for testing.
	body := strings.NewReader(test_value)
	// Creating request for controler.
	req, err := http.NewRequest("POST", "/cctweaked", body)
	if err != nil {
		t.Fatal(err)
	}
	// Creating recored to read output from controller.
	rr := httptest.NewRecorder()

	// Pushing test data to controller
	http.HandlerFunc(controllers.PostSendMessage).ServeHTTP(rr, req)

	// Creating a pull request for controller
	req, err = http.NewRequest("GET", "/cctweaked", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Pulling processed data from controller
	http.HandlerFunc(controllers.GetMessage).ServeHTTP(rr, req)

	expected := `{banan 10}`

	if strings.Trim(rr.Body.String(), "\n ") != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}

}

func FuzzPostMessage(f *testing.F) {
	f.Add(0, "banan")
	f.Fuzz(func(t *testing.T, msg_type int, content string) {

		if !utf8.Valid([]byte(content)) {
			return
		}

		body := strings.NewReader(fmt.Sprintf(`{"content":"%v","msg_type":%v}`,
			content, msg_type))

		req, err := http.NewRequest("POST", "/cctweaked", body)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()

		http.HandlerFunc(controllers.PostSendMessage).ServeHTTP(rr, req)

		req, err = http.NewRequest("GET", "/cctweaked", nil)
		if err != nil {
			t.Fatal(err)
		}
		http.HandlerFunc(controllers.GetMessage).ServeHTTP(rr, req)

		expected := fmt.Sprintf(`{%v %v}`, content, msg_type)

		if rr.Result().StatusCode == 400 {
			return
		}

		if strings.Trim(rr.Body.String(), "\n ") != expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), expected)
		}
	})
}
