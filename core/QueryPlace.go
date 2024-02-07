package core

import "strings"

type QueryPlace struct {
	CityName string
}

func NewQueryPlace(CityName string) QueryPlace {
	CityName = strings.TrimSpace(CityName)

	return QueryPlace{
		CityName,
	}
}