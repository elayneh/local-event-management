package controller

import (
	"backend/utils"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/machinebox/graphql"
)

var secretKey = []byte(os.Getenv("JWT_SECRET"))

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

func ResetPassword(w http.ResponseWriter, r *http.Request) {
	var req PasswordResetRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload: "+err.Error(), http.StatusBadRequest)
		return
	}

	email := req.Input.ResetRequestInput.Email
	if email == "" {
		http.Error(w, "Email is required", http.StatusBadRequest)
		return
	}

	resetToken, err := generateResetToken(email)
	if err != nil {
		http.Error(w, "Error generating reset token: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if err := EmailSender(false, email, resetToken); err != nil {
		log.Printf("Error sending reset email to %s: %v", email, err)
		http.Error(w, "Failed to send password reset email", http.StatusInternalServerError)
		return
	}

	reqGraphQL := graphql.NewRequest(`
        mutation UpdateResetToken($reset_token: String!, $email: String!) {
            update_users(where: {email: {_eq: $email}}, _set: {reset_token: $reset_token}) {
                affected_rows
            }
        }
    `)
	reqGraphQL.Var("email", email)
	reqGraphQL.Var("reset_token", resetToken)

	client := utils.Client()
	var respData struct {
		UpdateUsers struct {
			AffectedRows int `json:"affected_rows"`
		} `json:"update_users"`
	}

	if err := client.Run(context.Background(), reqGraphQL, &respData); err != nil {
		http.Error(w, "Error setting reset token: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if respData.UpdateUsers.AffectedRows == 0 {
		http.Error(w, "Invalid token or user not found", http.StatusBadRequest)
		return
	}

	// Respond with a success message
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Password reset email sent"})
}
func UpdatePassword(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	// Reset r.Body for further reading
	r.Body = io.NopCloser(bytes.NewBuffer(body))

	var req UpdatePasswordRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload: "+err.Error(), http.StatusBadRequest)
		return
	}

	newPassword := req.Input.UpdatePasswordInput.NewPassword
	resetToken := req.Input.UpdatePasswordInput.ResetToken

	if newPassword == "" || resetToken == "" {
		http.Error(w, "Password and token are required", http.StatusBadRequest)
		return
	}

	email, err := verifyResetToken(resetToken)
	if err != nil {
		http.Error(w, "Invalid reset token: "+err.Error(), http.StatusUnauthorized)
		return
	}

	log.Println("Email: ", email)

	hashedPassword, err := utils.HashPassword(newPassword)
	if err != nil {
		log.Println("Error hashing password: ", err)
		http.Error(w, "Error hashing password: "+err.Error(), http.StatusInternalServerError)
		return
	}

	reqGraphQL := graphql.NewRequest(`
        mutation UpdatePassword($email:String! $password:String!){
            update_users(where: {email: {_eq: $email}}, _set: {password: $password, reset_token: ""}) {
                affected_rows
            }
        }
    `)
	reqGraphQL.Var("email", email)
	reqGraphQL.Var("password", hashedPassword)

	client := utils.Client()
	var respData struct {
		UpdateUsers struct {
			AffectedRows int `json:"affected_rows"`
		} `json:"update_users"`
	}

	if err := client.Run(context.Background(), reqGraphQL, &respData); err != nil {
		http.Error(w, "Error updating password: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if respData.UpdateUsers.AffectedRows == 0 {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	log.Println("Response data: ", respData.UpdateUsers.AffectedRows)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Password updated successfully"})
}
