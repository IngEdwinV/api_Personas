/*
Paquete Principal, inicia el servidor Http
*/
package main

//Dependencias, librerias necesarias para la configuración
import (
	"log"
	"net/http"

	"github.com/IngEdwinV/api_Personas/src/server"
)

func main() {
	s := server.New()
	log.Fatal(http.ListenAndServe(":3000", s.Router()))
}
