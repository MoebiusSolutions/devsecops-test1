package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	myhndlr "devsecops-test1/app/services/helloworld-api/handlers"
)

func TestMain(t *testing.T) {
	t.Parallel()

	t.Run("greet", greet)
	t.Run("greetUser", greetUser)
}

func greet(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	h := http.HandlerFunc(myhndlr.Greet)

	h.ServeHTTP(w, r)

	// check HTTP response status code
	wantCode := http.StatusOK
	if s := w.Code; s != wantCode {
		t.Errorf("handler return wrong status code: got %v, want %v", s, wantCode)
	}

	// check HTTP response header content type
	wantContentType := "text/plain; charset=utf-8"
	if c := w.Header().Get("Content-type"); c != wantContentType {
		t.Errorf("handler return wrong status code: got %v, want %v", c, wantContentType)
	}

	// check HTTP response body
	wantBody := "Hello World! " + time.Now().UTC().Format("01-02-2006")
	if w.Body.String() != wantBody {
		t.Errorf("handler return wrong response body: got %v, want %v", w.Body.String(), wantBody)
	}
}

func greetUser(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/user", nil)
	w := httptest.NewRecorder()
	h := http.HandlerFunc(myhndlr.GreetUser)

	h.ServeHTTP(w, r)

	// check HTTP response status code
	wantCode := http.StatusOK
	if s := w.Code; s != wantCode {
		t.Errorf("handler return wrong status code: got %v, want %v", s, wantCode)
	}

	// check HTTP response header content type
	wantContentType := "text/plain; charset=utf-8"
	if c := w.Header().Get("Content-type"); c != wantContentType {
		t.Errorf("handler return wrong status code: got %v, want %v", c, wantContentType)
	}

	// check HTTP response body
	wantBody := "Hello World! Jonathan"
	if w.Body.String() != wantBody {
		t.Errorf("handler return wrong response body: got %v, want %v", w.Body.String(), wantBody)
	}
}
