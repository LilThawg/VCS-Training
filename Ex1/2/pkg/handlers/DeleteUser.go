package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/VCS-trainning/Ex1/2/crud/pkg/models"
	"github.com/gorilla/mux"
)

func (h handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	// Read the dynamic id parameter
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// find user by id
	var user models.User
	if result := h.DB.First(&user, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	// delete that user
	h.DB.Delete(&user)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Deleted")
}