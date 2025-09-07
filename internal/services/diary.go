package services

import (
	"log"

	"github.com/gunjanmistry08/diary-app/internal/database"
	"github.com/gunjanmistry08/diary-app/internal/models"
)

func DiaryEntry(userID uint, title, content string) error {

	result := database.DB.Create(&models.DiaryEntry{
		UserID:  userID,
		Title:   title,
		Content: content,
	})
	if result.Error != nil {
		log.Printf("failed to create diary entry: %v", result.Error)
		return result.Error
	}
	return nil
}

func GetDiaryEntriesByUser(userID uint) ([]models.DiaryEntry, error) {
	var entries []models.DiaryEntry
	result := database.DB.Where("user_id = ?", userID).Select("id, title, SUBSTR(content, 1, 50) as content").Order("created_at DESC").Find(&entries)
	if result.Error != nil {
		log.Printf("failed to get diary entries: %v", result.Error)
		return nil, result.Error
	}
	return entries, nil
}
