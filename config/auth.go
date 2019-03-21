package config

import (
	"strings"

	"github.com/spf13/viper"
)

var auth map[string]string

func loadAuth() {
	auth = viper.GetStringMapString("auth")
}

// GetAuthPassword returns the password of the username provided in the parameter
func GetAuthPassword(uname string) string {
	return auth[strings.ToLower(uname)]
}
