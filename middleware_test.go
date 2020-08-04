package middleware

import (
	"fmt"
	"github.com/google/uuid"
	"io/ioutil"
	"net/http"
	"testing"
)

var uuidgen = uuid.New().String()

type RequestIdHandler struct {}
func (h RequestIdHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.Header.Add("request-id", uuidgen)
}

type LogHandler struct {}
func (h LogHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Request=%v\n", r)
	defer fmt.Printf("Response=%v\n", w)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(r.Header.Get("request-id")))
}

func TestJoinHandlers(t *testing.T) {
	port := ":3000"
	go http.ListenAndServe(port, JoinHandlers(RequestIdHandler{}, LogHandler{}))
	client := http.Client{}
	response, err := client.Get(fmt.Sprintf("http://localhost%s", port))
	if err != nil {
		t.Errorf("unable to get response %s", err)
	}
	actualBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Errorf("unable to read response %s", err)
	}
	defer response.Body.Close()
	actualResponse := string(actualBytes)
	if actualResponse != uuidgen {
		t.Errorf("actualResponse=%s, expectedResponse=%s", actualResponse, uuidgen)
	}
}