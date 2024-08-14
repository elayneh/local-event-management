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

	"github.com/gin-gonic/gin"
)

func sendToken(ctx *gin.Context, role string, response AuthResponse) {
	// generate jwt token
	token, err := utils.GetToken(response.ID, role)

	if err != nil {
		fmt.Println(err.Error(), "when generating token")
		ctx.JSON(400, gin.H{"message": "something went wrong when creating token"})
		return
	}
	response.Role = role
	response.Token = token
	ctx.JSON(200, response)
}

func Login(ctx *gin.Context) {
	// Read and log the raw request body
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error reading request body"})
		return
	}
	fmt.Println("\n\n\nBODY: ", string(body))
	// Restore the body for later use by creating a new reader
	ctx.Request.Body = io.NopCloser(bytes.NewBuffer(body))

	// Unmarshal the request body into the loginRequest struct
	var loginRequest struct {
		Input struct {
			UserCredentials struct {
				Email    string `json:"email"`
				Password string `json:"password"`
			} `json:"userCredentials"`
		} `json:"input"`
	}
	if err := json.Unmarshal(body, &loginRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	// Extract the email and password
	loginCredentials := loginRequest.Input.UserCredentials
	fmt.Println("Login Credentials: ", loginCredentials)

	// Prepare the GraphQL query to fetch user data
	req := graphql.NewRequest(`
        query($email: String!) {
            users(where: {email: {_eq: $email}}) {
                id
                email
                password
                role
            }
        }
    `)

	// Set the email variable in the GraphQL query
	req.Var("email", loginCredentials.Email)

	// Create a response struct to hold the response data
	var respData struct {
		Users []struct {
			ID       string `json:"id"`
			Email    string `json:"email"`
			Password string `json:"password"`
			Role     string `json:"role"`
		} `json:"users"`
	}

	// Get the GraphQL client
	client := utils.Client()

	// Execute the GraphQL query
	err = client.Run(context.Background(), req, &respData)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error when querying user to login"})
		return
	}

	// Check if the user exists and the password is correct
	if len(respData.Users) > 0 && utils.ComparePasswords(respData.Users[0].Password, loginCredentials.Password) {
		// Prepare and send the authentication token
		var response AuthResponse
		response.ID = respData.Users[0].ID
		response.Role = respData.Users[0].Role
		sendToken(ctx, respData.Users[0].Role, response)
		return
	}

	// If credentials are invalid, return an error
	ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid email or password"})
}

func Register(ctx *gin.Context) {
	// Read and log the raw request body
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error reading request body"})
		return
	}

	// Restore the body for later use by creating a new reader
	ctx.Request.Body = io.NopCloser(bytes.NewBuffer(body))

	// Unmarshal the request body into the RequestBody struct
	var reqBody RequestBody
	if err := json.Unmarshal(body, &reqBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	// Extract user input from the request body
	newUser := reqBody.Input.UserInput
	// Hash the password
	password, err := utils.HashPassword(newUser.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Prepare the GraphQL mutation request
	req := graphql.NewRequest(`
		mutation($objects: [users_insert_input!]!) {
			insert_users(objects: $objects) {
				returning {
					id
				}
			}
		}
	`)

	// Set variables for the mutation
	req.Var("objects", []map[string]interface{}{
		{
			"first_name": newUser.FirstName,
			"last_name":  newUser.LastName,
			"role":       newUser.Role,
			"email":      newUser.Email,
			"password":   password,
		},
	})

	// Create a response struct to hold the response data
	var respData struct {
		InsertUsers struct {
			Returning []struct {
				ID string `json:"id"`
			} `json:"returning"`
		} `json:"insert_users"`
	}

	// Get the GraphQL client
	client := utils.Client()

	// Execute the GraphQL mutation
	err = client.Run(context.Background(), req, &respData)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println("Error executing GraphQL mutation: ", err)
		return
	}

	// Check if the user ID was returned
	if len(respData.InsertUsers.Returning) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "No user ID returned"})
		fmt.Println("Error: No user ID returned")
		return
	}

	// Prepare and send the authentication token
	var response AuthResponse
	response.ID = respData.InsertUsers.Returning[0].ID

	sendToken(ctx, newUser.Role, response)
}

func UpdateUser(ctx *gin.Context) {
	// 1. Get user data from request body
	var inputUser struct {
		Email       string `json:"email"`
		Password    string `json:"password"`
		NewPassword string `json:"newPassword"`
	}

	if err := ctx.ShouldBindJSON(&inputUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 2. Creating the GraphQL query
	req := graphql.NewRequest(`
		query($email: String!) {
			users(where: {email: {_eq: $email}}) {
				id
				password
				role
			}
		}
	`)

	req.Var("email", inputUser.Email)

	// 3. Execute the query and capture the response
	var queryResp struct {
		Users []struct {
			ID       string `json:"id"`
			Password string `json:"password"`
			Role     string `json:"role"`
		} `json:"users"`
	}

	client := utils.Client()
	err := client.Run(context.Background(), req, &queryResp)
	if err != nil {
		fmt.Println(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error when querying user"})
		return
	}

	// 4. Check if the user exists and the password is correct
	if len(queryResp.Users) > 0 && utils.ComparePasswords(queryResp.Users[0].Password, inputUser.Password) {

		var newPassword = queryResp.Users[0].Password

		if inputUser.NewPassword != "" {
			password, err := utils.HashPassword(inputUser.NewPassword)
			newPassword = password
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
		}

		// 5. Creating the GraphQL mutation
		mutationReq := graphql.NewRequest(`
			mutation($email: String!, $password: String!) {
				update_users(
					where: {email: {_eq: $email}}, 
					_set: {password: $password}
				) {
					returning {
						id
						role
						email
					}
				}
			}
		`)

		mutationReq.Var("email", inputUser.Email)
		mutationReq.Var("password", newPassword)

		// 6. Execute the mutation and capture the response
		var mutationResp struct {
			UpdateUsers struct {
				Returning []struct {
					ID    string `json:"id"`
					Role  string `json:"role"`
					Email string `json:"email"`
				} `json:"returning"`
			} `json:"update_users"`
		}

		err = client.Run(context.Background(), mutationReq, &mutationResp)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if len(mutationResp.UpdateUsers.Returning) == 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "No user ID returned"})
			return
		}

		var response AuthResponse
		response.ID = mutationResp.UpdateUsers.Returning[0].ID
		response.Role = mutationResp.UpdateUsers.Returning[0].Role

		// Generate and include the token
		sendToken(ctx, response.Role, response)
		return
	}

	ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid credentials"})
}
