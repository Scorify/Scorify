package users

type UserTag string

const (
	UserTagAdmin        UserTag = "administrator"
	UserTagMonitoring   UserTag = "monitoring"
	UserTagPolicymaker  UserTag = "policymaker"
	UserTagManagement   UserTag = "management"
	UserTagImpersonator UserTag = "impersonator"
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
