package utils

import "github.com/SimonVillalonIT/filehub/pkg/config"

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (c *Credentials) IsValidCredentials(config config.Config) bool {
	validCredentials := Credentials{Password: config.Credentials.Password, Username: config.Credentials.Username}
	return validCredentials == *c
}
