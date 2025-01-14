package handlers

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"github.com/delemike/recipe_api/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
	"net/http"
	"os"
	"time"
)

type AuthHandler struct {
	collection *mongo.Collection
	ctx        context.Context
}

func NewAuthHandler(ctx context.Context, collection *mongo.Collection) *AuthHandler {
	return &AuthHandler{
		collection: collection,
		ctx:        ctx,
	}
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type JWTOutput struct {
	Token   string    `json:"token"`
	Expires time.Time `json:"expires"`
}

func (handler *AuthHandler) SignInHandler(c *gin.Context) {

	var user models.User
	// parsing user --> deserialization
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hash the password
	h := sha256.New()
	h.Write([]byte(user.Password))
	passwordHash := hex.EncodeToString(h.Sum(nil))
	cur := handler.collection.FindOne(handler.ctx, bson.M{
		"username": user.Username,
		"password": passwordHash,
	})

	if cur.Err() != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid username or password"})
		return
	}

	// generate session token - will
	sessionToken := xid.New().String()
	session := sessions.Default(c)
	session.Set("username", user.Username)
	session.Set("token", sessionToken)
	err := session.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not save session token:" + err.Error()})
		return
	}

	// 10 minutes
	//expirationTime := time.Now().Add(10 * time.Minute)
	//claims := &Claims{
	//	Username: user.Username,
	//	StandardClaims: jwt.StandardClaims{
	//		ExpiresAt: expirationTime.Unix(),
	//	},
	//}
	//
	//token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	//if err != nil {
	//	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	//	return
	//}
	//
	//jwtOutput := JWTOutput{
	//	Token:   tokenString,
	//	Expires: expirationTime,
	//}
	c.JSON(http.StatusOK, gin.H{"message": "User signed in"})
}

func (handler *AuthHandler) SignOutHandler(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	err := session.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User signed out"})

}

func (handler *AuthHandler) RefreshHandler(c *gin.Context) {
	tokenValue := c.GetHeader("Authorization")
	claims := &Claims{}

	// Parse the token
	token, err := jwt.ParseWithClaims(tokenValue, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	// Specifically check for token expiration
	if err != nil {
		var ve *jwt.ValidationError
		if errors.As(err, &ve) {
			if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired, which is fine - we'll continue with refresh
			} else {
				// Some other validation error but NOT expiration like wrong signature, malformed token, etc.)
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
				return
			}
		}
	}
	// If token is not expired, check if it's close to expiration
	if token != nil && token.Valid {
		timeUntilExpiry := time.Unix(claims.ExpiresAt, 0).Sub(time.Now())
		if timeUntilExpiry > 30*time.Second {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Token is not expired yet"})
			return
		}
	}

	// Generate new token
	expirationTime := time.Now().Add(5 * time.Minute)
	claims.ExpiresAt = expirationTime.Unix()
	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := newToken.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	jwtOutput := JWTOutput{
		Token:   tokenString,
		Expires: expirationTime,
	}
	c.JSON(http.StatusOK, jwtOutput)
}

func AuthMiddleware() gin.HandlerFunc {

	//return func(c *gin.Context) {
	//	tokenValue := c.GetHeader("Authorization")
	//	if tokenValue == "" {
	//		slog.Error("Missing Authorization header")
	//		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization header"})
	//		return
	//	}
	//	claims := &Claims{}
	//
	//	tkn, err := jwt.ParseWithClaims(tokenValue, claims, func(token *jwt.Token) (interface{}, error) {
	//		return []byte(os.Getenv("JWT_SECRET")), nil
	//	})
	//
	//	// Check for errors or nil token
	//	if err != nil || tkn == nil {
	//		if err != nil {
	//			slog.Error("Token parsing error: " + err.Error())
	//		} else {
	//			slog.Error("Token is nil")
	//		}
	//		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid or missing token"})
	//		return
	//	}
	//
	//	// Check token validity
	//	if !tkn.Valid {
	//		slog.Error("Invalid token")
	//		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
	//		return
	//	}
	//
	//	// Token is valid, proceed to the next middleware/handler
	//	c.Next()
	//
	//}
	return func(c *gin.Context) {
		session := sessions.Default(c)
		sessionToken := session.Get("token")
		if sessionToken == nil {
			c.JSON(http.StatusForbidden, gin.H{"message": "not logged in"})
			c.Abort()
		}
		c.Next()
	}

}
