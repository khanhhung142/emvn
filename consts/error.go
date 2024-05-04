package consts

type errorDeatil struct {
	Code    int    `json:"code"` // Error code for the client to handle
	Message string `json:"message"`
}

type CustomError struct {
	errorDeatil
	HttpStatus int `json:"http"` // HTTP status code
}

// Implement the error interface
func (e CustomError) Error() string {
	return e.Message
}

func (e CustomError) Detail() errorDeatil {
	return e.errorDeatil
}

var (
	CodeInvalidToken       = CustomError{HttpStatus: 401, errorDeatil: errorDeatil{Code: 1001, Message: "Invalid token"}}
	CodeTokenExpired       = CustomError{HttpStatus: 401, errorDeatil: errorDeatil{Code: 1002, Message: "Token expired"}}
	CodeRedisKeyNotFound   = CustomError{HttpStatus: 500, errorDeatil: errorDeatil{Code: 1003, Message: "Redis key not found"}}
	CodeUserAlreadyExist   = CustomError{HttpStatus: 400, errorDeatil: errorDeatil{Code: 1004, Message: "User already exist"}}
	CodeWrongPassword      = CustomError{HttpStatus: 400, errorDeatil: errorDeatil{Code: 1005, Message: "Wrong password"}}
	CodeInternalError      = CustomError{HttpStatus: 500, errorDeatil: errorDeatil{Code: 1006, Message: "Internal error"}}
	CodeInvalidRequest     = CustomError{HttpStatus: 400, errorDeatil: errorDeatil{Code: 1007, Message: "Invalid request"}}
	CodeTokenRequired      = CustomError{HttpStatus: 401, errorDeatil: errorDeatil{Code: 1008, Message: "Token required"}}
	CodeUserAlreadyExists  = CustomError{HttpStatus: 400, errorDeatil: errorDeatil{Code: 1009, Message: "User already exists"}}
	CodeStorageError       = CustomError{HttpStatus: 500, errorDeatil: errorDeatil{Code: 1010, Message: "Storage error"}}
	CodeFileNotFound       = CustomError{HttpStatus: 404, errorDeatil: errorDeatil{Code: 1011, Message: "File not found"}}
	CodeFileInvalid        = CustomError{HttpStatus: 400, errorDeatil: errorDeatil{Code: 1012, Message: "Invalid file"}}
	CodeMusicTrackNotFound = CustomError{HttpStatus: 404, errorDeatil: errorDeatil{Code: 1013, Message: "Music track not found"}}
	CodeUserNotFound       = CustomError{HttpStatus: 404, errorDeatil: errorDeatil{Code: 1014, Message: "User not found"}}
)
