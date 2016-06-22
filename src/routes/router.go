package routes

import (
	"fmt"
	"net/http"
	"os"
	"log"
	"github.com/gorilla/mux"
)


type Router struct {
	routermultipler *mux.Router
	bindstring string
}



func NewRouter() *Router {
    r := new(Router)
    r.routermultipler = mux.NewRouter().StrictSlash(true)
    r.bindstring = fmt.Sprintf("%s:%s", os.Getenv("OPENSHIFT_GO_IP"), os.Getenv("OPENSHIFT_GO_PORT"))
    return r
}


 func (r *Router) CreateHandler() {
	
	r.routermultipler.HandleFunc("/", notImplemented).Methods("DELETE")
	r.routermultipler.HandleFunc("/", notImplemented).Methods("POST")
	r.routermultipler.HandleFunc("/", notImplemented).Methods("PUT")
	r.routermultipler.HandleFunc("/", getWeather).Methods("GET")
	
	log.Fatal(http.ListenAndServe(":8000", r.routermultipler))
}



func notImplemented(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	w.Header().Set("Content-Type","application/json")
	w.Write([]byte(`{"data": "not Implemented"}`))
	
}

func getWeather(w http.ResponseWriter, r *http.Request) {
	
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type","application/json")
	fmt.Fprintln(w, "Hello there!!")
}



