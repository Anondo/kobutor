package api

import (
	"context"
	"crypto/subtle"
	"fmt"
	"kobutor/config"
	"net/http"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/thedevsaddam/renderer"
)

// BasicAuth is the basic authentication middleware for the kobutor api
func BasicAuth(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		spew.Dump(r.BasicAuth())
		fmt.Println("kjsdfhkjdsfdskjfhdskjfhkdsfhdsfh")
		usr, pass, ok := r.BasicAuth()
		usr = strings.TrimSpace(usr)
		pass = strings.TrimSpace(pass)
		uPass := strings.TrimSpace(config.GetAuth(usr)) // password from user secret
		if !ok || usr == "" || pass == "" {
			renderer.New().JSON(w, http.StatusUnauthorized, renderer.M{
				"message": "Unauthorized",
			})
			return
		}
		if subtle.ConstantTimeCompare([]byte(uPass), []byte(pass)) != 1 {
			renderer.New().JSON(w, http.StatusUnauthorized, renderer.M{
				"message": "Unauthorized",
			})
			return
		}
		ctx := context.WithValue(r.Context(), "app-secret", usr)
		next.ServeHTTP(w, r.WithContext(ctx))
	}

	return http.HandlerFunc(fn)
}
