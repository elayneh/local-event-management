package controller

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"backend/utils"

	"github.com/machinebox/graphql"
)

func SendConfirmationEmail(email, token string) error {
	url := os.Getenv("EMAIL_API_URL")
	apiToken := os.Getenv("API_TOKEN")

	confirmationLink := fmt.Sprintf("http://localhost:3000/users/verify-email?token=%s", token)

	
	htmlContent := fmt.Sprintf(`
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>Email Confirmation</title>
			<style>
				body { font-family: Arial, sans-serif; }
				.container { max-width: 600px; margin: auto; padding: 20px; }
				.header { background-color: #f4f4f4; padding: 10px; text-align: center; }
				.content { margin-top: 20px; }
				.button { display: inline-block; padding: 10px 20px; font-size: 16px; color: #fff; background-color: #007bff; text-decoration: none; border-radius: 5px; }
				.button:hover { background-color: #0056b3; }
			</style>
		</head>
		<body>
			<div class="container">
				<div class="header">
					<h1>Email Confirmation</h1>
				</div>
				<div class="content">
					<p>Click the link below to confirm your email address:</p>
					<a href="%s" class="button">Confirm Email</a>
				</div>
			</div>
		</body>
		</html>
	`, confirmationLink)

	payload := map[string]interface{}{
		"from": map[string]string{
			"email": "mailtrap@example.com",
			"name":  "Mailtrap Test",
		},
		"to": []map[string]string{
			{
				"email": email,
			},
		},
		"subject":  "Confirm your Email",
		"html":     htmlContent,
		"category": "Integration Test",
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("error marshaling JSON: %v", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Authorization", "Bearer "+apiToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error sending request: %v", err)
	}
	fmt.Println("response: ", resp)
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to send email, status code: %d", resp.StatusCode)
	}

	return nil
}

func EmailConfirmationWebhook(w http.ResponseWriter, r *http.Request) {
	var event EventPayload
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		http.Error(w, "Invalid event payload", http.StatusBadRequest)
		return
	}

	newUser := event.Event.Data.New

	if err := SendConfirmationEmail(newUser.Email, newUser.Confirmation_Token); err != nil {
		http.Error(w, "Failed to send confirmation email", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Email confirmation sent successfully"))
}

func ConfirmEmail(w http.ResponseWriter, r *http.Request) {
	var confirmationRequest struct {
		Token string `json:"token"`
	}

	
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}
	if err := json.Unmarshal(body, &confirmationRequest); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	
	req := graphql.NewRequest(`
		mutation VerifyEmail($token: String!) {
			update_users(
				where: { confirmation_token: { _eq: $token } }
				_set: { is_email_verified: true }
			) {
				affected_rows
			}
		}
	`)
	req.Var("token", confirmationRequest.Token)

	
	client := utils.Client()
	var respData struct {
		UpdateUsers struct {
			AffectedRows int `json:"affected_rows"`
		} `json:"update_users"`
	}

	if err := client.Run(context.Background(), req, &respData); err != nil {
		http.Error(w, "Error confirming email", http.StatusInternalServerError)
		return
	}

	if respData.UpdateUsers.AffectedRows == 0 {
		http.Error(w, "Invalid token or user not found", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Email confirmed successfully"})
}
