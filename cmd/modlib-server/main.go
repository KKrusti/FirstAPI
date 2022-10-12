package modlib_server

import (
	"apirest/model"
	"apirest/pkg/services/db"
	"apirest/pkg/services/warrior"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type api struct {
	router     http.Handler
	warriorSrv warrior.ServiceInterface
}

type Server interface {
	Router() http.Handler
}

func New() Server {
	repository := db.New()
	a := &api{warriorSrv: warrior.New(repository)}

	r := mux.NewRouter()
	r.HandleFunc("/warriors", a.getWarriors).Methods(http.MethodGet)
	r.HandleFunc("/warrior", a.getWarriorById).Methods(http.MethodGet)
	r.HandleFunc("/warriorsByRace", a.getWarriorsByRace).Methods(http.MethodGet)
	r.HandleFunc("/warrior", a.warrior).Methods(http.MethodPost)
	a.router = r
	return a
}

func (a *api) Router() http.Handler {
	return a.router
}

func (a *api) getWarriorById(writer http.ResponseWriter, request *http.Request) {
	vars := request.URL.Query()["ID"]
	war := a.warriorSrv.FindById(vars[0])

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(war)
}

func (a *api) getWarriors(writer http.ResponseWriter, request *http.Request) {
	//TODO correct treatment of error
	warriors, _ := a.warriorSrv.GetAll()

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(warriors)
}

func (a *api) getWarriorsByRace(writer http.ResponseWriter, request *http.Request) {
	vars := request.URL.Query()["Race"]
	warriors := a.warriorSrv.FindByRace(vars[0])

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(warriors)
}

func (a *api) warrior(writer http.ResponseWriter, request *http.Request) {
	var war model.Warrior
	err := json.NewDecoder(request.Body).Decode(&war)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	a.warriorSrv.Add(war)
	writer.WriteHeader(http.StatusCreated)
}
