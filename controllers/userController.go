package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/daniel-z-johnson/journal-of-self/models"
)

// UserController handles requests
type UserController struct {
	us models.UserService
}

func NewUserController(us models.UserService) *UserController {
	return &UserController{
		us: us,
	}
}

func (uc *UserController) Signup(resp http.ResponseWriter, req *http.Request) {
	// new user
	var u models.User

	err := json.NewDecoder(req.Body).Decode(&u)
	if err != nil {
		http.Error(resp, err.Error(), http.StatusBadRequest)
		return
	}
	user, err := uc.us.Insert(u)
	if err != nil {
		http.Error(resp, "Issue creating user", http.StatusBadRequest)
	}
	fmt.Printf("%+v %+v\n", user, err)
}
