package graphql

import (
	"backend/hasura"
	"encoding/json"
	"net/http"
)

func GraphQLHandler(w http.ResponseWriter, r *http.Request) {
	var req hasura.GraphQLRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	response, err := hasura.SendGraphQLRequest(req.Query, req.Variables)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(map[string]interface{}{
		"data": response.Data,
	}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
