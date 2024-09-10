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
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>Email Confirmation</title>
    <style>
      body {
        font-family: "Helvetica Neue", Arial, sans-serif;
        background-color: #f9f9f9;
        margin: 0;
        padding: 0;
        color: #333;
      }

      .container {
        max-width: 600px;
        margin: 40px auto;
        padding: 20px;
        background-color: #fff;
        box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
        border-radius: 8px;
      }

      .header {
        background-color: #772f6e;
        padding: 20px;
        text-align: center;
        border-top-left-radius: 8px;
        border-top-right-radius: 8px;
        color: white;
      }

      .header h1 {
        margin: 0;
        font-size: 24px;
        font-weight: normal;
      }

      .header img {
        max-width: 100px;
        margin-bottom: 10px;
      }

      .content {
        padding: 20px;
        text-align: center;
      }

      .content p {
        font-size: 16px;
        line-height: 1.6;
        margin-bottom: 30px;
      }

      .button {
        display: inline-block;
        padding: 12px 30px;
        font-size: 18px;
        font-weight: bold;
        color: #fff;
        background-color: #cc1af0;
        text-decoration: none;
        border-radius: 50px;
        box-shadow: 0 4px 12px rgba(0, 123, 255, 0.3);
        transition: background-color 0.3s, box-shadow 0.3s;
      }

      .button:hover {
        background-color: #28af3a;
        box-shadow: 0 6px 15px rgba(55, 168, 21, 0.5);
      }

      .footer {
        text-align: center;
        padding: 10px;
        margin-top: 20px;
        font-size: 14px;
        color: #777;
        border-top: 1px solid #eee;
      }

      .footer a {
        color: #007bff;
        text-decoration: none;
        margin: 0 10px;
      }

      .footer a:hover {
        text-decoration: underline;
      }

      @media (max-width: 600px) {
        .container {
          margin: 20px;
          padding: 15px;
        }

        .content p {
          font-size: 14px;
        }

        .button {
          padding: 10px 20px;
          font-size: 16px;
        }
      }
    </style>
  </head>
  <body>
    <div class="container">
      <div class="header">
        <img src="https://via.placeholder.com/100x100" alt="Company Logo" />
        <h1>Email Confirmation</h1>
      </div>
      <div class="content">
        <p>
          Thank you for signing up! Please confirm your email address by
          clicking the button below:
        </p>
        <a href="%s" class="button">Confirm Email</a>
      </div>
      <div class="footer">
        <p>If you didn’t request this email, you can safely ignore it.</p>
        <p>
          Follow us on:
          <a href="https://www.facebook.com/begetaname">Facebook</a> |
          <a href="https://x.com/betenatefera">Twitter</a> |
          <a href="https://www.linkedin.com/in/belayneh-getachew-353a73222"
            >LinkedIn</a
          >
        </p>
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
	req.Header.Set("X-Hasura-Role", "admin")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error sending request: %v", err)
	}

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
