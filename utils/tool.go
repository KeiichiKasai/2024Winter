package utils

import "time"

func ParseBirthday(birthdayStr string) (time.Time, error) {
	layout := "2006-01-02"
	birthday, err := time.Parse(layout, birthdayStr)
	if err != nil {
		return time.Time{}, err
	}
	return birthday, nil
}
