package main

/*
flujo de la app

1. se ejecuta la funcion main():
registra la ruta "/" y el handler.

2. cuando alguien ingresa por el puerto 8080, se
ejecuta el handler asignado a la ruta

3. la funcion asigna el content-type por el header y
convierte el Response en json

4. lo envia

*/

import (
	"api/handlers"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlers.GetUsersHandler)
	http.HandleFunc("/userQ", handlers.GetQueryParam) // query params
	http.HandleFunc("/userP/", handlers.GetPathParam) // path params
	http.HandleFunc("/createUser/", handlers.CreateUserHandler)
	http.HandleFunc("/deleteUser", handlers.DeleteUserHandler) // query params

	// indica a go que va a escuchar y servir por el puerto 8080, nil significa que usamos el manejador por defecto de go
	port := 8080
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
