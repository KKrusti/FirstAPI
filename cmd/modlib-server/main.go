package modlib_server

import (
	"apirest/pkg/services/warriorService"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type api struct {
	router     http.Handler
	warriorSrv warriorService.WarriorService
}

type Server interface {
	Router() http.Handler
}

func New() Server {
	a := &api{warriorSrv: warriorService.NewWarriorService()}

	r := mux.NewRouter()
	r.HandleFunc("/warriors", a.getWarriors).Methods(http.MethodGet)
	r.HandleFunc("/warrior", a.getWarriorById).Methods(http.MethodGet)
	r.HandleFunc("/warriorsByRace", a.getWarriorsByRace).Methods(http.MethodGet)
	a.router = r
	return a
}

func (a *api) Router() http.Handler {
	return a.router
}

func (a *api) getWarriorById(writer http.ResponseWriter, request *http.Request) {
	vars := request.URL.Query()["ID"]
	warrior := a.warriorSrv.FindById(vars[0])

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(warrior)
}

func (a *api) getWarriors(writer http.ResponseWriter, request *http.Request) {
	warriors := a.warriorSrv.FindAll()

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(warriors)
}

func (a *api) getWarriorsByRace(writer http.ResponseWriter, request *http.Request) {
	vars := request.URL.Query()["Race"]
	warriors := a.warriorSrv.FindByRace(vars[0])

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(warriors)
}
