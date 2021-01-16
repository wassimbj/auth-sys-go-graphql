// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type User struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt string `json:"createdAt"`
}

type CreateAccData struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LogoutData struct {
	Out *bool `json:"out"`
}
