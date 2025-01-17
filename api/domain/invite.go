package domain

import (
	"gorm.io/gorm"
	"time"
)

type Invite struct {
	gorm.Model
	Name             string      `json:"name"`
	BgColor          string      `json:"bgColor"`
	Color            string      `json:"color"`
	ImageUrl         string      `json:"imageUrl"`
	MainMessage      string      `json:"mainMessage"`
	SecondaryMessage string      `json:"secondaryMessage"`
	DatesOptions     []time.Time `json:"datesOptions"`
	DatePicker       time.Time   `json:"datePicked"`
	OptionsPlaces    []string    `json:"optionsPlaces"`
	PickedPlaces     string      `json:"pickedPlaces"`
	CreationIp       string      `json:"creationIp"`
	VisitedIps       []string    `json:"watchedIp"`
	Accepted         bool        `json:"response"`
}

type InviteRepository interface{}

type InviteUseCase interface{}
