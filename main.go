package main

import (
	"fmt"
	"strings"
	"time"
	"users/api"
	concurentmodel "users/concurentModel"
	"users/database"

	"github.com/gin-gonic/gin"
)

func main() {
	var response string
	fmt.Print("Do you want to insert all users concurently? (yes/no): ")
	fmt.Scanln(&response)
	response = strings.ToLower(response)
	if response == "yes" {
		fmt.Println("inserting all users concurently please wait...")
		time.Sleep(3 * time.Second)
		concurentmodel.InsertUsersConcurently()
	}

	database.ConnectToDB()
	defer database.Close_database()

	fmt.Println("serving API...")
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	api.RegisterRoutes(router)
	router.Run(":8080")
}
