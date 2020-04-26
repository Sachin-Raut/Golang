package entities

import "time"

//Ranking is
type Ranking struct {
	PlayerID 		uint32
	RankingDate		time.Time
	RankingNumber	uint16
	RankingPoints	float64
}