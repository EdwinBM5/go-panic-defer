package main

import (
	"fmt"
	"io"
	"os"
)

/*
Un estudio contable necesita acceder a los datos de sus empleados para poder realizar distintas liquidaciones.
Para ello, cuentan con todo el detalle necesario en un archivo .txt.

Tendrás que desarrollar la funcionalidad para poder leer el archivo .txt que nos indica el cliente,
sin embargo, no han pasado el archivo a leer por nuestro programa.
Desarrollá el código necesario para leer los datos del archivo llamado “customers.txt”
(recordá lo visto sobre el pkg “os”).
Dado que no contamos con el archivo necesario, se obtendrá un error y, en tal caso,
el programa deberá arrojar un panic al intentar leer un archivo que no existe, mostrando el mensaje
“The indicated file was not found or is damaged”.

Sin perjuicio de ello, deberá siempre imprimirse por consola “ejecución finalizada”.
*/
func main() {
	ReadUserFile("customers.txt")
}

func ReadUserFile(filepath string) {
	file, err := os.Open(filepath)

	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
			fmt.Println("Ejecución finalizada")
		}
	}()

	defer file.Close()

	if err != nil {
		panic("The indicated file was not found or is damaged")
	}

	_, err = io.ReadAll(file)
	if err != nil {
		panic(err)
	}

}

/*
A continuación, vamos a crear un archivo “customers.txt” con información de los clientes del estudio.

Ahora que el archivo sí existe, el panic no debe ser lanzado.

Creamos el archivo “customers.txt” y le agregamos la información de los clientes.
Extendemos el código del punto uno para que podamos leer este archivo e imprimir los datos que contenga.
En el caso de no poder leerlo, se debe lanzar un “panic”.
Recordemos que siempre que termina la ejecución, independientemente del resultado,
debemos tener un “defer” que nos indique que la ejecución finalizó. También recordemos cerrar los archivos al finalizar su uso.
*/
