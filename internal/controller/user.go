package controller

import (
	"context"
	"log"
	"net/http"
	"pro/internal/global"
	"time"

	"github.com/golang-jwt/jwt/v4"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(c *gin.Context) {

	var newUser struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}


	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Error hashing password:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	result, err := global.DB.Exec(context.Background(), "INSERT INTO users (username, password) VALUES ($1, $2)", newUser.Username, hashedPassword)
	if err != nil {
		log.Println("Error inserting user:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	log.Println("User created successfully. Rows affected:", result.RowsAffected())
}

func RegisterUser(c *gin.Context) {
	CreateUser(c);
}

func LoginUser(c *gin.Context) {
	type Credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var credentials Credentials

	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	// Query the database to check if the user exists and the password matches
	var storedPassword string
	err := global.DB.QueryRow(context.Background(), "SELECT password FROM users WHERE username = $1", credentials.Username).Scan(&storedPassword)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials: C11"})
		return
	}


	err = bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(credentials.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials: KAA"})
		return
	}

	// Authentication successful, generate JWT token

	
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"usermane": credentials.Username,
		"password": credentials.Password,
		"exp": time.Now().Add(time.Hour*24*30).Unix(),
	})
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = credentials.Username
	claims["password"] = credentials.Password
	
	tokenString, _ := token.SignedString(global.JWT_KEY)

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "","", false, true)

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

// {
//     "username": "newuser2",
//     "password": "password123"
// }

// {
//     "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXNzd29yZCI6InBhc3N3b3JkMTIzIiwidXNlcm5hbWUiOiJuZXd1c2VyMiJ9.J5a_P65VcjHmzTSrWeDTZipKbVozMBrPUcOIni_-_RU"
// }


func GetUserInfo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Welcome to the protected route"})
	
}