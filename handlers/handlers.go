package handlers

import (
	"encoding/json"
	"github.com/golang-jwt/jwt/v5"
	"github.com/josestg/bitfield-rbac-demo/rbac"
	"net/http"
	"time"
)

func NewToken(secret []byte) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			Permissions []rbac.Permission `json:"permissions"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		exp := jwt.NewNumericDate(time.Now().Add(time.Minute * 5))
		tkn := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"permissions": req.Permissions, "exp": exp})
		str, err := tkn.SignedString(secret)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fakeResponse(w, r, str)
	})
}

func AddUsers() http.Handler  { return fakeHandler("Add Users") }
func DelUsers() http.Handler  { return fakeHandler("Del Users") }
func SeeUsers() http.Handler  { return fakeHandler("See Users") }
func AddRoles() http.Handler  { return fakeHandler("Add Roles") }
func DelRoles() http.Handler  { return fakeHandler("Del Roles") }
func SeeRoles() http.Handler  { return fakeHandler("See Roles") }
func AddEmails() http.Handler { return fakeHandler("Add Emails") }
func PutEmails() http.Handler { return fakeHandler("Put Emails") }
func SeeEmails() http.Handler { return fakeHandler("See Emails") }

func fakeHandler(desc string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fakeResponse(w, r, desc)
	})
}

func fakeResponse(w http.ResponseWriter, r *http.Request, data string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(map[string]string{
		"data":   data,
		"path":   r.URL.Path,
		"method": r.Method,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
