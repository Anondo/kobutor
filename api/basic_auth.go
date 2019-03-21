package api

import (
	"context"
	"kobutor/config"
	"net/http"
	"strings"

	"github.com/thedevsaddam/renderer"
)

type contextApp string

const (
	contextKeyApp contextApp = "app:name-from-secret"
)

// BasicAuth is the basic authentication middleware for the kobutor api
func BasicAuth(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		usr, pass, ok := r.BasicAuth()
		usr = strings.TrimSpace(usr)
		pass = strings.TrimSpace(pass)
		uPass := strings.TrimSpace(config.GetAuthPassword(usr)) // password from user
		if !ok || usr == "" || pass == "" {
			renderer.New().JSON(w, http.StatusUnauthorized, renderer.M{
				"message": "Unauthorized",
			})
			return
		}

		if usr != "" && pass != uPass {
			renderer.New().JSON(w, http.StatusUnauthorized, renderer.M{
				"message": "Unauthorized",
			})
			return
		}

		ctx := context.WithValue(r.Context(), contextKeyApp, usr)
		next.ServeHTTP(w, r.WithContext(ctx))
	}

	return http.HandlerFunc(fn)
}
