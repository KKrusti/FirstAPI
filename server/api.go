package server

import (
	gopher "apirest/model"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type api struct {
	router http.Handler
}

type Server interface {
	Router() http.Handler
}

func New() Server {
	a := &api{}

	r := mux.NewRouter()
	r.HandleFunc("/gophers", a.fetchGophersById).Methods(http.MethodGet)
	//r.HandleFunc("/gophers/{ID:[a-zA-Z0-9_]+}", a.fetchGopher).Methods(http.MethodGet)
	a.router = r
	return a
}

func (a *api) Router() http.Handler {
	return a.router
}

func (a *api) fetchGophersById(writer http.ResponseWriter, request *http.Request) {
	vars := request.URL.Query()["ID"]
	gophers := findGopherById(vars[0])

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(gophers)
}

func findGopherById(id string) gopher.Gopher {
	for _, value := range gophers {
		if value.ID == id {
			return value
		}
	}
	return gopher.Gopher{}
}

//func (a *api) fetchGopher(writer http.ResponseWriter, request *http.Request) {
//	vars := mux.Vars(r)
//	gopher, err := a.repository.FetchGopherByID(vars["ID"])
//	w.Header().Set("Content-Type", "application/json")
//	if err != nil {
//		w.WriteHeader(http.StatusNotFound) // We use not found for simplicity
//		json.NewEncoder(w).Encode("Gopher Not found")
//		return
//	}
//	json.NewEncoder(w).Encode(gopher)
//}

var gophers = []gopher.Gopher{
	{ID: "1", Name: "goku"},
	{ID: "2", Name: "vegeta"},
	{ID: "3", Name: "Krilin"},
}
