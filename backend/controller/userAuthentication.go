package controller

import (
	"backend/utils"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/machinebox/graphql"
)

func sendToken(w http.ResponseWriter, role string, response AuthResponse, is_email_verified bool) {
	token, err := utils.GetToken(response.ID, role, is_email_verified)
	if err != nil {
		http.Error(w, "Something went wrong when creating token", http.StatusBadRequest)
		return
	}

	response.Role = role
	response.Token = token
	response.IsEmailVerified = is_email_verified

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func Login(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	r.Body = io.NopCloser(bytes.NewBuffer(body))

	var loginRequest struct {
		Input struct {
			LoginCredentials struct {
				Email    string `json:"email"`
				Password string `json:"password"`
			} `json:"LoginCredentials"`
		} `json:"input"`
	}
	if err := json.Unmarshal(body, &loginRequest); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	client := utils.Client()

	loginCredentials := loginRequest.Input.LoginCredentials
	req := graphql.NewRequest(`
        query($email: String!) {
            users(where: {email: {_eq: $email}}) {
                id
                email
                password
                role
                is_email_verified
            }
        }
    `)

	req.Header.Set("x-hasura-role", "admin")
	req.Var("email", loginCredentials.Email)

	var respData struct {
		Users []struct {
			ID              string `json:"id"`
			Email           string `json:"email"`
			Password        string `json:"password"`
			Role            string `json:"role"`
			IsEmailVerified bool   `json:"is_email_verified"`
		} `json:"users"`
	}

	err = client.Run(context.Background(), req, &respData)
	if err != nil {
		http.Error(w, "Error when querying user to login", http.StatusBadRequest)
		return
	}

	if len(respData.Users) > 0 && utils.ComparePasswords(respData.Users[0].Password, loginCredentials.Password) && respData.Users[0].IsEmailVerified {
		var response AuthResponse
		response.ID = respData.Users[0].ID
		response.Role = respData.Users[0].Role
		response.IsEmailVerified = respData.Users[0].IsEmailVerified
		sendToken(w, respData.Users[0].Role, response, respData.Users[0].IsEmailVerified)
		return
	}

	http.Error(w, "Invalid email or password", http.StatusBadRequest)
}

func Register(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	r.Body = io.NopCloser(bytes.NewBuffer(body))

	var reqBody RequestBody
	if err := json.Unmarshal(body, &reqBody); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	newUser := reqBody.Input.UserInput

	password, err := utils.HashPassword(newUser.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Generate the confirmation token
	confirmationToken, err := utils.GetToken(newUser.Username, "user", false)
	if err != nil {
		http.Error(w, "Failed to generate confirmation token", http.StatusInternalServerError)
		return
	}

	req := graphql.NewRequest(`
		mutation($objects: [users_insert_input!]!) {
			insert_users(objects: $objects) {
				returning {
					id
				}
			}
		}
	`)
	req.Header.Set("x-hasura-role", "admin")

	req.Var("objects", []map[string]interface{}{
		{
			"username":           newUser.Username,
			"role":               newUser.Role,
			"email":              newUser.Email,
			"password":           password,
			"is_email_verified":  newUser.IsEmailVerified,
			"confirmation_token": confirmationToken,
		},
	})

	var respData struct {
		InsertUsers struct {
			Returning []struct {
				ID string `json:"id"`
			} `json:"returning"`
		} `json:"insert_users"`
	}

	client := utils.Client()

	err = client.Run(context.Background(), req, &respData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		fmt.Println("Error executing GraphQL mutation: ", err)
		return
	}

	if len(respData.InsertUsers.Returning) == 0 {
		http.Error(w, "No user ID returned", http.StatusBadRequest)
		fmt.Println("Error: No user ID returned")
		return
	}

	var response AuthResponse
	response.ID = respData.InsertUsers.Returning[0].ID
	sendToken(w, newUser.Role, response, newUser.IsEmailVerified)
}
