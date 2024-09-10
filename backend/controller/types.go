package controller

type ErrorResponse struct {
	Message string `json:"message"`
}

type AuthResponse struct {
	ID              string `json:"id"`
	Role            string `json:"role"`
	Token           string `json:"token"`
	IsEmailVerified bool   `json:"is_email_verified"`
}

type InputUser struct {
	Username        string `json:"username"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	Role            string `json:"role"`
	IsEmailVerified bool   `json:"is_email_verified"`
}

type RequestBody struct {
	Input struct {
		UserInput struct {
			Username           string `json:"username"`
			Email              string `json:"email"`
			Password           string `json:"password"`
			Role               string `json:"role"`
			IsEmailVerified    bool   `json:"is_email_verified"`
			Confirmation_Token string `json:"Confirmation_Token"`
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

type EventPayload struct {
	Event struct {
		Data struct {
			New struct {
				ID                 string `json:"id"`
				Email              string `json:"email"`
				Confirmation_Token string `json:"confirmation_token"`
			} `json:"new"`
		} `json:"data"`
	} `json:"event"`
}

type ActionRequest struct {
	Input PaymentInput `json:"input"`
}

type PaymentInput struct {
	PaymentCredentials PaymentCredentials `json:"paymentCredentials"`
}

type PaymentCredentials struct {
	Amount      int    `json:"amount"`
	Currency    string `json:"currency"`
	TxRef       string `json:"txRef"`
	CallbackURL string `json:"callbackURL"`
	ReturnURL   string `json:"returnURL"`
	UserId      string `json:"user_id"`
	EventId     string `json:"event_id"`
	UserName    string `json:"user_name"`
}

type ChapaResponse struct {
	Message string `json:"message"`
	Status  string `json:"status"`
	Data    struct {
		CheckoutURL string `json:"checkout_url"`
	} `json:"data"`
}

type PasswordResetRequest struct {
	Input struct {
		ResetRequestInput struct {
			Email string `json:"email"`
		} `json:"resetRequestInput"`
	} `json:"input"`
}

//	type UpdatePasswordRequest struct {
//		Input struct {
//			UpdatePasswordInput struct {
//				NewPassword string `json:"new_password"`
//				ResetToken  string `json:"reset_token"`
//			} `json:"updatePasswordInput`
//		} `json:"input"`
//	}
type UpdatePasswordRequest struct {
	Action struct {
		Name string `json:"name"`
	} `json:"action"`
	Input struct {
		UpdatePasswordInput struct {
			NewPassword string `json:"new_password"`
			ResetToken  string `json:"reset_token"`
		} `json:"updatePasswordInput"`
	} `json:"input"`
	RequestQuery     string `json:"request_query"`
	SessionVariables struct {
		Role string `json:"x-hasura-role"`
	} `json:"session_variables"`
}
