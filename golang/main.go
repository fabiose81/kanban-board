package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/joho/godotenv"

	"kanban-board/aws"
	"kanban-board/structs"
	"kanban-board/utils"
)

func main() {
	godotenv.Load()

	router := gin.New()

	router.SetTrustedProxies([]string{os.Getenv("TRUSTED_PROXIES")})

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{os.Getenv("ALLOW_ORIGINS")},
		AllowMethods: []string{"GET"},
		AllowHeaders: []string{"Content-Type", "Authorization"},
		MaxAge:       12 * time.Hour,
	}))

	router.POST("/save", save)
	router.GET("/get", get)

	router.Run(os.Getenv("PORT"))
}

func save(c *gin.Context) {
	status, statusCode, msg := validateJWT(c)
	if status {
		var board structs.BoardRequest
		if err := c.BindJSON(&board); err != nil {
			c.String(statusCode, fmt.Sprintf("'%s'", err))
		}
		board.UserId = msg
		result, err := aws.Save(board)
		if err == nil {
			c.JSON(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, fmt.Sprintf("'%s'", err))
		}
	} else {
		c.JSON(statusCode, fmt.Sprintf("'%s'", msg))
	}
}

func get(c *gin.Context) {
	status, statusCode, msg := validateJWT(c)
	if status {
		boards, err := aws.Get(msg)
		if err == nil {
			c.JSON(http.StatusOK, boards.Boards)
		} else {
			c.String(statusCode, fmt.Sprintf("'%s'", err))
		}
	} else {
		c.String(statusCode, fmt.Sprintf("'%s'", msg))
	}
}

func validateJWT(c *gin.Context) (bool, int, string) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return false, http.StatusUnauthorized, utils.MSG_HEADER_MISSING
	}

	jwtToken := strings.TrimPrefix(authHeader, "Bearer ")
	_, sub, err := aws.ValidateJWT(jwtToken)
	if err != nil {
		fmt.Printf(utils.MSG_JWT_ERROR_VALIDATE, err)
		return false, http.StatusBadRequest, fmt.Sprintf("'%s'", err)
	}

	return true, http.StatusOK, sub
}
