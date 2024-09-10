package controller

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"backend/utils"

	"github.com/machinebox/graphql"
)

func EmailConfirmationWebhook(w http.ResponseWriter, r *http.Request) {
	var event EventPayload
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		http.Error(w, "Invalid event payload", http.StatusBadRequest)
		return
	}

	newUser := event.Event.Data.New

	if err := EmailSender(true, newUser.Email, newUser.Confirmation_Token); err != nil {
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
