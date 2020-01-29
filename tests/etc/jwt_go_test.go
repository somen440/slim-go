package etc

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type TestJwtUser struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

var (
	expectToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJ7XCJpZFwiOjEyMyxcIm5hbWVcIjpcImhvZ2VcIixcImNyZWF0ZWRfYXRcIjpcIjAwMDEtMDEtMDFUMDA6MDA6MDBaXCJ9In0.cZBjTemxSE4bYwZWYRwiC5kQICe9li5QnOiKG5unpQ4"
	secret      = []byte("key")
)

func createTestUser() *TestJwtUser {
	createdAt, _ := time.Parse(time.RFC3339, "2012-11-11T12:23:34")
	return &TestJwtUser{
		ID:        123,
		Name:      "hoge",
		CreatedAt: createdAt,
	}
}

func TestNewWithClaims(t *testing.T) {
	userJSON, _ := json.Marshal(createTestUser())
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.StandardClaims{
		Subject: string(userJSON),
	})

	tokenString, err := token.SignedString(secret)
	if err != nil {
		t.Errorf(err.Error())
	}

	if expectToken != tokenString {
		t.Errorf("%v != %v", expectToken, tokenString)
	}
}

func TestParse(t *testing.T) {
	token, err := jwt.Parse(expectToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return secret, nil
	})
	if err != nil {
		t.Errorf(err.Error())
	}
	if !token.Valid {
		t.Errorf("%v", token)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		t.Errorf("not ok %v", token.Claims)
	}
	subject, ok := claims["sub"].(string)
	if !ok {
		t.Errorf("not ok %v", claims["sub"])
	}

	var testUser *TestJwtUser
	json.Unmarshal([]byte(subject), &testUser)

	expectUser := createTestUser()
	if !reflect.DeepEqual(expectUser, testUser) {
		t.Errorf("%v != %v", expectUser, testUser)
	}
}
