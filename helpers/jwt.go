package helpers

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang-server-android/config"
	"golang-server-android/models"
	"net/http"
	"time"
)

type JWTClaim struct {
	Email string `json:"email"`
	Id    uint   `json:"id"`
	Name  string `json:"name"`
	jwt.StandardClaims
}

func GenerateJWT(user models.User) (tokenString string, err error) {
	var (
		vConfig = config.GetConfig()
		jwtKey  = []byte(vConfig.GetString("jwt.secret"))
	)
	expirationTime := time.Now().Add(24 * 30 * time.Hour)
	claims := &JWTClaim{
		Email: user.Email,
		Id:    user.ID,
		Name:  user.Name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(jwtKey)
	return
}
func ValidateToken(c *gin.Context, signedToken string) (err error) {
	var (
		vConfig = config.GetConfig()
		jwtKey  = []byte(vConfig.GetString("jwt.secret"))
	)
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		},
	)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid token"})
		return
	}
	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Invalid token"})
		return
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	c.Set("id", claims.Id)
	return
}
