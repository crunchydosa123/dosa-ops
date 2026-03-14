package api

import (
	"context"

	"github.com/crunchydosa123/dosa-ops/internal/auth"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
)

type SignupRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(db *pgx.Conn) gin.HandlerFunc {
	return func(c *gin.Context) {

		var req LoginRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": "invalid request"})
			return
		}

		var userID string
		var hash string

		err := db.QueryRow(
			context.Background(),
			`SELECT id, password_hash FROM users WHERE email=$1`,
			req.Email,
		).Scan(&userID, &hash)

		if err != nil {
			c.JSON(401, gin.H{"error": "invalid fetching credentials"})
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
		token, err := auth.GenerateToken(userID)

		if err != nil {
			c.JSON(500, gin.H{"error": "failed to generate token"})
			return
		}

		c.JSON(200, gin.H{
			"token": token,
		})
	}
}

func Signup(db *pgx.Conn) gin.HandlerFunc {
	return func(c *gin.Context) {

		var req SignupRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": "invalid request"})
			return
		}

		hashedPassword, err := bcrypt.GenerateFromPassword(
			[]byte(req.Password),
			bcrypt.DefaultCost,
		)

		if err != nil {
			c.JSON(500, gin.H{"error": "failed to hash password"})
			return
		}

		// insert user
		_, err = db.Exec(
			context.Background(),
			`INSERT INTO users (name, email, password_hash)
			 VALUES ($1,$2,$3)`,
			req.Name,
			req.Email,
			string(hashedPassword),
		)

		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(201, gin.H{
			"message": "user created",
		})
	}
}
