/*
Paquete personalizado para generacion de endpoints y calculos
*/
package server

//Dependencias, librerias necesarias para la configuración
import (
	"encoding/json"
	"fmt"
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
	Position posicion `json:"Position"`
	Message  string   `json:"Message"`
}

type posicion struct {
	Coorx float32 `json:"x"`
	Coory float32 `json:"y"`
}

//Declaracion de los endpoints y funciones
func New() Server {
	a := &api{}

	r := mux.NewRouter()

	r.HandleFunc("/topsecret", a.topsecret).Methods(http.MethodPost)
	r.HandleFunc("/topsecret_split/{satellite_name}", a.topsecret_split).Methods(http.MethodPost)
	r.HandleFunc("/topsecret_split", a.topsecret_split_Get).Methods(http.MethodGet)
	r.HandleFunc("/topsecret_split", a.topsecret_split_Post).Methods(http.MethodPost)

	a.router = r
	return a
}

func (a *api) Router() http.Handler {
	return a.router
}

//endpoint: http://server_ip/topsecret
//metodo: Post
//Descripcion: funcion para leer json con los datos enviados por los tres satelites
//Input: json
//Output: json
func (a *api) topsecret(w http.ResponseWriter, r *http.Request) {

	var newData data
	var distances []float32
	var message [][]string

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}

	json.Unmarshal(reqBody, &newData)

	for i := 0; i < len(newData.Satellites); i++ {

		if newData.Satellites[i].Name == "kenobi" {

			newData.Satellites[i].CoorX = -500
			newData.Satellites[i].Coory = -200
			satelites = append(satelites, newData.Satellites[i])

			distances = append(distances, newData.Satellites[i].Distance)
			message = append(message, newData.Satellites[i].Message)

		} else if newData.Satellites[i].Name == "skywalker" {

			newData.Satellites[i].CoorX = 100
			newData.Satellites[i].Coory = -100
			satelites = append(satelites, newData.Satellites[i])

			distances = append(distances, newData.Satellites[i].Distance)
			message = append(message, newData.Satellites[i].Message)

		} else if newData.Satellites[i].Name == "sato" {

			newData.Satellites[i].CoorX = 500
			newData.Satellites[i].Coory = 100
			satelites = append(satelites, newData.Satellites[i])

			distances = append(distances, newData.Satellites[i].Distance)
			message = append(message, newData.Satellites[i].Message)

		}

	}

	msg := GetMessage(message...)
	x, y := GetLocation(distances...)

	var res response
	res.Position.Coorx = x
	res.Position.Coory = y
	res.Message = msg

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

//endpoint: http://server_ip/topsecret_split/satellite_name
//metodo: Post
//Descripcion: funcion para leer los datos enviados de un satelite
//input: Parametro "satellite_name" y json
func (a *api) topsecret_split(w http.ResponseWriter, r *http.Request) {

	var data satelite
	vars := mux.Vars(r)

	sateliteName := vars["satellite_name"]

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}

	json.Unmarshal(reqBody, &data)

	data.Name = sateliteName

	if sateliteName == "kenobi" {
		data.CoorX = -500
		data.Coory = -200
		satelites = append(satelites, data)
	} else if sateliteName == "skywalker" {
		data.CoorX = 100
		data.Coory = -100
		satelites = append(satelites, data)
	} else if sateliteName == "sato" {
		data.CoorX = 500
		data.Coory = 100
		satelites = append(satelites, data)
	}

	w.WriteHeader(http.StatusAccepted)
}

//endpoint: http://server_ip/topsecret_split
//metodo: Get
//Descripcion: funcion para retornar calculo de posicion y mensaje oculto
//Output: json
func (a *api) topsecret_split_Get(w http.ResponseWriter, r *http.Request) {

	var distances []float32
	var message [][]string

	for i, t := range satelites {
		if t.Distance >= 0 {
			fmt.Print(i)
			distances = append(distances, t.Distance)
			message = append(message, t.Message)
		}

	}

	if len(distances) >= 2 {
		msg := GetMessage(message...)
		x, y := GetLocation(distances...)

		var res response
		res.Position.Coorx = x
		res.Position.Coory = y
		res.Message = msg

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(res)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

//endpoint: http://server_ip/topsecret_split
//metodo: Post
//Descripcion: funcion que recibe los datos de una nave
//Input: json
//Output: json
func (a *api) topsecret_split_Post(w http.ResponseWriter, r *http.Request) {
	var data satelite
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}

	json.Unmarshal(reqBody, &data)
	if data.Name == "kenobi" {
		data.CoorX = -500
		data.Coory = -200
		satelites = append(satelites, data)
	} else if data.Name == "skywalker" {
		data.CoorX = 100
		data.Coory = -100
		satelites = append(satelites, data)
	} else if data.Name == "sato" {
		data.CoorX = 500
		data.Coory = 100
		satelites = append(satelites, data)
	}

	w.WriteHeader(http.StatusAccepted)
}
