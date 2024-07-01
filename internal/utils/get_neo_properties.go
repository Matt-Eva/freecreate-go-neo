package utils

type UserNodeProperties struct {
	Username string
	Name string
	Email string
	Password string
}

type UserHTTPResponseProperties struct {
	Username string `json:"username"`
	Name string `json:"name"`
}

