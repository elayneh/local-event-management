package controller

import (
	"backend/utils"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/google/uuid"
	"github.com/machinebox/graphql"
)

const chapaURL = "https://api.chapa.co/v1/transaction/initialize"
const chapaAPIKey = "CHASECK_TEST-VJKoOQEFoQc3TWHC9MwOzRTJmk6P5efo"
const hasuraURL = "http://hasura:8080/v1/graphql"

func generator() string {
	return uuid.New().String()
}

func InitializePayment(w http.ResponseWriter, r *http.Request) {
	TxRef := generator()
	CallbackURL := "https://webhook.site/077164d6-29cb-40df-ba29-8a00e59a7e60"

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	var actionRequest ActionRequest
	err = json.Unmarshal(body, &actionRequest)
	if err != nil {
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}

	requestBody := actionRequest.Input.PaymentCredentials
	if requestBody.Amount <= 0 || TxRef == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	payPayload := map[string]interface{}{
		"amount":       requestBody.Amount,
		"currency":     requestBody.Currency,
		"tx_ref":       TxRef,
		"callback_url": CallbackURL,
	}

	chapaPayload, err := json.Marshal(payPayload)
	if err != nil {
		http.Error(w, "Error preparing Chapa request", http.StatusInternalServerError)
		return
	}

	chapaReq, err := http.NewRequest("POST", chapaURL, bytes.NewBuffer(chapaPayload))
	if err != nil {
		http.Error(w, "Error creating Chapa request", http.StatusInternalServerError)
		return
	}
	chapaReq.Header.Set("Authorization", fmt.Sprintf("Bearer %s", chapaAPIKey))
	chapaReq.Header.Set("Content-Type", "application/json")

	chapaClient := &http.Client{}
	chapaResp, err := chapaClient.Do(chapaReq)
	if err != nil {
		http.Error(w, "Error making payment request to Chapa", http.StatusInternalServerError)
		return
	}
	defer chapaResp.Body.Close()

	if chapaResp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(chapaResp.Body)
		fmt.Println("Chapa error response: ", string(bodyBytes))
		http.Error(w, "Received non-200 response from Chapa", http.StatusInternalServerError)
		return
	}

	bodyBytes, _ := io.ReadAll(chapaResp.Body)
	fmt.Println("Raw Chapa response body: ", string(bodyBytes))

	var chapaResponse ChapaResponse
	err = json.Unmarshal(bodyBytes, &chapaResponse)
	if err != nil || chapaResponse.Status != "success" {
		fmt.Println("Error decoding JSON response: ", err)
		http.Error(w, "Failed to initialize payment", http.StatusInternalServerError)
		return
	}

	req := graphql.NewRequest(`
		mutation($objects: [tickets_insert_input!]!) {
			insert_tickets(objects: $objects) {
				returning {
					id
					status
				}
			}
		}
	`)

	req.Header.Set("x-hasura-role", "admin")
	req.Var("objects", []map[string]interface{}{
		{
			"amount":    requestBody.Amount,
			"currency":  requestBody.Currency,
			"event_id":  requestBody.EventId,
			"tx_ref":    TxRef,
			"user_id":   requestBody.UserId,
			"user_name": requestBody.UserName,
		},
	})

	var respData struct {
		InsertTickets struct {
			Returning []struct {
				ID     string `json:"id"`
				Status string `json:"status"`
			} `json:"returning"`
		} `json:"insert_tickets"`
	}

	client := utils.Client()
	err = client.Run(context.Background(), req, &respData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if len(respData.InsertTickets.Returning) == 0 {
		http.Error(w, "Failed to insert payment request", http.StatusBadRequest)
		return
	}

	graphqlResponse := map[string]interface{}{
		"status":  chapaResponse.Status,
		"message": "Hosted Link",
		"data": map[string]interface{}{
			"checkout_url": chapaResponse.Data.CheckoutURL,
		},
	}

	responseBytes, err := json.Marshal(graphqlResponse)
	if err != nil {
		http.Error(w, "Error preparing GraphQL response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseBytes)

	err = updatePaymentStatus(TxRef, "completed")
	if err != nil {
		fmt.Println("Error updating payment status: ", err)
	}

	http.Redirect(w, r, chapaResponse.Data.CheckoutURL, http.StatusSeeOther)
}

func updatePaymentStatus(txRef string, status string) error {
	mutation := `
	mutation UpdatePaymentStatus($txRef: String!, $status: String!) {
		update_tickets(where: {tx_ref: {_eq: $txRef}}, _set: {status: $status}) {
			affected_rows
		}
	}`

	requestBody := map[string]interface{}{
		"query": mutation,
		"variables": map[string]interface{}{
			"txRef":  txRef,
			"status": status,
		},
	}

	payload, err := json.Marshal(requestBody)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", hasuraURL, bytes.NewBuffer(payload))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-hasura-admin-secret", "myadminsecretkey")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to update payment status")
	}

	return nil
}
