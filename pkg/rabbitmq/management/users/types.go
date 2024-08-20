package users

type UserTag string

const (
	Admin        UserTag = "administrator"
	Monitoring   UserTag = "monitoring"
	Policymaker  UserTag = "policymaker"
	Management   UserTag = "management"
	Impersonator UserTag = "impersonator"
)

type createUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Tags     string `json:"tags"`
}

type userResponse struct {
	Name             string         `json:"name"`
	PasswordHash     string         `json:"password_hash"`
	HashingAlgorithm string         `json:"hashing_algorithm"`
	Tags             []string       `json:"tags"`
	Limits           map[string]int `json:"limits"`
}
