package api

import (
	"database/sql"

	"github.com/crunchydosa123/dosa-ops/internal/auth"
	"github.com/crunchydosa123/dosa-ops/internal/repository"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		var req AuthRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": "invalid request"})
			return
		}

		// get password hash from DB
		hash, err := repository.GetHashedPasswordByEmail(db, req.Email)
		if err != nil {
			c.JSON(401, gin.H{"error": "invalid credentials"})
			return
		}

		// compare password
		err = bcrypt.CompareHashAndPassword(
			[]byte(hash),
			[]byte(req.Password),
		)

		if err != nil {
			c.JSON(401, gin.H{"error": "invalid credentials"})
			return
		}

		// generate JWT
		token, err := auth.GenerateToken(1) // replace with actual userID later
		if err != nil {
			c.JSON(500, gin.H{"error": "failed to generate token"})
			return
		}

		c.JSON(200, gin.H{
			"token": token,
		})
	}
}

func Signup() gin.HandlerFunc {
	return func(c *gin.Context) {

		var req AuthRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": "invalid request"})
			return
		}

		// TODO
		// hash password
		// insert user into DB

		c.JSON(201, gin.H{
			"message": "user created",
		})
	}
}
