package auth

import (
	helpers "Auth-Service/app/response"
	"Auth-Service/interfaces"
	"encoding/json"
	"net/http"
)

type CreateUser struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreatedUserResponse struct {
	User  interfaces.User `json:"user"`
	Token string          `json:"token"`
}

func CreateUserService(r *http.Request, repo interfaces.AuthRepository) helpers.ApiResponse[CreatedUserResponse] {
	var user CreateUser
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		return *helpers.StandardApiResponse(false, http.StatusBadRequest, "Invalid request body", CreatedUserResponse{})
	}
	user = CreateUser{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
	createdUser, token, err := repo.CreateUser(user.Name, user.Email, user.Password)
	if err != nil {
		return *helpers.StandardApiResponse(false, http.StatusInternalServerError, err.Error(), CreatedUserResponse{})
	}
	createdUserResponse := CreatedUserResponse{
		User:  createdUser,
		Token: token,
	}
	return *helpers.StandardApiResponse(true, http.StatusCreated, "User created successfully", createdUserResponse)
}

func AuthLoginService(r *http.Request, repo interfaces.AuthRepository) helpers.ApiResponse[Login] {
	var loginUser Login
	json.NewDecoder(r.Body).Decode(&loginUser)
	_, err := repo.LoginUser(loginUser.Email, loginUser.Password)

	if err != nil {
		return *helpers.StandardApiResponse(false, http.StatusInternalServerError, "Login Failed", Login{})
	}
	return *helpers.StandardApiResponse(true, http.StatusOK, "Login successful", loginUser)
}
