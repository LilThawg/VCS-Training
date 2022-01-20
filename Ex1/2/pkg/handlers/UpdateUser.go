package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/VCS-trainning/Ex1/2/crud/pkg/models"
	"github.com/gorilla/mux"
)

func (h handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	// Read dynamic id parameter
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// Read Request body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var updatedUser models.User
	json.Unmarshal(body, &updatedUser)

	var user models.User

	if result := h.DB.First(&user, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	user.Id = updatedUser.Id
	user.Name = updatedUser.Name
	user.Dob = updatedUser.Dob
	user.Gender = updatedUser.Gender
	user.Email = updatedUser.Email

	h.DB.Save(&user)

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Updated")
}