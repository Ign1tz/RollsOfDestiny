package Server

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"log"
	"net/http"
	"time"
)

func verifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return token, nil
}

func createToken(userid string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"userid": userid,
			"exp":    time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func checkToken(w http.ResponseWriter, r *http.Request) (string, bool) {
	tokenString := r.Header.Get("Authorization")
	fmt.Println("tokenstrong", tokenString)
	if tokenString == "" {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Missing authorization header")
		return "", false
	}
	tokenString = tokenString[len("Bearer "):]

	token, err := verifyToken(tokenString)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Invalid token")
		return "", false
	}
	userid, err := getUserIDFromToken(token)
	if err != nil {
		log.Println(err)
	}
	return userid, true
}

func getUserIDFromToken(token *jwt.Token) (string, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", fmt.Errorf("invalid token claims")
	}

	userid, ok := claims["userid"].(string)
	if !ok {
		return "", fmt.Errorf("userid not found in token")
	}

	return userid, nil
}
