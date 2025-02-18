package utils

import (
	"fmt"
	"os"
)

func getServerName() string {
	return os.Getenv("SPACES_SERVER_NAME")
}

func GetServerURL() string {
	return fmt.Sprintf("https://%s", getServerName())
}

func GetFullId(objectType string, localId string) string {
	return fmt.Sprintf("%s/%s/%s", GetServerURL(), objectType, localId)
}
