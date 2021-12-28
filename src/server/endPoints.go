/*
Paquete personalizado para generacion de endpoints y calculos
*/
package server

//Dependencias, librerias necesarias para la configuración
import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

//Definicion de estrucuturas
type api struct {
	router http.Handler
}

type Server interface {
	Router() http.Handler
}

type response struct {
	Name     string `json:"Name"`
	LastName string `json:"LastName"`
	DNI      string `json:"DNI"`
	Empleado bool   `json:"Empleado"`
}

//Declaracion de los endpoints y funciones
func New() Server {
	a := &api{}

	r := mux.NewRouter()

	r.HandleFunc("/createPerson", a.createPerson).Methods(http.MethodPost)
	r.HandleFunc("/selectPerson/{Person_id}", a.selectPerson).Methods(http.MethodGet)
	r.HandleFunc("/getPersons", a.getPersons).Methods(http.MethodGet)
	r.HandleFunc("/deletePerson/{Person_id}", a.deletePerson).Methods(http.MethodDelete)
	r.HandleFunc("/updatePerson/{Person_id}", a.updatePerson).Methods(http.MethodPut)

	a.router = r
	return a
}

func (a *api) Router() http.Handler {
	return a.router
}

//endpoint: http://server_ip/createPerson
//metodo: Post
//Descripcion: funcion para leer json con los datos enviados para crear una persona
//Input: json
//Output: Code HTTP
func (a *api) createPerson(w http.ResponseWriter, r *http.Request) {

	var newData data

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}

	json.Unmarshal(reqBody, &newData)

	for i := 0; i < len(newData.Persons); i++ {

		persons = append(persons, newData.Persons[i])

	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
}

//endpoint: http://server_ip/selectPerson/person_id
//metodo: Get
//Descripcion: Funcion que retorna un json con la informacion de una persona
//input: Parametro "person_id"
//output: Json
func (a *api) selectPerson(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	person_id := vars["Person_id"]

	if person_id == "" {
		w.WriteHeader(http.StatusBadRequest)
	}

	name, lasName, DNI, empleado := getPerson(person_id)

	var res response
	res.Name = name
	res.LastName = lasName
	res.DNI = DNI
	res.Empleado = empleado

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)

}

//endpoint: http://server_ip/getPersons
//metodo: Get
//Descripcion: Funcion que retorna todos las personas creadas
//Output: json
func (a *api) getPersons(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(persons)
}

//endpoint: http://server_ip/deletePerson/person_id
//metodo: Delete
//Descripcion: Funcion que da de baja a una persona en el sistema
//Input: parametro person_id
//Output: code HTTP
func (a *api) deletePerson(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	person_id := vars["Person_id"]

	if person_id == "" {
		w.WriteHeader(http.StatusBadRequest)
	}

	res := deletePerson(person_id)

	if res {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}

}

//endpoint: http://server_ip/updatePerson/person_id
//metodo: PUT
//Descripcion: Funcion que actualiza una persona en el sistema
//Input: person_id, json
//Output: json
func (a *api) updatePerson(w http.ResponseWriter, r *http.Request) {

	var data person
	vars := mux.Vars(r)

	person_id := vars["Person_id"]

	if person_id == "" {
		w.WriteHeader(http.StatusBadRequest)
	}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}

	json.Unmarshal(reqBody, &data)

	res := updatePerson(person_id, data)

	if res {
		name, lasName, DNI, empleado := getPerson(person_id)

		var res response
		res.Name = name
		res.LastName = lasName
		res.DNI = DNI
		res.Empleado = empleado

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(res)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}
