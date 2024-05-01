package domain

import "time"

var MapOfBirthdays = map[string]time.Time{
	"Березуцкий Иван": time.Date(2004, time.May, 14, 0, 0, 0, 0, time.UTC),
	"Пупкин Кирилл":   time.Date(2004, time.January, 5, 0, 0, 0, 0, time.UTC),
	"Вася Петров":     time.Date(2002, time.September, 15, 0, 0, 0, 0, time.UTC),
	"Иванов Илья":     time.Date(2004, time.May, 1, 20, 49, 0, 0, time.UTC),
}
