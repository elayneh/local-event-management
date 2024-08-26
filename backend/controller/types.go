package controller

type AuthResponse struct {
	ID    string `json:"id"`
	Role  string `json:"role"`
	Token string `json:"token"`
}

type InputUser struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type RequestBody struct {
	Input struct {
		UserInput struct {
			Username string `json:"username"`
			Email    string `json:"email"`
			Password string `json:"password"`
			Role     string `json:"role"`
		} `json:"userInput"`
	} `json:"input"`
}

type UploadImagesInput struct {
	Files []string `json:"files"`
}

type UploadImagesResponse struct {
	ImageURLs []string `json:"imageUrls"`
}

type UploadImagesRequest struct {
	Action map[string]interface{} `json:"action"`
	Input  struct {
		UserInput UploadImagesInput `json:"input"`
	} `json:"input"`
	RequestQuery     string                 `json:"request_query"`
	SessionVariables map[string]interface{} `json:"session_variables"`
}
