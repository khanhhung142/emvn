package auth_usecase

// Define the input and output of the usecase layer here

type SignInOutput struct {
	Token string `json:"token"`
	Exp   int64  `json:"exp_time"`
}
