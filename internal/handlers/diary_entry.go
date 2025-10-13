package handlers

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gunjanmistry08/diary-app/internal/services"
)

type CreateDiaryEntryRequest struct {
	UserID  uint   `json:"user_id" binding:"required"`
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type DiaryEntryResponse struct {
	ID      uint   `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func CreateDiaryEntry(c *gin.Context) {
	var req CreateDiaryEntryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("Error creating diary entry:", err)
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}
	err := services.DiaryEntry(req.UserID, req.Title, req.Content)
	if err != nil {
		log.Println("Error creating diary entry:", err)
		c.JSON(500, gin.H{"error": "Failed to create diary entry"})
		return
	}
	c.JSON(200, gin.H{"message": "Diary entry created"})
}

func GetDiaryEntriesByUser(c *gin.Context) {
	userIDParam := c.Param("user_id")
	var userID uint
	_, err := fmt.Sscan(userIDParam, &userID)
	if err != nil {
		log.Println("Error getting user ID:", err)
		c.JSON(400, gin.H{"error": "Invalid user ID"})
		return
	}
	entries, err := services.GetDiaryEntriesByUser(userID)

	var response []DiaryEntryResponse
	for _, entry := range entries {
		response = append(response, DiaryEntryResponse{
			ID:      entry.ID,
			Title:   entry.Title,
			Content: entry.Content,
		})
	}
	if err != nil {
		log.Println("Error getting diary entries:", err)
		c.JSON(500, gin.H{"error": "Failed to get diary entries"})
		return
	}
	c.JSON(200, gin.H{"entries": response})
}

func RegisterDiaryRoutes(g *gin.RouterGroup) {
	// Use route group for diary entries
	g.POST("/entries", CreateDiaryEntry)
	g.GET("/entries/user/:user_id", GetDiaryEntriesByUser)
}
