package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	// Default initializes an engine with a Logger() and Recovery() middleware
	router := gin.Default()

	// usersGroup = router.Group("users")

	router.Run(":3000")
}
