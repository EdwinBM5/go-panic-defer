package main

import (
	"errors"
	"fmt"
)

/*
El mismo estudio del ejercicio anterior, solicita una funcionalidad para poder registrar datos de nuevos clientes.
Los datos requeridos son:

File
Name
ID
Phone number
Home

Tarea 1: Antes de registrar a un cliente, debés verificar si el mismo ya existe.
Para ello, necesitás leer los datos de un array. En caso de que esté repetido,
debes manipular adecuadamente el error como hemos visto hasta aquí. Ese error deberá:

1.- Generar un panic;

2.- lanzar por consola el mensaje: “Error: client already exists”, y continuar con la ejecución del programa normalmente.

Tarea 2: Luego de intentar verificar si el cliente a registrar ya existe,
desarrollá una función para validar que todos los datos a registrar de un cliente contienen
un valor distinto de cero. Esta función debe retornar, al menos, dos valores.
Uno de ellos tendrá que ser del tipo error para el caso de que se ingrese por parámetro algún valor cero
(recordá los valores cero de cada tipo de dato, ej: 0, “”, nil).

Tarea 3: Antes de finalizar la ejecución, incluso si surgen panics, se deberán imprimir por
consola los siguientes mensajes: “End of execution” y “Several errors were detected at runtime”.
Utilizá defer para cumplir con este requerimiento.

Requerimientos generales:

Utilizá recover para recuperar el valor de los panics que puedan surgir
Recordá realizar las validaciones necesarias para cada retorno que pueda contener un valor error.
Generá algún error, personalizandolo a tu gusto utilizando alguna de las funciones de Go (realiza también la validación pertinente para el caso de error retornado).
*/

type ClientInfo struct {
	File        string
	Name        string
	ID          int
	PhoneNumber int
	Home        string
}

func (c ClientInfo) CheckClientInfo() (bool, error) {
	if c.File == "" || c.Name == "" || c.ID == 0 || c.PhoneNumber == 0 || c.Home == "" {
		return false, errors.New("Error: client fields are empty")

	}

	return true, nil
}

type ClientList struct {
	Clients []ClientInfo
}

func (c ClientList) CheckClient(client ClientInfo) {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println(r)
		}
	}()

	for _, c := range c.Clients {
		if c.ID == client.ID {
			panic("Error: client already exists")
		}
	}
}

var clients = ClientList{
	Clients: []ClientInfo{
		{
			File:        "juan.txt",
			Name:        "Juan",
			ID:          1,
			PhoneNumber: 123456789,
			Home:        "Calle 123",
		},
		{
			File:        "rodrigo.txt",
			Name:        "Rodrigo",
			ID:          2,
			PhoneNumber: 987654321,
			Home:        "Calle 456",
		},
		{
			File:        "maria.txt",
			Name:        "Maria",
			ID:          3,
			PhoneNumber: 123456789,
			Home:        "Calle 789",
		},
	},
}

func main() {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println(r)
		}

		fmt.Println("End of execution")
		fmt.Println("Several errors were detected at runtime")
	}()

	var client = ClientInfo{
		File:        "dsadsa",
		Name:        "Juan",
		ID:          1,
		PhoneNumber: 123456789,
		Home:        "Calle 123",
	}

	clients.CheckClient(client) // Etapa 1

	if isValidClient, err := client.CheckClientInfo(); !isValidClient { // Etapa 2
		fmt.Println(err)
		return
	}

	fmt.Println("Client is valid.")

}
