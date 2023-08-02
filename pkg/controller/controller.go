package controller

import (
	"fmt"
	"math/rand"
	"net/http"
	"spotbuzz/pkg/database"
	"spotbuzz/pkg/model"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// AddPlayer adds a new player to the database.
func AddPlayer(c *gin.Context) {
	var player model.Player

	// Bind the JSON data to the "player" variable
	if err := c.ShouldBindJSON(&player); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, fmt.Errorf("error during binding:%v", err))
		return
	}

	// Insert the player into the database
	result := database.DB.Create(&player)
	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, fmt.Errorf("error during insert data:%v", result.Error))
		return
	}

	c.JSON(http.StatusOK, player)
}

// UpdatePlayer updates an existing player's information (name and score) based on the player's ID passed as a query parameter.
func UpdatePlayer(c *gin.Context) {
	var player model.Player

	// Parse the player ID from the query parameter
	id := c.Query("id")
	uintId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, fmt.Errorf("error during convert into uint:%v", err))
		return
	}

	// Fetch the existing player from the database using the ID
	err = database.DB.Where("id=?", uintId).First(&player).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, fmt.Errorf("player not found:%v", err))
		return
	}

	// Bind the JSON data to the "player" variable
	if err := c.ShouldBindJSON(&player); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, fmt.Errorf("error during binding:%v", err))
		return
	}

	// Update the player in the database
	err = database.DB.Save(&player).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, fmt.Errorf("error during updating:%v", err))
		return
	}

	// Respond with the updated player information
	c.JSON(http.StatusOK, player)
}

// DeletePlayer deletes a player from the database based on the player's ID passed as a query parameter.
func DeletePlayer(c *gin.Context) {
	var player model.Player

	// Parse the player ID from the query parameter
	id := c.Query("id")
	uintId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, fmt.Errorf("error during convert into uint:%v", err))
		return
	}

	// Delete the player from the database
	err = database.DB.Where("id=?", uintId).Delete(&player).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, fmt.Errorf("error during deleting:%v", err))
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "delete successfully"})
}

// GetPlayers fetches and returns a list of all players in descending order of their IDs.
func GetPlayers(c *gin.Context) {
	var players []model.Player

	// Fetch all players from the database in descending order of their IDs
	err := database.DB.Order("id DESC").Find(&players).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, fmt.Errorf("error during getting all players:%v", err))
		return
	}

	c.JSON(http.StatusOK, players)
}

// GetPlayerByRank fetches and returns the player ranked by a given rank value (val).
func GetPlayerByRank(c *gin.Context) {
	var player model.Player

	// Parse the rank value from the query parameter
	rankStr := c.Query("val")
	rank, err := strconv.Atoi(rankStr)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, fmt.Errorf("error during convert into int:%v", err))
		return
	}

	// Fetch the player based on the rank from the database
	err = database.DB.Order("score DESC, id ASC").Offset(rank - 1).Limit(1).Find(&player).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, fmt.Errorf("error during get player by rank:%v", err))
		return
	}

	c.JSON(http.StatusOK, player)
}

// GetRandomPlayer fetches and returns a random player from the database.
func GetRandomPlayer(c *gin.Context) {
	var player model.Player
	var playerCount int64

	// Get the total count of players in the database
	database.DB.Model(&player).Count(&playerCount)

	// Generate a random player ID
	rand.Seed(time.Now().UnixNano())
	randomID := rand.Int63n(int64(playerCount)) + 1

	// Fetch the random player from the database
	database.DB.Where("id=?", int(randomID)).First(&player)

	c.JSON(http.StatusOK, player)
}
