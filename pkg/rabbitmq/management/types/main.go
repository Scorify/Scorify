package types

type ErrorResponse struct {
	Error  string `json:"error"`
	Reason string `json:"reason"`
}
