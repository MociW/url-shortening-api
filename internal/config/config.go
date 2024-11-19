package config

import "github.com/spf13/viper"

var Salt = viper.GetString("JWT_SALT")
