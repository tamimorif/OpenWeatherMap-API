package models

import "time"

type City struct {
	ID        uint
	Name      string
	CreatedAt time.Time
}

type Marks struct {
	ID          uint
	CitiesID    uint
	Temperature float64
	WindSpeed   float64
	Visibility  uint
	CreatedAt   time.Time
	//ChanceOfRain float64bh
}

//type Hourly struct {
//	Duration    []time.Duration
//	Tempareture []time.Time
//}
