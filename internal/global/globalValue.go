package global

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

var DB *pgx.Conn

var Gin = gin.Default()

var JWT_KEY = []byte("FFFRRRTTT_4324")
