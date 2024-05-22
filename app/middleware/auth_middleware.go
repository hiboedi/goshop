package middleware

import (
	"fmt"
	"net/http"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}
func (middleware *AuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Get the token from the Authorization header
	tokenString := r.Header.Get("Authorization")

	// Jika ini adalah permintaan login, lewati autentikasi
	if r.URL.Path == "/api/users/login" && r.Method == "POST" {
		middleware.Handler.ServeHTTP(w, r)
		return
	}

	if r.URL.Path == "/api/users" && r.Method == "POST" {
		middleware.Handler.ServeHTTP(w, r)
		return
	}

	// Pastikan header Authorization tersedia
	if tokenString == "" {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Missing authorization header")
		return
	}

	// Memeriksa apakah token valid
	tokenString = tokenString[len("Bearer "):]
	if err := VerifyToken(tokenString); err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, err.Error()) // Memberikan pesan kesalahan yang sesungguhnya
		return
	}

	// Jika token valid, lanjutkan pemrosesan permintaan
	middleware.Handler.ServeHTTP(w, r)
}
