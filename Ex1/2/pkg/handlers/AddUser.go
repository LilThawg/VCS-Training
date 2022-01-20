package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/VCS-trainning/Ex1/2/crud/pkg/models"
)

func (h handler) AddUser(w http.ResponseWriter, r *http.Request){
	// Read to request body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var user models.User
	json.Unmarshal(body, &user)

	// Append to the User table
	if result := h.DB.Create(&user); result.Error != nil {
		fmt.Println(result.Error)
	}

	// Send a 201 created response
	w.WriteHeader(http.StatusCreated)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Created")
}