package admin

type LoginResponse struct {
	Token string `json:"token"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

var (
	// InternalServerError 500
	InternalServerError = ErrorResponse{
		Error: "Internal Server Error",
	}
	// BadRequest 400
	BadRequest = ErrorResponse{
		Error: "Bad Request",
	}
	// Unauthorized 401
	Unauthorized = ErrorResponse{
		Error: "Unauthorized",
	}
)
