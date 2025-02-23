package user

type RegisterRequest struct {
	Username       string `json:"username"`
	Password       string `json:"password"`
	Phone          string `json:"phone"`
	EmergencyPhone string `json:"emergency_phone"`
	Letter         string `json:"letter"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
