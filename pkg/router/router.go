package router

import (
	"spotbuzz/pkg/controller"

	"github.com/gin-gonic/gin"
)

// Router initializes and returns a new Gin router with all the required routes.
func Router() *gin.Engine {
	// Create a new Gin router with default middleware (logger, recovery, etc.)
	r := gin.Default()

	// Define the routes and map them to the corresponding controller functions

	// POST /players - Add a new player
	r.POST("/players", controller.AddPlayer)

	// PUT /players - Update player attributes (name and score)
	r.PUT("/players", controller.UpdatePlayer)

	// DELETE /players - Delete a player by ID
	r.DELETE("/players", controller.DeletePlayer)

	// GET /players - Fetch a list of all players in descending order
	r.GET("/players", controller.GetPlayers)

	// GET /players/rank - Fetch the player ranked by a given rank value (val)
	r.GET("/players/rank", controller.GetPlayerByRank)

	// GET /players/random - Fetch a random player
	r.GET("/players/random", controller.GetRandomPlayer)

	return r
}
