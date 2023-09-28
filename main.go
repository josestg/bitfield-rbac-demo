package main

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/josestg/bitfield-rbac-demo/handlers"
	"github.com/josestg/bitfield-rbac-demo/rbac"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"os"
)

var secret = []byte(os.Getenv("SECRET"))

func main() {
	mux := httprouter.New()

	mux.Handler(http.MethodPost, "/token", handlers.NewToken(secret))

	// protected routes
	mux.Handler(http.MethodGet, "/users", WithPermission(handlers.SeeUsers(), rbac.SeeUsers))
	mux.Handler(http.MethodPost, "/users", WithPermission(handlers.AddUsers(), rbac.AddUsers))
	mux.Handler(http.MethodDelete, "/users", WithPermission(handlers.DelUsers(), rbac.DelUsers))
	mux.Handler(http.MethodGet, "/roles", WithPermission(handlers.SeeRoles(), rbac.SeeRoles))
	mux.Handler(http.MethodPost, "/roles", WithPermission(handlers.AddRoles(), rbac.AddRoles))
	mux.Handler(http.MethodDelete, "/roles", WithPermission(handlers.DelRoles(), rbac.DelRoles))
	mux.Handler(http.MethodGet, "/emails", WithPermission(handlers.SeeEmails(), rbac.SeeEmails))
	mux.Handler(http.MethodPost, "/emails", WithPermission(handlers.AddEmails(), rbac.AddEmails))
	mux.Handler(http.MethodPut, "/emails", WithPermission(handlers.PutEmails(), rbac.PutEmails))

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatalf("could not listen on port 8080 %v", err)
	}
}

func WithPermission(next http.Handler, p rbac.Permission) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bearer := r.Header.Get("Authorization")
		if bearer == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		tokstr := bearer[len("Bearer "):]
		var claims struct {
			jwt.RegisteredClaims
			Permissions []rbac.Permission `json:"permissions"`
		}

		_, err := jwt.ParseWithClaims(tokstr, &claims, func(token *jwt.Token) (interface{}, error) {
			return secret, nil
		})

		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		role := rbac.NewRole(claims.Permissions...)
		if !role.HasPermission(p) {
			http.Error(w, "You don't have permission", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}
