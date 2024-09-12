package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

const (
	confirmationEmailSubject = "Confirm Your Email"
	resetPasswordSubject     = "Reset Your Password"
	contentTypeHeader        = "Content-Type"
	authorizationHeader      = "Authorization"
	contentTypeJSON          = "application/json"
)

func EmailSender(isConfirmation bool, email string, token string) error {
	apiURL := os.Getenv("EMAIL_API_URL")
	apiToken := os.Getenv("EMAIL_API_TOKEN")
	if apiURL == "" || apiToken == "" {
		return fmt.Errorf("missing email API configuration")
	}

	emailData, err := prepareEmailData(isConfirmation, email, token)
	if err != nil {
		return fmt.Errorf("failed to prepare email data: %w", err)
	}

	err = sendEmailRequest(apiURL, apiToken, emailData)
	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}

func sendEmailRequest(url, apiToken string, emailData map[string]interface{}) error {
	jsonData, err := json.Marshal(emailData)
	if err != nil {
		return fmt.Errorf("failed to marshal email data: %w", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set(contentTypeHeader, contentTypeJSON)
	req.Header.Set(authorizationHeader, "Bearer "+apiToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("failed to send email, status code: %d, response: %s", resp.StatusCode, body)
	}

	return nil
}

func prepareEmailData(isConfirmationEmail bool, email, token string) (map[string]interface{}, error) {
	var subject, htmlContent string

	if isConfirmationEmail {
		subject = confirmationEmailSubject
		htmlContent = generateConfirmationEmailHTML(token)
	} else {
		subject = resetPasswordSubject
		htmlContent = generatePasswordResetEmailHTML(token)
	}

	emailData := map[string]interface{}{
		"from": map[string]string{
			"email": "mailtrap@example.com",
			"name":  "Mailtrap Test",
		},
		"to": []map[string]string{
			{
				"email": email,
			},
		},
		"subject": subject,
		"html":    htmlContent,
	}

	return emailData, nil
}

// generateConfirmationEmailHTML generates the HTML content for email confirmation
func generateConfirmationEmailHTML(token string) string {
	confirmationLink := fmt.Sprintf("http://localhost:3000/users/verify-email?token=%s", token)
	return fmt.Sprintf(`
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8" />
			<meta name="viewport" content="width=device-width, initial-scale=1.0" />
			<title>Email Confirmation</title>
			<style>
				/* Styles omitted for brevity */
			</style>
		</head>
		<body>
			<div class="container">
				<div class="header">
					<img src="https://via.placeholder.com/100x100" alt="Company Logo" />
					<h1>Email Confirmation</h1>
				</div>
				<div class="content">
					<p>Thank you for signing up! Please confirm your email address by clicking the button below:</p>
					<a href="%s" class="button">Confirm Email</a>
				</div>
			</div>
		</body>
		</html>
	`, confirmationLink)
}

func generatePasswordResetEmailHTML(reset_token string) string {
	passwordResetLink := fmt.Sprintf("http://localhost:3000/users/reset-password?reset_token=%s", reset_token)
	return fmt.Sprintf(`
		<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>Reset Password</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            background-color: #f4f4f4;
            margin: 0;
            padding: 0;
        }
        .container {
            max-width: 600px;
            margin: 20px auto;
            background-color: #ffffff;
            border-radius: 8px;
            box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
            overflow: hidden;
        }
        .header {
            background-color: #4a90e2;
            color: #ffffff;
            padding: 20px;
            text-align: center;
        }
        .header img {
            width: 100px;
            height: auto;
            margin-bottom: 10px;
        }
        .header h1 {
            margin: 0;
            font-size: 24px;
        }
        .content {
            padding: 20px;
            text-align: center;
        }
        .content p {
            font-size: 16px;
            color: #333333;
            margin-bottom: 20px;
        }
        .button {
            display: inline-block;
            background-color: #4a90e2;
            color: #ffffff;
            padding: 15px 25px;
            border-radius: 5px;
            text-decoration: none;
            font-weight: bold;
            transition: background-color 0.3s;
        }
        .button:hover {
            background-color: #357ab8;
        }
        .footer {
            padding: 20px;
            text-align: center;
            font-size: 12px;
            color: #777777;
            background-color: #f9f9f9;
        }
    </style>
</head>
<body>
    <div class="container">
        <div class="header">
            <img src="https://via.placeholder.com/100x100" alt="Company Logo" />
            <h1>Reset Password</h1>
        </div>
        <div class="content">
            <p>Please reset your password by clicking the button below:</p>
            <a href="%s" class="button">Reset Password</a>
        </div>
        <div class="footer">
            <p>If you did not request this email, please ignore it.</p>
            <p>&copy; 2024 Your Company Name. All rights reserved.</p>
        </div>
    </div>
</body>
</html>
	`, passwordResetLink)
}

func generateResetToken(email string) (string, error) {
	claims := jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(15 * time.Minute).Unix(), // Shorter expiry for better security
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

func verifyResetToken(token string) (string, error) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return "", err
	}

	email, ok := claims["email"].(string)
	if !ok {
		return "", fmt.Errorf("invalid email in token")
	}

	return email, nil
}
