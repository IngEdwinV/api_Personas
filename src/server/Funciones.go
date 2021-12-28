/*
Paquete personalizado para generacion de endpoints y calculos
*/
package server

//Dependencias, librerias necesarias para la configuración
import (
	"math"
)

//Definicion de estrucuturas
type data struct {
	Satellites []satelite `json:"satellites"`
}

type satelite struct {
	Name     string   `json:"Name"`
	Distance float32  `json:"Distance"`
	Message  []string `json:"Message"`
	CoorX    float32  `json:"CoorX"`
	Coory    float32  `json:"Coory"`
}

type allsatelite []satelite

var satelites = allsatelite{}

// input: distancia al emisor tal cual se recibe en cada satélite
// output: las coordenadas ‘x’ e ‘y’ del emisor del mensaje
func GetLocation(distances ...float32) (x, y float32) {

	var R1 = distances[1]
	var R2 = distances[2]
	var d = float32(100)
	var i = float32(500)
	var j = float32(100)

	x = (float32(math.Pow(float64(R1), 2)) - float32(math.Pow(float64(R2), 2)) + float32(math.Pow(float64(d), 2))) / (2 * d)

	y = ((float32(math.Pow(float64(R1), 2))-float32(math.Pow(float64(R2), 2))+float32(math.Pow(float64(i), 2))+float32(math.Pow(float64(j), 2)))/(2*j) - (i/j)*x)

	return x, y
}

// input: el mensaje tal cual es recibido en cada satélite
// output: el mensaje tal cual lo genera el emisor del mensaje
func GetMessage(messages ...[]string) (msg string) {
	var mensajeFinal [225]string

	for i := 0; i < len(messages); i++ {
		for j := 0; j < len(messages[i]); j++ {
			var dato = messages[i][j]
			if dato != "" {
				mensajeFinal[j] = dato
			}
		}
	}

	for t := 0; t < len(mensajeFinal); t++ {
		if mensajeFinal[t] != "" {
			msg = msg + " " + mensajeFinal[t]
		}

	}
	return msg
}
