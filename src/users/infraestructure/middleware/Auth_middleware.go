package middleware

import (
    "net/http"
    "strings"
    "github.com/gin-gonic/gin"
    "github.com/dgrijalva/jwt-go"
    "os"
	"fmt"
)

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        tokenString := c.GetHeader("Authorization")
        if tokenString == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization Header"})
            c.Abort()
            return
        }
        if strings.HasPrefix(tokenString, "Bearer ") {
            tokenString = strings.TrimPrefix(tokenString, "Bearer ")
        }

        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
            }
            return []byte(os.Getenv("JWT_SECRET")), nil
        })

        if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
            c.Set("username", claims["username"])
        } else {
            c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
            c.Abort()
            return
        }

        c.Next()
    }
}
