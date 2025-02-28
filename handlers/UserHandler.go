package handlers

import (
	"api/db"
	"api/models"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Metodo no permitido", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(db.Users)
}

func GetQueryParam(w http.ResponseWriter, r *http.Request) {
	age := r.URL.Query().Get("age")

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(models.Message{Text: age})
}

func GetPathParam(w http.ResponseWriter, r *http.Request) {
	age := strings.TrimPrefix(r.URL.Path, "/userP/")
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(models.Message{Text: age})

}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Metodo no permitido", http.StatusMethodNotAllowed)
		return
	}

	defer r.Body.Close() // el defer dice que se va a ejecutar en la ultima linea antes de que se acabe la funcion

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Error al parsear json", http.StatusBadRequest)
		return
	}

	db.Users = append(db.Users, user)

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)

}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Metodo no permitido", http.StatusMethodNotAllowed)
		return
	}
	age, err := strconv.Atoi(r.URL.Query().Get("age"))

	if err != nil {
		return
	}

	for i, user := range db.Users {
		if user.Age == age {
			w.Header().Set("Content-type", "application/json")
			w.WriteHeader(http.StatusNoContent)
			json.NewEncoder(w).Encode(db.Users[i])
			db.Users = append(db.Users[:i], db.Users[i+1:]...)
		}
	}
}

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Metodo no permitido", http.StatusMethodNotAllowed)
		return
	}
	var userUpdated models.User

	json.NewDecoder(r.Body).Decode(&userUpdated)

	for i, user := range db.Users {
		if user.Age == userUpdated.Age {
			db.Users[i] = userUpdated
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(db.Users[i])
			return

		}

	}
	w.WriteHeader(http.StatusNotFound)
}
