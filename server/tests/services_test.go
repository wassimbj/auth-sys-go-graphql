package tests

import (
	"auth-sys/services"
	"testing"
)

func TestServices(t *testing.T) {
	t.Run("Create user account", func(t *testing.T) {
		s := services.CreateAccount("wassim", "wassim@gmail.com", "12345")

		if s != true {
			t.FailNow()
		}
	})

	t.Run("Login user to his account", func(t *testing.T) {
		s := services.LoginUser("wassim@gmail.com", "12345")

		if s != true {
			t.FailNow()
		}
	})
}
