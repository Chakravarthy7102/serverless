package http

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	Status int         `json:"status"`
	Result interface{} `json:"result"`
}

func newResponse(data interface{}, status int) *Response {
	return &Response{
		Status: status,
		Result: data,
	}
}

func (res *Response) bytes() []byte {
	data, _ := json.Marshal(res)
	return data
}

func (res *Response) string() string {
	return string(res.bytes())
}

func (res *Response) sendResponse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(res.Status)
	_, _ = w.Write(res.bytes())
	log.Println(res.string())
}

//200
func StatusOK(w http.ResponseWriter, r *http.Request, data interface{}) {
	newResponse(data, http.StatusOK).sendResponse(w, r)
}

//204
func StatusNoContent(w http.ResponseWriter, r *http.Request) {
	newResponse(nil, http.StatusNoContent).sendResponse(w, r)
}

//404
func StatusNotFound(w http.ResponseWriter, r *http.Request, err error) {

	data := map[string]interface{}{"error": err.Error()}
	newResponse(data, http.StatusNotFound).sendResponse(w, r)
}

//400
func StatusBadRequest(w http.ResponseWriter, r *http.Request, err error) {
	data := map[string]interface{}{
		"error": err.Error(),
	}
	newResponse(data, http.StatusBadRequest).sendResponse(w, r)
}

//405
func StatusMethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	newResponse(nil, http.StatusMethodNotAllowed).sendResponse(w, r)
}

//409
func StatusConflict(w http.ResponseWriter, r *http.Request, err error) {
	data := map[string]interface{}{"error": err.Error()}
	newResponse(data, http.StatusConflict).sendResponse(w, r)
}

//500
func StatusInternalServerError(w http.ResponseWriter, r *http.Request, err error) {
	data := map[string]interface{}{"error": err.Error()}
	newResponse(data, http.StatusInternalServerError).sendResponse(w, r)
}
