package errortype

// Define error codes.
type ErrorCode int

// error code
const (
	ServerError ErrorCode = iota
	Unauthorized
	Success
	UserNotFound
	PosNotFound
	NoToken
	UnExpectedSign
	InvalidToken
	ClaimError
	BadRequestPayload
)

// Map error codes to error messages.
var errorMessages = map[ErrorCode]string{
	ServerError:       "Internal Server Error",
	Unauthorized:      "Unauthorized",
	Success:           "Response Success",
	UserNotFound:      "user not found or wrong username/password or user not active",
	PosNotFound:       "posclientid not found",
	NoToken:           "no token provided",
	UnExpectedSign:    "unexpected signing method",
	InvalidToken:      "invalid token",
	ClaimError:        "Error during claiming token",
	BadRequestPayload: "Payload structure does not match",
}

// Helper function to get error message.
func ErrorMessage(code ErrorCode) string {
	return errorMessages[code]
}

//วิธี map จาก error code ไปเป็น message
// func main() {
//     fmt.Println(ServerError, ErrorMessage(ServerError))  // "500 Server Internal Error"
//     fmt.Println(Unauthorized, ErrorMessage(Unauthorized))  // "501 Error: Unauthorized"
//     fmt.Println(Success, ErrorMessage(Success))  // "502 Response Success"
// }
