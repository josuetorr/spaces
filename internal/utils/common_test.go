package utils

import (
	"fmt"
	"os"
	"testing"
)

func TestGetServerURL(t *testing.T) {
	Load(".env")
	expected := fmt.Sprintf("https://%s", os.Getenv("SPACES_SERVER_NAME"))
	result := GetServerURL()
	t.Log(expected)
	t.Log(result)

	if expected != result {
		t.Errorf("Expected: %s. Got %s", expected, result)
	}
}

func TestFullId(t *testing.T) {
	expected := fmt.Sprintf("https://%s/users/bob", getServerName())
	result := GetFullId("users", "bob")

	if expected != result {
		t.Errorf("Expected: %s. Got %s", expected, result)
	}
}
