package controllers

import (
	"encoding/json"
	"github.com/viky1sr/go_cache.git/app/traits"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/viky1sr/go_cache.git/app/models"
	"github.com/viky1sr/go_cache.git/app/services"
)

var responseTrait = traits.ResponseTrait{}

type UserController struct {
	userService services.UserService
}

func NewUserController(userService services.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (c *UserController) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := c.userService.GetAllUsers()
	if err != nil {
		responseTrait.RespondWithFailure(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func (c *UserController) GetUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		responseTrait.RespondWithFailure(w, http.StatusBadRequest, err.Error())
		return
	}

	user, err := c.userService.GetUserByID(id)
	if err != nil {
		responseTrait.RespondWithFailure(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (c *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		responseTrait := traits.ResponseTrait{}
		responseTrait.RespondWithFailure(w, http.StatusBadRequest, "Body cant be null")
		return
	}
	err = c.userService.CreateUser(&user)
	if err != nil {
		responseTrait.RespondWithFailure(w, http.StatusBadRequest, err.Error())
		return
	}
	resultUser := map[string]interface{}{
		"name":  user.Name,
		"email": user.Email,
	}
	responseTrait.RespondWithSuccess(w, http.StatusCreated, "Success create user", resultUser)
	return
}

func (c *UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		responseTrait.RespondWithFailure(w, http.StatusBadRequest, err.Error())
		return
	}

	var user models.User
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		responseTrait.RespondWithFailure(w, http.StatusBadRequest, err.Error())
		return
	}

	user.ID = uint(id)

	err = c.userService.UpdateUser(id, &user)
	if err != nil {
		responseTrait.RespondWithFailure(w, http.StatusInternalServerError, err.Error())
		return
	}

	resultUser := map[string]interface{}{
		"name":  user.Name,
		"email": user.Email,
	}
	responseTrait.RespondWithSuccess(w, http.StatusCreated, "Success update user", resultUser)
	return

}

func (c *UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		responseTrait.RespondWithFailure(w, http.StatusBadRequest, err.Error())
		return
	}

	err = c.userService.DeleteUser(id)
	if err != nil {
		responseTrait.RespondWithFailure(w, http.StatusInternalServerError, err.Error())
		return
	}

	responseTrait.RespondWithSuccess(w, http.StatusNoContent, "Success delete user", "")
	return
}
