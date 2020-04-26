package entities

import "time"

//Player is
type Player struct {
	ID				uint32
	FirstName		string
	LastName		string
	IsRightHanded	bool
	BirthDate		time.Time
	CountryCode		string
}