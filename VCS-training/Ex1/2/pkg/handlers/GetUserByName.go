package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/VCS-trainning/Ex1/2/crud/pkg/models"
	"github.com/gorilla/mux"
)

func (h handler) GetUserByName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	//find user by name
	var users []models.User
	if result := h.DB.Where("name = ?",name).Find(&users); result.Error != nil {
		fmt.Println(result.Error)
	} 

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
