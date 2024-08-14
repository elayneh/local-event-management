package controller

type AuthResponse struct {
	ID    string `json:"id"`
	Role  string `json:"role"`
	Token string `json:"token"`
}

type InputUser struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Role  string `json:"role"`
}

type RequestBody struct {
	Action map[string]interface{} `json:"action"`
	Input  struct {
		UserInput InputUser `json:"userInput"`
	} `json:"input"`
	RequestQuery     string                 `json:"request_query"`
	SessionVariables map[string]interface{} `json:"session_variables"`
}
