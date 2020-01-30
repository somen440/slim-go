package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

type AuthMiddleware struct {
	path    string
	secret  string
	timeout int
}

func NewAuthMiddleware() (*AuthMiddleware, error) {
	timeout, err := strconv.Atoi(os.Getenv("AUTH_TIMEOUT"))
	if err != nil {
		return nil, err
	}
	return &AuthMiddleware{
		path:    os.Getenv("AUTH_PATH"),
		secret:  os.Getenv("AUTH_SECRET"),
		timeout: timeout,
	}, nil
}

func (m *AuthMiddleware) Middleware(h http.Handler) http.Handler {
	os.Stdout.Write([]byte("AuthMiddleware Middleware\n"))

	os.Stdout.Write([]byte(fmt.Sprintf("  path: %s \n", m.path)))
	os.Stdout.Write([]byte(fmt.Sprintf("  secret: %s \n", m.secret)))
	os.Stdout.Write([]byte(fmt.Sprintf("  timeout: %d \n", m.timeout)))

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.URL.Path, m.path) {
			h.ServeHTTP(w, r)
			return
		}

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			w.WriteHeader(http.StatusForbidden)
			fmt.Fprintf(w, "Forbidden")
			return
		}

		authPair := strings.Fields(authHeader)
		if !(len(authPair) == 2 && authPair[0] == "Bearer") {
			w.WriteHeader(http.StatusForbidden)
			fmt.Fprintf(w, "Forbidden")
			return
		}

		tokenString := authPair[1]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return m.secret, nil
		})
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			fmt.Fprintf(w, "invalid token: %s", err.Error())
			return
		}
		if !token.Valid {
			w.WriteHeader(http.StatusForbidden)
			fmt.Fprintf(w, "invalid token: %s", tokenString)
			return
		}

		h.ServeHTTP(w, r)
	})
}
