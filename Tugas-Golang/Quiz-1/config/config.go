package config

import (
	"fmt"
)

const (
	DBUser     = "root"
	DBPassword = ""
	DBName     = "quiz_1"
)

func GetDBConnectionString() string {
	return fmt.Sprintf("%s:%s@/%s", DBUser, DBPassword, DBName)
}
