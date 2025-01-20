package api

import (
	"fmt"
	"net/http"
	"users/database"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	userId := c.Param("id")
	res, err := database.GetUserByID(userId)
	if err != nil {
		fmt.Println("Err: ", err.Error())
		c.JSON(http.StatusOK, gin.H{
			"user_id": userId,
			"message": "User not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": res,
	})
}
