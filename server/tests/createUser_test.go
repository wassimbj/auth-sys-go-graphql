package tests

import (
	"auth-sys/services"
	"testing"
)

func TestCreateUser(t *testing.T) {
	s := services.CreateAccount("khmais", "khmais@gmail.com", "12345")

	if s != true {
		t.FailNow()
	}
}
