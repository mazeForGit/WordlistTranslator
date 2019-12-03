package routers

import (
	"testing"
	"net/http/httptest"
	"net/http"
	"github.com/gin-gonic/gin"
	"encoding/json"
	data "github.com/mazeForGit/WordlistExtractor/data"
	"bytes"
)

func Test_WordsGET_emptyData(t *testing.T) {
	router := gin.Default()
	router.GET("/words", WordsGET)

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/words?format=csv", nil)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("You received a %v error.", w.Code)
	}

	resstring := string(w.Body.Bytes())

	if len(resstring) != 0 {
		t.Errorf("expected empty response string")
	}
}
func Test_WordsGET_wrongFormat(t *testing.T) {
	router := gin.Default()
	router.GET("/words", WordsGET)

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/words?format=wrong", nil)
	router.ServeHTTP(w, req)

	var s data.Status
	json.Unmarshal(w.Body.Bytes(), &s)

	if s.Code != 422 && s.Text != "unknown format = wrong" {
		t.Errorf("got wrong response .. code=%v, text=%s", s.Code, s.Text)
	}
}
func Test_WordsPOST_validData(t *testing.T) {
	router := gin.Default()
	router.POST("/words", WordsPOST)

	wrd := data.Word{Id: 0, Name: "test", New: true, Occurance: 0, Tests: nil}
	payload, err := json.Marshal(wrd)

	if err != nil {
		w := httptest.NewRecorder()
		req := httptest.NewRequest("POST", "/words", bytes.NewBuffer(payload))
		router.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Fatalf("You received a %v error.", w.Code)
		}

		var s data.Status
		json.Unmarshal(w.Body.Bytes(), &s)

		if s.Code != 200 && s.Text != "entity added" {
			t.Errorf("got wrong response .. code=%v, text=%s", s.Code, s.Text)
		}
	}
}
func Test_WordsDELETE_validData(t *testing.T) {
	router := gin.Default()
	router.POST("/words", WordsPOST)
	router.DELETE("/words", WordsByIdDELETE)

	wrd := data.Word{Id: 0, Name: "test", New: true, Occurance: 0, Tests: nil}
	payload, err := json.Marshal(wrd)

	if err != nil {
		w := httptest.NewRecorder()
		req := httptest.NewRequest("POST", "/words", bytes.NewBuffer(payload))
		router.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Fatalf("You received a %v error.", w.Code)
		}

		var s data.Status
		json.Unmarshal(w.Body.Bytes(), &s)

		if s.Code != 200 && s.Text != "entity added" {
			t.Errorf("got wrong response .. code=%v, text=%s", s.Code, s.Text)
		}
	}
}
func Test_WordsPOST_invalidData(t *testing.T) {
	router := gin.Default()
	router.POST("/words", WordsPOST)

	wrd := data.Word{Id: 0, Name: "test", New: false, Occurance: 0, Tests: nil}
	payload, err := json.Marshal(wrd)

	if err != nil {
		w := httptest.NewRecorder()
		req := httptest.NewRequest("POST", "/words", bytes.NewBuffer(payload))
		router.ServeHTTP(w, req)

		var s data.Status
		json.Unmarshal(w.Body.Bytes(), &s)

		if s.Code != 422 && s.Text != "unprocessable entity" {
			t.Errorf("got wrong response .. code=%v, text=%s", s.Code, s.Text)
		}
	}
}
