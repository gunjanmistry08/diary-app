package models

import (
	"time"

	"gorm.io/gorm"
)

type HabitStrategyString string

const (
	FrequencyDaily                 HabitStrategyString = "daily"
	FrequencyRepeating             HabitStrategyString = "repeating"
	FrequencyNumberOfDaysPerPeriod HabitStrategyString = "number_of_days_per_period"
	FrequencySomeDaysOfTheWeek     HabitStrategyString = "some_days_of_the_week"
)

type HabitType string

const (
	Counting HabitType = "counting"
	Boolean  HabitType = "boolean"
)

type HabitStrategy struct {
	gorm.Model
	Strategy  string    `json:"strategy"`
	Frequency string    `json:"frequency"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Habit struct {
	gorm.Model
	Name            string        `json:"title"`
	Description     string        `json:"content"`
	UserID          uint          `json:"user_id"`
	User            User          `gorm:"foreignKey:UserID" json:"-"`
	HabitStrategyID uint          `json:"habit_strategy_id"`
	HabitStrategy   HabitStrategy `gorm:"foreignKey:HabitStrategyID" json:"-"`
	Type            HabitType     `gorm:"type:enum('counting', 'boolean')" json:"type"`
	CreatedAt       time.Time     `json:"created_at"`
	UpdatedAt       time.Time     `json:"updated_at"`
}

type HabitEntry struct {
	gorm.Model
	HabitID   uint      `json:"habit_id"`
	Habit     Habit     `gorm:"foreignKey:HabitID" json:"-"`
	EntryDate time.Time `json:"entry_date"`
	Note      string    `json:"note"`
	Count     uint      `json:"count"`
	UserID    uint      `json:"user_id"`
	User      User      `gorm:"foreignKey:UserID" json:"-"`
	CreatedAt time.Time `json:"created_at"`
}
