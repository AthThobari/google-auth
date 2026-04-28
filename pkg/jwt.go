package pkg

import (
    "time"
    "github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(userID, username, email, role string, secret string) (string, error) {
    claims := jwt.MapClaims{
        "id": userID,
        "email": email,
        "username": username,
        "role": role,
        "iss": "auth-service",
        "exp":   time.Now().Add(24 * time.Hour).Unix(),
        "iat": time.Now().Unix(),
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte(secret))
}