package config

import (
	"net/http"

	"github.com/gorilla/sessions"
)

/*
	you can use session store databases, in our case we could used "github.com/antonlindstrom/pgstore"
	there is a plenty mysql, redis, mongo, check the docs
	? https://github.com/gorilla/sessions#store-implementations
*/

var store = sessions.NewCookieStore([]byte(AppConfig().SESSION_ID))

// store.Options.MaxAge = 60 * 60 * 24 * 10 // 10 days, MaxAge is in seconds

func SaveSession(userId interface{}, w http.ResponseWriter, r *http.Request) error {
	store.Options.MaxAge = 60 * 60 * 24 * 2 // 2 days, MaxAge is in seconds
	//------------------------ session-key/name
	session, _ := store.Get(r, "user")
	// Set some session values.
	session.Values["id"] = userId
	// Save it before we write to the response/return from the handler.
	err := session.Save(r, w)
	if err != nil {
		return err
	}

	return nil
}

func GetSession(getKey string, r *http.Request) interface{} {
	session, _ := store.Get(r, "user")
	return session.Values[getKey]
}
