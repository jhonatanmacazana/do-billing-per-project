package utils

import "time"

func MonthsCountSince(createdAtTime time.Time) int {
	now := time.Now()
	months := 0
	month := createdAtTime.Month()
	for createdAtTime.Before(now) {
		createdAtTime = createdAtTime.Add(time.Hour * 24)
		nextMonth := createdAtTime.Month()
		if nextMonth != month {
			months++
		}
		month = nextMonth
	}

	return months
}

func ParseISOString(s string) (time.Time, error) {
	// layout := "2021-10-21T04:04:53.131546Z"
	t, err := time.Parse(time.RFC3339Nano, s)

	if err != nil {
		return time.Time{}, err
	}

	return t, nil
}
