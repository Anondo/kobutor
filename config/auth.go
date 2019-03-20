package config

import (
	"strings"

	"github.com/spf13/viper"
)

var auth map[string]string

func loadAuth() {
	auth = viper.GetStringMapString("auth")
}

// GetAuth returns the authentications map
func GetAuth(id string) string {
	return auth[strings.ToLower(id)]
}
