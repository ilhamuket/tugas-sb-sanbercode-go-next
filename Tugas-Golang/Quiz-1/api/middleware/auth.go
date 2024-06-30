package middleware

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var users = map[string]string{
	"admin":   "password",
	"editor":  "secret",
	"trainer": "rahasia",
}

func BasicAuth(h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		username, password, ok := r.BasicAuth()
		if !ok {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		if pwd, userExists := users[username]; !userExists || pwd != password {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		h(w, r, ps)
	}
}
