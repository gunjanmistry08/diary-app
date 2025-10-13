package models

import (
	"time"

	"gorm.io/gorm"
)

type HabitStrategyString string

const (
	FrequencyDaily                 HabitStrategyString = "DAILY"
	FrequencyRepeating             HabitStrategyString = "REPEATING"
	FrequencyNumberOfDaysPerPeriod HabitStrategyString = "NUMBER_OF_DAYS_PER_PERIOD"
	FrequencySomeDaysOfTheWeek     HabitStrategyString = "SOME_DAYS_OF_THE_WEEK"
)

type HabitRepeating struct {
	gorm.Model
	HabitID  uint  `json:"habit_id"`
	Habit    Habit `gorm:"foreignKey:HabitID" json:"-"`
	Interval uint  `json:"interval"` // in days
}

type HabitNumberOfDaysPerPeriod struct {
	gorm.Model
	HabitID uint  `json:"habit_id"`
	Habit   Habit `gorm:"foreignKey:HabitID" json:"-"`
	Number  uint  `json:"number"`
	Period  uint  `json:"period"` // in days
}

type HabitDaysOftheWeek struct {
	gorm.Model
	HabitID   uint  `json:"habit_id"`
	Habit     Habit `gorm:"foreignKey:HabitID" json:"-"`
	Monday    bool  `json:"monday"`
	Tuesday   bool  `json:"tuesday"`
	Wednesday bool  `json:"wednesday"`
	Thursday  bool  `json:"thursday"`
	Friday    bool  `json:"friday"`
	Saturday  bool  `json:"saturday"`
	Sunday    bool  `json:"sunday"`
}

type HabitType string

const (
	Counting HabitType = "COUNTING"
	Boolean  HabitType = "BOOLEAN"
)

type HabitStrategy struct {
	gorm.Model
	Strategy  string              `json:"strategy"`
	Frequency HabitStrategyString `json:"frequency"`
	CreatedAt time.Time           `json:"created_at"`
	UpdatedAt time.Time           `json:"updated_at"`
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

func SeedHabitStrategies(db *gorm.DB) error {
	strategies := []HabitStrategy{
		{
			Strategy:  "Daily",
			Frequency: FrequencyDaily,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Strategy:  "Repeating",
			Frequency: FrequencyRepeating,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Strategy:  "Number of Days Per Period",
			Frequency: FrequencyNumberOfDaysPerPeriod,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Strategy:  "Some Days of the Week",
			Frequency: FrequencySomeDaysOfTheWeek,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	for _, s := range strategies {
		if err := db.Create(&s).Error; err != nil {
			return err
		}
	}
	return nil
}
