package models

import "time"

type City struct {
	ID        uint      `json:"id" gorm:"primarykey"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type Marks struct {
	ID          uint      `json:"id" gorm:"primarykey"`
	CityID      uint      `json:"city_id"`
	Temperature float64   `json:"temperature"`
	WindSpeed   float64   `json:"wind_speed"`
	Visibility  uint      `json:"visibility"`
	CreatedAt   time.Time `json:"created_at"`
	//ChanceOfRain float64bh
}

//type Hourly struct {
//	Duration    []time.Duration
//	Tempareture []time.Time
//}
