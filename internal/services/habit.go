package services

import (
	"errors"
	"log"

	// "slices"
	"time"

	"github.com/gunjanmistry08/diary-app/internal/database"
	"github.com/gunjanmistry08/diary-app/internal/models"
)

func CreateHabit(userID, habitStrategyID, numberofDays, period uint, title, description string, type_ models.HabitType, days []string) error {
	var habit models.Habit
	result := database.DB.Where("user_id = ? AND title = ?", userID, title).First(&habit)
	if result.Error == nil {
		return errors.New("habit already exists")
	}

	if type_ != models.Counting && type_ != models.Boolean {
		return errors.New("invalid habit type")
	}

	// check the strategy exists

	var strategy models.HabitStrategy
	if err := database.DB.First(&strategy, habitStrategyID).Error; err != nil {
		log.Printf("error finding habit strategy: %v", err)
		return errors.New("habit strategy not found")
	}

	habit = models.Habit{
		Name:            title,
		Description:     description,
		UserID:          userID,
		HabitStrategyID: habitStrategyID,
		Type:            type_,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	tx := database.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Create(&habit).Error; err != nil {
		log.Printf("error creating habit: %v", err)
		return err
	}

	if strategy.Frequency == models.FrequencyNumberOfDaysPerPeriod {
		if numberofDays == 0 || period == 0 {
			tx.Rollback()
			return errors.New("number of days and period must be greater than 0 for this strategy")
		}
		habitNumberOfDays := models.HabitNumberOfDaysPerPeriod{
			HabitID: habit.ID,
			Number:  numberofDays,
			Period:  period,
		}

		if err := tx.Create(&habitNumberOfDays).Error; err != nil {
			tx.Rollback()
			log.Printf("error creating habit number of days: %v", err)
			return err
		}
	}

	if strategy.Frequency == models.FrequencySomeDaysOfTheWeek {
		if len(days) == 0 {
			tx.Rollback()
			return errors.New("at least one day must be selected for this strategy")
		}

		dayMap := map[string]*bool{
			"monday":    new(bool),
			"tuesday":   new(bool),
			"wednesday": new(bool),
			"thursday":  new(bool),
			"friday":    new(bool),
			"saturday":  new(bool),
			"sunday":    new(bool),
		}
		for _, day := range days {
			if b, ok := dayMap[day]; ok {
				*b = true
			}
		}
		habitDay := models.HabitDaysOftheWeek{
			HabitID:   habit.ID,
			Monday:    *dayMap["monday"],
			Tuesday:   *dayMap["tuesday"],
			Wednesday: *dayMap["wednesday"],
			Thursday:  *dayMap["thursday"],
			Friday:    *dayMap["friday"],
			Saturday:  *dayMap["saturday"],
			Sunday:    *dayMap["sunday"],
		}

		// habitDay := models.HabitDaysOftheWeek{
		// 	HabitID:   habit.ID,
		// 	Monday:    slices.Contains(days, "monday"),
		// 	Tuesday:   slices.Contains(days, "tuesday"),
		// 	Wednesday: slices.Contains(days, "wednesday"),
		// 	Thursday:  slices.Contains(days, "thursday"),
		// 	Friday:    slices.Contains(days, "friday"),
		// 	Saturday:  slices.Contains(days, "saturday"),
		// 	Sunday:    slices.Contains(days, "sunday"),
		// }

		if err := tx.Create(&habitDay).Error; err != nil {
			tx.Rollback()
			log.Printf("error creating habit days of the week: %v", err)
			return err
		}
	}

	if strategy.Frequency == models.FrequencyRepeating {
		if numberofDays == 0 {
			tx.Rollback()
			return errors.New("interval must be greater than 0 for this strategy")
		}
		habitRepeating := models.HabitRepeating{
			HabitID:  habit.ID,
			Interval: numberofDays,
		}

		if err := tx.Create(&habitRepeating).Error; err != nil {
			tx.Rollback()
			log.Printf("error creating habit repeating: %v", err)
			return err
		}
	}

	if strategy.Frequency == models.FrequencyDaily {
		// no additional fields to set
	}

	tx.Commit()
	return nil
}

func CreateHabitEntry(userID, habitID uint, date time.Time, completed bool, count uint) error {
	return nil
}
