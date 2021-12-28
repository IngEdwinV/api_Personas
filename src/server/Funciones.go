/*
Paquete personalizado para generacion de endpoints y calculos
*/
package server

//Dependencias, librerias necesarias para la configuración

//Definicion de estrucuturas
type data struct {
	Persons []person `json:"Persons"`
}

type person struct {
	Name     string `json:"Name"`
	LastName string `json:"LastName"`
	DNI      string `json:"DNI"`
	Empleado bool   `json:"Empleado"`
}

type allsPerson []person

var persons = allsPerson{}

// input:El codido identificativo de la persona (DNI)
// output: Datos de la persona
func getPerson(person_id string) (name, lasName, DNI string, empleado bool) {

	for _, person := range persons {
		if person.DNI == person_id {
			return person.Name, person.LastName, person.DNI, person.Empleado
		}
	}
	return "Null", "Null", "Null", false
}

// input:El codido identificativo de la persona (DNI) y la informacion de la persona
// output: Respuesta true o false que indica si se ejecuto correctamente el proceso
func updatePerson(person_id string, person person) (res bool) {

	for i, t := range persons {
		if t.DNI == person_id {
			persons = append(persons[:i], persons[i+1:]...)

			person.DNI = t.DNI
			persons = append(persons, person)
			return true

		}
	}

	return false
}

// input:El codido identificativo de la persona (DNI)
// output: Respuesta true o false que indica si se ejecuto correctamente el proceso
func deletePerson(person_id string) (res bool) {

	for i, t := range persons {
		if t.DNI == person_id {
			persons = append(persons[:i], persons[i+1:]...)
			return true
		}
	}

	return false
}
