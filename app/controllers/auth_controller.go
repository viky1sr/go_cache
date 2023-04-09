package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/viky1sr/go_cache.git/app/services"
	"github.com/viky1sr/go_cache.git/app/traits"
)

type AuthController struct {
	authService services.AuthService
}

func NewAuthController(authService services.AuthService) *AuthController {
	return &AuthController{authService}
}

func (c *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	// Parse request body
	var reqBody struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		responseTrait := traits.ResponseTrait{}
		responseTrait.RespondWithFailure(w, http.StatusBadRequest, "Invalid Request Body")
		return
	}

	// Authenticate user
	token, err := c.authService.Login(reqBody.Email, reqBody.Password)
	if err != nil {
		responseTrait := traits.ResponseTrait{}
		responseTrait.RespondWithFailure(w, http.StatusUnauthorized, err.Error())
		return
	}

	fmt.Print(token, err)

	// Return token
	responseTrait := traits.ResponseTrait{}
	responseTrait.RespondWithSuccess(w, http.StatusOK, "Success login", token)
}

func (c *AuthController) Logout(w http.ResponseWriter, r *http.Request) {
	// Parse request body
	var reqBody struct {
		Token string `json:"token"`
	}
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		responseTrait := traits.ResponseTrait{}
		responseTrait.RespondWithFailure(w, http.StatusBadRequest, "Invalid Request Body")
		return
	}

	// Invalidate token
	if err := c.authService.Logout(reqBody.Token); err != nil {
		responseTrait := traits.ResponseTrait{}
		responseTrait.RespondWithFailure(w, http.StatusInternalServerError, "Failed to logout")
		return
	}

	// Return success response
	responseTrait := traits.ResponseTrait{}
	responseTrait.RespondWithSuccess(w, http.StatusOK, "Logged out successfully", "")
}
