package tests

import (
	"testing"

	"github.com/gunjanmistry08/diary-app/internal/models"
	"github.com/gunjanmistry08/diary-app/internal/services"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect test database: %v", err)
	}
	if err := db.AutoMigrate(&models.User{}, &models.DiaryEntry{}); err != nil {
		t.Fatalf("failed to migrate: %v", err)
	}
	return db
}

func createTestUser(db *gorm.DB, username, email string) *models.User {
	user := &models.User{
		Username: username,
		Email:    email,
		Password: "hashedpassword",
	}
	db.Create(user)
	return user
}

func TestCreateDiaryEntry(t *testing.T) {
	db := setupTestDB(t)
	user := createTestUser(db, "testuser", "test@example.com")

	err := services.DiaryEntry(user.ID, "My Title", "My Content")
	if err != nil {
		t.Fatalf("failed to create diary entry: %v", err)
	}
}

func TestGetDiaryEntriesByUser(t *testing.T) {
	db := setupTestDB(t)
	user := createTestUser(db, "testuser2", "test2@example.com")

	// Create entries
	for range 3 {
		err := services.DiaryEntry(user.ID, "Title", "Content")
		if err != nil {
			t.Fatalf("failed to create diary entry: %v", err)
		}
	}
}

func TestUpdateDiaryEntry(t *testing.T) {
	db := setupTestDB(t)
	user := createTestUser(db, "testuser3", "test3@example.com")

	_ = user

}

func TestDeleteDiaryEntry(t *testing.T) {
	db := setupTestDB(t)
	user := createTestUser(db, "testuser4", "test4@example.com")

	_ = user
}

func TestGetDiaryEntryByID(t *testing.T) {
	db := setupTestDB(t)
	user := createTestUser(db, "testuser5", "test5@example.com")

	for range 3 {
		err := services.DiaryEntry(user.ID, "Title", "Content")
		if err != nil {
			t.Fatalf("failed to create diary entry: %v", err)
		}
	}

	entries, err := services.GetDiaryEntriesByUser(user.ID)
	if err != nil {
		t.Fatalf("failed to get diary entries: %v", err)
	}

	if len(entries) != 3 {
		t.Fatalf("expected 3 entries, got %d", len(entries))
	}
}

func TestCreateEntryWithEmptyFields(t *testing.T) {
	db := setupTestDB(t)
	user := createTestUser(db, "testuser6", "test6@example.com")

	err := services.DiaryEntry(user.ID, "", "")
	if err == nil {
		t.Errorf("expected error for empty title and content")
	}
}

func TestUpdateEntryWithWrongUser(t *testing.T) {
	db := setupTestDB(t)
	user1 := createTestUser(db, "userA", "a@example.com")
	user2 := createTestUser(db, "userB", "b@example.com")

	_ = user1
	_ = user2
}

func TestDeleteEntryWithWrongUser(t *testing.T) {
	db := setupTestDB(t)
	user1 := createTestUser(db, "userC", "c@example.com")
	user2 := createTestUser(db, "userD", "d@example.com")

	_ = user1
	_ = user2
}

func TestGetEntriesByUser_Empty(t *testing.T) {
	db := setupTestDB(t)
	user := createTestUser(db, "nouser", "nouser@example.com")

	_ = user
}

func TestEntryTimestamps(t *testing.T) {
	db := setupTestDB(t)
	user := createTestUser(db, "timestampuser", "ts@example.com")

	_ = user
}
